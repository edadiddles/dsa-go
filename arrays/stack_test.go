package arrays

import (
    "testing"
)

func PushTest(t *testing.T) {
    var s Stack[int]

    s = Push(s, 1)

    if s.Vals[0] != 1 {
        t.Fatalf("val expected 1 received %v", s.Vals[0])
    }
        
}
