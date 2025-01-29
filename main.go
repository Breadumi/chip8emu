package main

import "time"

type chip8 struct {
	opcode uint16     // stores current opcode
	memory [4096]byte // memory locations
	// 0x000 - 0x1FF - Chip 8 interpreter (contains font set in emu)
	// 0x050 - 0x0A0 - Used for the built-in 4x5 pixel font set (0-F)
	// 0x200 - 0xFFF - Program ROM and work RAM
	register    [16]byte   // V0 to VF registers (VF doubles as carry flag)
	idx         uint16     // index register
	pc          uint16     // program counter
	delay_timer byte       // 60 Hz
	sound_timer byte       // 60 Hz (sound buzzer at 0)
	stack       [16]uint16 // 16 level stack
	sp          uint16     // stack pointer
	key         [16]byte   // HEX based keypad (0x0 - 0xF)
}

func main() {

	// Set up render system and register input callbacks
	// setupGraphics()
	// setupInput()

	// Initialize the Chip8 system and load the game into memory
	chip8 := initialize()
	chip8.loadFontSet()
	// myChip8.loadGame("pong")
	ticker := time.NewTicker(time.Second / 60)
	defer ticker.Stop()
	done := make(chan struct{})

	// Emulation loop
	// for begin
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			chip8.emulateCycle()

			// if draw flag set, update the screen
			// if myChip8.drawFlag {
			//     drawGraphics()
			// }

			// Store key press state (Press and Release)
			// myChip8.setKeys()
		}

		// for end
	}
}
