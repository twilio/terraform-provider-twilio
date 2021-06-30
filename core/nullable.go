package core

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type NullableString struct {
	Valid bool
	Value string
}

type NullableBool struct {
	Valid bool
	Value bool
}

type NullableInt struct {
	Valid bool
	Value int64
}

type NullableFloat struct {
	Valid bool
	Value float64
}

// ---

// DecoratedBasicTypeInterface
func (nullable NullableString) Get() (interface{}, bool) {
	if !nullable.Valid {
		return nil, false
	}
	return "[" + nullable.Value + "]", true
}

func (nullable NullableString) GetNativePresentation() (interface{}, bool) {
	if !nullable.Valid {
		return nil, false
	}
	return nullable.Value, true
}

func (nullable *NullableString) Set(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return CreateErrorGeneric(fmt.Sprintf("Invalid type for Set NullablString '%v'", reflect.TypeOf(value)))
	}
	if str == "" {
		nullable.Valid = false
		nullable.Value = ""
		return nil
	}
	if len(str) < 2 || str[0] != '[' || str[len(str)-1] != ']' {
		nullable.Valid = false
		nullable.Value = ""
		return CreateErrorGeneric(fmt.Sprintf("Invalid value for Set NullableString '%s'", str))
	}
	nullable.Valid = true
	nullable.Value = str[1 : len(str)-1]
	return nil
}

// json marshaling
func (nullable NullableString) MarshalJSON() ([]byte, error) {
	if !nullable.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nullable.Value)
}

func (nullable *NullableString) UnmarshalJSON(buffer []byte) error {
	nullable.Valid = false
	nullable.Value = ""

	var value interface{}
	if err := json.Unmarshal(buffer, &value); err != nil {
		return err
	}
	if value == nil {
		return nil
	}
	if conv, ok := value.(string); ok {
		nullable.Valid = true
		nullable.Value = conv
		return nil
	}
	return CreateErrorGeneric(fmt.Sprintf("Invalid type for UnmarshalJSON NullableString '%v'", reflect.TypeOf(value)))
}

// ---

func (nullable NullableBool) Get() (interface{}, bool) {
	if !nullable.Valid {
		return nil, false
	}
	return strconv.FormatBool(nullable.Value), true
}

func (nullable NullableBool) GetNativePresentation() (interface{}, bool) {
	if !nullable.Valid {
		return nil, false
	}
	return nullable.Value, true
}

func (nullable *NullableBool) Set(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return CreateErrorGeneric(fmt.Sprintf("Invalid type for Set NullableBool '%v'", reflect.TypeOf(value)))
	}
	if str == "" {
		nullable.Valid = false
		nullable.Value = false
		return nil
	}
	if str == "false" {
		nullable.Valid = true
		nullable.Value = false
		return nil

	}
	if str == "true" {
		nullable.Valid = true
		nullable.Value = true
		return nil

	}
	return CreateErrorGeneric(fmt.Sprintf("Invalid value for Set NullableBool '%s'", str))
}

// json marshaling
func (nullable NullableBool) MarshalJSON() ([]byte, error) {
	if !nullable.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nullable.Value)
}

func (nullable *NullableBool) UnmarshalJSON(buffer []byte) error {
	nullable.Valid = false
	nullable.Value = false

	var value interface{}
	if err := json.Unmarshal(buffer, &value); err != nil {
		return err
	}
	if value == nil {
		return nil
	}
	if conv, ok := value.(bool); ok {
		nullable.Valid = true
		nullable.Value = conv
		return nil
	}
	return CreateErrorGeneric(fmt.Sprintf("Invalid type for UnmarshalJSON NullableBool '%v'", reflect.TypeOf(value)))
}

// ---

func (nullable NullableInt) Get() (interface{}, bool) {
	if !nullable.Valid {
		return nil, false
	}
	return strconv.FormatInt(nullable.Value, 10), true
}

func (nullable NullableInt) GetNativePresentation() (interface{}, bool) {
	if !nullable.Valid {
		return nil, false
	}
	return nullable.Value, true
}

func (nullable *NullableInt) Set(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return CreateErrorGeneric(fmt.Sprintf("Invalid type for Set NullableInt '%v'", reflect.TypeOf(value)))
	}
	if str == "" {
		nullable.Valid = false
		nullable.Value = 0
		return nil
	}
	conv, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		nullable.Valid = false
		nullable.Value = 0
		return WrapErrorGeneric(err, "Failed to convert integer value for NullableInt")
	}
	nullable.Valid = true
	nullable.Value = conv
	return nil
}

// json marshaling
func (nullable NullableInt) MarshalJSON() ([]byte, error) {
	if !nullable.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nullable.Value)
}

func (nullable *NullableInt) UnmarshalJSON(buffer []byte) error {
	nullable.Valid = false
	nullable.Value = 0

	var value interface{}
	if err := json.Unmarshal(buffer, &value); err != nil {
		return err
	}
	if value == nil {
		return nil
	}
	if conv, ok := value.(float64); ok {
		nullable.Valid = true
		nullable.Value = (int64)(conv)
		return nil
	}
	return CreateErrorGeneric(fmt.Sprintf("Invalid type for UnmarshalJSON NullableInt '%v'", reflect.TypeOf(value)))
}

// ---

func (nullable NullableFloat) Get() (interface{}, bool) {
	if !nullable.Valid {
		return nil, false
	}
	return fmt.Sprintf("%v", nullable.Value), true
}

func (nullable NullableFloat) GetNativePresentation() (interface{}, bool) {
	if !nullable.Valid {
		return nil, false
	}
	return nullable.Value, true
}

func (nullable *NullableFloat) Set(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return CreateErrorGeneric(fmt.Sprintf("Invalid type for Set NullableFloat '%v'", reflect.TypeOf(value)))
	}
	if str == "" {
		nullable.Valid = false
		nullable.Value = 0
		return nil
	}
	conv, err := strconv.ParseFloat(str, 64)
	if err != nil {
		nullable.Valid = false
		nullable.Value = 0
		return WrapErrorGeneric(err, "Failed to convert integer value for NullableFloat")
	}
	nullable.Valid = true
	nullable.Value = conv
	return nil
}

// json marshaling
func (nullable NullableFloat) MarshalJSON() ([]byte, error) {
	if !nullable.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nullable.Value)
}

func (nullable *NullableFloat) UnmarshalJSON(buffer []byte) error {
	nullable.Valid = false
	nullable.Value = 0

	var value interface{}
	if err := json.Unmarshal(buffer, &value); err != nil {
		return err
	}
	if value == nil {
		return nil
	}
	if conv, ok := value.(float64); ok {
		nullable.Valid = true
		nullable.Value = conv
		return nil
	}
	return CreateErrorGeneric(fmt.Sprintf("Invalid type for UnmarshalJSON NullableFloat '%v'", reflect.TypeOf(value)))
}
