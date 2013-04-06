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

