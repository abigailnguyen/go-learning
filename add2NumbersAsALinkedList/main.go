package add2NumbersAsALinkedList

import ("fmt")
/*
*	Specify the recursive node by using pointer *Node
* 	Declare next as Node type will yield an error of invalid recursive type
*/
type Node struct {
	value 	int
	next 	*Node 
}

func add2LinkedList(l1, l2 *Node) *Node {
	return add2NumbersIteratively(l1, l2)
}

func add2NumbersIteratively(l1, l2 *Node) *Node {
	var (
		resultNode = new(Node)
	)
	for node1, node2, node3, prevNode, c := l1, l2, resultNode, new(Node), 0;
		node1 != nil || node2 != nil || c > 0; {
		if (node3 == nil) {        
			node3 = prevNode.next
			// only initialize a new node if there are values to hold,
			// to avoid initialize a zero value struct for a empty node. it will be set to nil.
			node3.next = new(Node)
			node3 = node3.next
		}
		val := c
		if node1 != nil {
			val += node1.value
			node1 = node1.next
		}
		if node2 != nil {
			val += node2.value
			node2 = node2.next
		}
		c = int(val / 10)
		node3.value = val % 10
		// Go does not allow a variable to be declared to assign the pointer memory only.
		// needs to use one of its member to store value of node3 so that error that the variable is not considered unused.
		prevNode.next = node3   
		// if node3.next is nil, and we set node3 to nil, we loose the link to memory pointer to the previous node.
		// so we need to save the memory pointer to node3 in prevNode.next
		node3 = node3.next 
	}
	return resultNode
}

func add2NumbersRecursively(l1 *Node, l2 *Node, c int) *Node {
	var (
		nextnode1, nextnode2 *Node
		val = c
		resultNode = new(Node)
	)
	if l1 != nil {
		val += l1.value
		nextnode1 = l1.next
	} 
	if l2 != nil {
		val += l2.value
		nextnode2 = l2.next
	}
	c = int(val / 10)
	fmt.Printf("carried over is %v", c)
	resultNode.value = val % 10
	
	if nextnode1 != nil || nextnode2 != nil || c > 0 {
		resultNode.next = add2NumbersRecursively(nextnode1, nextnode2, c)
	}
	return resultNode 
}