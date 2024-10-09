package arrays

type LinkedNode[T any] struct {
    NextNode *LinkedNode[T]
    PreviousNode *LinkedNode[T]
    Val T
}

func Insert[T any](n LinkedNode[T], val T) LinkedNode[T] {
    node := LinkedNode[T] {
        NextNode: nil,
        PreviousNode: nil,
        Val: val,
    }
    for n.NextNode != nil {}
    n.NextNode = &node

    return n
}




