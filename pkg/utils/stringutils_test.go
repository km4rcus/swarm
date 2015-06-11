package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringListSplit(t *testing.T) {
	list := "0-2,5,8"

	s := StringListSplit(list)
	assert.Equal(t, s[0], "0")
	assert.Equal(t, s[1], "1")
	assert.Equal(t, s[2], "2")
	assert.Equal(t, s[3], "5")
	assert.Equal(t, s[4], "8")
}