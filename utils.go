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
		fmt.Printf(format, asInterfaceSlice(cols)...)
	}
}

func asStringSlice(o []interface{}) []string {
	s := make([]string, len(o))
	for i, v := range o {
		s[i] = fmt.Sprint(v)
	}
	return s
}

func asStringMatrix(o [][]interface{}) [][]string {
	m := make([][]string, len(o))
	for i, a := range o {
		m[i] = make([]string, len(a))
		for j, b := range a {
			m[i][j] = fmt.Sprint(b)
		}
	}
	return m
}

func asInterfaceSlice(o []string) []interface{} {
	s := make([]interface{}, len(o))
	for i, v := range o {
		s[i] = v
	}
	return s
}

func asInterfaceMatrix(o [][]string) [][]interface{} {
	m := make([][]interface{}, len(o))
	for i, a := range o {
		m[i] = make([]interface{}, len(a))
		for j, b := range a {
			m[i][j] = fmt.Sprint(b)
		}
	}
	return m
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
