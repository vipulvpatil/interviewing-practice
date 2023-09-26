package queue

type Queue[T any] struct {
	values []T
}

func (s *Queue[T]) Enqueue(x T) {
	s.values = append(s.values, x)
}

func (s *Queue[T]) Dequeue() *T {
	if s.IsEmpty() {
		return nil
	}
	value := s.values[0]
	s.values = s.values[1:]
	return &value
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.values) == 0
}
