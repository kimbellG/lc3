package lc3

// #cgo CFLAGS: -g -Wno-implicit
// #include <stdlib.h>
// #include <unistd.h>
// #include <fcntl.h>
// #include <sys/time.h>
// #include <sys/types.h>
// #include <sys/termios.h>
// #include <sys/mman.h>
//
// struct termios original_tio;
//
// void disable_input_buffering()
// {
//     tcgetattr(STDIN_FILENO, &original_tio);
//     struct termios new_tio = original_tio;
//     new_tio.c_lflag &= ~ICANON & ~ECHO;
//     tcsetattr(STDIN_FILENO, TCSANOW, &new_tio);
// }
//
// void restore_input_buffering()
// {
//     tcsetattr(STDIN_FILENO, TCSANOW, &original_tio);
// }
//
// u_int16_t check_key()
// {
//     fd_set readfds;
//     FD_ZERO(&readfds);
//     FD_SET(STDIN_FILENO, &readfds);
//
//     struct timeval timeout;
//     timeout.tv_sec = 0;
//     timeout.tv_usec = 0;
//     return select(1, &readfds, NULL, NULL, &timeout) != 0;
// }
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
)

func Run(images ...string) error {
	for _, image := range images {
		if err := readImage(image); err != nil {
			return err
		}
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)

	C.disable_input_buffering()
	go func() {
		<-sigs

		C.restore_input_buffering()
		os.Exit(0)
	}()

	registers[R_COND] = FL_ZRO

	const PC_START = 0x3000

	registers[R_PC] = PC_START

	for {
		instruction := memory.ReadInstruction()

		instruction.Execute()
	}
}

func checkKey() bool {
	ret := C.check_key()

	return ret != 0
}

func readImage(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open %s: %w", path, err)
	}
	defer file.Close()

	if err := readProgram(file); err != nil {
		return fmt.Errorf("read image: %w", err)
	}

	return nil
}

func readProgram(file io.Reader) error {
	var origin uint16

	originData := make([]byte, 2)

	if _, err := file.Read(originData); err != nil {
		return fmt.Errorf("read origin: %w", err)
	}

	buf := bytes.NewBuffer(originData)

	err := binary.Read(buf, binary.BigEndian, &origin)
	if err != nil {
		return fmt.Errorf("translate to binary failed: %w", err)
	}

	log.Printf("read program: received origin: 0x%x", origin)

	programData := make([]byte, MemoryMax-int(origin))
	if _, err := file.Read(programData); err != nil {
		return fmt.Errorf("read program: %w", err)
	}

	intByte := make([]byte, 0, 2)
	for _, b := range programData {
		if len(intByte) == 2 {
			val := binary.BigEndian.Uint16(intByte)
			memory[origin] = val

			origin++

			intByte = intByte[:0]
		}

		intByte = append(intByte, b)
	}

	return nil
}
