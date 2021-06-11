package common

import (
	"reflect"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestTagNaming(t *testing.T) {
	type test1 struct {
		T1 string `json:"t1"`
		T2 string
		T3 string `json:"t3a" provider:"t3b"`
		T4 string `provider:"t4"`
		T5 string `json:"t5" provider:",id"`
		T6 string `provider:",id"`
		T7 string `provider:",ignore"`
		T8 string `provider:",flatten"`
	}
	var field reflect.StructField
	var tag tagInfo
	field, _ = reflect.TypeOf(test1{}).FieldByName("T1")
	tag = getNameFromTag(field)
	assert.Equal(t, tag.name, "t1", "wrong name")
	assert.Equal(t, tag.isId, false, "Id set")
	assert.Equal(t, tag.flatten, false, "Flatten set")
	assert.Equal(t, tag.ignore, false, "ignore set")
	field, _ = reflect.TypeOf(test1{}).FieldByName("T2")
	tag = getNameFromTag(field)
	assert.Equal(t, tag.name, "T2", "wrong name")
	assert.Equal(t, tag.isId, false, "Id set")
	assert.Equal(t, tag.flatten, false, "Flatten set")
	assert.Equal(t, tag.ignore, false, "ignore set")
	field, _ = reflect.TypeOf(test1{}).FieldByName("T3")
	tag = getNameFromTag(field)
	assert.Equal(t, tag.name, "t3b", "wrong name")
	assert.Equal(t, tag.isId, false, "Id set")
	assert.Equal(t, tag.flatten, false, "Flatten set")
	assert.Equal(t, tag.ignore, false, "ignore set")
	field, _ = reflect.TypeOf(test1{}).FieldByName("T4")
	tag = getNameFromTag(field)
	assert.Equal(t, tag.name, "t4", "wrong name")
	assert.Equal(t, tag.isId, false, "Id set")
	assert.Equal(t, tag.flatten, false, "Flatten set")
	assert.Equal(t, tag.ignore, false, "ignore set")
	field, _ = reflect.TypeOf(test1{}).FieldByName("T5")
	tag = getNameFromTag(field)
	assert.Equal(t, tag.name, "t5", "wrong name")
	assert.Equal(t, tag.isId, true, "Id not set")
	assert.Equal(t, tag.flatten, false, "Flatten set")
	assert.Equal(t, tag.ignore, false, "ignore set")
	field, _ = reflect.TypeOf(test1{}).FieldByName("T6")
	tag = getNameFromTag(field)
	assert.Equal(t, tag.name, "T6", "wrong name")
	assert.Equal(t, tag.isId, true, "Id not set")
	assert.Equal(t, tag.flatten, false, "Flatten set")
	assert.Equal(t, tag.ignore, false, "ignore set")
	field, _ = reflect.TypeOf(test1{}).FieldByName("T7")
	tag = getNameFromTag(field)
	assert.Equal(t, tag.name, "T7", "wrong name")
	assert.Equal(t, tag.isId, false, "Id set")
	assert.Equal(t, tag.flatten, false, "Flatten set")
	assert.Equal(t, tag.ignore, true, "ignore not set")
	field, _ = reflect.TypeOf(test1{}).FieldByName("T8")
	tag = getNameFromTag(field)
	assert.Equal(t, tag.name, "T8", "wrong name")
	assert.Equal(t, tag.isId, false, "Id set")
	assert.Equal(t, tag.flatten, true, "Flatten not set")
	assert.Equal(t, tag.ignore, false, "ignore set")
}

func TestSimpleUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"T2": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"T3": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T4": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
	v := map[string]interface{}{
		"T1": true,
		"T2": 1,
		"T3": "t3",
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T1 bool
		T2 int
		T3 string
		T4 *int
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, testStruct.T1, true, "T1 did not unmarshal")
	assert.Equal(t, testStruct.T2, 1, "T2 did not unmarshal")
	assert.Equal(t, testStruct.T3, "t3", "T3 did not unmarshal")
	assert.Nil(t, testStruct.T4, "T4 should be nil")
}

func TestComplexUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T2": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},
		"T3": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T4": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"T5": {
			Type:     schema.TypeFloat,
			Required: true,
		},
		"T6": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"T7": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
	v := map[string]interface{}{
		"T1": "AC00112233445566778899aabbccddeeff",
		"T2": []interface{}{"t2a", "t2b"},
		"T3": "t3",
		"T4": 1,
		"T5": 1.0,
		"T6": "AC00112233445566778899aabbccddeeff",
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)
	resourceData.SetId("t0")

	type test1 struct {
		T0 *string `provider:",id"`
		T1 AccountSid
		T2 []string
		T3 *string
		T4 *int
		T5 *float64
		T6 *AccountSid
		T7 *int
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, *testStruct.T0, "t0", "T0 did not unmarshal")
	assert.Equal(t, testStruct.T1.String(), "AC00112233445566778899aabbccddeeff", "T1 did not unmarshal")
	assert.Equal(t, len(testStruct.T2), 2, "T2 did not unmarshal")
	assert.Equal(t, testStruct.T2[0], "t2a", "T2 did not unmarshal")
	assert.Equal(t, testStruct.T2[1], "t2b", "T2 did not unmarshal")
	assert.Equal(t, *testStruct.T3, "t3", "T3 did not unmarshal")
	assert.Equal(t, *testStruct.T4, 1, "T4 did not unmarshal")
	assert.Equal(t, *testStruct.T5, 1.0, "T5 did not unmarshal")
	assert.Equal(t, testStruct.T6.String(), "AC00112233445566778899aabbccddeeff", "T6 did not unmarshal")
	assert.Nil(t, testStruct.T7, "T7 did not unmarshal")
}

func TestTimeUnMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T0": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	v := map[string]interface{}{
		"T0": "2021-05-17T01:35:33Z",
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)
	err := resourceData.Set("T0", "2021-05-17T01:35:33Z")

	type test1 struct {
		T0 *time.Time
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}

	assert.Nil(t, err, "Date T0 did not unmarshal")
	assert.Equal(t, resourceData.Get("T0"), "2021-05-17T01:35:33Z", "Date T0 did not unmarshal")
}

func TestUnmarshalNilValueToPointer(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T2": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
	v := map[string]interface{}{
		"T1": nil,
		"T2": nil,
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T1 *string
		T2 *[]string
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Nil(t, testStruct.T1, "T1 did not unmarshal")
	assert.Nil(t, testStruct.T2, "T2 did not unmarshal")
}

func TestFlattenUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T2": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	v := map[string]interface{}{
		"T1": "t1",
		"T2": "t2",
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)
	resourceData.SetId("t0")

	type test2 struct {
		T1 string
		T2 string
	}

	type test1 struct {
		T0     *string `provider:",id"`
		Nested test2   `provider:",flatten"`
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, *testStruct.T0, "t0", "T0 did not unmarshal")
	assert.Equal(t, testStruct.Nested.T1, "t1", "T1 did not unmarshal")
	assert.Equal(t, testStruct.Nested.T2, "t2", "T2 did not unmarshal")
}

func TestListUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"T1a": {
						Type:     schema.TypeString,
						Required: true,
					},
					"T1b": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Required: true,
		},
	}
	v := map[string]interface{}{
		"T1": []interface{}{
			map[string]interface{}{"T1a": "r1", "T1b": "r2"},
			map[string]interface{}{"T1a": "r3", "T1b": "r4"},
		},
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)
	resourceData.SetId("t0")

	type test2 struct {
		T1a string
		T1b string
	}

	type test1 struct {
		T1 []test2
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, len(testStruct.T1), 2, "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1[0].T1a, "r1", "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1[0].T1b, "r2", "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1[1].T1a, "r3", "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1[1].T1b, "r4", "T1 did not unmarshal")
}

func TestSidListUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Required: true,
		},
	}
	v := map[string]interface{}{
		"T1": []interface{}{"AC00112233445566778899aabbccddeefe", "AC00112233445566778899aabbccddeeff"},
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)
	resourceData.SetId("t0")

	type test1 struct {
		T1 []AccountSid
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, len(testStruct.T1), 2, "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1[0].String(), "AC00112233445566778899aabbccddeefe", "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1[1].String(), "AC00112233445566778899aabbccddeeff", "T1 did not unmarshal")
}

func TestMapUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeMap,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Required: true,
		},
	}
	v := map[string]interface{}{
		"T1": map[string]interface{}{"T1a": "r1", "T1b": "r2"},
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)
	resourceData.SetId("t0")

	type test1 struct {
		T1 map[string]string
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, len(testStruct.T1), 2, "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1["T1a"], "r1", "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1["T1b"], "r2", "T1 did not unmarshal")
}

func TestSidMapUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeMap,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Required: true,
		},
	}
	v := map[string]interface{}{
		"T1": map[string]interface{}{"T1a": "AC00112233445566778899aabbccddeefe", "T1b": "AC00112233445566778899aabbccddeeff"},
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)
	resourceData.SetId("t0")

	type test1 struct {
		T1 map[string]AccountSid
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, len(testStruct.T1), 2, "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1["T1a"].String(), "AC00112233445566778899aabbccddeefe", "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1["T1b"].String(), "AC00112233445566778899aabbccddeeff", "T1 did not unmarshal")
}

func TestComplexMapUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	v := map[string]interface{}{
		"T1": "{\"M0\":{\"T1a\":\"r1\",\"T1b\":true},\"M1\":{\"T1a\":\"r2\",\"T1b\":false}}",
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)
	resourceData.SetId("t0")

	type test2 struct {
		T1a string
		T1b bool
	}

	type test1 struct {
		T1 map[string]test2
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, len(testStruct.T1), 2, "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1["M0"].T1a, "r1", "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1["M0"].T1b, true, "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1["M1"].T1a, "r2", "T1 did not unmarshal")
	assert.Equal(t, testStruct.T1["M1"].T1b, false, "T1 did not unmarshal")
}

func TestPureJsonUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	v := map[string]interface{}{
		"T1": "{\"test\":\"test_value\"}",
	}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T1 interface{}
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, testStruct.T1.(map[string]interface{})["test"], "test_value", "T1 did not unmarshal")
}

func TestJsonEncodedNilUnMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test struct {
		T1 *interface{}
	}

	testStruct := test{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Nil(t, testStruct.T1, "T2 did not unmarshal")
}

func TestOptionalFlattenUnmarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T2": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)
	resourceData.SetId("t0")

	type test2 struct {
		T1 *string
		T2 *string
	}

	type test1 struct {
		T0     *string `provider:",id"`
		Nested test2   `provider:",flatten"`
	}

	testStruct := test1{}
	if err := UnmarshalSchema(&testStruct, resourceData); err != nil {
		t.Errorf("Unmarshall failed: result '%v'", err)
	}
	assert.Equal(t, *testStruct.T0, "t0", "T0 did not unmarshal")
	assert.Nil(t, testStruct.Nested.T1, "T1 did not unmarshal")
	assert.Nil(t, testStruct.Nested.T2, "T2 did not unmarshal")
}

func TestSimpleMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"T2": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"T3": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T4": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"T5": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"T6": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T0 string `provider:",id"`
		T1 bool
		T2 int
		T3 string
		T4 *int
		T5 *interface{}
		T6 *interface{}
	}

	testStruct := test1{T0: "t0", T1: true, T2: 1, T3: "t3", T5: nil}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.Equal(t, resourceData.Id(), "t0", "Id did not marshal")
	assert.Equal(t, resourceData.Get("T1"), true, "T1 did not marshal")
	assert.Equal(t, resourceData.Get("T2"), 1, "T2 did not marshal")
	assert.Equal(t, resourceData.Get("T3"), "t3", "T3 did not marshal")
	assert.Equal(t, 0, resourceData.Get("T4"), "T4 did not marshal")
	assert.Equal(t, "", resourceData.Get("T5"), "T4 did not marshal")
	assert.Equal(t, "", resourceData.Get("T6"), "T4 did not marshal")
}

func TestComplexMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T2": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},
		"T3": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T4": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T0 *string `provider:",id"`
		T1 AccountSid
		T2 []string
		T3 *string
		T4 *int
	}

	testString := "t0"
	testSid := AccountSid{}
	_ = testSid.Set("AC00112233445566778899aabbccddeeff")
	testString2 := "2010-04-01"
	testInt := 1

	testStruct := test1{T0: &testString, T1: testSid, T2: []string{"t2a", "t2b"}, T3: &testString2, T4: &testInt}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.Equal(t, resourceData.Id(), "t0", "Id did not marshal")
	assert.Equal(t, resourceData.Get("T1"), "AC00112233445566778899aabbccddeeff", "T1 did not marshal")
	assert.Equal(t, len(resourceData.Get("T2").([]interface{})), 2, "T2 did not unmarshal")
	assert.Equal(t, resourceData.Get("T2.0"), "t2a", "T2 did not unmarshal")
	assert.Equal(t, resourceData.Get("T2.1"), "t2b", "T2 did not unmarshal")
	assert.Equal(t, resourceData.Get("T3"), "2010-04-01", "T3 did not marshal")
	assert.Equal(t, resourceData.Get("T4"), 1, "T4 did not marshal")
}

func TestObjectMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"Errors": {
			Type: schema.TypeString,
		},
		"Links": {
			Type: schema.TypeString,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		Links  map[string]interface{}
		Errors []map[string]interface{}
	}

	testStruct := test1{
		Links: map[string]interface{}{
			"test_users": "https://studio.twilio.com/v2/Flows/FWXX/TestUsers",
			"revisions":  "https://studio.twilio.com/v2/Flows/FWXX/Revisions",
			"executions": "https://studio.twilio.com/v2/Flows/FWXX/Executions",
		},
		Errors: []map[string]interface{}{
			{
				"message":       "some message",
				"property_path": "some property path",
			},
			{
				"message":       "some message 2",
				"property_path": "some property path 2",
			},
		},
	}

	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}

	//  "{\"message\":\"is missing but it is required\",\"property_path\":\"#/description\"}"

	testMe := resourceData.Get("Links")
	testMe2 := resourceData.Get("Errors")
	assert.Equal(t, testMe, "{\"executions\":\"https://studio.twilio.com/v2/Flows/FWXX/Executions\",\"revisions\":\"https://studio.twilio.com/v2/Flows/FWXX/Revisions\",\"test_users\":\"https://studio.twilio.com/v2/Flows/FWXX/TestUsers\"}")
	assert.Equal(t, testMe2, "[{\\message\\:\\some message\\,\\property_path\\:\\some property path\\},{\\message\\:\\some message 2\\,\\property_path\\:\\some property path 2\\}]")
}

func TestTimeMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T0": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T0 *time.Time
	}

	testDate, _ := time.Parse(time.RFC3339, "2021-05-17T01:35:33Z")

	testStruct := test1{T0: &testDate}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.Equal(t, resourceData.Get("T0"), "2021-05-17T01:35:33Z", "Date T0 did not marshal")
}

func TestFlattenMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
		"T2": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test2 struct {
		T1 string
		T2 string
	}

	type test1 struct {
		T0     *string `provider:",id"`
		Nested test2   `provider:",flatten"`
	}

	testString := "t0"

	testStruct := test1{T0: &testString, Nested: test2{T1: "t1", T2: "t2"}}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.Equal(t, resourceData.Id(), "t0", "Id did not marshal")
	assert.Equal(t, resourceData.Get("T1"), "t1", "T1 did not marshal")
	assert.Equal(t, resourceData.Get("T2"), "t2", "T2 did not marshal")
}

func TestListMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"T1a": {
						Type:     schema.TypeString,
						Required: true,
					},
					"T1b": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test2 struct {
		T1a string
		T1b string
	}

	type test1 struct {
		T1 []test2
	}

	testStruct := test1{T1: []test2{{T1a: "r1", T1b: "r2"}, {T1a: "r3", T1b: "r4"}}}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.Equal(t, len(resourceData.Get("T1").([]interface{})), 2, "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.0.T1a"), "r1", "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.0.T1b"), "r2", "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.1.T1a"), "r3", "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.1.T1b"), "r4", "T1 did not unmarshal")
}

func TestSidListMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T1 []AccountSid
	}

	sid1, _ := CreateAccountSid("AC00112233445566778899aabbccddeefe")
	sid2, _ := CreateAccountSid("AC00112233445566778899aabbccddeeff")

	testStruct := test1{T1: []AccountSid{sid1, sid2}}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.Equal(t, len(resourceData.Get("T1").([]interface{})), 2, "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.0"), "AC00112233445566778899aabbccddeefe", "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.1"), "AC00112233445566778899aabbccddeeff", "T1 did not unmarshal")
}

func TestMapMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeMap,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T1 map[string]string
	}

	testStruct := test1{T1: map[string]string{"T1a": "r1", "T1b": "r2"}}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.Equal(t, len(resourceData.Get("T1").(map[string]interface{})), 2, "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.T1a"), "r1", "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.T1b"), "r2", "T1 did not unmarshal")
}

func TestSidMapMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeMap,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T1 map[string]AccountSid
	}

	sid1, _ := CreateAccountSid("AC00112233445566778899aabbccddeefe")
	sid2, _ := CreateAccountSid("AC00112233445566778899aabbccddeeff")

	testStruct := test1{T1: map[string]AccountSid{"T1a": sid1, "T1b": sid2}}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.Equal(t, len(resourceData.Get("T1").(map[string]interface{})), 2, "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.T1a"), "AC00112233445566778899aabbccddeefe", "T1 did not unmarshal")
	assert.Equal(t, resourceData.Get("T1.T1b"), "AC00112233445566778899aabbccddeeff", "T1 did not unmarshal")
}

func TestComplexMapMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test2 struct {
		T1a string
		T1b bool
	}

	type test1 struct {
		T1 map[string]test2
	}

	testStruct := test1{T1: map[string]test2{"M0": {T1a: "r1", T1b: true}, "M1": {T1a: "r2", T1b: false}}}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.Equal(t, resourceData.Get("T1"), "{\"M0\":{\"T1a\":\"r1\",\"T1b\":true},\"M1\":{\"T1a\":\"r2\",\"T1b\":false}}", "T1 did not unmarshal")
}

func TestPureJsonMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	type test1 struct {
		T1 interface{}
	}

	testStruct := test1{T1: map[string]interface{}{"test": "test_value"}}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}

	assert.Equal(t, resourceData.Get("T1"), "{\"test\":\"test_value\"}", "T1 did not unmarshal")
}

func TestJsonEncodedListOfObjectsMarshal(t *testing.T) {
	s := map[string]*schema.Schema{
		"T1": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	type test2 struct {
		T1 []interface{}
	}

	v := map[string]interface{}{}
	resourceData := schema.TestResourceDataRaw(t, s, v)

	m1 := map[string]interface{}{
		"foo": "bar1",
	}
	m2 := map[string]interface{}{
		"foo": "bar2",
	}
	testStruct := test2{
		T1: []interface{}{m1, m2},
	}
	if err := MarshalSchema(resourceData, &testStruct); err != nil {
		t.Errorf("Marshall failed: result '%v'", err)
	}
	assert.EqualValues(t, "{\"foo\":\"bar1\"}", resourceData.Get("T1.0"), "T1 did not unmarshal")
	assert.EqualValues(t, "{\"foo\":\"bar2\"}", resourceData.Get("T1.1"), "T1 did not unmarshal")
}
