package main

import "fmt"

// integer Node for Simplicity 
type Tree interface {
	search(val int) Node
	insert(val int)
	traverse(r chan int)
}

type TreeImpl struct {
	wurzel Node
}

type Node struct {
	value int
	left  *Node
	rigth *Node
}

func (t TreeImpl) search(val int) Node {
	act := t.wurzel
	done := false

	for !done {
		if act.value == val {
			done = true
			break
		}

		if act.value > val {
			if act.left == nil {
				done = true
				break
			}
			act = *act.left
		} else {
			if act.rigth == nil {
				done = true
				break
			}
			act = *act.rigth
		}

	}

	return act

}

func (t TreeImpl) insert(val int) {
	act := t.wurzel
	done := false

	for !done {

		if act.value > val {
			if act.left == nil {
				tmp := new(Node)
				tmp.value = val
				act.left = tmp
				done = true

				break
			}
			act = *act.left
		} else {
			if act.rigth == nil {
				tmp := new(Node)
				tmp.value = val
				act.rigth = tmp

				done = true
				break
			}
			act = *act.rigth
		}

	}

}

func main() {
	t := new(TreeImpl)
	t.insert(5)
	t.insert(6)
	t.insert(4)
	fmt.Println(t)

	fmt.Println("Hello, 世界")
}
