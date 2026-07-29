package generics

import (
	"github.com/google/go-cmp/cmp"
)

type List[T any] struct {
	Items []T
}

func (list *List[T]) Add(item T) {
	list.Items = append(list.Items, item)
}

func (list *List[T]) InsertAt(item T, index int) {
	list.Items = append(list.Items, item)
	copy(list.Items[index+1:], list.Items[index:])
	list.Items[index] = item
}

func (list *List[T]) RemoveItem(index int) {
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
}

func (list *List[T]) Remove(item T) {
	index := list.FindItem(item)
	if index != -1 {
		list.RemoveItem(index)
	}

}

func (list *List[T]) FindItem(item T) int {

	for i, v := range list.Items {
		if cmp.Equal(v, item) {
			return i
		}
	}
	return -1
}
