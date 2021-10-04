package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
