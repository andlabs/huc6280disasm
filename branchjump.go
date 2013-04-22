// 21 april 2013
package main

import (
	"fmt"
	"os"
)

func dobranch(pos uint32) (label string, newpos uint32) {
	b := bytes[pos]
	pos++
	offset := int32(int8(b))
	// TODO does not properly handle jumps across page boundaries
	bpos := uint32(int32(pos) + offset)
	label = mklabel(bpos, "loc")
	queueDisassemble(bpos)
	return label, pos
}

// xxx label
func op_branch(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		label, pos := dobranch(pos)
		return fmt.Sprintf("%s\t%s", m, label), pos, false
	}
}

// xxx #nn,zz,label
func op_zpbitbr(m string, n int) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		b := bytes[pos]
		pos++
		label, pos := dobranch(pos)
		return fmt.Sprintf("%s\t#%d,$%02X,%s", m, n, b, label), pos, false
	}
}

// jmp hhll
func jmp_absolute(pos uint32) (disassembled string, newpos uint32, done bool) {
	w, pos := getword(pos)
	phys, err := physical(w)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get physical address for jmp to $%04X: %v\n", w, err)
		return fmt.Sprintf("jmp\t$%04X", w), pos, true
	}
	label := mklabel(phys, "loc")
	queueDisassemble(phys)
	return fmt.Sprintf("jmp\t%s", label), pos, true
}

// jmp hhll,x
func jmp_absolutex(pos uint32) (disassembled string, newpos uint32, done bool) {
	w, pos := getword(pos)
	return fmt.Sprintf("jmp\t$%04X,x", w), pos, true
}

// jmp (hhll)
func jmp_absoluteindirect(pos uint32) (disassembled string, newpos uint32, done bool) {
	w, pos := getword(pos)
	return fmt.Sprintf("jmp\t($%04X)", w), pos, true
}

// jsr hhll
func jsr_absolute(pos uint32) (disassembled string, newpos uint32, done bool) {
	w, pos := getword(pos)
	phys, err := physical(w)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get physical address for jsr to $%04X: %v\n", w, err)
		return fmt.Sprintf("jsr\t$%04X", w), pos, true
	}
	label := mklabel(phys, "sub")
	queueDisassemble(phys)
	return fmt.Sprintf("jsr\t%s", label), pos, false
}

// rti
func rti_noarguments(pos uint32) (disassembled string, newpos uint32, done bool) {
	return "rti", pos, true
}

// rts
func rts_noarguments(pos uint32) (disassembled string, newpos uint32, done bool) {
	return "rts", pos, true
}
