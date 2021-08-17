package add2NumbersAsALinkedList

import (
	"fmt"
	"testing"
)

func Test2numberlist(t *testing.T) {
	l1 := new(Node)
	l2 := new(Node)
	l3 := new(Node)
	l1.value = 1
	l1.next = l2
	l2.value = 2
	l2.next = l3
	l3.value = 5

	l4 := &Node{
		2,
		&Node{
			3,
			&Node{
				value: 9,
			},
		},
	}

	result := add2LinkedList(l1, l4)
	actual := &Node{
		3,
		&Node{
			5,
			&Node{
				4,
				&Node{
					value: 1,
				},
			},
		},
	}

	if compareLinkedListIteratively(result, actual) {
		fmt.Println("equals")
	} else {
		fmt.Printf("Expected %v, Got %v\n", *actual, *result)
		t.Fail()
	}
}

func Test2numberlist2(t *testing.T) {
	l1 := new(Node)
	l2 := new(Node)
	l1.value = 6
	l1.next = l2
	l2.value = 5

	l4 := &Node{
		4,
		&Node{
			3,
			&Node{
				value: 9,
			},
		},
	}

	result := add2LinkedList(l1, l4)
	actual := &Node{
		0,
		&Node{
			9,
			&Node{
				value: 9,
			},
		},
	}

	if compareLinkedListIteratively(result, actual) {
		fmt.Println("equals")
	} else {
		fmt.Printf("Expected %v, Got %v\n", *actual, *result)
		t.Fail()
	}
}

func compareLinkedList(l1, l2 *Node) bool {
	return l1.value == l2.value &&
		compareLinkedList(l1.next, l2.next)
}

func compareLinkedListIteratively(l1, l2 *Node) (result bool) {
	for l1 != nil && l2 != nil {
		if result = l1.value == l2.value; !result {
			fmt.Printf("l1: %v, l2: %v\n", l1, l2)
			return
		}
		fmt.Printf("l1: %v, l2: %v\n", l1, l2)

		l1, l2 = l1.next, l2.next
	}
	fmt.Printf("l1: %v, l2: %v\n", l1, l2)
	result = l1 == l2 // nil comparison is acceptable for the same type, but not for different types
	return
}

func TestNil(t *testing.T) {
	l1 := new(Node)
	l2 := new(Node)
	// expect 2 zero struct to be not equal
	fmt.Printf("l1 %v , l2 %v\n", l1, l2)
	if l1 == l2 {
		t.Fail()
	}
}
