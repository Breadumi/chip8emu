package main

import (
	"fmt"
	"log"
	"strconv"
)

func (c *chip8) emulateCycle() {
	// Fetch Opcode
	c.opcode = b2i(c.memory[c.pc], c.memory[c.pc+1])
	// Decode Opcode

	// Execute Opcode

	// Update timers
}

// concatenate two bytes together and convert to uint16
func b2i(b1, b2 byte) uint16 {

	bits := fmt.Sprintf("%08b%08b", b1, b2)
	num, err := strconv.ParseUint(bits, 2, 16)
	if err != nil {
		log.Fatalf("could not convert bytes to uint16")
	}
	return uint16(num)
}
