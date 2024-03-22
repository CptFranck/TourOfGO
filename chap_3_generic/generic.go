package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	genericFunction()
	
	genericType()
	
	addExample()
	
	mapValuesExample()
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

type test struct{
	String string `default:"pouet"`; 
	StringBis string `default:"pouet"`; 
}

func genericFunction() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	sst := []test{
		{"foo","foo"},
		{"bar","foo"},
		{"baz","foo"},
	}
	fmt.Println(Index(sst, test{"baz","foo"}))
}

////////////////////////////////////////////////

type List[T any] struct {
	next *List[T]
	val  T
}

func findLinkListIndexOf[T comparable](v T, list *List[T]) int{
	currentNode := list
	i := 0
	for currentNode.val != v{
		if currentNode.next != nil {
			return -1
		}
		currentNode = currentNode.next
		i++
	}
	return i
}

func (list *List[T]) addToList(v T) {
	currentNode := list
	for currentNode.next != nil{
		currentNode = currentNode.next
	}
	currentNode.next = &List[T]{}
	currentNode.val = v
}


func (list *List[T]) addToListAt(v T, i int) {
	currentNode := list
	for j := 0; j != (i-1) && j < (i-1); j++ {
		currentNode = currentNode.next
	}
	fmt.Printf("node : %v \n", currentNode)
	currentNode.next = &List[T]{val: v, next: currentNode.next}
}

func deleteInList[T comparable](v T, list *List[T]) {
	currentNode := list
	for currentNode.next != nil && currentNode.next.val != v{
		currentNode = currentNode.next
	}
	fmt.Printf("node : %v \n", currentNode)
	currentNode.next = currentNode.next.next
}

func deleteInListAt[T comparable](i int, list *List[T]) {
	currentNode := list
	for j := 0; j != (i-1) && j < (i-1); j++ {
		currentNode = currentNode.next
	}
	fmt.Printf("node : %v \n", currentNode)
	currentNode.next = currentNode.next.next	
}

func (list *List[T]) printList() {
	currentNode := list
	fmt.Printf("List : \n", )
	for currentNode.next != nil{
		fmt.Printf("%v, ", currentNode.val)
		currentNode = currentNode.next
	}
	fmt.Printf("\n\n")
}

func genericType()  {
	fmt.Printf("\n")
	intList := List[int]{}

	for i := 0; i < 10; i++ {
		intList.addToList(i)
	}
	intList.printList()
	deleteInList(5, &intList)
	intList.printList()
	intList.addToListAt(10, 5)
	intList.printList()
	fmt.Printf("index of 10 should be 5 and it find : %v", findLinkListIndexOf(10, &intList))
	deleteInListAt(8, &intList)
	intList.printList()
}
////////////////////////////////////////////////

// type Num interface {
// 	int8 | int16 | int32 | int64 | float32 | float64
// }

// func Add[T int | float](a T, b T) T {
// func Add[T Num](a T, b T) T {
func Add[T constraints.Ordered](a T, b T) T {
	return a + b
}

type userID int
func AddUserID[T ~int](a T, b T) T {
	return a + b
}

func addExample()  {
	fmt.Println(Add(1,2.3))

	a:= userID(1)
	b:= userID(2)
	// works due to "~" marker
	fmt.Println(AddUserID(a,b))
}

//////////////

func MapValues[T constraints.Ordered](values []T, mapFunc func(t T) T )[]T{
	var newValues []T
	for _, v := range values{
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues 
}

func mapValuesExample()  {
	result := MapValues([]int{1,2,3}, func(a int) int{return a * 2})
	println(result)
}
