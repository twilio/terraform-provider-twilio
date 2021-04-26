package common

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestRequiredSchema(t *testing.T) {
	schema := AsInt(SchemaRequired)
	assert.Equal(t, false, schema.Optional, "Invalid options")
	assert.Equal(t, true, schema.Required, "Invalid options")
	assert.Equal(t, false, schema.Computed, "Invalid options")
	assert.Equal(t, false, schema.ForceNew, "Invalid options")
}

func TestOptionalSchema(t *testing.T) {
	schema := AsInt(SchemaOptional)
	assert.Equal(t, true, schema.Optional, "Invalid options")
	assert.Equal(t, false, schema.Required, "Invalid options")
	assert.Equal(t, false, schema.Computed, "Invalid options")
	assert.Equal(t, false, schema.ForceNew, "Invalid options")
}

func TestComputedSchema(t *testing.T) {
	schema := AsInt(SchemaComputed)
	assert.Equal(t, false, schema.Optional, "Invalid options")
	assert.Equal(t, false, schema.Required, "Invalid options")
	assert.Equal(t, true, schema.Computed, "Invalid options")
	assert.Equal(t, false, schema.ForceNew, "Invalid options")
}

func TestForceNewRequiredSchemaSchema(t *testing.T) {
	schema := AsInt(SchemaForceNewRequired)
	assert.Equal(t, false, schema.Optional, "Invalid options")
	assert.Equal(t, true, schema.Required, "Invalid options")
	assert.Equal(t, false, schema.Computed, "Invalid options")
	assert.Equal(t, true, schema.ForceNew, "Invalid options")
}

func TestForceNewRequiredOptionalSchema(t *testing.T) {
	schema := AsInt(SchemaForceNewOptional)
	assert.Equal(t, true, schema.Optional, "Invalid options")
	assert.Equal(t, false, schema.Required, "Invalid options")
	assert.Equal(t, false, schema.Computed, "Invalid options")
	assert.Equal(t, true, schema.ForceNew, "Invalid options")
}

func TestInvalidEmptySchema(t *testing.T) {
	defer func() { recover() }()
	AsInt(&options{})
	t.Errorf("Invalid schema allowed")
}

func TestInvalidComputedRequiredSchema(t *testing.T) {
	defer func() { recover() }()
	AsInt(&options{Required: true, Computed: true})
	t.Errorf("Invalid schema allowed")
}

func TestInvalidComputedOptionalSchema(t *testing.T) {
	defer func() { recover() }()
	AsInt(&options{Optional: true, Computed: true})
	t.Errorf("Invalid schema allowed")
}

func TestInvalidComputedForcenewSchema(t *testing.T) {
	defer func() { recover() }()
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

	_, errs := s.ValidateFunc("XX00112233445566778899aabbccddeeff", "test")
	assert.Empty(t, errs, "Sid validate errored")

	_, errs = s.ValidateFunc("abc", "test")
	assert.NotEmpty(t, errs, "Sid validate did not error")
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
