package arrays

import (
    "testing"
)

type BasicStruct struct {
    a int
}


func TestInsertInt(t *testing.T) {
    node := LinkedNode[int] {
        NextNode: nil,
        PreviousNode: nil,
        Val: 0,
    }

    node = Insert(node, 1)

    n := node.NextNode

    if n.Val != 1 {
        t.Fatalf("val expected 1 received %v", node.Val)
    }
}


func TestInsertString(t *testing.T) {
    node := LinkedNode[string] {
        NextNode: nil,
        PreviousNode: nil,
        Val: "",
    }

    node = Insert(node, "Hello test")

    n := node.NextNode

    if n.Val != "Hello test" {
        t.Fatalf("val expected 'Hello test' received '%v'", node.Val)
    }
}

func TestInsertFloat32(t *testing.T) {
    node := LinkedNode[float32] {
        NextNode: nil,
        PreviousNode: nil,
        Val: 0.0,
    }

    node = Insert(node, 3.14)

    n := node.NextNode

    if n.Val != 3.14 {
        t.Fatalf("val expected 3.14 received '%v'", node.Val)
    }
}

func TestInsertStuct(t *testing.T) {
    tStruct := BasicStruct{
        a: 1,
    }
    node := LinkedNode[BasicStruct] {
        NextNode: nil,
        PreviousNode: nil,
        Val: BasicStruct{a:0},
    }

    node = Insert(node, tStruct)

    n := node.NextNode

    if n.Val.a != 1 {
        t.Fatalf("val expected to be struct with a parameter has value 1 received '%v'", node.Val.a)
    }
}

