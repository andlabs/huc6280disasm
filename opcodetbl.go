// 5 april 2013
package main

// opcode reference: http://shu.emuunlim.com/download/pcedocs/pce_cpu.html

// TODO adjust immediates so that they have effect on a
// TODO figure out which have no effect on a and make them not mark a as invalid

type opcode func(pos uint32) (disassembled string, newpos uint32, done bool)

var opcodes [0x100]opcode

func init() {
	// this must be in init() to avoid a compile-time initialization loop
	opcodes = [0x100]opcode{
		// adc: add with carry
		0x69:	op_immediate("adc"),	// adc #nn
		0x65:	op_zeropage("adc"),		// adc zz
		0x75:	op_zeropagex("adc"),	// adc zz,x
		0x72:	op_indirect("adc"),		// adc (zz)
		0x61:	op_indirectx("adc"),		// adc (zz,x)
		0x71:	op_indirecty("adc"),		// adc (zz),y
		0x6D:	op_absolute("adc"),		// adc hhll
		0x7D:	op_absolutex("adc"),	// adc hhll,x
		0x79:	op_absolutey("adc"),	// adc hhll,y

		// and: bitwise and
		0x29:	op_immediate("and"),	// and #nn
		0x25:	op_zeropage("and"),	// and zz
		0x35:	op_zeropagex("and"),	// and zz,x
		0x32:	op_indirect("and"),		// and (zz)
		0x21:	op_indirectx("and"),		// and (zz,x)
		0x31:	op_indirecty("and"),		// and (zz),y
		0x2D:	op_absolute("and"),		// and hhll
		0x3D:	op_absolutex("and"),	// and hhll,x
		0x39:	op_absolutey("and"),	// and hhll,y

		// asl: arithmetic shift left
		0x06:	op_zeropage("asl"),		// asl zz
		0x16:	op_zeropagex("asl"),	// asl zz,x
		0x0E:	op_absolute("asl"),		// asl hhll
		0x1E:	op_absolutex("asl"),		// asl hhll,x
		0x0A:	op_accumulator("asl"),	// asl a

		// bbr: branch on bit clear (reset)
		0x0F:	op_zpbitbr("bbr", 0),	// bbr #0,zz,hhll
		0x1F:	op_zpbitbr("bbr", 1),	// bbr #1,zz,hhll
		0x2F:	op_zpbitbr("bbr", 2),	// bbr #2,zz,hhll
		0x3F:	op_zpbitbr("bbr", 3),	// bbr #3,zz,hhll
		0x4F:	op_zpbitbr("bbr", 4),	// bbr #4,zz,hhll
		0x5F:	op_zpbitbr("bbr", 5),	// bbr #5,zz,hhll
		0x6F:	op_zpbitbr("bbr", 6),	// bbr #6,zz,hhll
		0x7F:	op_zpbitbr("bbr", 7),	// bbr #7,zz,hhll

		// bcc: branch on carry clear
		0x90:	op_branch("bcc"),		// bcc hhll

		// bbs: branch on bit set
		0x8F:	op_zpbitbr("bbs", 0),	// bbs #0,zz,hhll
		0x9F:	op_zpbitbr("bbs", 1),	// bbs #1,zz,hhll
		0xAF:	op_zpbitbr("bbs", 2),	// bbs #2,zz,hhll
		0xBF:	op_zpbitbr("bbs", 3),	// bbs #3,zz,hhll
		0xCF:	op_zpbitbr("bbs", 4),	// bbs #4,zz,hhll
		0xDF:	op_zpbitbr("bbs", 5),	// bbs #5,zz,hhll
		0xEF:	op_zpbitbr("bbs", 6),	// bbs #6,zz,hhll
		0xFF:	op_zpbitbr("bbs", 7),	// bbs #7,zz,hhll

		// bcs: branch on carry set
		0xB0:	op_branch("bcs"),		// bcs hhll

		// beq: branch on equal
		0xF0:	op_branch("beq"),		// beq hhll

		// bit: test bit of accumulator
		0x89:	op_immediate("bit"),	// bit #nn
		0x24:	op_zeropage("bit"),		// bit zz
		0x34:	op_zeropagex("bit"),	// bit zz,x
		0x2C:	op_absolute("bit"),		// bit hhll
		0x3C:	op_absolutex("bit"),		// bit hhll,x

		// bmi: branch on minus
		0x30:	op_branch("bmi"),		// bmi hhll

		// bne: branch on not equal
		0xD0:	op_branch("bne"),		// bne hhll

		// bpl: branch on plus
		0x10:	op_branch("bpl"),		// bpl hhll

		// bra: branch
		0x80:	op_branch("bra"),		// bra hhll

		// brk: software break
		0x00:	op_noarguments("brk"),	// brk

		// bsr: branch to subroutine
		0x44:	op_branch("bsr"),		// bsr hhll
		// TODO make it produce a sub_ label

		// bvs: branch on overflow set
		0x70:	op_branch("bvs"),		// bvs hhll

		// bvc: branch on overflow clear
		0x50:	op_branch("bvc"),		// bvc hhll

		// clc: clear carry flag
		0x18:	op_noarguments("clc"),	// clc

		// cla: clear accumulator
		0x62:	op_noarguments("cla"),	// cla

		// cld: clear decimal flag
		0xD8:	op_noarguments("cld"),	// cld

		// cli: ENABLE interrupts (clears interrupt disable flag)
		0x58:	op_noarguments("cli"),	// cli

		// clv: clear overflow flag
		0xB8:	op_noarguments("clv"),	// clv

		// cly: clear y register
		0xC2:	op_noarguments("cly"),	// cly

		// clx: clear x register
		0x82:	op_noarguments("clx"),	// clx

		// cpx: compare x
		0xE0:	op_immediate("cpx"),	// cpx #nn
		0xE4:	op_zeropage("cpx"),	// cpx zz
		0xEC:	op_absolute("cpx"),		// cpx hhll

		// csh: set CPU speed to the higher speed
		0xD4:	op_noarguments("csh"),	// csh

		// csl: set CPU speed to the lower speed
		0x54:	op_noarguments("csl"),	// csl

		// cmp: compare a
		0xC9:	op_immediate("cmp"),	// cmp #nn
		0xC5:	op_zeropage("cmp"),	// cmp zz
		0xD5:	op_zeropagex("cmp"),	// cmp zz,x
		0xD2:	op_indirect("cmp"),		// cmp (zz)
		0xC1:	op_indirectx("cmp"),	// cmp (zz,x)
		0xD1:	op_indirecty("cmp"),	// cmp (zz),y
		0xCD:	op_absolute("cmp"),		// cmp hhll
		0xDD:	op_absolutex("cmp"),	// cmp hhll,x
		0xD9:	op_absolutey("cmp"),	// cmp hhll,y

		// dex: decrement x
		0xCA:	op_noarguments("dex"),	// dex

		// dec: decrement
		0xC6:	op_zeropage("dec"),		// dec zz
		0xD6:	op_zeropagex("dec"),	// dec zz,x
		0xCE:	op_absolute("dec"),		// dec hhll
		0xDE:	op_absolutex("dec"),	// dec hhll,x
		0x3A:	dec_accumulator,		// dec a

		// cpy: compare y
		0xC0:	op_immediate("cpy"),	// cpy #nn
		0xC4:	op_zeropage("cpy"),		// cpy zz
		0xCC:	op_absolute("cpy"),		// cpy hhll

		// eor: exclusive or
		0x49:	op_immediate("eor"),	// eor #nn
		0x45:	op_zeropage("eor"),		// eor zz
		0x55:	op_zeropagex("eor"),	// eor zz,x
		0x52:	op_indirect("eor"),		// eor (zz)
		0x41:	op_indirectx("eor"),		// eor (zz,x)
		0x51:	op_indirecty("eor"),		// eor (zz),y
		0x4D:	op_absolute("eor"),		// eor hhll
		0x5D:	op_absolutex("eor"),	// eor hhll,x
		0x59:	op_absolutey("eor"),		// eor hhll,y

		// inc: increment
		0xE6:	op_zeropage("inc"),		// inc zz
		0xF6:	op_zeropagex("inc"),	// inc zz,x
		0xEE:	op_absolute("inc"),		// inc hhll
		0xFE:	op_absolutex("inc"),		// inc hhll,x
		0x1A:	inc_accumulator,		// inc a

		// inx: increment x
		0xE8:	op_noarguments("inx"),	// inx

		// dey: decrement y
		0x88:	op_noarguments("dey"),	// dey

		// iny: increment y
		0xC8:	op_noarguments("iny"),	// iny

		// jmp: jump
		0x4C:	jmp_absolute,			// jmp hhll
		0x6C:	jmp_absoluteindirect,	// jmp (hhll)
		0x7C:	jmp_absolutex,		// jmp hhll,x

		// jsr: jump to subroutine
		0x20:	jsr_absolute,			// jsr hhll

		// lda: load to a
		0xA9:	lda_immediate,		// lda #nn
		0xA5:	op_zeropage("lda"),		// lda zz
		0xB5:	op_zeropagex("lda"),	// lda zz,x
		0xB2:	op_indirect("lda"),		// lda (zz)
		0xA1:	op_indirectx("lda"),		// lda (zz,x)
		0xB1:	op_indirecty("lda"),		// lda (zz),y
		0xAD:	op_absolute("lda"),		// lda hhll
		0xBD:	op_absolutex("lda"),		// lda hhll,x
		0xB9:	op_absolutey("lda"),		// lda hhll,y

		// ldx: load to x
		0xA2:	op_immediate("ldx"),	// ldx #nn
		0xA6:	op_zeropage("ldx"),		// ldx zz
		0xB6:	op_zeropagey("ldx"),	// ldx zz,y
		0xAE:	op_absolute("ldx"),		// ldx hhll
		0xBE:	op_absolutey("ldx"),		// ldx hhll,y

		// ldy: load to y
		// TODO - assuming x for the two indexed ones based on http://www6.atpages.jp/~appsouko/work/PCE/6280op.html as my original source for the opcodes say x but show y in its examples; what is correct, x or y?
		0xA0:	op_immediate("ldy"),	// ldy #nn
		0xA4:	op_zeropage("ldy"),		// ldy zz
		0xB4:	op_zeropagex("ldy"),	// ldy zz,x
		0xAC:	op_absolute("ldy"),		// ldy hhll
		0xBC:	op_absolutex("ldy"),		// ldy hhll,x

		// lsr: logical shift right
		0x46:	op_zeropage("lsr"),		// lsr zz
		0x56:	op_zeropagex("lsr"),	// lsr zz,x
		0x4E:	op_absolute("lsr"),		// lsr hhll
		0x5E:	op_absolutex("lsr"),		// lsr hhll,x
		0x4A:	op_accumulator("lsr"),	// lsr a

		// ora: bitwise or
		0x09:	op_immediate("ora"),	// ora #nn
		0x05:	op_zeropage("ora"),		// ora zz
		0x15:	op_zeropagex("ora"),	// ora zz,x
		0x12:	op_indirect("ora"),		// ora (zz)
		0x01:	op_indirectx("ora"),		// ora (zz,x)
		0x11:	op_indirecty("ora"),		// ora (zz),y
		0x0D:	op_absolute("ora"),		// ora hhll
		0x1D:	op_absolutex("ora"),	// ora hhll,x
		0x19:	op_absolutey("ora"),		// ora hhll,y

		// nop: no operation
		0xEA:	op_noarguments("nop"),	// nop

		// pha: push a
		0x48:	pha_noarguments,		// pha

		// php: push p (status register)
		0x08:	op_push("php"),		// php

		// phx: push x
		0xDA:	op_push("phx"),		// phx

		// phy: push y
		0x5A:	op_push("phy"),		// phy

		// pla: pop a
		0x68:	pla_noarguments,		// pla

		// plp: pop p
		0x28:	op_pop("plp"),			// plp

		// plx: pop x
		0xFA:	op_pop("plx"),			// plx

		// ply: pop y
		0x7A:	op_pop("ply"),			// ply

		// rol: rotate left
		0x26:	op_zeropage("rol"),		// rol zz
		0x36:	op_zeropagex("rol"),	// rol zz,x
		0x2E:	op_absolute("rol"),		// rol hhll
		0x3E:	op_absolutex("rol"),		// rol hhll,x
		0x2A:	op_accumulator("rol"),	// rol a

		// rmb: clear (reset) bit
		0x07:	op_zpbit("rmb", 0),		// rmb #0,zz
		0x17:	op_zpbit("rmb", 1),		// rmb #1,zz
		0x27:	op_zpbit("rmb", 2),		// rmb #2,zz
		0x37:	op_zpbit("rmb", 3),		// rmb #3,zz
		0x47:	op_zpbit("rmb", 4),		// rmb #4,zz
		0x57:	op_zpbit("rmb", 5),		// rmb #5,zz
		0x67:	op_zpbit("rmb", 6),		// rmb #6,zz
		0x77:	op_zpbit("rmb", 7),		// rmb #7,zz

		// ror: rotate right
		0x66:	op_zeropage("ror"),		// ror zz
		0x76:	op_zeropagex("ror"),	// ror zz,x
		0x6E:	op_absolute("ror"),		// ror hhll
		0x7E:	op_absolutex("ror"),		// ror hhll,x
		0x6A:	op_accumulator("ror"),	// ror a

		// rti: return from interrupt
		0x40:	rti_noarguments,		// rti

		// rts: return from subroutine
		0x60:	rts_noarguments,		// rts

		// sax: swap a and x
		0x22:	op_noarguments("sax"),	// sax

		// say: swap a and y
		0x42:	op_noarguments("say"),	// say

		// sbc: subtract with borrow (carry)
		0xE9:	op_immediate("sbc"),	// sbc #nn
		0xE5:	op_zeropage("sbc"),		// sbc zz
		0xF5:	op_zeropagex("sbc"),	// sbc zz,x
		0xF2:	op_indirect("sbc"),		// sbc (zz)
		0xE1:	op_indirectx("sbc"),		// sbc (zz,x)
		0xF1:	op_indirecty("sbc"),		// sbc (zz),y
		0xED:	op_absolute("sbc"),		// sbc hhll
		0xFD:	op_absolutex("sbc"),	// sbc hhll,x
		0xF9:	op_absolutey("sbc"),	// sbc hhll,y

		// sed: set decimal flag
		0xF8:	op_noarguments("sed"),	// sed

		// sec: set carry flag
		0x38:	op_noarguments("sec"),	// sec

		// sei: DISABLE interrupts (sets interrupt disable flag)
		0x78:	op_noarguments("sei"),	// sei

		// set: set T flag (changes next opcode that operates on a to operate on the zero page address pointed to by x instead)
		0xF4:	op_noarguments("set"),	// set

		// st0: store in HuC6270 address register
		0x03:	op_immediate("st0"),	// st0 #nn

		// st1: store in HuC6270 data register low
		0x13:	op_immediate("st1"),	// st1 #nn

		// st2: store in HuC6270 data register high
		0x23:	op_immediate("st2"),	// st2 #nn

		// smb: set bit
		0x87:	op_zpbit("smb", 0),		// smb #0,zz
		0x97:	op_zpbit("smb", 1),		// smb #1,zz
		0xA7:	op_zpbit("smb", 2),		// smb #2,zz
		0xB7:	op_zpbit("smb", 3),		// smb #3,zz
		0xC7:	op_zpbit("smb", 4),		// smb #4,zz
		0xD7:	op_zpbit("smb", 5),		// smb #5,zz
		0xE7:	op_zpbit("smb", 6),		// smb #6,zz
		0xF7:	op_zpbit("smb", 7),		// smb #7,zz

		// sta: store a
		0x85:	op_zeropage("sta"),		// sta zz
		0x95:	op_zeropagex("sta"),	// sta zz,x
		0x92:	op_indirect("sta"),		// sta (zz)
		0x81:	op_indirectx("sta"),		// sta (zz,x)
		0x91:	op_indirecty("sta"),		// sta (zz),y
		0x8D:	op_absolute("sta"),		// sta hhll
		0x9D:	op_absolutex("sta"),		// sta hhll,x
		0x99:	op_absolutey("sta"),		// sta hhll,y

		// stx: store x
		0x86:	op_zeropage("stx"),		// stx zz
		0x96:	op_zeropagey("stx"),	// stx zz,y
		0x8E:	op_absolute("stx"),		// stx hhll

		// sty: store y
		0x84:	op_zeropage("sty"),		// sty zz
		0x94:	op_zeropagex("sty"),	// sty zz,x
		0x8C:	op_absolute("sty"),		// sty hhll

		// stz: store zero
		0x64:	op_zeropage("stz"),		// stz zz
		0x74:	op_zeropagex("stz"),	// stz zz,x
		0x9C:	op_absolute("stz"),		// stz hhll
		0x9E:	op_absolutex("stz"),		// stz hhll,x

		// tai: fill destination buffer with source word
		0xF3:	op_transfer("tai"),		// tai hhll,hhll,hhll

		// sxy: swap x and y
		0x02:	op_noarguments("sxy"),	// sxy

		// tam: transfer a to memory page register(s)
		0x53:	tam_pageregs,			// tam #nn,...

		// tax: transfer a to x
		0xAA:	op_noarguments("tax"),	// tax

		// tay: transfer a to y
		0xA8:	op_noarguments("tay"),	// tay

		// tia: copy source buffer to destination word (for instance, if destination word is a memory-mapped data port)
		0xE3:	op_transfer("tia"),		// tia hhll,hhll,hhll

		// tdd: copy source buffer GOING DOWN to destination buffer GOING DOWN
		0xC3:	op_transfer("tdd"),		// tdd hhll,hhll,hhll

		// tin: copy source buffer to destination byte
		0xD3:	op_transfer("tin"),		// tin hhll,hhll,hhll

		// tii: copy source buffer to destination buffer
		0x73:	op_transfer("tii"),		// tii hhll,hhll,hhll

		// tma: transfer memory page register to a
		0x43:	tma_pageregs,			// tma #nn

		// trb: test value against bits set in a, then clear (reset) those bits in the value
		0x14:	op_zeropage("trb"),		// trb zz
		0x1C:	op_absolute("trb"),		// trb hhll

		// tsb: test value against bits set in a, then sets those bits in the value
		0x04:	op_zeropage("tsb"),		// tsb zz
		0x0C:	op_absolute("tsb"),		// tsb hhll

		// tst: test bits (and also clears them? I don't quite understand what's going on here)
		0x83:	tst_zeropage,			// tst #nn,zz
		0xA3:	tst_zeropagex,			// tst #nn,zz,x
		0x93:	tst_absolute,			// tst #nn,hhll
		0xB3:	tst_absolutex,			// tst #nn,hhll,x

		// tsx: transfer s to x
		0xBA:	op_noarguments("tsx"),	// tsx

		// txa: transfer x to a
		0x8A:	op_noarguments("txa"),	// txa

		// tya: transfer y to a
		0x98:	op_noarguments("tya"),	// tya

		// txs: transfer x to s
		0x9A:	op_noarguments("txs"),	// txs
	}
}
