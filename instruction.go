package lc3

import "math"

type Instruction uint16

func (i Instruction) GetOP() uint16 {
	return uint16(i) >> 12
}

func (i Instruction) Execute() {
	operation, ok := operations[i.GetOP()]
	if !ok {

		return
	}

	operation(i)
}

func (i Instruction) ExtractDRSRStruct() (dr, arg1, arg2 uint16) {
	const (
		immediate = 1
	)

	dr = uint16((i >> 9) & 0x7)
	arg1 = registers[uint16((i>>6)&0x7)]

	mode := uint16((i >> 5) & 0x1)

	if mode == immediate {
		arg2 = signExtend(uint16(i&0x1F), 5)
	} else {
		arg2 = registers[uint16(i&0x7)]
	}

	return dr, arg1, arg2
}

func (i Instruction) ExtractPCOffset9Struct() (r, address uint16) {
	return i.ExtractPCOffsetStruct(9)
}

func (i Instruction) ExtractPCOffset11Struct() (f, address uint16) {
	return i.ExtractPCOffsetStruct(11)
}

func (i Instruction) ExtractPCOffsetStruct(PCOffsetSize int) (other, address uint16) {
	other = uint16((i.uint16() >> PCOffsetSize) & i.createOneMask(i.Size()-i.OPSize()-PCOffsetSize))

	offset := signExtend(uint16(i.uint16()&i.createOneMask(PCOffsetSize)), PCOffsetSize)

	return other, registers[R_PC] + offset
}

func (i Instruction) uint16() uint16 {
	return uint16(i)
}

func (i Instruction) ExtractBits(start, count int) uint16 {
	return (i.uint16() >> start) & i.createOneMask(count)
}

// mask format example: bitCount = 3: 0000 0000 0000 0111.
func (i Instruction) createOneMask(bitCount int) uint16 {
	return uint16(math.Pow(2, float64(bitCount)) - 1)
}

func (i Instruction) Size() int {
	return 16
}

func (i Instruction) OPSize() int {
	return 4
}
