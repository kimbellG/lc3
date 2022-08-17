package lc3

// #cgo CFLAGS: -g -Wno-implicit
// #include <termios.h>
// #include <unistd.h>
// #include <stdio.h>
// int getch(void)
// {
//   struct termios oldattr, newattr;
//   int ch;
//   tcgetattr(STDIN_FILENO, &oldattr);
//   newattr = oldattr;
//    newattr.c_lflag &= ~(ICANON | ECHO);
//    tcsetattr(STDIN_FILENO, TCSANOW, &newattr);
//    ch = getchar();
//    tcsetattr(STDIN_FILENO, TCSANOW, &oldattr);
//    return ch;
//}
import "C"
import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var screen *bytes.Buffer = new(bytes.Buffer)
var output *bufio.Writer = bufio.NewWriter(os.Stdin)

type TrapInstruction uint16

func moveCursor(x, y int) {
	fmt.Fprintf(screen, "\033[%d;%dH", x, y)
}

func getchar() uint16 {
	b := C.getchar()

	return uint16(b)
}

func clearTerminal() {
	output.WriteString("\033[2J")
}

func (instr TrapInstruction) Execute() {
	switch instr & 0xFF {
	case TRAP_GETC:
		TrapGets()
	case TRAP_OUT:
		TrapOut()
	case TRAP_PUTS:
		TrapPuts()
	case TRAP_IN:
		TrapIn()
	case TRAP_PUTSP:
		TrapPutsp()
	case TRAP_HALT:
		TrapHalt()
	}
}

func TrapGets() {
	ch := C.getch()

	registers[R_R0] = uint16(ch)

	registers.UpdateFlags(R_R0)
}

func TrapOut() {
	ch := byte(registers[R_R0])

	os.Stdout.Write([]byte{ch})
}

func TrapPuts() {
	var (
		str = make([]byte, 0)
	)

	for i, c := 0, memory[registers[R_R0]]; c != 0; i++ {
		c = memory[int(registers[R_R0])+i]

		str = append(str, byte(c))
	}

	os.Stdout.Write(str)
}

func TrapIn() {
	fmt.Print("Execute character: ")
	ch := make([]byte, 1)

	os.Stdin.Read(ch)
	os.Stdout.Write(ch)

	registers[R_R0] = uint16(ch[0])

	registers.UpdateFlags(R_R0)
}

func TrapPutsp() {
	var (
		str = make([]byte, 0)
	)

	for i, c := 0, memory[registers[R_R0]]; c != 0; i++ {
		c = memory[int(registers[R_R0])+i]

		str = append(str, byte(c&0xFF), byte(c>>8))
	}

	os.Stdout.Write(str)
}

func TrapHalt() {
	fmt.Println("Halt")
	os.Exit(0)
}
