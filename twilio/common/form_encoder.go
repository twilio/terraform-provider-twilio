package common

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func processParamValue(paramType string, form *url.Values, field reflect.StructField, value reflect.Value) error {
	tag := field.Tag.Get(paramType)
	name, opts := ParseTag(tag)
	if name == "" {
		name = field.Name
	}

	if opts.Contains("ignore") {
		return nil
	}

	if value.Kind() == reflect.Ptr && value.IsNil() {
		if opts.Contains("omitempty") {
			return nil
		} else {
			return CreateErrorGeneric(fmt.Sprintf("Field: %s value cannot be empty or nil", name))
		}
	}

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Bool:
		form.Add(name, strconv.FormatBool(value.Bool()))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		form.Add(name, strconv.FormatInt(value.Int(), 10))

	case reflect.Float32, reflect.Float64:
		form.Add(name, fmt.Sprintf("%v", value))

	case reflect.String:
		form.Add(name, value.String())

	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			if err := processParamValue(paramType, form, field, value.Index(i)); err != nil {
				return nil
			}
		}

	case reflect.Interface:
		bytes, err := json.Marshal(value.Interface())
		if err != nil {
			return WrapErrorGeneric(err, "Error marshalling "+name)
		}
		form.Add(name, string(bytes))

	case reflect.Struct:
		if basic, ok := value.Interface().(DecoratedBasicTypeGetterInterface); ok {
			if value, ok := basic.GetNativePresentation(); ok {
				if err := processParamValue(paramType, form, field, reflect.ValueOf(value)); err != nil {
					return nil
				}
			}
		} else if opts.Contains("flatten") {
			for i := 0; i < field.Type.NumField(); i++ {
				subField := field.Type.Field(i)
				subValue := value.Field(i)

				if err := processParamValue(paramType, form, subField, subValue); err != nil {
					return err
				}
			}
		} else {
			return CreateErrorGeneric(fmt.Sprintf("Field: %s is of invalid struct type", name))
		}

	default:
		return CreateErrorGeneric(fmt.Sprintf("Field: %s is of invalid type", name))
	}
	return nil
}

func FormEncoder(src interface{}) (url.Values, error) {
	if src == nil {
		return nil, CreateErrorGeneric("Nil form provided")
	}

	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Struct {
		return nil, CreateErrorGeneric(fmt.Sprintf("FormEncoder expects struct input. Got %v", srcType.Kind()))
	}

	form := url.Values{}
	srcValue := reflect.ValueOf(src)
	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)
		fieldValue := srcValue.Field(i)

		if err := processParamValue("form", &form, field, fieldValue); err != nil {
			return nil, err
		}
	}
	return form, nil
}
