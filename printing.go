package main

import "fmt"

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
