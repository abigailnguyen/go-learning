package main
import (
	"math"
	"time"
	"fmt"
	"os"
	// "sync"
)
/*
* 	O(n) time and O(n) space
*/

type Node struct {
	value	int64
	left	*Node
	right	*Node
}
// min  max
// 	\   /
//	  5
//  /   \
//  4   7
func isValid(node *Node, low, high int64) bool {
	if node == nil { return true }
	if node.value > low && node.value < high {
		return true && 
			isValid(node.left, low, node.value) && 
			isValid(node.right, node.value, high)
	}
	return false
}

func checkValidBST(n *Node) bool {
	// entry to the loop, the left's low is always math.Min, and the right's high is always math.Max
	return isValid(n, math.MinInt64, math.MaxInt64)  
}

func checkValidBSTConcurrent(n *Node) bool {
	// entry to the loop, the left's low is always math.Min, and the right's high is always math.Max
	failCheck 	:= make(chan bool)
	timeoutCh 	:= time.After(20 * time.Second)
	done := make(chan bool)

	go isValidConcurrent(n, math.MinInt64, math.MaxInt64, failCheck, done)

	select {
	case <-timeoutCh:
		fmt.Println("took too long to complete the task")
		return false
 	case <-failCheck:
		fmt.Println("Is BST: False")
		return false
	case <-done:
		fmt.Println("Is BST: True")
		return true
	}
}
// out = the same channel passing down
// done = the same channel passing up
func isValidConcurrent(node *Node, low, high int64, out chan<- bool, done chan<- bool) {	
	if node == nil { 
		close(done) // tip of tree
		return 
	}
	if node.value < low ||  node.value > high {
		out <- false // send channel, immediately terminate program when a false has been seen
	}

	leftCh := make(chan bool)
	rightCh := make(chan bool)

	go func() {
		isValidConcurrent(node.left, low, node.value, out, leftCh)
	}()

	go func() {
		isValidConcurrent(node.right, node.value, high, out, rightCh)
	}()
	<-leftCh // wait for left comparison to complete
	<-rightCh // wait for right comparison to complete
	close(done)
}

func main() {
	defer fmt.Println("main() terminated.")
	root := new(Node)
	root.value = 5
	root.left = new(Node)
	root.left.value = 4
	root.right = new(Node)
	root.right.value = 7
	root.left.left = new(Node)
	root.left.left.value = 2
	root.right.left = new(Node)
	root.right.left.value = 6
	checkValidBSTConcurrent(root)
}

func writeToFile(f *os.File, s string) {
	_, err := f.WriteString(s + "\n")
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
