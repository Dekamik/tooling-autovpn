package main

import (
    "errors"
    "reflect"
)

func asInterfaceSlice(origin []string) []interface{} {
    s := make([]interface{}, len(origin))
    for i, v := range origin {
        s[i] = v
    }
    return s
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
    panic(errors.New("Not a slice "))
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
