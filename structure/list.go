package structure

type List[T any] struct {
	data []T
}

func NewList[T any]() List[T] {
	return List[T]{
		data: []T{},
	}
}

func (l *List[T]) Add(e T) {
	l.data = append(l.data, e)
}

func (l *List[T]) AsSlice() []T {
	return l.data
}
