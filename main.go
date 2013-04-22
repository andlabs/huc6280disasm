// 5 april 2013
// from tms34010disasm - 17 march 2013
package main

import (
	"fmt"
	"os"
	"flag"
	"io/ioutil"
)

var bytes []byte
var instructions map[uint32]string
var labels map[uint32]string
var comments map[uint32]string

var vectorLocs = map[uint32]string{
	0x1FFE:	"EntryPoint",
	0x1FFC:	"NMI",
	0x1FFA:	"TimerInterrupt",
	0x1FF8:	"IRQ1",
	0x1FF6:	"IRQ2_BRK",
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s ROM", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var err error

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

	instructions = map[uint32]string{}
	labels = map[uint32]string{}
	comments = map[uint32]string{}

	// autoanalyze vectors
	for addr, label := range vectorLocs {
		posw, _ := getword(addr)
		pos, err := physical(posw)
		if err != nil {
			errorf("internal error: could not get physical address for %s vector (meaning something is up with the paging or the game actually does have the vector outside page 7): %v\n", label, err)
		}
		if labels[pos] != "" {		// if already defined as a different vector, concatenate the labels to make sure everything is represented
			// TODO because this uses a map, it will not be in vector order
			labels[pos] = labels[pos] + "_" + label
		} else {
			labels[pos] = label
		}
		queueDisassemble(pos)
	}
	disassemble()

	// TODO read additional starts from standard input

	print()
}
