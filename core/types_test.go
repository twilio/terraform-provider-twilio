package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToString(t *testing.T) {
	assert.Equal(t, "123", IntToString(123))
	assert.Equal(t, "0", IntToString(0))
	assert.Equal(t, "-123", IntToString(-123))
}

func TestStringToInt(t *testing.T) {
	value, err := StringToInt("123")
	assert.Nil(t, err)
	assert.Equal(t, 123, value)

	value, err = StringToInt("0")
	assert.Nil(t, err)
	assert.Equal(t, 0, value)

	value, err = StringToInt("-123")
	assert.Nil(t, err)
	assert.Equal(t, -123, value)

	_, err = StringToInt("blurg")
	assert.NotNil(t, err)
}
