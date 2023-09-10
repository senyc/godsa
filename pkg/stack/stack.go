package stack

type Stack[T any] []T

func (s *Stack[T]) Push(t T) {
	*s = append(*s, t)
}

// Add error here
func (s *Stack[T]) Pop() T {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}
