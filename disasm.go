// 21 april 2013
package main

import (
	"fmt"
	"os"
)

const operandString = "---"

func disassemble(pos uint32) {
	for {
		if _, already := instructions[pos]; already {
			break		// reached a point we previously reached
		}
		b := bytes[pos]
		op := opcodes[b]
		if op == nil {
			// TODO make a comment
			fmt.Fprintf(os.Stderr, "illegal opcode at $%X\n", pos)
			break
		}
		s, newpos, done := op(pos + 1)
		instructions[pos] = s
		for i := pos + 1; i < newpos; i++ {
			instructions[i] = operandString
		}
		if done {
			break
		}
		pos = newpos
	}
}

func getword(pos uint32) (w uint16, newpos uint32) {
	w = uint16(byte[pos])
	pos++
	w |= uint16(byte[pos]) << 8
	pos++
	return w, pos
}
