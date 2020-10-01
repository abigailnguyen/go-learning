package main
import (
	"fmt"
	"testing"
	// "os"
	// "time"
	// "sync"
)
//   5
// /   \
// 4   7  
// true
func TestValidBST_1(t *testing.T) {
	root := new(Node)
	root.value = 5
	root.left = new(Node)
	root.left.value = 4
	root.right = new(Node)
	root.right.value = 7
	result := checkValidBST(root) 
	fmt.Printf("Is a Binary Search Tree: %v\n", result)
	if !result {
		t.Fail()
	} 
}

func TestValidBST_1_Concurrent(t *testing.T) {
	root := new(Node)
	root.value = 5
	root.left = new(Node)
	root.left.value = 4
	root.right = new(Node)
	root.right.value = 7
	result := checkValidBSTConcurrent(root) 
	fmt.Printf("Is a Binary Search Tree: %v\n", result)
	if !result {
		t.Fail()
	} 
}

//   	  5
// 		/   \
// 	   4     7
//	  /  \ /  \
//	 2     6      
//  true

func TestValidBST_2(t *testing.T) {
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
	result := checkValidBST(root)
	fmt.Printf("Is a Binary Search Tree: %v\n", result)

	if !result {
		t.Fail()
	} 
}

func TestValidBST_2_Concurrent(t *testing.T) {
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
	result := checkValidBSTConcurrent(root)
	fmt.Printf("Is a Binary Search Tree: %v\n", result)

	if !result {
		t.Fail()
	} 
}

//   	  5
// 		/   \
// 	   4     7
//	       /  \
//	       2
// false
func TestInvalidBST_1(t *testing.T) {
	root := new(Node)
	root.value = 5
	root.left = new(Node)
	root.left.value = 4
	root.right = new(Node)
	root.right.value = 7
	root.right.left = new(Node)
	root.right.left.value = 2
	result := checkValidBST(root)
	fmt.Printf("Is a Binary Search Tree: %v\n", result)

	if result {
		t.Fail()
	} 
}

func TestInvalidBST_1_Concurrent(t *testing.T) {
	root := new(Node)
	root.value = 5
	root.left = new(Node)
	root.left.value = 4
	root.right = new(Node)
	root.right.value = 7
	root.right.left = new(Node)
	root.right.left.value = 2
	result := checkValidBSTConcurrent(root)
	fmt.Printf("Is a Binary Search Tree: %v\n", result)

	if result {
		t.Fail()
	} 
}

//   	  5
// 		/   \
// 	   4     7
//	 /  \
//	     8   
// false
func TestInvalidBST_2(t *testing.T) {
	root := new(Node)
	root.value = 5
	root.left = new(Node)
	root.left.value = 4
	root.right = new(Node)
	root.right.value = 7
	root.left.right = new(Node) //false
	root.left.right.value = 8
	result := checkValidBST(root)
	fmt.Printf("Is a Binary Search Tree: %v\n", result)

	if result {
		t.Fail()
	} 
}


func TestInvalidBST_2_Concurrent(t *testing.T) {
	root := new(Node)
	root.value = 5
	root.left = new(Node)
	root.left.value = 4
	root.right = new(Node)
	root.right.value = 7
	root.left.right = new(Node) //false
	root.left.right.value = 8
	result := checkValidBSTConcurrent(root)
	fmt.Printf("Is a Binary Search Tree: %v\n", result)

	if result {
		t.Fail()
	} 
}