package common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DecoratedBasicTypeInterface interface {
	Set(value interface{}) error
	GetNativePresentation() (interface{}, bool)
	Get() (interface{}, bool)
}

type DecoratedBasicTypeGetterInterface interface {
	GetNativePresentation() (interface{}, bool)
	Get() (interface{}, bool)
}

type tagInfo struct {
	name    string
	isId    bool
	flatten bool
	ignore  bool
}

// Returns string for name, boolean whether this is id field, boolean whether this struct needs to be flattened to parent namespace
func getNameFromTag(field reflect.StructField) tagInfo {
	fieldName, tags := ParseTag(field.Tag.Get("provider"))
	if fieldName == "" {
		rawFieldName, _ := ParseTag(field.Tag.Get("json"))
		fieldName = ToSnakeCase(rawFieldName)
	}
	if fieldName == "" {
		fieldName = field.Name
	}
	return tagInfo{
		name:    fieldName,
		isId:    tags.Contains("id"),
		flatten: tags.Contains("flatten"),
		ignore:  tags.Contains("ignore"),
	}
}

func isBasic(elemType reflect.Type) bool {
	return reflect.PtrTo(elemType).Implements(reflect.TypeOf((*DecoratedBasicTypeInterface)(nil)).Elem())
}

func unmarshallNode(fieldType reflect.Type, fieldValue reflect.Value, getter func(tagInfo, string) (interface{}, bool), tag tagInfo, fieldName string) error {
	switch fieldType.Kind() {
	case reflect.Bool:
		srcValue, exists := getter(tag, fieldName)
		if !exists {
			return nil
		}
		boolValue, ok := srcValue.(bool)
		if !ok {
			return CreateErrorGeneric("Wrong type for bool field " + fieldName)
		}
		fieldValue.SetBool(boolValue)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		srcValue, exists := getter(tag, fieldName)
		if !exists {
			return nil
		}
		intValue := reflect.ValueOf(srcValue).Int()
		fieldValue.SetInt(intValue)

	case reflect.Float32, reflect.Float64:
		srcValue, exists := getter(tag, fieldName)
		if !exists {
			return nil
		}
		floatValue := reflect.ValueOf(srcValue).Float()
		fieldValue.SetFloat(floatValue)

	case reflect.String:
		srcValue, exists := getter(tag, fieldName)
		if !exists {
			return nil
		}
		stringValue, ok := srcValue.(string)
		if !ok {
			return CreateErrorGeneric("Wrong type for string field " + fieldName)
		}
		fieldValue.SetString(stringValue)

	case reflect.Ptr:
		// srcValue itself is not a pointer, it is coming from Terraform
		// This is how we map existing keys to non-null pointers
		_, exists := getter(tag, fieldName)
		if !exists {
			return nil
		}
		object := reflect.New(fieldType.Elem())
		if err := unmarshallNode(fieldType.Elem(), object.Elem(), getter, tag, fieldName); err != nil {
			return err
		}

		fieldValue.Set(object)

	case reflect.Struct:
		if _, ok := fieldValue.Interface().(time.Time); ok {
			srcValue, exists := getter(tag, fieldName)
			if !exists {
				return nil
			}

			timeValue, err := time.Parse(time.RFC3339, srcValue.(string))
			if err != nil {
				return CreateErrorGeneric("Wrong type for time field " + fieldName)
			}

			fieldValue.Set(reflect.ValueOf(timeValue))
		} else {
			if basic, ok := fieldValue.Addr().Interface().(DecoratedBasicTypeInterface); ok {
				srcValue, exists := getter(tag, fieldName)
				if !exists {
					return nil
				}
				return basic.Set(srcValue)
			}
			for i := 0; i < fieldType.NumField(); i++ {
				subType := fieldType.Field(i).Type
				subValue := fieldValue.Field(i)
				subTag := getNameFromTag(fieldType.Field(i))
				structName := fieldName
				if tag.flatten {
					structName = fieldName[:strings.LastIndex(fieldName, ".")+1]
				}
				if len(structName) > 0 && structName[len(structName)-1:] != "." {
					structName += "."
				}
				structName += subTag.name
				if err := unmarshallNode(subType, subValue, getter, subTag, structName); err != nil {
					return err
				}
			}
		}

	case reflect.Slice:
		srcValue, exists := getter(tag, fieldName)
		if !exists {
			return nil
		}
		sliceValue, ok := srcValue.([]interface{})
		if !ok {
			return CreateErrorGeneric("Wrong type for slice field " + fieldName)
		}
		elemType := fieldType.Elem()
		fieldValue.Set(reflect.MakeSlice(fieldType, len(sliceValue), len(sliceValue)))
		for i := range sliceValue {
			elemValue := fieldValue.Index(i)
			if err := unmarshallNode(elemType, elemValue, getter, tag, fieldName+"."+strconv.Itoa(i)); err != nil {
				return err
			}
		}

	case reflect.Map:
		// Terraform only supports maps of basic types, thus we need to deal maps of complex types as json
		// we only support maps with string keys as a simplification
		elemType := fieldType.Elem()
		elemKind := elemType.Kind()
		srcValue, exists := getter(tag, fieldName)
		if !exists {
			return nil
		}
		if elemKind == reflect.Bool || elemKind == reflect.Int || elemKind == reflect.String || isBasic(elemType) {
			// input is map
			srcMap, ok := srcValue.(map[string]interface{})
			if !ok {
				return CreateErrorGeneric("Wrong type for map field " + fieldName)
			}
			mapValue := reflect.MakeMap(fieldType)
			for key := range srcMap {
				elemValue := reflect.New(elemType).Elem()
				if err := unmarshallNode(elemType, elemValue, getter, tag, fieldName+"."+key); err != nil {
					return err
				}
				mapValue.SetMapIndex(reflect.ValueOf(key), elemValue)
			}
			fieldValue.Set(mapValue)
		} else {
			// input is json
			stringValue, ok := srcValue.(string)
			if !ok {
				return CreateErrorGeneric("Wrong type for map (other type) field " + fieldName)
			}
			mapValue := reflect.New(fieldType)
			if err := json.Unmarshal(([]byte)(stringValue), mapValue.Interface()); err != nil {
				return WrapErrorGeneric(err, "json Unmarshal failed for field "+fieldName)
			}
			fieldValue.Set(mapValue.Elem())
		}

	case reflect.Interface:
		// free form json data, it is usable as long as it is a string of json which unmarshalls into map[string]interface{}
		srcValue, exists := getter(tag, fieldName)
		if !exists {
			return nil
		}
		stringValue, ok := srcValue.(string)
		if !ok {
			return CreateErrorGeneric("Wrong type for interface field " + fieldName)
		}
		var jsonValue map[string]interface{}
		if err := json.Unmarshal(([]byte)(stringValue), &jsonValue); err != nil {
			return WrapErrorGeneric(err, "json Unmarshal for interface field '"+fieldName+"' failed from '"+stringValue+"'")
		}
		fieldValue.Set(reflect.ValueOf(jsonValue))

	default:
		return CreateErrorGeneric("Unknown field type for unmarshalling " + fieldName)
	}

	return nil
}

// From structure (client) presentation to provider (resourceData)
func UnmarshalSchema(dest interface{}, resourceData *schema.ResourceData) error {
	destType := reflect.TypeOf(dest).Elem()
	destValue := reflect.Indirect(reflect.ValueOf(dest))

	for i := 0; i < destType.NumField(); i++ {
		field := destType.Field(i)

		topLevelGetter := func(tag tagInfo, name string) (interface{}, bool) {
			if tag.isId {
				srcValue := resourceData.Id()
				// empty id (resource not created)
				if srcValue == "" {
					return nil, false
				}
				return srcValue, true
			} else {
				srcValue, exists := resourceData.GetOk(name)
				if !exists {
					// empty/nil values will be defaults
					return nil, false
				}
				return srcValue, true
			}
		}

		tag := getNameFromTag(field)
		if tag.ignore {
			continue
		}
		fieldValue := destValue.Field(i)
		if !fieldValue.CanSet() {
			return CreateErrorGeneric("Unmarshalling failed for field " + tag.name)
		}

		if err := unmarshallNode(field.Type, fieldValue, topLevelGetter, tag, tag.name); err != nil {
			return err
		}
	}

	return nil
}

func marshallNode(setter func(string, interface{}) error, fieldName string, fieldType reflect.Type, fieldValue reflect.Value, flatten bool) error {

	switch fieldType.Kind() {
	case reflect.Bool:
		if err := setter(fieldName, fieldValue.Bool()); err != nil {
			return err
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if err := setter(fieldName, fieldValue.Int()); err != nil {
			return err
		}

	case reflect.Float32, reflect.Float64:
		if err := setter(fieldName, fieldValue.Float()); err != nil {
			return err
		}

	case reflect.String:
		if err := setter(fieldName, fieldValue.String()); err != nil {
			return err
		}

	case reflect.Ptr:
		if fieldValue.IsNil() {
			return setter(fieldName, nil)
		}
		ptrSetter := func(name string, value interface{}) error {
			// pointer to the type of value
			newValue := reflect.New(reflect.TypeOf(value))
			// set the value of newValue to value of 'value' interface
			newValue.Elem().Set(reflect.ValueOf(value))
			return setter(name, newValue.Interface())
		}

		if err := marshallNode(ptrSetter, fieldName, fieldType.Elem(), reflect.Indirect(fieldValue), flatten); err != nil {
			return err
		}

	case reflect.Struct:
		if _, ok := fieldValue.Interface().(time.Time); ok {
			if err := setter(fieldName, fieldValue.Interface().(time.Time).Format(time.RFC3339)); err != nil {
				return err
			}
		} else {
			if basic, ok := fieldValue.Addr().Interface().(DecoratedBasicTypeInterface); ok {
				if value, ok := basic.Get(); ok {
					return setter(fieldName, value)
				} else {
					// TF expects typed nil!
					var value *interface{} = nil
					return setter(fieldName, value)
				}
			}

			mapTarget := make(map[string]interface{})
			var subSetter func(string, interface{}) error
			if flatten {
				subSetter = setter
			} else {
				subSetter = func(name string, value interface{}) error {
					mapTarget[name[strings.LastIndex(name, ".")+1:]] = value
					return nil
				}
			}

			for i := 0; i < fieldType.NumField(); i++ {
				subType := fieldType.Field(i).Type
				subValue := fieldValue.Field(i)
				tag := getNameFromTag(fieldType.Field(i))
				if tag.ignore {
					continue
				}
				var combinedName string
				if flatten {
					dotIndex := strings.LastIndex(fieldName, ".")
					if dotIndex == -1 {
						dotIndex = 0
					}
					combinedName = fieldName[0:dotIndex]
					if len(combinedName) != 0 {
						combinedName += "."
					}
					combinedName += tag.name
				} else {
					combinedName = fieldName + "." + tag.name
				}
				if err := marshallNode(subSetter, combinedName, subType, subValue, tag.flatten); err != nil {
					return err
				}
			}
			if flatten {
				return nil
			}
			return setter(fieldName, mapTarget)
		}

	case reflect.Slice:
		if fieldValue.IsNil() {
			return setter(fieldName, nil)
		}
		sliceTarget := make([]interface{}, fieldValue.Len(), fieldValue.Len()) //nolint
		sliceSetter := func(name string, value interface{}) error {
			index, err := strconv.ParseInt(name[strings.LastIndex(name, ".")+1:], 10, 32)
			if err != nil {
				return CreateErrorGeneric("Invalid index for field " + name)
			}

			sliceTarget[index] = value
			return nil
		}
		elemType := fieldType.Elem()
		for i := 0; i < fieldValue.Len(); i++ {
			if err := marshallNode(sliceSetter, fieldName+"."+strconv.Itoa(i), elemType, fieldValue.Index(i), flatten); err != nil {
				return err
			}
		}
		return setter(fieldName, sliceTarget)

	case reflect.Map:
		if fieldValue.IsNil() {
			return setter(fieldName, nil)
		}
		// same as in unmarshalling, we can only use map for the basic types
		elemKind := fieldType.Elem().Kind()
		if elemKind == reflect.Bool || elemKind == reflect.Int || elemKind == reflect.String || isBasic(fieldType.Elem()) {
			mapTarget := make(map[string]interface{})
			mapSetter := func(name string, value interface{}) error {
				mapTarget[name[strings.LastIndex(name, ".")+1:]] = value
				return nil
			}
			elemType := fieldType.Elem()
			iter := fieldValue.MapRange()
			for iter.Next() {
				// Map items are not addressable. and this is needed for complex types -> make a copy
				value := reflect.New(elemType).Elem()
				value.Set(iter.Value())
				if err := marshallNode(mapSetter, fieldName+"."+iter.Key().String(), elemType, value, flatten); err != nil {
					return err
				}
			}
			return setter(fieldName, mapTarget)
		} else {
			byteValue, err := json.Marshal(fieldValue.Interface())
			if err != nil {
				return WrapErrorGeneric(err, "Json marshaling failed for field "+fieldName)
			}
			if err := setter(fieldName, string(byteValue)); err != nil {
				return err
			}
		}

	case reflect.Interface:
		if fieldValue.IsNil() {
			return setter(fieldName, nil)
		}
		// free form json
		byteValue, err := json.Marshal(fieldValue.Interface().(map[string]interface{}))
		if err != nil {
			return WrapErrorGeneric(err, "Json marshaling failed for field "+fieldName)
		}
		if err := setter(fieldName, string(byteValue)); err != nil {
			return err
		}

	default:
		errStr := fmt.Sprint("Unknown field type for marshalling, fieldName: ", fieldName, "\n fieldType: ", fieldType.Kind(), ", fieldValue: ", fieldValue)
		return CreateErrorGeneric(errStr)
	}
	return nil
}

//From provider (resourceData) to structure (client) presentation
func MarshalSchema(resourceData *schema.ResourceData, src interface{}) error {
	srcType := reflect.TypeOf(src).Elem()
	srcValue := reflect.Indirect(reflect.ValueOf(src))
	for i := 0; i < srcType.NumField(); i++ {
		tag := getNameFromTag(srcType.Field(i))
		if tag.ignore {
			continue
		}
		fieldType := srcType.Field(i).Type
		field := reflect.Indirect(srcValue).Field(i)
		topLevelSetter := func(name string, value interface{}) error {
			if tag.isId {
				if stringValue, ok := value.(*string); ok {
					resourceData.SetId(*stringValue)
					return nil
				}
				if stringValue, ok := value.(string); ok {
					resourceData.SetId(stringValue)
					return nil
				}
				return CreateErrorGeneric("Wrong type of id field")
			} else {
				val := resourceData.Get(name)

				// If the resource data does not contain the property, ignore it.
				if reflect.TypeOf(val) == nil {
					return nil
				}

				// if the resource data expected type is a string and the actual data type is not, json encode it
				if reflect.TypeOf(val).Kind() == reflect.String && reflect.TypeOf(value) != nil && reflect.TypeOf(value).Kind() != reflect.String {
					marshaledVal, _ := json.Marshal(value)
					// clean up the encoded string
					stringEncodedVal := strings.Replace(string(marshaledVal), "\"", "", -1)
					return resourceData.Set(name, stringEncodedVal)
				}

				return resourceData.Set(name, value)
			}
		}

		if err := marshallNode(topLevelSetter, tag.name, fieldType, field, tag.flatten); err != nil {
			return err
		}
	}
	return nil
}

var removeSpecial = regexp.MustCompile("[^a-zA-Z0-9]+")
var matchFirstCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := removeSpecial.ReplaceAllString(str, "_")
	snake = matchFirstCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
