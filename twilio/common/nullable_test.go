package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullableStringGet(t *testing.T) {

	T1 := NullableString{Valid: true, Value: "str"}

	value, ok := T1.Get()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "[str]", value, "got value is not valid")
}

func TestNullableStringGetEmpty(t *testing.T) {

	T1 := NullableString{Valid: true, Value: ""}

	value, ok := T1.Get()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "[]", value, "got value is not valid")
}

func TestNullableStringGetNil(t *testing.T) {

	T1 := NullableString{Valid: false, Value: ""}

	_, ok := T1.Get()
	assert.Equal(t, false, ok, "got value is valid")
}

func TestNullableStringGetNativePresentation(t *testing.T) {

	T1 := NullableString{Valid: true, Value: "str"}

	value, ok := T1.GetNativePresentation()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "str", value, "got value is not valid")
}

func TestNullableStringGetNativePresentationEmpty(t *testing.T) {

	T1 := NullableString{Valid: true, Value: ""}

	value, ok := T1.GetNativePresentation()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "", value, "got value is not valid")
}

func TestNullableStringGetNativePresentationNil(t *testing.T) {

	T1 := NullableString{Valid: false, Value: ""}

	_, ok := T1.GetNativePresentation()
	assert.Equal(t, false, ok, "got value is valid")
}

func TestNullableStringSet(t *testing.T) {
	var test NullableString

	err := test.Set("[str]")
	assert.Nil(t, err, "test did not set")
	assert.True(t, test.Valid, "marshaled value is not valid")
	assert.Equal(t, "str", test.Value, "marshaled value is not valid")
}

func TestNullableStringSetEmpty(t *testing.T) {
	var test NullableString

	err := test.Set("[]")
	assert.Nil(t, err, "test did not set")
	assert.True(t, test.Valid, "marshaled value is not valid")
	assert.Equal(t, "", test.Value, "marshaled value is not valid")
}

func TestNullableStringSetNil(t *testing.T) {
	var test NullableString

	err := test.Set("")
	assert.Nil(t, err, "test did not set")
	assert.False(t, test.Valid, "marshaled value is valid")
}

func TestNullableStringMarshal(t *testing.T) {

	type test struct {
		T1 NullableString
	}

	jsonValue := test{T1: NullableString{Valid: true, Value: "str"}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":\"str\"}", string(jsonString), "marshaled value is not valid")
}

func TestNullableStringMarshalEmpty(t *testing.T) {

	type test struct {
		T1 NullableString
	}

	jsonValue := test{T1: NullableString{Valid: true, Value: ""}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":\"\"}", string(jsonString), "marshaled value is not valid")
}

func TestNullableStringMarshalNil(t *testing.T) {

	type test struct {
		T1 NullableString
	}

	jsonValue := test{T1: NullableString{Valid: false, Value: ""}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":null}", string(jsonString), "marshaled value is not valid")
}

func TestNullableStringUmarshal(t *testing.T) {

	type test struct {
		T1 NullableString
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":\"str\"}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.True(t, jsonValue.T1.Valid, "marshaled value is not valid")
	assert.Equal(t, "str", jsonValue.T1.Value, "marshaled value is not valid")
}

func TestNullableStringUmarshalEmpty(t *testing.T) {

	type test struct {
		T1 NullableString
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":\"\"}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.True(t, jsonValue.T1.Valid, "marshaled value is not valid")
	assert.Equal(t, "", jsonValue.T1.Value, "marshaled value is not valid")
}

func TestNullableStringUmarshalNil(t *testing.T) {

	type test struct {
		T1 NullableString
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":null}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.False(t, jsonValue.T1.Valid, "marshaled value is valid")
}

// ---

func TestNullableBoolGet(t *testing.T) {

	T1 := NullableBool{Valid: true, Value: true}

	value, ok := T1.Get()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "true", value, "got value is not valid")
}

func TestNullableBoolGetFalse(t *testing.T) {

	T1 := NullableBool{Valid: true, Value: false}

	value, ok := T1.Get()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "false", value, "got value is not valid")
}

func TestNullableBoolGetNil(t *testing.T) {

	T1 := NullableBool{Valid: false, Value: false}

	_, ok := T1.Get()
	assert.Equal(t, false, ok, "got value is valid")
}

func TestNullableBoolGetNativePresentation(t *testing.T) {

	T1 := NullableBool{Valid: true, Value: true}

	value, ok := T1.GetNativePresentation()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, true, value, "got value is not valid")
}

func TestNullableBoolGetNativePresentationFalse(t *testing.T) {

	T1 := NullableBool{Valid: true, Value: false}

	value, ok := T1.GetNativePresentation()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, false, value, "got value is not valid")
}

func TestNullableBoolGetNativePresentationNil(t *testing.T) {

	T1 := NullableBool{Valid: false, Value: false}

	_, ok := T1.GetNativePresentation()
	assert.Equal(t, false, ok, "got value is valid")
}

func TestNullableBoolSet(t *testing.T) {
	var test NullableBool

	err := test.Set("true")
	assert.Nil(t, err, "test did not set")
	assert.True(t, test.Valid, "marshaled value is not valid")
	assert.Equal(t, true, test.Value, "marshaled value is not valid")
}

func TestNullableBoolSetFalse(t *testing.T) {
	var test NullableBool

	err := test.Set("false")
	assert.Nil(t, err, "test did not set")
	assert.True(t, test.Valid, "marshaled value is not valid")
	assert.Equal(t, false, test.Value, "marshaled value is not valid")
}

func TestNullableBoolSetNil(t *testing.T) {
	var test NullableBool

	err := test.Set("")
	assert.Nil(t, err, "test did not set")
	assert.False(t, test.Valid, "marshaled value is valid")
}

func TestNullableBoolMarshal(t *testing.T) {

	type test struct {
		T1 NullableBool
	}

	jsonValue := test{T1: NullableBool{Valid: true, Value: true}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":true}", string(jsonString), "marshaled value is not valid")
}

func TestNullableBoolMarshalFalse(t *testing.T) {

	type test struct {
		T1 NullableBool
	}

	jsonValue := test{T1: NullableBool{Valid: true, Value: false}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":false}", string(jsonString), "marshaled value is not valid")
}

func TestNullableBoolMarshalNil(t *testing.T) {

	type test struct {
		T1 NullableBool
	}

	jsonValue := test{T1: NullableBool{Valid: false, Value: false}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":null}", string(jsonString), "marshaled value is not valid")
}

func TestNullableBoolUmarshal(t *testing.T) {

	type test struct {
		T1 NullableBool
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":true}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.True(t, jsonValue.T1.Valid, "marshaled value is not valid")
	assert.Equal(t, true, jsonValue.T1.Value, "marshaled value is not valid")
}

func TestNullableBoolUmarshalFalse(t *testing.T) {

	type test struct {
		T1 NullableBool
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":false}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.True(t, jsonValue.T1.Valid, "marshaled value is not valid")
	assert.Equal(t, false, jsonValue.T1.Value, "marshaled value is not valid")
}

func TestNullableBoolUmarshalNil(t *testing.T) {

	type test struct {
		T1 NullableBool
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":null}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.False(t, jsonValue.T1.Valid, "marshaled value is valid")
}

// ---

func TestNullableIntGet(t *testing.T) {

	T1 := NullableInt{Valid: true, Value: 1}

	value, ok := T1.Get()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "1", value, "got value is not valid")
}

func TestNullableIntGet0(t *testing.T) {

	T1 := NullableInt{Valid: true, Value: 0}

	value, ok := T1.Get()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "0", value, "got value is not valid")
}

func TestNullableIntGetNil(t *testing.T) {

	T1 := NullableInt{Valid: false, Value: 0}

	_, ok := T1.Get()
	assert.Equal(t, false, ok, "got value is valid")
}

func TestNullableIntGetNativePresentation(t *testing.T) {

	T1 := NullableInt{Valid: true, Value: 1}

	value, ok := T1.GetNativePresentation()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, int64(1), value, "got value is not valid")
}

func TestNullableIntGetNativePresentation0(t *testing.T) {

	T1 := NullableInt{Valid: true, Value: 0}

	value, ok := T1.GetNativePresentation()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, int64(0), value, "got value is not valid")
}

func TestNullableIntGetNativePresentationNil(t *testing.T) {

	T1 := NullableInt{Valid: false, Value: 0}

	_, ok := T1.GetNativePresentation()
	assert.Equal(t, false, ok, "got value is valid")
}

func TestNullableIntSet(t *testing.T) {
	var test NullableInt

	err := test.Set("1")
	assert.Nil(t, err, "test did not set")
	assert.True(t, test.Valid, "marshaled value is not valid")
	assert.Equal(t, int64(1), test.Value, "marshaled value is not valid")
}

func TestNullableIntSet0(t *testing.T) {
	var test NullableInt

	err := test.Set("0")
	assert.Nil(t, err, "test did not set")
	assert.True(t, test.Valid, "marshaled value is not valid")
	assert.Equal(t, int64(0), test.Value, "marshaled value is not valid")
}

func TestNullableIntSetNil(t *testing.T) {
	var test NullableInt

	err := test.Set("")
	assert.Nil(t, err, "test did not set")
	assert.False(t, test.Valid, "marshaled value is valid")
}

func TestNullableIntMarshal(t *testing.T) {

	type test struct {
		T1 NullableInt
	}

	jsonValue := test{T1: NullableInt{Valid: true, Value: 1}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":1}", string(jsonString), "marshaled value is not valid")
}

func TestNullableIntMarshal0(t *testing.T) {

	type test struct {
		T1 NullableInt
	}

	jsonValue := test{T1: NullableInt{Valid: true, Value: 0}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":0}", string(jsonString), "marshaled value is not valid")
}

func TestNullableIntMarshalNil(t *testing.T) {

	type test struct {
		T1 NullableInt
	}

	jsonValue := test{T1: NullableInt{Valid: false, Value: 0}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":null}", string(jsonString), "marshaled value is not valid")
}

func TestNullableIntUmarshal(t *testing.T) {

	type test struct {
		T1 NullableInt
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":1}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.True(t, jsonValue.T1.Valid, "marshaled value is not valid")
	assert.Equal(t, int64(1), jsonValue.T1.Value, "marshaled value is not valid")
}

func TestNullableIntUmarshal0(t *testing.T) {

	type test struct {
		T1 NullableInt
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":0}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.True(t, jsonValue.T1.Valid, "marshaled value is not valid")
	assert.Equal(t, int64(0), jsonValue.T1.Value, "marshaled value is not valid")
}

func TestNullableIntUmarshalNil(t *testing.T) {

	type test struct {
		T1 NullableInt
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":null}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.False(t, jsonValue.T1.Valid, "marshaled value is valid")
}

// ---

func TestNullableFloatGet(t *testing.T) {

	T1 := NullableFloat{Valid: true, Value: 1}

	value, ok := T1.Get()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "1", value, "got value is not valid")
}

func TestNullableFloatGet0(t *testing.T) {

	T1 := NullableFloat{Valid: true, Value: 0}

	value, ok := T1.Get()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, "0", value, "got value is not valid")
}

func TestNullableFloatGetNil(t *testing.T) {

	T1 := NullableFloat{Valid: false, Value: 0}

	_, ok := T1.Get()
	assert.Equal(t, false, ok, "got value is valid")
}

func TestNullableFloatGetNativePresentation(t *testing.T) {

	T1 := NullableFloat{Valid: true, Value: 1}

	value, ok := T1.GetNativePresentation()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, 1., value, "got value is not valid")
}

func TestNullableFloatGetNativePresentation0(t *testing.T) {

	T1 := NullableFloat{Valid: true, Value: 0}

	value, ok := T1.GetNativePresentation()
	assert.Equal(t, true, ok, "got value is not valid")
	assert.Equal(t, 0., value, "got value is not valid")
}

func TestNullableFloatGetNativePresentationNil(t *testing.T) {

	T1 := NullableFloat{Valid: false, Value: 0}

	_, ok := T1.GetNativePresentation()
	assert.Equal(t, false, ok, "got value is valid")
}

func TestNullableFloatSet(t *testing.T) {
	var test NullableFloat

	err := test.Set("1")
	assert.Nil(t, err, "test did not set")
	assert.True(t, test.Valid, "marshaled value is not valid")
	assert.Equal(t, 1., test.Value, "marshaled value is not valid")
}

func TestNullableFloatSet0(t *testing.T) {
	var test NullableFloat

	err := test.Set("0")
	assert.Nil(t, err, "test did not set")
	assert.True(t, test.Valid, "marshaled value is not valid")
	assert.Equal(t, 0., test.Value, "marshaled value is not valid")
}

func TestNullableFloatSetNil(t *testing.T) {
	var test NullableFloat

	err := test.Set("")
	assert.Nil(t, err, "test did not set")
	assert.False(t, test.Valid, "marshaled value is valid")
}

func TestNullableFloatMarshal(t *testing.T) {

	type test struct {
		T1 NullableFloat
	}

	jsonValue := test{T1: NullableFloat{Valid: true, Value: 1}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":1}", string(jsonString), "marshaled value is not valid")
}

func TestNullableFloatMarshal0(t *testing.T) {

	type test struct {
		T1 NullableFloat
	}

	jsonValue := test{T1: NullableFloat{Valid: true, Value: 0}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":0}", string(jsonString), "marshaled value is not valid")
}

func TestNullableFloatMarshalNil(t *testing.T) {

	type test struct {
		T1 NullableFloat
	}

	jsonValue := test{T1: NullableFloat{Valid: false, Value: 0}}

	jsonString, err := json.Marshal(jsonValue)
	assert.Nil(t, err, "test did not marshal")
	assert.Equal(t, "{\"T1\":null}", string(jsonString), "marshaled value is not valid")
}

func TestNullableFloatUmarshal(t *testing.T) {

	type test struct {
		T1 NullableFloat
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":1}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.True(t, jsonValue.T1.Valid, "marshaled value is not valid")
	assert.Equal(t, 1., jsonValue.T1.Value, "marshaled value is not valid")
}

func TestNullableFloatUmarshal0(t *testing.T) {

	type test struct {
		T1 NullableFloat
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":0}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.True(t, jsonValue.T1.Valid, "marshaled value is not valid")
	assert.Equal(t, 0., jsonValue.T1.Value, "marshaled value is not valid")
}

func TestNullableFloatUmarshalNil(t *testing.T) {

	type test struct {
		T1 NullableFloat
	}

	var jsonValue test

	err := json.Unmarshal([]byte("{\"T1\":null}"), &jsonValue)
	assert.Nil(t, err, "test did not unmarshal")
	assert.False(t, jsonValue.T1.Valid, "marshaled value is valid")
}
