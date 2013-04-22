// 21 april 2013
package main

import (
	"fmt"
)

func print() {
	lbu32 := uint32(len(bytes))
	for i := uint32(0); i < lbu32; i++ {
		if label, ok := labels[i]; ok {
			fmt.Printf("%s:\n", label)
		}
		if instruction, ok := instructions[i]; ok && instruction != operandString {
			fmt.Printf("\t%s\t\t; $%X\n", instruction, i)
		}
	}
}
