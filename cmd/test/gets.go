package main

import (
	"log"

	"github.com/kimbellG/lc3"
)

func main() {
	start := lc3.GetRegister(lc3.R_R0)
	lc3.TrapIn()
	end := lc3.GetRegister(lc3.R_R0)

	log.Printf("%c %c", start, end)
}
