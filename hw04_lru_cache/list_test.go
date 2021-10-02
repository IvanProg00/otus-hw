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

		l.Remove(l.Back())
		require.Equal(t, []int{10, 80, 60, 40}, getListInt(l))
		require.Equal(t, 4, l.Len())
		l.Remove(l.Front())
		require.Equal(t, []int{80, 60, 40}, getListInt(l))
		require.Equal(t, 3, l.Len())
		l.Remove(l.Back())
		require.Equal(t, []int{80, 60}, getListInt(l))
		require.Equal(t, 2, l.Len())
		l.Remove(l.Front())
		require.Equal(t, []int{60}, getListInt(l))
		require.Equal(t, 1, l.Len())
		l.Remove(l.Back())
		require.Equal(t, []int{}, getListInt(l))
		require.Equal(t, 0, l.Len())

	})

	t.Run("prev next value", func(t *testing.T) {
		l := NewList()
		i1 := l.PushFront(10) // [10]
		require.Equal(t, 1, l.Len())
		require.Equal(t, 10, i1.Value)
		require.Nil(t, i1.Next)
		require.Nil(t, i1.Prev)

		l.Remove(l.Back())
		i1 = l.PushBack(20) // [20]
		require.Equal(t, 1, l.Len())
		require.Equal(t, 20, i1.Value)
		require.Nil(t, i1.Next)
		require.Nil(t, i1.Prev)

		i2 := l.PushBack(30) // [20, 30]
		require.Equal(t, 2, l.Len())
		require.Equal(t, 20, i1.Value)
		require.NotNil(t, i1.Next)
		require.Nil(t, i1.Prev)
		require.Equal(t, 30, i2.Value)
		require.Nil(t, i2.Next)
		require.NotNil(t, i2.Prev)

		i3 := l.PushFront(40) // [40, 20, 30]
		require.Equal(t, 3, l.Len())
		require.Equal(t, 20, i1.Value)
		require.NotNil(t, i1.Next)
		require.NotNil(t, i1.Prev)
		require.Equal(t, 30, i2.Value)
		require.Nil(t, i2.Next)
		require.NotNil(t, i2.Prev)
		require.Equal(t, 40, i3.Value)
		require.NotNil(t, i3.Next)
		require.Nil(t, i3.Prev)
	})
}

func getListInt(l List) []int {
	elems := make([]int, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(int))
	}
	return elems
}
