package common

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessFormValueBool(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		a bool
	}

	testValue := testStruct{a: true}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"true"}, form["a"])
}

func TestProcessFormValueInt(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		a int
	}

	testValue := testStruct{a: 1}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"1"}, form["a"])
}

func TestProcessFormValueFloat(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		a float64
	}

	testValue := testStruct{a: 1}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"1"}, form["a"])
}

func TestProcessFormValueString(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		a string
	}

	testValue := testStruct{a: "str"}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"str"}, form["a"])
}

func TestProcessFormValueIntList(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		a []int
	}

	testValue := testStruct{a: []int{1, 2, 3}}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"1", "2", "3"}, form["a"])
}

func TestProcessFormValueArray(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		a [1]int
	}

	testValue := testStruct{a: [1]int{1}}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.NotNil(t, err, "ProcessFormValue Succeeded (failure expected)")
}

func TestProcessFormValueBasic(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		A Sid
	}

	sid, _ := CreateSid("XX00112233445566778899aabbccddeeff")
	testValue := testStruct{A: sid}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"XX00112233445566778899aabbccddeeff"}, form["A"])
}

func TestProcessFormValuePtrNil(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		A *Sid
	}

	testValue := testStruct{A: nil}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.NotNil(t, err, "ProcessFormValue Succeeded (failure expected)")
}

func TestProcessFormValuePtrBasic(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		A *Sid
	}

	sid, _ := CreateSid("XX00112233445566778899aabbccddeeff")
	testValue := testStruct{A: &sid}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"XX00112233445566778899aabbccddeeff"}, form["A"])
}

func TestProcessFormValueJson(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		A interface{}
	}

	testValue := testStruct{A: map[string]interface{}{"a": 1, "b": "str"}}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"{\"a\":1,\"b\":\"str\"}"}, form["A"])
}

func TestProcessFormValueName(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		a int `form:"b"`
	}

	testValue := testStruct{a: 1}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"1"}, form["b"])
}

func TestProcessFormValueIgnore(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		a int `form:",ignore"`
	}

	testValue := testStruct{a: 1}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	_, ok := form["a"]
	assert.Equal(t, false, ok, "form is not empty")
}

func TestProcessFormValueOmitEmpty(t *testing.T) {
	form := url.Values{}

	type testStruct struct {
		A *Sid `form:",omitempty"`
	}

	testValue := testStruct{A: nil}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	_, ok := form["A"]
	assert.Equal(t, false, ok, "form is not empty")
}

func TestProcessFormValueFlatten(t *testing.T) {
	form := url.Values{}

	type testStruct2 struct {
		B int
	}

	type testStruct struct {
		A testStruct2 `form:",flatten"`
	}

	testValue := testStruct{A: testStruct2{B: 3}}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.Nil(t, err, "ProcessFormValue failed")
	assert.Equal(t, []string{"3"}, form["B"])
}

func TestProcessFormValueNoFlatten(t *testing.T) {
	form := url.Values{}

	type testStruct2 struct {
		B int
	}

	type testStruct struct {
		A testStruct2
	}

	testValue := testStruct{A: testStruct2{B: 3}}

	err := processParamValue("form", &form, reflect.TypeOf(testValue).Field(0), reflect.ValueOf(testValue).Field(0))
	assert.NotNil(t, err, "ProcessFormValue Succeeded")
}

func TestFormEncoderSimple(t *testing.T) {

	type testStruct struct {
		a int
		b int
	}

	testValue := testStruct{a: 2, b: 3}

	form, err := FormEncoder(testValue)
	assert.Nil(t, err, "FormEncoder failed")
	assert.Equal(t, []string{"2"}, form["a"])
	assert.Equal(t, []string{"3"}, form["b"])
}

func TestFormEncoderNil(t *testing.T) {
	_, err := FormEncoder(nil)
	assert.NotNil(t, err, "FormEncoder succeeded")
}

func TestFormEncoderArray(t *testing.T) {
	testValue := []int{1, 2}
	_, err := FormEncoder(testValue)
	assert.NotNil(t, err, "FormEncoder succeeded")
}
