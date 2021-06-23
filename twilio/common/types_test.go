package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToString(t *testing.T) {
	assert.Equal(t, "123", IntToString(123))
	assert.Equal(t, "0", IntToString(0))
	assert.Equal(t, "-123", IntToString(-123))
}
