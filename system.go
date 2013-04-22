// 5 april 2013
package main

import (
	"fmt"
)

// TODO change "valid" to "known"

var a byte
var a_valid bool

var pages [8]struct{
	which	byte
	valid		bool
}

var carryflag int
var carryflag_valid bool

func init() {
	// TODO verify all this
	a_valid = false
	for i := 0; i < 7; i++ {
		pages[i].valid = false
	}
	pages[7].which = 0x00		// we need the vectors at startup
	pages[7].valid = true
	carryflag = 0
	carryflag_valid = false
}

func physical(logical uint16) (uint32, error) {
	page := (logical & 0xE000) >> 13
	if !pages[page].valid {
		return 0, fmt.Errorf("attempt to get physical address of logical $%X, but the page has not yet been initialized", logical)
	}
	physical := uint32(logical) &^ 0xE000
	physical |= 0x2000 * uint32(pages[page].which)
	return physical, nil
}

func invalidate() {
	a_valid = false
	carryflag_valid = false
}
