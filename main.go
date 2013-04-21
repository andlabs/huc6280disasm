// 5 april 2013
// from tms34010disasm - 17 march 2013
package main

import (
	"fmt"
	"os"
	"flag"
	"io/ioutil"
	bytespkg "bytes"
	"encoding/binary"
)

var bytes []byte
var words []uint16
var instructions map[uint32]string
var labels map[uint32]string

var vectorLocs = map[uint32]string{
	0x1FFE:	"EntryPoint",
	0x1FFC:	"NMI",
	0x1FFA:	"TimerInterrupt",
	0x1FF8:	"IRQ1",
	0x1FF6:	"IRQ2_BRK",
}

func errorf(format string, args ...interface{}) {
	fmg.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s ROM", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}

	filename := flag.Arg(0)

	bytes, err = ioutil.ReadFile(filename)
	if err != nil {
		errorf("error reading input file %s: %v", filename, err)
	}
	if len(bytes) < 0x2000 {
		errorf("given input file %s does not provide a complete interrupt vector table (this restriction may be lifted in the future)", filename)
	}
	if len(bytes) >= 0x1F0000 {
		errorf("given input file %s too large (this restriction may be lifted in the future)", filename)
	}
	words = make([]uint16, len(bytes) / 2)
	err = binary.Read(bytespkg.NewReader(b), binary.LittleEndian, &words)
	if err != nil {
		errorf("error building words array from input byte stream: %v", err)
	}
	// TODO do we just assume the slice was filled properly?

	instructions = map[uint32]string{}
	labels = map[uint32]string{}

	// autoanalyze vectors
	for addr, label := range vectorLocs {
		if labels[addr] != "" {		// if already defined as a different vector, concatenate the labels to make sure everything is represented
			// TODO because this uses a map, it will not be in vector order
			labels[addr] = labels[addr] + "_" + label
		} else {
			labels[addr] = label
		}
		addr /= 2		// words
		pos := uint32(words[addr])
		// TODO handle bad addresses here or in disassemble()?
		disassemble(pos)
	}

	// TODO read additional starts from standard input
	// TODO print
}
