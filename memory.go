package lc3

import "log"

type Memory [MemoryMax]uint16

func (m Memory) Read(address uint16) uint16 {
	if address == MR_KBSR {
		if checkKey() {
			memory[MR_KBSR] = 1 << 15
			memory[MR_KBDR] = getchar()
		} else {
			memory[MR_KBSR] = 0
		}
	}

	return m[address]
}

func (m Memory) Write(address uint16, val uint16) {
	m[address] = val
}

func (m Memory) ReadInstruction() Instruction {
	instruction := m.Read(registers[R_PC])
	registers[R_PC]++

	log.Printf("pc: 0x%x, instruction: %b", registers[R_PC], Instruction(instruction).GetOP())

	return Instruction(instruction)
}
