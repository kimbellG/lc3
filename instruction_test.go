package lc3

import "testing"

func TestExtractPCOffsetStruct(t *testing.T) {
	tt := []struct {
		instr                  Instruction
		PCOffsetSize           int
		wantOther, wantAddress uint16
	}{
		{
			instr:        Instruction(0x08FF),
			PCOffsetSize: 11,
			wantOther:    0x1,
			wantAddress:  0xFF,
		},
		{
			instr:        Instruction(0x08FF),
			PCOffsetSize: 9,
			wantOther:    0x4,
			wantAddress:  0xFF,
		},
	}

	for _, tc := range tt {
		other, address := tc.instr.ExtractPCOffsetStruct(tc.PCOffsetSize)

		if other != tc.wantOther {
			t.Errorf("other isn't equal: want: %X: got: %X", tc.wantOther, other)
		}

		if address != tc.wantAddress {
			t.Errorf("address isn't equal: want: %X: got: %X", tc.wantAddress, address)
		}

	}
}
