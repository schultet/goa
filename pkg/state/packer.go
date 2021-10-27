package state

import (
	"fmt"
	"unsafe"
)

const bitsPerInt = int(unsafe.Sizeof(int(0))) * 8 // 32/64 depending on system
const bitsPerBin = int(unsafe.Sizeof(Bin(0))) * 8 // 32/64 depending on system

// Packer can compress (pack) states (slice of ints) and unpack them
type Packer interface {
	Pack(State) PackedState
	Unpack(PackedState) State
	ValueOf(s PackedState, v int) int
}

// compact representation of a state. Should only be used by the stateRegistry
// to store states.
type PackedState []Bin

// Two packed states are equal iff they have the same values
func (s PackedState) Equals(other PackedState) bool {
	if len(s) != len(other) {
		return false
	}
	for i := range s {
		if s[i] != other[i] {
			return false
		}
	}
	return true
}

// PhonyPacker does not perform any packing, instead just returns a packedState
// that equals the unpacked state and vice versa. Intended as a placeholder
// until IntPacker is implemented.
type PhonyPacker struct {
}

// NewPhonyPacker returns a new PhonyPacker
func NewPhonyPacker() *PhonyPacker {
	return &PhonyPacker{}
}

func (p *PhonyPacker) Pack(s State) PackedState {
	packedState := make([]Bin, len(s))
	for i, v := range s {
		packedState[i] = Bin(v)
	}
	return packedState
}
func (p *PhonyPacker) Unpack(ps PackedState) State {
	state := make(State, len(ps))
	for i, v := range ps {
		state[i] = int(v)
	}
	return state
}
func (p *PhonyPacker) ValueOf(s PackedState, i int) int { return int(s[i]) }

// Bin is the bin type/size values are packed into
// Note: must be unsigned and large enough to represent the biggest value
type Bin uint

// packable object/item
type PackageInfo struct {
	volume    int // volume
	binIndex  int
	shift     uint
	readMask  Bin
	clearMask Bin
}

func NewPackageInfo(volume, binIndex, shift int) PackageInfo {
	readMask := bitMask(shift, shift+requiredBitSize(volume))
	return PackageInfo{
		volume:    volume,
		binIndex:  binIndex,
		shift:     uint(shift),
		readMask:  readMask,
		clearMask: ^readMask,
	}
}

func DefaultPackageInfo() PackageInfo {
	return PackageInfo{binIndex: -1, shift: 0, readMask: 0, clearMask: 0}
}

func (v PackageInfo) String() string {
	return fmt.Sprintf("{%d, %d, %d}", v.volume, v.binIndex, v.shift)
	//return fmt.Sprintf("%d", v.volume)
}

func (v PackageInfo) get(buffer []Bin) int {
	return int((buffer[v.binIndex] & v.readMask) >> v.shift)
}

// value >= 0, value < v.domRange
func (v PackageInfo) set(buffer []Bin, value int) {
	buffer[v.binIndex] = (buffer[v.binIndex] & v.clearMask) |
		Bin(value<<v.shift)
}

func requiredBitSize(numValues int) int {
	var bits uint = 0
	for (1 << bits) < uint(numValues) {
		bits++
	}
	return int(bits)
}

// from >= 0, to >= from, to <= bitsPerBin
func bitMask(from, to int) Bin {
	length := to - from
	if length == bitsPerBin {
		return ^Bin(0)
	} else {
		return ((Bin(1) << uint(length)) - 1) << uint(from)
	}
}

// IntPacker packs a state with variadic domain ranges into as few ints as
// possible, using a greedy bin-packing strategy.
type IntPacker struct {
	packageInfos []PackageInfo
	numBins      int
}

// NewIntPacker returns a new IntPacker
func NewIntPacker(ranges []int) *IntPacker {
	packer := &IntPacker{numBins: 0}
	packer.packBins(ranges)
	return packer
}

func (p *IntPacker) ValueOf(s PackedState, v int) int {
	return p.get(s, v)
}

func (p *IntPacker) Pack(s State) PackedState {
	buf := make([]Bin, p.numBins)
	for i, v := range s {
		if v < 0 {
			fmt.Printf("%+v %d\n", s, i)
			panic("baa")
		}
		p.set(buf, i, v)
	}
	return PackedState(buf)
}

// PackAppend packs a state into a Bin slice
func (p *IntPacker) PackAppend(buf []Bin, s State) {
	for i, v := range s {
		if v < 0 {
			fmt.Printf("%+v %d\n", s, i)
			panic("baa")
		}
		p.set(buf, i, v)
	}
}

func (p *IntPacker) Unpack(ps PackedState) State {
	res := make(State, len(p.packageInfos))
	for i := range p.packageInfos {
		res[i] = p.get(ps, i)
		if res[i] < 0 {
			fmt.Printf("%+v %d\n%+v", ps, i, p.packageInfos)
			panic("waa")
		}
	}
	return res
}

func (p *IntPacker) getNumBins() int { return p.numBins }

func (p *IntPacker) get(buf []Bin, v int) int {
	return p.packageInfos[v].get(buf)
}

func (p *IntPacker) set(buf []Bin, v, val int) {
	p.packageInfos[v].set(buf, val)
}

// packBin packs a bin in a greedy fashion, always adding the largest variable
// that still fits. Returns the number of variables added to the bin.
func (p *IntPacker) packBin(ranges []int, bitsToVars [][]int) int {
	p.numBins += 1
	binIndex, usedBits, numVarsInBin := p.numBins-1, 0, 0
	for {
		// get largest var fitting in the bin
		bits := bitsPerBin - usedBits
		for bits > 0 && len(bitsToVars[bits]) == 0 {
			bits--
		}
		if bits == 0 { // bin full or all vars packed
			return numVarsInBin
		}
		bestFit := bitsToVars[bits]
		v := bestFit[len(bestFit)-1]
		bestFit = bestFit[:len(bestFit)-1]
		bitsToVars[bits] = bestFit
		p.packageInfos[v] = NewPackageInfo(ranges[v], binIndex, usedBits)
		usedBits += bits
		numVarsInBin++
	}
}

// packBins
// len(p.variableInfos) == 0
func (p *IntPacker) packBins(ranges []int) {
	numVars := len(ranges)
	p.packageInfos = make([]PackageInfo, numVars)
	bitsToVars := make([][]int, bitsPerBin+1)
	for v := numVars - 1; v >= 0; v-- {
		bits := requiredBitSize(ranges[v])
		bitsToVars[bits] = append(bitsToVars[bits], v)
	}

	packedVars := 0
	for packedVars != numVars {
		packedVars += p.packBin(ranges, bitsToVars)
	}
}

//// BinaryPacker assumes all variables have a binary domain
//type BinaryPacker struct {
//	stateSize int
//}
//
//func NewBinaryPacker(stateSize int) *BinaryPacker {
//	return &BinaryPacker{stateSize: stateSize}
//}
//
//func (b BinaryPacker) Pack(values State) PackedState {
//	var s PackedState
//	n := len(values)
//	if n%bitsPerInt == 0 {
//		s = make(PackedState, int(len(values)/bitsPerInt))
//	} else {
//		s = make(PackedState, int(len(values)/bitsPerInt+1))
//	}
//
//	v := 0
//	for i := 0; i < len(s)-1; i++ {
//		for j := 0; j < bitsPerInt; j++ {
//			s[i] = s[i] ^ (values[j+v] << uint(j))
//		}
//		v += bitsPerInt
//	}
//	remainder := n - v
//	for j := 0; j < remainder; j++ {
//		s[len(s)-1] = s[len(s)-1] ^ (values[j+v] << uint(j))
//	}
//
//	return s
//}
//
//func (b BinaryPacker) Unpack(s PackedState) State {
//	res := make(State, b.stateSize)
//	v := 0
//	for i := 0; i < len(s)-1; i++ {
//		for j := 0; j < bitsPerInt; j++ {
//			res[v] = (s[i] & (1 << uint(j))) >> uint(j)
//			v++
//		}
//	}
//	remainder := b.stateSize % bitsPerInt
//	for j := 0; j < remainder; j++ {
//		res[v] = (s[len(s)-1] & (1 << uint(j))) >> uint(j)
//		v++
//	}
//	return res
//}
//
//func (b BinaryPacker) ValueOf(s PackedState, v int) int {
//	var pos int
//	relPos := v % bitsPerInt
//	pos = v / bitsPerInt
//	return (s[pos] & (1 << uint(relPos))) >> uint(relPos)
//}
//
//func (b BinaryPacker) PutValue(s PackedState, v, d int) {
//	// TODO implement
//}
