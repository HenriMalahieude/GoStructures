# Go Structures!
Just GoLang Data Structures (with generics to handle all data types).

Nothing more, nothing less.

Headers have more in-depth documentation. Please wait while I finish more structures for more documentation.

## Status: ![Go Tests](https://github.com/HenriMalahieude/GoStructures/actions/workflows/go.yml/badge.svg)

## Completed Structures:
1. Trees
    * Binary Search Trees
        * Console Output
        * Visualization (using Dotty by Graphviz)
2. Linked Lists
    * Single
    * Double
    * Circular
3. Stacks
    * Linked List Based
    * Array/Slice Based
4. Queues
    * Non-Priority Linked List Style

## Planned Structures:
1. Graphs
2. ~~Vectors (?)~~
3. Balance Trees (?)
4. Heaps (Min and Max)


# Documention

## Lists

``LinkedList``:
* `SinglyLinkedList`
* `DoublyLinkedList`
* `CircularList`

Usable Methods:
```golang
//This applies to NewDoublyLinkedList and NewCircularList
container := list.NewSinglyLinkedList[int]()

container.PushFront(1)

container.PushBack(2)

container.Insert(2, 3) //Pos, Value
container.Insert(400, 4) //Special Case, just inserted at the back

element, err := container.Get(0)

list.RemoveFirst(func(v int) bool{
    return v == 2
})

position, err := container.Search(func(v int) bool{
    return v == 1;
})

noElements := container.Empty()
```

## Queues

`Queue`:
* `LinkedQueue`

Usable Methods:
```golang
line := queue.NewLinkedQueue[string]()

line.Enqueue("John Doe")

frontOfLine := line.Peak()

nowServing := line.Dequeue()

arePeopleWaiting := line.Empty()

```

## Stack

`Stack`:
* `LinkedStack`
* `ArrangedStack`

Usable Methods:
```golang
//Applies to all stacks
example := stack.NewLinkedStack[float64]()

example.Push(3.1415)
example.Push(2.7878)

top := example.Peak()

using := example.Pop()

anythingLeft := example.Empty() //Bool
```

## Tree

`Tree`:
* `BinaryTree`

Usable Methods:
```golang
pine := tree.NewBinarySearchTree[string](
// Compares left and right to see which one is "bigger" (Used in Insertion/Deletion)
    func(l, r string) bool {
        return len(l) > len(r);
    },
// This Checks for Equality (Used in Search/Deletion)
    func(a, b string) bool {
        return len(a) == len(b)
    })

pine.Insert("A")
pine.Insert("BB")
pine.Insert("CCC")
pine.Insert("DDDD")

pine.Remove("A")

inList := pine.Search("A") //bool

//Useful Functions for Binary Search Trees (Using fmt.Printf("%v")):
func (b BinaryTree[T]) InorderPrint()
func (b BinaryTree[T]) PreorderPrint()
func (b BinaryTree[T]) PostorderPrint()

//If graphviz/dotty is installed:
func (b BinaryTree[T]) VisualizeDotty(fileName string) //Creates a .dot file, that dotty will use to make a .jpg file

```