package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushBack(100)
		l.PushFront(110)
		l.PushFront(120)
		require.Equal(t, []int{120, 110, 100}, getListInt(l))

		l = NewList()
		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, []int{10, 20, 30}, getListInt(l))

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, getListInt(l))

		l.Remove(l.Back())
		require.Equal(t, []int{70, 80, 60, 40, 10, 30}, getListInt(l))
		l.Remove(l.Front())
		require.Equal(t, []int{80, 60, 40, 10, 30}, getListInt(l))

		prevLast := l.Back().Prev
		l.MoveToFront(prevLast)
		require.Equal(t, []int{10, 80, 60, 40, 30}, getListInt(l))
	})
}

func getListInt(l List) []int {
	elems := make([]int, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(int))
	}
	return elems
}
