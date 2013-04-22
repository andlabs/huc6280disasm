// 21 april 2013
package main

import (
	"fmt"
	"os"
)

const operandString = "---"

var toDisassemble []uint32

func queueDisassemble(physical uint32) {
	toDisassemble = append(toDisassemble, physical)
}

func doDisassemble() {
	if len(toDisassemble) == 0 {
		return
	}
	pos := toDisassemble[0]
	toDisassemble = toDisassemble[1:]
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

func disassemble() {
	for len(toDisassemble) != 0 {
		doDisassemble()
	}
}

func getword(pos uint32) (w uint16, newpos uint32) {
	w = uint16(bytes[pos])
	pos++
	w |= uint16(bytes[pos]) << 8
	pos++
	return w, pos
}

// TODO watch for labels that cross into multi-byte instructions (that's what operandString is for)
func mklabel(bpos uint32, prefix string) (label string) {
	if labels[bpos] == "" {
		labels[bpos] = fmt.Sprintf("%s_%X", prefix, bpos)
	}
	return labels[bpos]
}
