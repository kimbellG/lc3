package lc3

const MemoryMax = 1 << 16

var (
	memory    Memory    // Machine RAM
	registers Registers // Space for all registers
)

// Machine register list.
const (
	R_R0 = iota // default register from 0 to 7
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC   // program counter
	R_COND // condition register
	R_COUNT
)

// Condition flags. It means sign of the previous calculation
const (
	FL_POS = 1 << iota // > 0
	FL_ZRO             // = 0
	FL_NEG             // < 0

)

// Machine commands
const (
	OP_BR   uint16 = iota // branch
	OP_ADD                // add
	OP_LD                 // load
	OP_ST                 // store
	OP_JSR                // jump register
	OP_AND                // bitwise and
	OP_LDR                // load register
	OP_STR                // store register
	OP_RTI                // unused
	OP_NOT                // bitwise not
	OP_LDI                // load indirect
	OP_STI                // store indirect
	OP_JMP                // jump
	OP_RES                // reserved
	OP_LEA                // load effective address
	OP_TRAP               // execute trap
)

// TRAP commands control codes
const (
	TRAP_GETC  = 0x20 + iota // get char from keyboard, not echoed onto the terminal
	TRAP_OUT                 // output a character
	TRAP_PUTS                // output a word string
	TRAP_IN                  // get char from keyboard, echoed onto the terminal
	TRAP_PUTSP               // output a byte string
	TRAP_HALT                // halt a program
)

// memory registers
const (
	MR_KBSR = 0xFE00
	MR_KBDR = 0xFE02
)
