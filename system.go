// 5 april 2013
package main

import (
	"fmt"
)

// TODO change "valid" to "known"

type validbyte struct {
	value	byte
	valid		bool
}

type envt struct {
	a		validbyte
	pages	[8]validbyte
	carryflag	validbyte
	stack	[]validbyte
}

var env *envt

func newenv() *evnt {
	e := new(envt)
	// TODO verify all this
	e.a.valid = false
	for i := 0; i < 7; i++ {
		e.pages[i].valid = false
	}
	e.pages[7].value = 0x00		// we need the vectors at startup
	e.pages[7].valid = true
	e.carryflag.value = 0
	e.carryflag.valid = false
	return e
}

func init() {
	env = newenv()
}

func physical(logical uint16) (uint32, error) {
	page := (logical & 0xE000) >> 13
	if !env.pages[page].valid {
		return 0, fmt.Errorf("attempt to get physical address of logical $%X, but the page has not yet been initialized", logical)
	}
	physical := uint32(logical) &^ 0xE000
	physical |= 0x2000 * uint32(env.pages[page].value)
	return physical, nil
}

func invalidate() {
	env.a.valid = false
	env.carryflag.valid = false
}

func push(value byte, valid bool) {
	env.stack = append(env.stack, validbyte{
		value:	value,
		valid:	valid,
	})
}

func pop() (value byte, valid bool) {
	if len(env.stack) == 0 {
		return 0, false	// TODO correct?
	}
	t := env.stack[len(env.stack) - 1]
	stack = env.stack[:len(env.stack) - 1]
	return t.value, t.valid
}

func pusha() {
	push(env.a, env.a.valid)
}

func pushinvalid() {
	push(env.a, false)		// value of a irrelevant
}

func popa() {
	env.a, env.a.valid = pop()
}

func saveenv() *envt {
	e := new(envt)
	e.a = env.a
	e.pages = env.pages
	e.carryflag = env.carryflag
	e.stack = make([]validbyte, len(env.stack))
	copy(e.stack, env.stack)
	return e
}

func restoreenv(e *envt) {
	env = e
}
