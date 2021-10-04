package main

import (
	"fmt"
)

func verbose(str string) {
	if options.Verbose {
		fmt.Print(str)
	}
}

func verboseln(str string) {
	if options.Verbose {
		fmt.Println(str)
	}
}
