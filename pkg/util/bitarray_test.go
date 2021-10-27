package util

import (
	"testing"
)

func TestNewBitArray(t *testing.T) {
	ba := NewBitArray(1000, 1)
	_ = NewBitArray(1000, 8)
	if ba.capacity != 1000 || ba.bitsPerWord != 1 {
		t.Errorf("Error creating new BitArray")
	}
}

func TestGetB(t *testing.T) {
	ba := NewBitArray(100, 1)
	ba.bytePool[0] = byte(25)
	ba.bytePool[3] = byte(12)
	setBits := []uint32{3, 4, 7, 28, 29}
	for _, pos := range setBits {
		if ba.GetB(pos) != 1 {
			t.Errorf("Bit retrieval error.\nexp:%d\nwas:%d", 1, ba.GetB(pos))
		}
	}
	unsetBits := []uint32{0, 1, 2, 5, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
		18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 30, 31, 32, 33, 34, 36}
	for _, pos := range unsetBits {
		if ba.GetB(pos) != 0 {
			t.Errorf("BitArray.SetB error.\nexp:%d\nwas:%d", 0, ba.GetB(pos))
		}
	}
}

func TestSetB(t *testing.T) {
	ba := NewBitArray(100, 1)
	ba.SetB(3, byte(1))
	ba.SetB(4, byte(1))
	ba.SetB(7, byte(1))
	ba.SetB(28, byte(1))
	ba.SetB(29, byte(1))
	ba.SetB(35, byte(1))
	setBits := []uint32{3, 4, 7, 28, 29, 35}
	for _, pos := range setBits {
		if ba.GetB(pos) != 1 {
			t.Errorf("BitArray.SetB error.\nexp:%d\nwas:%d", 1, ba.GetB(pos))
		}
	}
	unsetBits := []uint32{0, 1, 2, 5, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
		18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 30, 31, 32, 33, 34, 36}
	for _, pos := range unsetBits {
		if ba.GetB(pos) != 0 {
			t.Errorf("BitArray.SetB error.\nexp:%d\nwas:%d", 0, ba.GetB(pos))
		}
	}
}
