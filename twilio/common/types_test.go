package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt32ToString(t *testing.T) {
	assert.Equal(t, "123", Int32ToString(int32(123)))
	assert.Equal(t, "0", Int32ToString(int32(0)))
	assert.Equal(t, "-123", Int32ToString(int32(-123)))
}
