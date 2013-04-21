// 21 april 2013
package main

import (
	"fmt"
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

// tam #nn,...
func tam_pageregs(pos uint32) (disassembled string, newpos uint32, done bool) {
	b := bytes[pos]
	pos++
	prstring := ""
	curpage := 0
	for i := 0; i < 8; i++ {
		if b & 1 != 0 {		// mark this one
			if !a_valid {
				fmt.Fprintf(os.Stderr, "cannot apply new page because a is not valid at $%X", pos - 2)
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
		prstring = prstring[:len(prstring) - 2]	// strip trailing comma
	}
	return fmt.Sprintf("tam\t%s", prstring), pos, false
}

var tmapages = map[byte]string{
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
	a = pages[page].value
	a_valid = pages[page].valid
	return fmt.Sprintf("tma\t#%d", page), pos, false
}
