package arrays

import (
)

type Stack[T any] struct {
    Vals []T
}

//TODO: refactor! this is super hacky
func Push[T any](s Stack[T], val T) Stack[T] {
    tVal := []T{val}
    s.Vals = append(tVal, s.Vals...)
    return s
}


