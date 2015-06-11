package bitmap

const size = 64

type Bitmap struct {
	bitset []uint64
	n      uint64
}

// set the lenght of the bitmap to len and its value to 0...00
func ZeroBitmap(len uint64) Bitmap {
	var b Bitmap
	n := (len + size - 1) / size
	b.bitset = make([]uint64, n, n)
	b.n = len
	for i := range b.bitset {
		b.bitset[i] = uint64(0)
	}
	return b
}

// set the lenght of the bitmap to len and its value to 1...11
func OneBitmap(len uint64) Bitmap {
	var b Bitmap
	n := (len + size - 1 ) / size
	b.bitset = make([]uint64, n, n)
	b.n = len
	for i := range b.bitset {
		b.bitset[i] = 0xFFFFFFFFFFFFFFFF
	}
	return b
}

func SetBit(b *Bitmap, i uint64) {
	b.bitset[i/size] |= (1 << (i % size))
}

func UnsetBit(b *Bitmap, i uint64) {
	b.bitset[i/size] &= ^(1 << (i % size))
}

func ToggleBit(b *Bitmap, i uint64) {
	b.bitset[i/size] ^= (1 << (i % size))
}

func GetBit(b Bitmap, i uint64) uint64 {
	return b.bitset[i/size] & (1 << (i % size))
}

func ToString(b Bitmap) string {
    var i uint64
    s := ""
    for i=0; i < b.n; i++ {
        if GetBit(b,i) != uint64(0) {
            s += "1"
        } else {
            s += "0"
        }
    }
    return s
}

func IsZero(b Bitmap) bool {
	for _,v := range b.bitset {
	   if v != uint64(0) { 
		return false
	   }
	}
	return true
}

// bL and bR must be of the same length
func XOR(bL Bitmap, bR Bitmap) Bitmap {
	var res Bitmap
	res.n = bL.n
	res.bitset = make([]uint64, bL.n, bL.n)
	for i := range bL.bitset {
		res.bitset[i] = bL.bitset[i] ^ bR.bitset[i]
	}
	return res
}

// bL and bR must be of the same length
func AND(bL Bitmap, bR Bitmap) Bitmap {
	var res Bitmap
	res.n = bL.n
	res.bitset = make([]uint64, bL.n, bL.n)
	for i := range bL.bitset {
		res.bitset[i] = bL.bitset[i] & bR.bitset[i]
	}
	return res
}

// bL and bR must be of the same length
func OR(bL Bitmap, bR Bitmap) Bitmap {
	var res Bitmap
	res.n = bL.n
	res.bitset = make([]uint64, bL.n, bL.n)
	for i := range bL.bitset {
		res.bitset[i] = bL.bitset[i] | bR.bitset[i]
	}
	return res
}