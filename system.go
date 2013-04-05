// 5 april 2013
package main

import (
	// ...
)

var a byte
var a_valid bool

var pages [8]struct{
	which	byte
	valid		bool
}

var carryflag int

func init() {
	// TODO verify all this
	a_valid = false
	for i := 0; i < 7; i++ {
		pages[i].valid = false
	}
	pages[7].which = 0x00		// we need the vectors at startup
	pages[7].valid = true
	carryflag = 0
}

func physical(logical uint16) uint32 {
	page := (logical & 0xE000) >> 14
	if !pages[page].valid {
		errorf("attempt to get physical address of logical $%X, but the page has not yet been initialized", logical)
	}
	physical := uint32(logical) &^ 0xE000
	physical |= 0x2000 * uint32(pages[page].which)
	return physical
}