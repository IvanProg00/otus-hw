package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	c := NewCache(5)

	wasInCache := c.Set("aaa", 100)
	require.False(t, wasInCache)

	wasInCache = c.Set("bbb", 200)
	require.False(t, wasInCache)

	val, ok := c.Get("aaa")
	require.True(t, ok)
	require.Equal(t, 100, val)

	val, ok = c.Get("bbb")
	require.True(t, ok)
	require.Equal(t, 200, val)

	wasInCache = c.Set("aaa", 300)
	require.True(t, wasInCache)

	val, ok = c.Get("aaa")
	require.True(t, ok)
	require.Equal(t, 300, val)

	val, ok = c.Get("ccc")
	require.False(t, ok)
	require.Nil(t, val)
}

func TestCacheEmpty(t *testing.T) {
	c := NewCache(10)

	_, ok := c.Get("aaa")
	require.False(t, ok)

	_, ok = c.Get("bbb")
	require.False(t, ok)
}

func TestCachePurgeLogic(t *testing.T) {
	c := NewCache(1)

	wasInCache := c.Set("a", 1)
	require.False(t, wasInCache)
	val, wasInCache := c.Get("a")
	require.True(t, wasInCache)
	require.Equal(t, 1, val)

	wasInCache = c.Set("b", 2)
	require.False(t, wasInCache)
	val, wasInCache = c.Get("b")
	require.True(t, wasInCache)
	require.Equal(t, 2, val)
	val, wasInCache = c.Get("a")
	require.False(t, wasInCache)
	require.Nil(t, val)

	c.Clear()
	val, wasInCache = c.Get("a")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("b")
	require.False(t, wasInCache)
	require.Nil(t, val)
}

func cacheSetGet(t *testing.T, c Cache) {
	t.Helper()
	wasInCache := c.Set("1", 1) // {"1": 1}
	require.False(t, wasInCache)
	wasInCache = c.Set("2", 2) // {"2": 2, "1": 1}
	require.False(t, wasInCache)
	wasInCache = c.Set("3", 3) // {"3": 3, "2": 2, "1": 1}
	require.False(t, wasInCache)
	wasInCache = c.Set("4", 4) // {"4": 4, "3": 3, "2": 2, "1": 1}
	require.False(t, wasInCache)

	wasInCache = c.Set("5", 5) // {"5": 5, "4": 4, "3": 3, "2": 2}
	require.False(t, wasInCache)
	val, wasInCache := c.Get("2")
	require.True(t, wasInCache)
	require.Equal(t, 2, val)
	val, wasInCache = c.Get("3")
	require.True(t, wasInCache)
	require.Equal(t, 3, val)
	val, wasInCache = c.Get("4")
	require.True(t, wasInCache)
	require.Equal(t, 4, val)
	val, wasInCache = c.Get("5")
	require.True(t, wasInCache)
	require.Equal(t, 5, val)
	val, wasInCache = c.Get("1")
	require.False(t, wasInCache)
	require.Nil(t, val)

	wasInCache = c.Set("2", 22) // {"2": 22, "5": 5, "4": 4, "3": 3}
	require.True(t, wasInCache)
	wasInCache = c.Set("6", 6) // {"6": 6, "2": 22, "5": 5, "4": 4}
	require.False(t, wasInCache)
	val, wasInCache = c.Get("2")
	require.True(t, wasInCache)
	require.Equal(t, 22, val)
	val, wasInCache = c.Get("3")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("4")
	require.True(t, wasInCache)
	require.Equal(t, 4, val)
	val, wasInCache = c.Get("5")
	require.True(t, wasInCache)
	require.Equal(t, 5, val)
	val, wasInCache = c.Get("6")
	require.True(t, wasInCache)
	require.Equal(t, 6, val)
}

func cacheReplaceGet(t *testing.T, c Cache) {
	t.Helper()
	wasInCache := c.Set("2", 200) // {"2": 200, "6": 6, "5": 5, "4": 4}
	require.True(t, wasInCache)
	wasInCache = c.Set("5", 50) // {"5": 50, "2": 200, "6": 6, "4": 4}
	require.True(t, wasInCache)
	wasInCache = c.Set("6", 60) // {"6": 60, "5": 50, "2": 200, "4": 4}
	require.True(t, wasInCache)
	wasInCache = c.Set("4", 40) // {"4": 40, "6": 60, "5": 50, "2": 200}
	require.True(t, wasInCache)
	wasInCache = c.Set("4", 44) // {"4": 44, "6": 60, "5": 50, "2": 200}
	require.True(t, wasInCache)
	wasInCache = c.Set("5", 55) // {"5": 55, "4": 44, "6": 60, "2": 200}
	require.True(t, wasInCache)
	wasInCache = c.Set("6", 66) // {"6": 66, "5": 55, "4": 44, "2": 200}
	require.True(t, wasInCache)

	c.Set("100", 100) // {"100": 100, "6": 66, "5": 55, "4": 44}
	val, wasInCache := c.Get("2")
	require.False(t, wasInCache)
	require.Nil(t, val)
	c.Set("200", 200) // {"100": 100, "200": 200, "6": 66, "5": 55}
	val, wasInCache = c.Get("4")
	require.False(t, wasInCache)
	require.Nil(t, val)
	c.Set("300", 300) // {"100": 100, "200": 200, "300": 300, "6": 66}
	val, wasInCache = c.Get("5")
	require.False(t, wasInCache)
	require.Nil(t, val)
	c.Set("400", 400) // {"100": 100, "200": 200, "300": 300, "400": 400}
	val, wasInCache = c.Get("6")
	require.False(t, wasInCache)
	require.Nil(t, val)

	val, wasInCache = c.Get("100") // {"100": 100, "200": 200, "300": 300, "400": 400}
	require.True(t, wasInCache)
	require.Equal(t, 100, val)
	val, wasInCache = c.Get("300") // {"300": 300, "100": 100, "200": 200, "400": 400}
	require.True(t, wasInCache)
	require.Equal(t, 300, val)
	val, wasInCache = c.Get("100") // {"100": 100, "300": 300, "200": 200, "400": 400}
	require.True(t, wasInCache)
	require.Equal(t, 100, val)
	val, wasInCache = c.Get("400") // {"400": 400, "100": 100, "300": 300, "200": 200}
	require.True(t, wasInCache)
	require.Equal(t, 400, val)
	val, wasInCache = c.Get("200") // {"200": 200, "400": 400, "100": 100, "300": 300}
	require.True(t, wasInCache)
	require.Equal(t, 200, val)
	val, wasInCache = c.Get("300") // {"300": 300, "200": 200, "400": 400, "100": 100}
	require.True(t, wasInCache)
	require.Equal(t, 300, val)
	val, wasInCache = c.Get("400") // {"400": 400, "300": 300, "200": 200, "100": 100}
	require.True(t, wasInCache)
	require.Equal(t, 400, val)
	val, wasInCache = c.Get("200") // {"200": 200, "400": 400, "300": 300, "100": 100}
	require.True(t, wasInCache)
	require.Equal(t, 200, val)
}

func cachePopGet(t *testing.T, c Cache) {
	t.Helper()
	wasInCache := c.Set("1", 1) // {"1": 1, "200": 200, "400": 400, "300": 300}
	require.False(t, wasInCache)
	val, wasInCache := c.Get("100")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("300")
	require.True(t, wasInCache)
	require.Equal(t, 300, val)
	val, wasInCache = c.Get("400")
	require.True(t, wasInCache)
	require.Equal(t, 400, val)
	val, wasInCache = c.Get("200")
	require.True(t, wasInCache)
	require.Equal(t, 200, val)
	val, wasInCache = c.Get("1")
	require.True(t, wasInCache)
	require.Equal(t, 1, val)

	wasInCache = c.Set("2", 2) // {"2": 2, "1": 1, "200": 200, "400": 400}
	require.False(t, wasInCache)
	val, wasInCache = c.Get("300")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("400")
	require.True(t, wasInCache)
	require.Equal(t, 400, val)
	val, wasInCache = c.Get("200")
	require.True(t, wasInCache)
	require.Equal(t, 200, val)
	val, wasInCache = c.Get("1")
	require.True(t, wasInCache)
	require.Equal(t, 1, val)
	val, wasInCache = c.Get("2")
	require.True(t, wasInCache)
	require.Equal(t, 2, val)

	wasInCache = c.Set("3", 3) // {"3": 3, "2": 2, "1": 1, "200": 200}
	require.False(t, wasInCache)
	wasInCache = c.Set("4", 4) // {"4": 4, "3": 3, "2": 2, "1": 1}
	require.False(t, wasInCache)
	val, wasInCache = c.Get("300")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("200")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("1")
	require.True(t, wasInCache)
	require.Equal(t, 1, val)
	val, wasInCache = c.Get("2")
	require.True(t, wasInCache)
	require.Equal(t, 2, val)
	val, wasInCache = c.Get("3")
	require.True(t, wasInCache)
	require.Equal(t, 3, val)
	val, wasInCache = c.Get("4")
	require.True(t, wasInCache)
	require.Equal(t, 4, val)
}

func cacheClearSetGet(t *testing.T, c Cache) {
	t.Helper()
	c.Clear()
	val, wasInCache := c.Get("4")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("3")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("2")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("1")
	require.False(t, wasInCache)
	require.Nil(t, val)

	wasInCache = c.Set("1", 1)
	require.False(t, wasInCache)
	wasInCache = c.Set("2", 2)
	require.False(t, wasInCache)
	wasInCache = c.Set("3", 3)
	require.False(t, wasInCache)
	wasInCache = c.Set("4", 4)
	require.False(t, wasInCache)
	wasInCache = c.Set("5", 5)
	require.False(t, wasInCache)
	// {"5": 5, "4": 4, "3": 3, "2": 2}

	val, wasInCache = c.Get("1")
	require.False(t, wasInCache)
	require.Nil(t, val)
	val, wasInCache = c.Get("2")
	require.True(t, wasInCache)
	require.Equal(t, 2, val)
	val, wasInCache = c.Get("3")
	require.True(t, wasInCache)
	require.Equal(t, 3, val)
	val, wasInCache = c.Get("4")
	require.True(t, wasInCache)
	require.Equal(t, 4, val)
	val, wasInCache = c.Get("5")
	require.True(t, wasInCache)
	require.Equal(t, 5, val)
}

func TestCacheComplexLogic(t *testing.T) {
	c := NewCache(4)
	cacheSetGet(t, c)
	cacheReplaceGet(t, c)
	cachePopGet(t, c)
	cacheClearSetGet(t, c)
}

func TestCachePurge(t *testing.T) {
	c := NewCache(3)
	c.Set("1", 1) // {"1": 1}
	c.Set("2", 2) // {"2": 2, "1": 1}
	c.Set("3", 3) // {"3": 3, "2": 2, "1": 1}
	c.Set("4", 4) // {"4": 4, "3": 3, "2": 2}

	c.Clear()
	val, wasInCache := c.Get("1")
	require.False(t, wasInCache)
	require.Nil(nil, val)
	val, wasInCache = c.Get("2")
	require.False(t, wasInCache)
	require.Nil(nil, val)
	val, wasInCache = c.Get("3")
	require.False(t, wasInCache)
	require.Nil(nil, val)
	val, wasInCache = c.Get("4")
	require.False(t, wasInCache)
	require.Nil(nil, val)
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
