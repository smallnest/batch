package batch

import (
	"testing"
	"time"
)

func TestBatch(t *testing.T) {
	in := make(chan int, 100)
	go func() {
		for i := 0; i < 50; i++ {
			for j := 0; j < 100; j++ {
				in <- i*1000 + j
			}
			time.Sleep(10 * time.Millisecond)
		}
		close(in)
	}()

	count := 0
	Batch(in, 80, func(items []int) {
		count += len(items)
	})

	if count != 5000 {
		t.Errorf("expect 500 but got %d", count)
	}
}

func TestBatchTimeout(t *testing.T) {
	in := make(chan int, 100)
	go func() {
		for i := 0; i < 50; i++ {
			for j := 0; j < 100; j++ {
				in <- i*1000 + j
			}
			time.Sleep(10 * time.Millisecond)

		}

		close(in)
	}()

	count := 0
	BatchTimeout(in, 80, 10*time.Millisecond, func(items []int) {
		count += len(items)
	})

	if count != 5000 {
		t.Errorf("expect 500 but got %d", count)
	}
}
