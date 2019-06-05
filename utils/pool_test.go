package utils

import (
	"sync/atomic"
	"testing"
)

func TestPool(t *testing.T) {
	wgp := NewWaitGroupPool(10)

	number := 1000
	var loop int32
	for i := 0; i < number; i++ {
		wgp.Add()
		go func(loop *int32) {
			defer wgp.Done()
			atomic.AddInt32(loop, 1)
		}(&loop)
	}
	wgp.Wait()

	if int(loop) != number {
		t.Fatal("Pool test is failure")
	}
}
