package main

import (
	"encoding/hex"
	"testing"
)

func TestEmulateCycle(t *testing.T) {
	chip8 := initialize()
	if v, err := hex.DecodeString("B2"); err == nil {
		chip8.memory[0] = v[0]
	}
	if v, err := hex.DecodeString("04"); err == nil {
		chip8.memory[1] = v[0]
	}
	chip8.emulateCycle()
	if chip8.opcode != uint16(45572) {
		t.Errorf("Opcode does not match - Expected: %v, Actual: %v", 45572, chip8.opcode)
	}
}
