package stack

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(x T) {
	s.values = append(s.values, x)
}

func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		return nil
	}
	length := len(s.values)
	value := s.values[length-1]
	s.values = s.values[:length-1]
	return &value
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
