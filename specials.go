// 21 april 2013
package main

import (
	"fmt"
	"os"
)

// paging reference: http://turbo.mindrec.com/tginternals/hw/

// lda #nn
func lda_immediate(pos uint32) (disassembled string, newpos uint32, done bool) {
	b := bytes[pos]
	pos++
	a = b
	a_valid = true
	return fmt.Sprintf("lda\t#$%02X", b), pos, false
}

// inc a
func inc_accumulator(pos uint32) (disassembled string, newpos uint32, done bool) {
	// whether or not a is valid does not matter here (if a is invalid the value will not be used anyway)
	a++
	return fmt.Sprintf("inc\ta"), pos, false
}

// dec a
func dec_accumulator(pos uint32) (disassembled string, newpos uint32, done bool) {
	// whether or not a is valid does not matter here (if a is invalid the value will not be used anyway)
	a--
	return fmt.Sprintf("dec\ta"), pos, false
}

// pha
func pha_noarguments(pos uint32) (disassembled string, newpos uint32, done bool) {
//	pusha()
	return "pha", pos, false
}

// php, phx, phy
func op_push(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
//		pushinvalid()
		return fmt.Sprintf("%s", m), pos, false
	}
}

// pla
func pla_noarguments(pos uint32) (disassembled string, newpos uint32, done bool) {
//	popa()
	invalidate()
	return "pla", pos, false
}

// plp, plx, ply
func op_pop(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
//		pop()
		return fmt.Sprintf("%s", m), pos, false
	}
}

// tam #nn,...
func tam_pageregs(pos uint32) (disassembled string, newpos uint32, done bool) {
	b := bytes[pos]
	pos++
	prstring := ""
	curpage := 0
	for i := 0; i < 8; i++ {
		if b & 1 != 0 {		// mark this one
			if !a_valid {
				addcomment(pos - 2, "(!) cannot apply new page because a is not valid")
			} else {
				pages[curpage].which = a
				pages[curpage].valid = true
			}
			prstring += fmt.Sprintf("#%d,", curpage)
		}
		b >>= 1
		curpage++
	}
	if prstring == "" {
		fmt.Fprintf(os.Stderr, "tam defining nothing at $%X\n", pos - 2)
		prstring = "<nothing>"
	} else {
		prstring = prstring[:len(prstring) - 1]	// strip trailing comma
	}
	return fmt.Sprintf("tam\t%s", prstring), pos, false
}

var tmapages = map[byte]int{
	0x01:	0,
	0x02:	1,
	0x04:	2,
	0x08:	3,
	0x10:	4,
	0x20:	5,
	0x40:	6,
	0x80:	7,
}

// tma #nn
func tma_pageregs(pos uint32) (disassembled string, newpos uint32, done bool) {
	b := bytes[pos]
	pos++
	if _, ok := tmapages[b]; !ok {
		fmt.Fprintf(os.Stderr, "tma with invalid argument $%02X specified\n", b)
		invalidate()		// don't know what to do
		return fmt.Sprintf("tma\t<invalid $%02X>", b), pos, false
	}
	page := tmapages[b]
	a = pages[page].which
	a_valid = pages[page].valid
	return fmt.Sprintf("tma\t#%d", page), pos, false
}

// these are only special because of their unique formats
// TODO do any of them touch a?

// tst #nn,zz
func tst_zeropage(pos uint32) (disassembled string, newpos uint32, done bool) {
	invalidate()
	b := bytes[pos]
	pos++
	z := bytes[pos]
	pos++
	addoperandcomment(pos - 3, uint16(z))
	return fmt.Sprintf("tst\t#$%02X,%02X", b, z), pos, false
}

// tst #nn,zz,x
func tst_zeropagex(pos uint32) (disassembled string, newpos uint32, done bool) {
	invalidate()
	b := bytes[pos]
	pos++
	z := bytes[pos]
	pos++
	addoperandcomment(pos - 3, uint16(z))
	return fmt.Sprintf("tst\t#$%02X,%02X,x", b, z), pos, false
}

// tst #nn,hhll
func tst_absolute(pos uint32) (disassembled string, newpos uint32, done bool) {
	invalidate()
	b := bytes[pos]
	pos++
	w, pos := getword(pos)
	addoperandcomment(pos - 4, w)
	return fmt.Sprintf("tst\t#$%02X,%04X", b, w), pos, false
}

// tst #nn,hhll,x
func tst_absolutex(pos uint32) (disassembled string, newpos uint32, done bool) {
	invalidate()
	b := bytes[pos]
	pos++
	w, pos := getword(pos)
	addoperandcomment(pos - 4, w)
	return fmt.Sprintf("tst\t#$%02X,%04X,x", b, w), pos, false
}
