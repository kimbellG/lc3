package lc3

func signExtend(n uint16, bitCount int) uint16 {
	if (n>>(bitCount-1))&1 == 1 {
		n |= (0xFFFF << bitCount)
	}

	return n
}
