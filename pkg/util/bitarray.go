package util

import (
	"fmt"
	"log"
)

var supportedBitsPerWord = []byte{1, 2, 4, 8}

// This struct implements a BitArray with variable bit width.  Values in
// BitArray can use either 1, 2, 4 or 8 bits (respectively 8, 4, 2 or 1 values
// can be represented using a single byte).
type BitArray struct {
	bytePool     []byte // byte array
	bitsPerWord  byte   // supported values: 1,2,4,8
	wordsPerByte byte   // #values encoded per byte
	capacity     uint32 // #values to be stored in bytePool
}

// Creates a new BitArray that can hold 'capacity' many words of size bitsPerWord
func NewBitArray(capacity uint32, bitsPerWord byte) *BitArray {
	validBitsPerWord := false
	for _, bpw := range supportedBitsPerWord {
		if bpw == bitsPerWord {
			validBitsPerWord = true
		}
	}
	if !validBitsPerWord {
		panic(fmt.Sprintf("BitArray only supports %v bits per word",
			supportedBitsPerWord))
	}
	wordsPerByte := byte(8) / bitsPerWord
	return &BitArray{
		make([]byte, capacity/uint32(wordsPerByte)+1),
		bitsPerWord,
		wordsPerByte,
		capacity,
	}
}

func (b *BitArray) Len() int          { return len(b.bytePool) }
func (b *BitArray) Bytes() []byte     { return b.bytePool }
func (b *BitArray) BitsPerWord() byte { return b.bitsPerWord }

// Sets multiple bytes (to insert a search state into BitArray)
func (b *BitArray) SetBytes(pos uint32, bs []byte) {
	// TODO: implement
	log.Fatalf("")
}

// Returns all values between pos and pos+n (to fetch a complete search state)
func (b *BitArray) GetBytes(pos uint32, n uint32) []byte {
	// TODO: implement
	return []byte{}
}

// Returns the value stored at position pos
func (b *BitArray) GetB(pos uint32) byte {
	theByte := pos / uint32(b.wordsPerByte)
	posInByte := byte(pos % uint32(b.wordsPerByte))
	bpw := b.bitsPerWord
	oo := (byte(0xFF<<(8-bpw)) >> (posInByte * bpw))
	oorr := b.bytePool[theByte] & oo
	return oorr >> (8 - (posInByte+1)*bpw)
}

// Sets the value at pos to val
func (b *BitArray) SetB(pos uint32, val byte) {
	theByte := pos / uint32(b.wordsPerByte)
	posInByte := byte(pos % uint32(b.wordsPerByte))
	bpw := b.bitsPerWord
	oo := (byte(0xFF<<(8-bpw)) >> (posInByte * bpw)) ^ 0xFF
	zr := b.bytePool[theByte] & oo
	sr := byte(val<<(8-bpw)) >> (posInByte * bpw)
	b.bytePool[theByte] = zr | sr
}

//
// TODO: maybe it will be useful to represent values with more than one byte.
// Currently a byte that holds a single value can represent the range 1-255
// (size of byte). It would be nice to store greater values in multiple bytes,
// or maybe use uint16,unit32,... instead of byte to be able to store bigger
// values.
//
