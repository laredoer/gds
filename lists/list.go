package lists

type List[T any] interface {
	Get(index int) (T, bool)
	Remove(index int)
	Add(values ...T)
	Contains(values ...T) bool
	Sort()
	Swap(index1, index2 int)
	Insert(index int, values ...T)
	Set(index int, value T)

	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}
