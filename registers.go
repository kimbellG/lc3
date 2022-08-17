package lc3

type Registers [R_COUNT]uint16

func (r Registers) UpdateFlags(reg uint16) {
	switch {
	case r[reg] == 0:
		r.SetCondition(FL_ZRO)
	case r[reg]>>15 == 1:
		r.SetCondition(FL_NEG)
	default:
		r.SetCondition(FL_POS)
	}
}

func (r Registers) SetCondition(status uint16) {
	r[R_COND] = status
}

func GetRegister(r uint16) uint16 {
	return registers[r]
}
