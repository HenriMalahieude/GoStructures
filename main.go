package main

import (
	"fmt"

	"github.com/HenriMalahieude/GoStructures/queue"
	"github.com/HenriMalahieude/GoStructures/stack"
	"github.com/HenriMalahieude/GoStructures/tree"
)

func catch(origin string){
	if r := recover(); r != nil{
		fmt.Println("Recovered an error in " + origin + ":", r)
	}	
}

func doCheck(tNum *int, out string, ff func() bool){
	if ff() {
		(*tNum)++;
	}else{
		fmt.Println(out)
	}
}

func testBST() {
	defer catch("BST")

	var total *int = new(int);
	(*total) = 0;

	tbst1 := tree.NewBinarySearchTree(func(a, b int) bool{
		return a < b
	}, func(a, b int) bool{
		return a == b
	})

	tbst1.Insert(5)

	doCheck(total, "Insert/Search (Case1): Not Passed, root case", func() bool{
		return tbst1.Search(5)
	})

	tbst1.Insert(3)
	tbst1.Insert(7)

	doCheck(total, "Insert/Search (Case2): Not Passed, leaf case", func() bool{
		return tbst1.Search(3) && tbst1.Search(7)
	})

	tbst1.Insert(2)
	tbst1.Insert(4)
	tbst1.Insert(6)
	tbst1.Insert(8)

	doCheck(total, "Insert/Search (Case3): Not Passed, internal case", func() bool{
		return tbst1.Search(3) && tbst1.Search(7)
	})

	doCheck(total, "Insert/Search (Case4): Not Passed, leaf trip case", func() bool{
		return tbst1.Search(2) && tbst1.Search(6)
	})
	

	tbst1.Remove(2)

	doCheck(total, "Remove/Search (Case1): Not Passed, leaf case", func() bool{
		return !tbst1.Search(2)
	})

	tbst1.Remove(7)
	
	doCheck(total, "Remove/Search (Case2): Not Passed, internal case", func() bool{
		return !tbst1.Search(7)
	})

	tbst1.Remove(5)
	
	doCheck(total, "Remove/Search (Case3): Not Passed, root case", func() bool{
		return !tbst1.Search(5)
	})

	//There may need more tests, but this is comprehensive enough atm
	fmt.Print("Test Passed on BST: ", *total, " / 7\n\n")
}

func testStack() {
	defer catch("Stack")

	var total int = 0;

	tstack1 := stack.NewLinkedStack[int]()

	if !tstack1.Empty(){
		fmt.Println("Empty (Case1): Not Passed")
	}else{
		total++
	}

	tstack1.Push(1)

	if tstack1.Peak() != 1 {
		fmt.Println("Push/Peak (Case1): Not Passed, did not push/peak correctly")
	}else{
		total++
	}

	if tstack1.Empty() {
		fmt.Println("Empty (Case2): Not passed")
	}else{
		total++
	}

	tstack1.Push(2)
	tstack1.Push(3)

	if tstack1.Peak() != 3 {
		fmt.Println("Push/Peak (Case2): Not Passed, did not push/peak correctly")
	}else{
		total++
	}

	rem := tstack1.Pop()

	if tstack1.Peak() != 2 {
		fmt.Println("Pop/Peak (Case1): Not Passed, did not pop/peak correctly")
	}else{
		total++
	}

	if rem != 3 {
		fmt.Println("Pop (Case1): Not Passed, incorrect return")
	}else{
		total++
	}

	tstack1.Pop()
	tstack1.Pop()

	if !tstack1.Empty(){
		fmt.Println("Empty/Pop (Case1): Not Passed")
	}else{
		total++
	}

	fmt.Print("Test Passed on Stack: ", total, " / 7\n\n")
}

func testQueue() {
	defer catch("Queue")

	var total int = 0;
	tqueue1 := queue.NewLinkedQueue[int]()

	if !tqueue1.Empty() {
		fmt.Println("Empty (Case1): Not passed")
	}else{
		total++
	}

	tqueue1.Enqueue(1)

	if tqueue1.Peak() != 1 {
		fmt.Println("Enqueue/Peak (Case1): Not passed, incorrect peak value")
	}else{
		total++
	}

	if tqueue1.Empty() {
		fmt.Println("Empty (Case2): Not passed")
	}else{
		total++
	}

	tqueue1.Enqueue(2)
	tqueue1.Enqueue(3)

	if tqueue1.Peak() != 1 {
		fmt.Println("Enqueue/Peak (Case2): Not passed, incorrect peak value")
	}else{
		total++
	}

	val := tqueue1.Dequeue()

	if val != 1 {
		fmt.Println("Dequeue (Case1): Not passed, incorrect dequeue return")
	}else{
		total++
	}

	if tqueue1.Peak() != 2 {
		fmt.Println("Dequeue (Case1.5): Not passed, next value incorrect")
	}else{
		total++
	}

	tqueue1.Dequeue()
	tqueue1.Dequeue()

	if !tqueue1.Empty() {
		fmt.Println("Dequeue/Empty (Case1): not functioning properly")
	}else{
		total++
	}

	fmt.Print("Tests Passed on Queue: ", total, " / 7\n\n")
}

func testSingleList() {

}

func testDoubleList() {

}

func testCircleList() {

}

func main() {
	testStack()
	testQueue()
	testBST()
}