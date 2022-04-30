package list

import (
	"fmt"
	"testing"
)

func tList(t *testing.T, listTest LinkedList[int], suite string) {
	//Empty + Push Back/Front + Get
	if !listTest.Empty() {
		t.Error(suite + ": Empty (Case1)")
	}

	listTest.PushFront(2) // List: 2,

	if listTest.Empty() {
		t.Error(suite + ": Empty (Case2)")
	}

	if item, err := listTest.Get(0); item != 2 || err != nil {
		if err != nil {
			t.Fatal(err)
		}

		t.Fatal(suite + ": Get/PushFront (Case1)\nExpected 2, got " + fmt.Sprint(item))
	}

	listTest.PushFront(1) // List: 1, 2,

	if item, err := listTest.Get(0); item != 1 || err != nil {
		if err != nil {
			t.Fatal(err)
		}

		t.Fatal(suite + ": Get/PushFront (Case2)\nExpected 1, got " + fmt.Sprint(item))
	}

	listTest.PushBack(3) //List: 1, 2, 3,


	if item, err := listTest.Get(2); item != 3 || err != nil {
		if err != nil {
			t.Fatal(err)
		}

		t.Fatal(suite + ": Get/PushBack (Case1)\nExpected 3, got " + fmt.Sprint(item))
	}

	//Insert
	listTest.Insert(0, 0) //List: 0, 1, 2, 3,

	if item, err := listTest.Get(0); item != 0 || err != nil {
		if err != nil {
			t.Fatal(err)
		}

		t.Fatal(suite + ": Get/Insert (Case1)\nExpected 0, got " + fmt.Sprint(item))
	}

	listTest.Insert(4, 5)


	if item, err := listTest.Get(4); item != 5 || err != nil {
		if err != nil {
			t.Fatal(err)
		}

		t.Fatal(suite + ": Get/Insert (Case2)\nExpected 5, got " + fmt.Sprint(item))
	}

	listTest.Insert(10, 6) //this long also to accord the circular list

	if item, err := listTest.Get(5); item != 6 || err != nil {
		if err != nil {
			t.Fatal(err)
		}

		t.Fatal(suite + ": Get/Insert (Case3)\nExpected 6, got " + fmt.Sprint(item))
	}

	listTest.Insert(4, 4)

	if item, err := listTest.Get(4); item != 4 || err != nil {
		if err != nil {
			t.Fatal(err)
		}

		t.Fatal(suite + ": Get/Insert (Case4)\nExpected 4, got " + fmt.Sprint(item))
	}

	//Remove Cases

	//Search Cases
}

func TestLists(t *testing.T){
	sing := NewSinglyLinkedList[int]()
	doub := NewDoublyLinkedList[int]()
	circ := NewCircularList[int]()

	tList(t, sing, "Single")
	//Any Extra Tests

	tList(t, doub, "Double")
	//Any Extra Tests

	tList(t, circ, "Circle")
	//Any Extra Tests
}