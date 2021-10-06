package main

import "fmt"

func verbose(a ...interface{}) {
    if options.Verbose {
        fmt.Print(a...)
    }
}

func verbosef(format string, a ...interface{}) {
    if options.Verbose {
        fmt.Printf(format, a...)
    }
}

func verboseln(a ...interface{}) {
    if options.Verbose {
        fmt.Println(a...)
    }
}

func printTable(m [][]string) {
    var format = ""
    for col := 0; col < len(m[0]); col++ {
        colValues := Map(m, func(item interface{}) interface {} { return item.([]string)[col] })
        format += fmt.Sprintf("%%-%ds ", maxlen(colValues))
    }
    format += "\n"
    for _, cols := range m {
        fmt.Printf(format, asInterfaceSlice(cols)...)
    }
}
