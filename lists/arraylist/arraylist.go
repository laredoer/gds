package arraylist

type List[T any] struct {
	elements []T
	size     int
}

func New[T any](values ...T) *List[T] {

	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}

	return list
}

func (list *List[T]) Add(values ...T) {
	list.elements = append(list.elements, values...)
}

func (list *List[T]) Get(index int) (T, bool) {
	return list.elements[index], true
}

func (list *List[T]) Remove(index int) {
	var t T
	list.elements[index] = t
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
}
