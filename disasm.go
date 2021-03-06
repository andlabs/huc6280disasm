// 21 april 2013
package main

import (
	"fmt"
	"os"
)

const operandString = "---"

func disassemble(pos uint32) {
	if pos >= uint32(len(bytes)) {
		fmt.Fprintf(os.Stderr, "cannot disassemble at $%X as it is past ROM (size $%X bytes)\n", pos, len(bytes))
		return
	}
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
	w = uint16(bytes[pos])
	pos++
	w |= uint16(bytes[pos]) << 8
	pos++
	return w, pos
}

// TODO watch for labels that cross into multi-byte instructions (that's what operandString is for)
func mklabel(bpos uint32, prefix string, priority int) (label string) {
	mk := false
	if labels[bpos] == "" {					// new label
		mk = true
	} else if labelpriorities[bpos] <= priority {		// higher (or same) priority label
		mk = true
	}
	if mk {
		labels[bpos] = fmt.Sprintf("%s_%X", prefix, bpos)
		labelpriorities[bpos] = priority
	}
	return labels[bpos]
}

func addcomment(pos uint32, format string, args ...interface{}) {
	c := fmt.Sprintf(format, args...)
	if comments[pos] != "" {
		comments[pos] += " | "
	}
	comments[pos] += c
}

func addoperandcomment(pos uint32, logical uint16) {
	physical, err := physical(logical)
	if err != nil {
		addcomment(pos, "$%04X - cannot get physical address (%v)", logical, err)
		return
	}
	addcomment(pos, "$%04X -> $%X", logical, physical)
}
