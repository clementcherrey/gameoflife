package bitmap

import "encoding/binary"

// MinNodes represents the minimum block of cells
// that we need to compute the next step.
//
// It is a 4x4 square of bits. For example:
// o 1 o o
// 1 1 o o
// o 1 o 1
// 1 o o o
//
type MinNodes uint16

// NewMinNodes creates a MinNodes pointer
func NewMinNodes(u uint16) *MinNodes {
	mn := MinNodes(u)
	return &mn
}

// Sets the bit at pos in MinNodes
func (mn *MinNodes) SetBit(pos uint) {
	*mn |= 1 << pos
}

// Similar to SetBit but for a single byte
func setBit(b *byte, pos uint) {
	*b |= 1 << pos
}

// Clears the bit at pos in MinNodes
func (mn *MinNodes) UnSetBit(pos uint) {
	*mn |= ^(1 << pos)
}

// GetBit returns the value of bit i of byte b.
// The bit index must be between 0 and 7.
func (mn *MinNodes) GetBit(pos uint) bool {
	return *mn&(1<<pos) != 0
}

// o o a o      o a o o
// o b c d -->  b c d o
// o o e o      o e o o
// 0 o o o      o o o o
// TODO describe what's crossPos
func (mn *MinNodes) MoveTopLeft(crossPos uint) MinNodes {
	switch crossPos {
	case 0:
		crossPos = 0
	case 1:
		crossPos = 1
	case 2:
		crossPos = 4
	case 3:
		crossPos = 5
	default:
		panic("crosspos invalid")
	}
	return *mn << crossPos
}

// TODO: finish test
// TODO: create a separate type for result?
func (mn *MinNodes) ApplyGoFRule() byte {
	var result byte // we only care about the last 4 bits
	for _, i := range []uint{0, 1, 2, 3} {
		if mn.MoveTopLeft(i).ApplyGoFRuleToTopLeft() {
			setBit(&result, 3-i) // TODO: set using opposite order?
		}
	}
	return result
}

// TODO: describe
func (mn MinNodes) ApplyGoFRuleToTopLeft() bool {
	surroundingCoordinates := []MinCoordinate{
		{1, 0},
		{0, 1},
		{2, 1},
		{1, 2},
	}
	cnt := 0
	for _, c := range surroundingCoordinates {
		if mn.GetBit(coordToPos(c)) {
			cnt++
		}
	}

	isAlive := mn.GetBit(rawCoordToPos(1, 1))

	if isAlive && cnt > 1 && cnt < 4 {
		return true
	}
	if !isAlive && cnt == 3 {
		return true
	}
	return false
}

// apply basic rule
// o 1 o o
// 1 1 1 o --> 0100 1110  0100 0000
// o 1 o o
// o o o o
// TODO drop
func (mn *MinNodes) getInnerCrosses() []uint16 {
	// mask to get each cross
	masks := []uint16{0x4E40} // TODO: add 3 more masks

	crosses := make([]uint16, 0, 4)
	for _, mask := range masks {
		crosses = append(crosses, uint16(*mn)&mask)
	}
	return crosses
}

// coordinate for MinNodes
// this represent the intuitive way of thinking of it
// example
// o o o o
// o o o o --> b column 2, line 3
// o a o o --> a column 1, line 1
// o o b o
type MinCoordinate struct {
	column, row uint
}

// convert MinCoordinate to position (~index)
//
// TODO: make sure this is what we have and what we want
// IMPORTANT: index starts at the bottom right and goes from left-right, bottom-up
// so we have:
// (0,0) --> 15
// (2,3) --> 1
// (3,2) --> 4
// TODO add test
func coordToPos(c MinCoordinate) uint {
	return (3 - c.column) + (4 * (3 - c.row))
}
func rawCoordToPos(column, row uint) uint {
	return (3 - column) + (4 * (3 - row))
}

// TODO
func posToCoord(pos uint) MinCoordinate {
	return MinCoordinate{0, 0}
}

// Aggregates creates a new MinNodes given 4 results
//
// the input r is an slice of 4 bytes that represents 4 results
// Note: order matter
//
// How it works ?
// - first we build 2 lines that each contains 2 squares of 4 bits
//   each line is build this way:
//   0000 aaaa + 0000 bbbb --> a a b b  (i.e. aabb aabb)
//                             a a b b
// - then we aggregate the 2 lines to a MinNodes
//
func Aggregate(r []byte) *MinNodes {
	l1 := ((r[0] & 12) << 4) | ((r[0] & 3) << 2) | ((r[1] & 12) << 2) | (r[1] & 3)
	l2 := ((r[2] & 12) << 4) | ((r[2] & 3) << 2) | ((r[3] & 12) << 2) | (r[3] & 3)

	var square uint16
	binary.LittleEndian.PutUint16([]byte{l1, l2}, square)
	return NewMinNodes(square)
}
