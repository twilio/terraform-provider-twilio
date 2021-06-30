package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryEncoderSimple(t *testing.T) {

	type testStruct struct {
		a int
		b int
	}

	testValue := testStruct{a: 2, b: 3}

	queryString, err := QueryEncoder(testValue)
	assert.Nil(t, err, "QueryEncoder failed")
	assert.Contains(t, []string{"?a=2&b=3", "?b=3&a=2"}, queryString)
}

func TestQueryEncoderNil(t *testing.T) {
	_, err := FormEncoder(nil)
	assert.NotNil(t, err, "FormEncoder succeeded")
}

func TestQueryEncoderArray(t *testing.T) {
	testValue := []int{1, 2}
	_, err := FormEncoder(testValue)
	assert.NotNil(t, err, "FormEncoder succeeded")
}
