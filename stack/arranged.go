package stack

func NewArrangedStack[T any]() *ArrangedStack[T] {
	return &ArrangedStack[T]{}
}

func (as *ArrangedStack[T]) Push(value T) {
	as.arrangement = append(as.arrangement, value)
}

func (as *ArrangedStack[T]) Pop() T {
	val := as.arrangement[len(as.arrangement)-1]
	as.arrangement = as.arrangement[0:(len(as.arrangement) - 1)]

	return val
}

func (as *ArrangedStack[T]) Peak() T {
	return as.arrangement[len(as.arrangement)-1]
}

func (as *ArrangedStack[T]) Empty() bool {
	return len(as.arrangement) <= 0
}
