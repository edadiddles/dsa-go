package arrays

import (
    "testing"
)

func PushTest(t *testing.T) {
    var s Stack[int]

    Push(s, 1)
    if s.Vals[0] != 1 {
        t.Fatalf("val expected 1 received %v", s.Vals[0])
    }       
}


func TopTest(t *testing.T) {
    var s Stack[int]
    Push(s, 2)
    Push(s, 3)
    Push(s, 4)
    Push(s, 5)

    v := Top(s)

    if v != 2 {
        t.Fatalf("val expected 2 received %v", v)
    }
}

func PopTest(t *testing.T) {
    var s Stack[int]
    Push(s, 2)
    Push(s, 3)
    Push(s, 4)
    Push(s, 5)

    v := Pop(s)

    if v != 2 {
        t.Fatalf("val expected 2 received %v", v)
    } else if s.Vals[0] != 3 {
        t.Fatalf("val expected 3 received %v", s.Vals[0])
    }
}
