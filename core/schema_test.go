package core

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestRequiredSchema(t *testing.T) {
	testSchema := AsInt(SchemaRequired)
	assert.Equal(t, false, testSchema.Optional, "Invalid options")
	assert.Equal(t, true, testSchema.Required, "Invalid options")
	assert.Equal(t, false, testSchema.Computed, "Invalid options")
	assert.Equal(t, false, testSchema.ForceNew, "Invalid options")
}

func TestOptionalSchema(t *testing.T) {
	testSchema := AsInt(SchemaOptional)
	assert.Equal(t, true, testSchema.Optional, "Invalid options")
	assert.Equal(t, false, testSchema.Required, "Invalid options")
	assert.Equal(t, false, testSchema.Computed, "Invalid options")
	assert.Equal(t, false, testSchema.ForceNew, "Invalid options")
	assert.Equal(t, false, testSchema.Sensitive, "Invalid options")
}

func TestComputedSchema(t *testing.T) {
	testSchema := AsInt(SchemaComputed)
	assert.Equal(t, false, testSchema.Optional, "Invalid options")
	assert.Equal(t, false, testSchema.Required, "Invalid options")
	assert.Equal(t, true, testSchema.Computed, "Invalid options")
	assert.Equal(t, false, testSchema.ForceNew, "Invalid options")
	assert.Equal(t, false, testSchema.Sensitive, "Invalid options")
}

func TestComputedOptionalSchema(t *testing.T) {
	testSchema := AsInt(SchemaComputedOptional)
	assert.Equal(t, true, testSchema.Optional, "Invalid options")
	assert.Equal(t, false, testSchema.Required, "Invalid options")
	assert.Equal(t, true, testSchema.Computed, "Invalid options")
	assert.Equal(t, false, testSchema.ForceNew, "Invalid options")
	assert.Equal(t, false, testSchema.Sensitive, "Invalid options")
}

func TestComputedSensitiveSchema(t *testing.T) {
	testSchema := AsInt(SchemaComputedSensitive)
	assert.Equal(t, false, testSchema.Optional, "Invalid options")
	assert.Equal(t, false, testSchema.Required, "Invalid options")
	assert.Equal(t, true, testSchema.Computed, "Invalid options")
	assert.Equal(t, false, testSchema.ForceNew, "Invalid options")
	assert.Equal(t, true, testSchema.Sensitive, "Invalid options")
}

func TestForceNewRequiredSchemaSchema(t *testing.T) {
	testSchema := AsInt(SchemaForceNewRequired)
	assert.Equal(t, false, testSchema.Optional, "Invalid options")
	assert.Equal(t, true, testSchema.Required, "Invalid options")
	assert.Equal(t, false, testSchema.Computed, "Invalid options")
	assert.Equal(t, true, testSchema.ForceNew, "Invalid options")
	assert.Equal(t, false, testSchema.Sensitive, "Invalid options")
}

func TestForceNewRequiredOptionalSchema(t *testing.T) {
	testSchema := AsInt(SchemaForceNewOptional)
	assert.Equal(t, true, testSchema.Optional, "Invalid options")
	assert.Equal(t, false, testSchema.Required, "Invalid options")
	assert.Equal(t, false, testSchema.Computed, "Invalid options")
	assert.Equal(t, true, testSchema.ForceNew, "Invalid options")
	assert.Equal(t, false, testSchema.Sensitive, "Invalid options")
}

func TestInvalidEmptySchema(t *testing.T) {
	defer func() { _ = recover() }()
	AsInt(&options{})
	t.Errorf("Invalid schema allowed")
}

func TestInvalidComputedRequiredSchema(t *testing.T) {
	defer func() { _ = recover() }()
	AsInt(&options{Required: true, Computed: true})
	t.Errorf("Invalid schema allowed")
}

func TestInvalidComputedForcenewSchema(t *testing.T) {
	defer func() { _ = recover() }()
	AsInt(&options{ForceNew: true, Computed: true})
	t.Errorf("Invalid schema allowed")
}

func TestStringSchema(t *testing.T) {
	s := AsString(SchemaRequired)
	assert.Equal(t, schema.TypeString, s.Type, "Invalid type")
}

func TestStringBool(t *testing.T) {
	s := AsBool(SchemaRequired)
	assert.Equal(t, schema.TypeBool, s.Type, "Invalid type")
}

func TestIntSchema(t *testing.T) {
	s := AsInt(SchemaRequired)
	assert.Equal(t, schema.TypeInt, s.Type, "Invalid type")
}

func TestFloatSchema(t *testing.T) {
	s := AsFloat(SchemaRequired)
	assert.Equal(t, schema.TypeFloat, s.Type, "Invalid type")
}

func TestSidSchema(t *testing.T) {
	s := AsSid(&Sid{}, SchemaRequired)
	assert.Equal(t, schema.TypeString, s.Type, "Invalid type")

	err := s.ValidateDiagFunc("XX00112233445566778899aabbccddeeff", nil)
	assert.Empty(t, err, "Sid errored")

	err = s.ValidateDiagFunc("abc", nil)
	assert.NotEmpty(t, err, "Sid validate did not error")
}

func TestListSchemaScalar(t *testing.T) {
	s := AsList(AsInt(SchemaRequired), SchemaRequired)
	assert.Equal(t, schema.TypeList, s.Type, "Invalid type")
	assert.Equal(t, schema.TypeInt, s.Elem.(*schema.Schema).Type, "Invalid type")
}

func TestListSchemaComplex(t *testing.T) {
	mapSchema := map[string]*schema.Schema{"test1": AsInt(SchemaRequired)}
	s := AsList(mapSchema, SchemaRequired)
	assert.Equal(t, schema.TypeList, s.Type, "Invalid type")
	assert.Equal(t, mapSchema, s.Elem.(*schema.Resource).Schema, "Invalid type")
}
