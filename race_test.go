package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
)

func TestDataRaceCondition(t *testing.T) {
	var state int32
	mu := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			state += int32(i)
			mu.Unlock()
		}(i)
	}

}

func TestAtomicDataRaceCondition(t *testing.T) {
	var state int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&state, int32(i))
		}(i)
	}

}

func TestThreadSafe(t *testing.T) {

	// var m map[int]int

	var m sync.Map
	m.Store("key1", "value1")

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Store(i, i-rand.Int())
			// m[i] = i - rand.Int()
		}(i)
	}
}
