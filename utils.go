package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
