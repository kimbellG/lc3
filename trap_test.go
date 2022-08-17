package lc3

import "testing"

func TestGets(t *testing.T) {
	TrapGets()

	ch := registers[R_R0]

	t.Logf("%b", ch)
}
