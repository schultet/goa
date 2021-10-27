package state

import (
	"math/rand"
	"testing"
	"time"
)

var (
	benchN                   = 3500
	benchPacker, benchStates = setupPackerBench(benchN)
)

func TestIntPackerSet(t *testing.T) {
	//ranges := []int{2048, 1024, 1024, 2048, 2048, 2048, 2048, 1024, 4, 4, 4}
	ranges := []int{
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
	}
	//s := []int{44, 35, 77, 100, 355, 0, 7, 2, 1, 2, 3}
	s := []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1, 1, 1,
		0, 0, 0, 0, 0, 1, 0, 0,
		0, 1, 1, 1, 1, 1, 0, 0,
	}
	exp := []Bin{1042349824}

	packer := NewIntPacker(ranges)
	buf := make([]Bin, packer.numBins)
	for i, val := range s {
		packer.set(buf, i, val)
	}
	for i := range buf {
		if buf[i] != exp[i] {
			t.Errorf("ERR exp: %d was: %d\n", exp[i], buf[i])
		}
	}
}

func setupPackerBench(numStates int) (*IntPacker, []State) {
	rand.Seed(time.Now().UTC().UnixNano())
	ranges := []int{
		10, 10, 10, 10, 10, 10, 10, 10,
		20, 20, 20, 20, 20, 20, 20, 20,
		2, 2, 2, 2, 2, 2, 2, 2,
		5, 5, 5, 5, 5, 5, 5, 5,
	}
	packer := NewIntPacker(ranges)

	states := make([]State, numStates)
	for i := range states {
		states[i] = randomState(ranges)
	}

	return packer, states
}

func randomState(ranges []int) State {
	s := make(State, len(ranges))
	for i, r := range ranges {
		s[i] = rand.Intn(r)
	}
	return s
}

func BenchmarkIntPackerPack(b *testing.B) {
	n := 3500
	packer, states, n := benchPacker, benchStates, benchN
	pool := make([]PackedState, n)

	for i := 0; i < b.N; i++ {
		for i, s := range states {
			pool[i] = packer.Pack(s)
		}
	}
}

func BenchmarkIntPackerPackAppend(b *testing.B) {
	packer, states, n := benchPacker, benchStates, benchN
	numBins := packer.numBins
	pool := make([]Bin, n*numBins)

	for i := 0; i < b.N; i++ {
		for i, s := range states {
			packer.PackAppend(pool[i*numBins:], s)
		}
	}
}

// TODO: test with non-binary ranges/values
func TestIntPackerGet(t *testing.T) {
	ranges := []int{
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2,
	}
	exp := []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1, 1, 1,
		0, 0, 0, 0, 0, 1, 0, 0,
		0, 1, 1, 1, 1, 1, 0, 0,
	}
	buf := []Bin{1042349824}

	packer := NewIntPacker(ranges)
	for i, v := range exp {
		if packer.get(buf, i) != v {
			t.Errorf("ERR exp: %d was: %d\n", v, packer.get(buf, i))
		}
	}

	ranges2 := []int{
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
		2, 2, 2, 11,
	}
	buf2 := []Bin{18446040386265677232, 543296577406}
	index2 := 59
	packer2 := NewIntPacker(ranges2)
	if res := packer2.get(buf2, index2); res != 1 {
		t.Errorf("ERR exp: %d was: %d\n", 1, res)
	}
	for i := range ranges2 {
		if res := packer2.get(buf2, i); res < 0 {
			t.Errorf("ERR exp: %d was: %d\n", 1, res)
		}
	}
}

//func TestBinaryPackerPack(t *testing.T) {
//	tests := [][]int{
//		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
//			1, 1, 1, 1, 1, 1, 1, 1, 1},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
//			0, 0, 0, 0, 0, 0, 0, 0, 0},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
//			0, 0, 0, 0, 0, 0, 0, 0, 1},
//		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
//			0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
//		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
//			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
//		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
//			0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
//		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
//			0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
//		{1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0,
//			0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1},
//		{0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1,
//			1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0,
//			1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
//	}
//
//	for _, tc := range tests {
//		packer := BinaryPacker{len(tc)}
//		x := packer.Pack(tc)
//		y := packer.Unpack(x)
//		if !slice.EqualIntArrays(y, tc) {
//			t.Errorf("\nexp: %v\nwas: %v\n", tc, y)
//		}
//	}
//}
//
//func TestBinaryPackerValueOf(t *testing.T) {
//	tests := [][]int{
//		{0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
//			0, 0, 0, 0, 0, 0, 0, 0, 1},
//	}
//	for _, tc := range tests {
//		packer := BinaryPacker{len(tc)}
//		x := packer.Pack(tc)
//		y := make([]int, 0)
//		for i := range tc {
//			y = append(y, packer.ValueOf(x, i))
//		}
//		if !slice.EqualIntArrays(y, tc) {
//			t.Errorf("\nexp: %v\nwas: %v\n", tc, y)
//		}
//	}
//}
