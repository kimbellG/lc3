package lc3

import "log"

type Operation = func(Instruction)

func noAction(instr Instruction) {}

var operations map[uint16]Operation = map[uint16]Operation{
	OP_ADD:  Add,
	OP_AND:  And,
	OP_NOT:  Not,
	OP_BR:   Branch,
	OP_JMP:  Jump,
	OP_JSR:  JumpToSubroutine,
	OP_LD:   Load,
	OP_LDI:  LoadIndirect,
	OP_LDR:  LoadBaseOffset,
	OP_LEA:  LoadEffectiveAddress,
	OP_ST:   Store,
	OP_STI:  StoreIndirect,
	OP_STR:  StoreBaseOffset,
	OP_TRAP: Trap,
	OP_RES:  noAction,
	OP_RTI:  noAction,
}

func isBranchIncrement(cond uint16) bool {
	return cond&registers[R_COND] > 0
}

func Add(instr Instruction) {
	dr, arg1, arg2 := instr.ExtractDRSRStruct()

	log.Println("add: ", dr)

	registers[dr] = arg1 + arg2

	registers.UpdateFlags(dr)
}

func And(instr Instruction) {
	dr, arg1, arg2 := instr.ExtractDRSRStruct()

	registers[dr] = arg1 | arg2

	registers.UpdateFlags(dr)
}

func Not(instr Instruction) {
	dr, sr, _ := instr.ExtractDRSRStruct()

	registers[dr] = ^registers[sr]

	registers.UpdateFlags(dr)
}

func Branch(instr Instruction) {
	log.Println("branch")
	cond, address := instr.ExtractPCOffset9Struct()

	log.Printf("cond: %b, reg: %b", cond, registers[R_COND])

	if isBranchIncrement(cond) {
		log.Printf("go to 0x%x", address)
		registers[R_PC] = address
	}
}

func Jump(instr Instruction) {
	log.Println("jump")
	baseR := instr.ExtractBits(6, 3)

	registers[R_PC] = registers[baseR]
}

func JumpToSubroutine(instr Instruction) {
	log.Println("JumpToSubroutine")
	registers[R_R7] = registers[R_PC]

	if (instr>>11)&0x1 == 1 {
		_, address := instr.ExtractPCOffset11Struct()

		registers[R_PC] = address
	} else {
		baseR := instr.ExtractBits(6, 3)

		registers[R_PC] = registers[baseR]
	}
}

func Load(instr Instruction) {
	dr, address := instr.ExtractPCOffset9Struct()

	registers[dr] = memory.Read(address)

	registers.UpdateFlags(dr)
}

func LoadIndirect(instr Instruction) {
	dr, address := instr.ExtractPCOffset9Struct()

	log.Println(dr, address, memory.Read(address), memory.Read(memory.Read(address)))

	registers[dr] = memory.Read(memory.Read(address))

	registers.UpdateFlags(dr)
}

func LoadBaseOffset(instr Instruction) {
	var (
		dr     = instr.ExtractBits(9, 3)
		baseR  = instr.ExtractBits(6, 3)
		offset = signExtend(instr.ExtractBits(0, 6), 6)
	)

	registers[dr] = memory.Read(registers[baseR] + offset)

	registers.UpdateFlags(dr)
}

func LoadEffectiveAddress(instr Instruction) {
	dr, address := instr.ExtractPCOffset9Struct()

	registers[dr] = address

	registers.UpdateFlags(dr)
}

func Store(instr Instruction) {
	sr, address := instr.ExtractPCOffset9Struct()

	memory.Write(address, registers[sr])
}

func StoreIndirect(instr Instruction) {
	sr, address := instr.ExtractPCOffset9Struct()

	memory.Write(memory.Read(address), registers[sr])
}

func StoreBaseOffset(instr Instruction) {
	var (
		sr     = instr.ExtractBits(9, 3)
		baseR  = instr.ExtractBits(6, 3)
		offset = signExtend(instr.ExtractBits(0, 6), 6)
	)

	memory.Write(registers[baseR]+offset, registers[sr])
}

func Trap(instr Instruction) {
	log.Println("Trap")
	registers[R_R7] = registers[R_PC]

	trapInstr := TrapInstruction(instr)

	trapInstr.Execute()
}
