package main

import (
	"errors"
	"fmt"
	"reflect"
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



func printMatrix(m [][]string) {
	var format = ""
	for col := 0; col < len(m[0]); col++ {
		format += fmt.Sprintf("%%-%ds ", maxlen(Map(m, func(item interface{}) interface {} { return item.([]string)[col] })))
	}
	format += "\n"
	for _, cols := range m {
		fmt.Printf(format, cols...)
	}
}

func stringMatrix(t [][]interface{}) [][]string {
	if reflect.TypeOf(t).Kind() == reflect.Slice {
		s := reflect.ValueOf(t)
		matrix := make([][]string, s.Len())
		for i := 0; i < s.Len(); i++ {
			o := s.Index(i).Elem()
			arr := make([]string, o.Len())
			for j := 0; j < o.Len(); j++ {
				arr[j] = o.Index(j).String()
			}
			matrix[i] = arr
		}
		return matrix
	}
	return nil
}

func Map(t interface{}, f func(interface{}) interface{}) []interface{} {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		arr := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			arr[i] = f(s.Index(i).Interface())
		}
		return arr
	}
	return nil
}

func maxlen(t interface{}) int {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		var max = 0
		for i := 0; i < s.Len(); i++ {
			length := s.Index(i).Elem().Len()
			if length > max {
				max = length
			}
		}
		return max
	}
	panic(errors.New("Not a slice "))
}
