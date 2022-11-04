package ano

import "fmt"

type node[T any] struct {
	value T
	left  *node[T]
	right *node[T]
}

func (n *node[T]) Add(Value T, Compare func(A, B T) bool) {
	dir := Compare(n.value, Value)
	if !dir {
		if n.left == nil {
			n.left = &node[T]{value: Value}
			return
		}
		n.left.Add(Value, Compare)
	}
	if dir {
		if n.right == nil {
			n.right = &node[T]{value: Value}
			return
		}
		n.right.Add(Value, Compare)
	}
}

func (n *node[T]) GetList() []T {
	list := []T{}
	if n.left != nil {
		l := n.left.GetList()
		list = append(list, l...)
	}
	list = append(list, n.value)
	if n.right != nil {
		l := n.right.GetList()
		list = append(list, l...)
	}
	return list
}

func (n *node[T]) Print() {
	if n.left != nil {
		n.left.Print()
	}
	fmt.Printf("%+v\n", n)
	if n.right != nil {
		n.right.Print()
	}
}
