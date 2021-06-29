package core

import (
	"fmt"
	"net/url"
	"reflect"
)

func QueryEncoder(src interface{}) (string, error) {
	if src == nil {
		return "", CreateErrorGeneric("Nil query provided")
	}

	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Struct {
		return "", CreateErrorGeneric(fmt.Sprintf("QueryEncoder expects struct input. Got %v", srcType.Kind()))
	}

	values := url.Values{}
	srcValue := reflect.ValueOf(src)
	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)
		fieldValue := srcValue.Field(i)

		if err := processParamValue("query", &values, field, fieldValue); err != nil {
			return "", err
		}
	}

	paramString := ""
	for key, value := range values {
		if len(paramString) > 0 {
			paramString += "&"
		} else {
			paramString = "?"
		}
		paramString += key + "=" + url.PathEscape(value[0])
	}
	return paramString, nil
}
