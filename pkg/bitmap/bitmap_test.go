package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroBitmap(t *testing.T) {

	b := ZeroBitmap(100)

	assert.Equal(t, b.n, uint64(100))
	assert.Equal(t, b.bitset[0], uint64(0))
	assert.Equal(t, b.bitset[1], uint64(0))
}

func TestOneBitmap(t *testing.T) {
	b := OneBitmap(100)

	assert.Equal(t, b.n, uint64(100))
	assert.Equal(t, b.bitset[0], uint64(0xFFFFFFFFFFFFFFFF))
	assert.Equal(t, b.bitset[1], uint64(0xFFFFFFFFFFFFFFFF))
}

func TestSetBit(t *testing.T) {
	
	b1 := ZeroBitmap(8) 
	assert.Equal(t, b1.bitset[0], uint64(0))

	SetBit(&b1, 7)
	assert.Equal(t, b1.bitset[0], uint64(128))

    UnsetBit(&b1, 7)
	assert.Equal(t, b1.bitset[0], uint64(0))

	SetBit(&b1, 0)
    assert.Equal(t, b1.bitset[0], uint64(1))

	UnsetBit(&b1, 0)
	assert.Equal(t, b1.bitset[0], uint64(0))
}

func TestAND(t *testing.T) {

	b1 := ZeroBitmap(8)
    SetBit(&b1, 0)
	b2 := OneBitmap(8)

	b := AND(b1, b2) 
	assert.Equal(t, b.bitset[0], uint64(1))
}

func TestIsZero(t *testing.T) {
	b1 := ZeroBitmap(25)
	assert.True(t, IsZero(b1))

	b2 := OneBitmap(25)
	assert.False(t, IsZero(b2))
}

func TestToString(t *testing.T) {
	b1 := ZeroBitmap(16)
	SetBit(&b1,15)
	SetBit(&b1,7)
	SetBit(&b1,1)
	assert.Equal(t, b1.n, uint64(16))
	assert.Equal(t, ToString(b1), "0100000100000001")
}