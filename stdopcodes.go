// 21 april 2013

import (
	"fmt"
)

// xxx #nn
func op_immediate(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		b := bytes[pos]
		pos++
		return fmt.Sprintf("%s\t#$%02X", m, b), pos, false
	}
}

// xxx zz
func op_zeropage(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		b := bytes[pos]
		pos++
		return fmt.Sprintf("%s\t$%02X", m, b), pos, false
	}
}

// xxx zz,x
func op_zeropagex(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		b := bytes[pos]
		pos++
		return fmt.Sprintf("%s\t$%02X,x", m, b), pos, false
	}
}

// xxx zz,y
func op_zeropagey(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		b := bytes[pos]
		pos++
		return fmt.Sprintf("%s\t$%02X,y", m, b), pos, false
	}
}

// xxx (zz)
func op_indirect(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		b := bytes[pos]
		pos++
		return fmt.Sprintf("%s\t($%02X)", m, b), pos, false
	}
}

// xxx (zz,x)
func op_indirectx(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		b := bytes[pos]
		pos++
		return fmt.Sprintf("%s\t($%02X,x)", m, b), pos, false
	}
}

// xxx (zz),y
func op_indirecty(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		b := bytes[pos]
		pos++
		return fmt.Sprintf("%s\t($%02X),y", m, b), pos, false
	}
}

// xxx hhll
func op_absolute(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		w, pos := getword(pos)
		return fmt.Sprintf("%s\t$%04X", m, w), pos, false
	}
}

// xxx hhll,x
func op_absolutex(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		w, pos := getword(pos)
		return fmt.Sprintf("%s\t$%04X,x", m, w), pos, false
	}
}

// xxx hhll,y
func op_absolutey(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		w, pos := getword(pos)
		return fmt.Sprintf("%s\t$%04X,y", m, w), pos, false
	}
}

// xxx a
func op_accumulator(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		return fmt.Sprintf("%s\ta", m), pos, false
	}
}

// xxx
func op_noarguments(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		return fmt.Sprintf("%s", m), pos, false
	}
}

// xxx #nn,zz
func op_zpbit(m string, n int) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		b := bytes[pos]
		pos++
		return fmt.Sprintf("%s\t#%d,$%02X", m, n, b), pos, false
	}
}

// xxx hhll,hhll,hhll
func op_transfer(m string) opcode {
	return func(pos uint32) (disassembled string, newpos uint32, done bool) {
		invalidate()
		src, pos := getword(pos)		// TODO make labels for src and dest?
		dest, pos := getword(pos)
		len, pos := getword(pos)
		return fmt.Sprintf("%s\t$%04X,$%04X,$%04X", m, src, dest, len), pos, false
	}
}
