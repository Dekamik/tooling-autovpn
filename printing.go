package main

import "fmt"

func printTable(matrix [][]string) {
    var format = ""
    for col := 0; col < len(matrix[0]); col++ {
        colValues := Map(matrix, func(item interface{}) interface {} { return item.([]string)[col] })
        format += fmt.Sprintf("%%-%ds ", maxlen(colValues))
    }
    format += "\n"
    for _, cols := range matrix {
        fmt.Printf(format, asInterfaceSlice(cols)...)
    }
}
