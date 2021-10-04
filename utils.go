package main

import (
	"fmt"
)

func verbose(str string) {
	if config.Verbose {
		fmt.Print(str)
	}
}

func verboseln(str string) {
	if config.Verbose {
		fmt.Println(str)
	}
}
