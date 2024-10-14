package arrays

import (
)

type Stack[T any] struct {
    Vals []T
}

//TODO: refactor! this is super hacky
func Push[T any](s Stack[T], val T) {
    tVal := []T{val}
    s.Vals = append(tVal, s.Vals...)
}

func Top[T any](s Stack[T]) T {
    return s.Vals[0]
}

func Pop[T any](s Stack[T]) T {
    val := s.Vals[0]
    s.Vals = s.Vals[1:len(s.Vals)-1]
    return val
}
