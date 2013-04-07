// 5 april 2013
package main

// TODO adjust immediates so that they have effect on a

var opcodetbl [0x100]xxxxx = {
	// adc: add with carry
	0x69:	op_immediate("adc")	// adc #nn
	0x65:	op_zeropage("adc")		// adc zz
	0x75:	op_zeropagex("adc")	// adc zz,x
	0x72:	op_indirect("adc")		// adc (zz)
	0x61:	op_indirectx("adc")		// adc (zz,x)
	0x71:	op_indirecty("adc")		// adc (zz),y
	0x6D:	op_absolute("adc")		// adc hhll
	0x7D:	op_absolutex("adc")		// adc hhll,x
	0x79:	op_absolutey("adc")		// adc hhll,y

	// and: bitwise and
	0x29:	op_immediate("and")	// and #nn
	0x25:	op_zeropage("and")		// and zz
	0x35:	op_zeropagex("and")	// and zz,x
	0x32:	op_indirect("and")		// and (zz)
	0x21:	op_indirectx("and")		// and (zz,x)
	0x31:	op_indirecty("and")		// and (zz),y
	0x2D:	op_absolute("and")		// and hhll
	0x3D:	op_absolutex("and")		// and hhll,x
	0x39:	op_absolutey("and")		// and hhll,y

	// asl: arithmetic shift left
	0x06:	op_zeropage("asl")		// asl zz
	0x16:	op_zeropagex("asl")		// asl zz,x
	0x0E:	op_absolute("asl")		// asl hhll
	0x1E:	op_absolutex("asl")		// asl hhll,x
	0x0A:	op_accumulator("asl")	// asl a

	// bbr: branch on bit clear (reset)
	0x0F:	op_zpbitbr("bbr", 0)		// bbr #0,zz,hhll
	0x1F:	op_zpbitbr("bbr", 1)		// bbr #1,zz,hhll
	0x2F:	op_zpbitbr("bbr", 2)		// bbr #2,zz,hhll
	0x3F:	op_zpbitbr("bbr", 3)		// bbr #3,zz,hhll
	0x4F:	op_zpbitbr("bbr", 4)		// bbr #4,zz,hhll
	0x5F:	op_zpbitbr("bbr", 5)		// bbr #5,zz,hhll
	0x6F:	op_zpbitbr("bbr", 6)		// bbr #6,zz,hhll
	0x7F:	op_zpbitbr("bbr", 7)		// bbr #7,zz,hhll

	// bcc: branch on carry clear
	0x90:	op_branch("bcc")		// bcc hhll

	// bbs: branch on bit set
	0x8F:	op_zpbitbr("bbs", 0)		// bbs #0,zz,hhll
	0x9F:	op_zpbitbr("bbs", 1)		// bbs #1,zz,hhll
	0xAF:	op_zpbitbr("bbs", 2)		// bbs #2,zz,hhll
	0xBF:	op_zpbitbr("bbs", 3)		// bbs #3,zz,hhll
	0xCF:	op_zpbitbr("bbs", 4)		// bbs #4,zz,hhll
	0xDF:	op_zpbitbr("bbs", 5)		// bbs #5,zz,hhll
	0xEF:	op_zpbitbr("bbs", 6)		// bbs #6,zz,hhll
	0xFF:	op_zpbitbr("bbs", 7)		// bbs #7,zz,hhll

	// bcs: branch on carry set
	0xB0:	op_branch("bcs")		// bcs hhll

	// beq: branch on equal
	0xF0:	op_branch("beq")		// beq hhll

	// bit: test bit of accumulator
	0x89:	op_immediate("bit")		// bit #nn
	0x24:	op_zeropage("bit")		// bit zz
	0x34:	op_zeropagex("bit")		// bit zz,x
	0x2C:	op_absolute("bit")		// bit hhll
	0x3C:	op_absolutex("bit")		// bit hhll,x

	// bmi: branch on minus
	0x30:	op_branch("bmi")		// bmi hhll

	// bne: branch on not equal
	0xD0:	op_branch("bne")		// bne hhll

