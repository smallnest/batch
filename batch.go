package batch

import (
	"errors"
	"time"
)


// Batch reads items from in channel, and prepares a batch of items to the handler.
// If you close the channel, Batch will drain all items and exit without any error.
// If there is no sufficient items, it handles current read items and waits for the next batch.
func Batch[T any](in chan T, batchSize int, fn func(items []T)) error {
	if batchSize <= 1 {
		return errors.New("batch is unnecessary")
	}
	var items = make([]T, 0, batchSize)

	for {
		items = items[:0]
		item, ok := <-in
		if !ok {
			return nil
		}
		items = append(items, item)

	prepare_batch:
		for {
			select {
			case item, ok := <-in:
				if !ok {
					break prepare_batch
				}
				items = append(items, item)
				if len(items) == batchSize {
					break prepare_batch
				}
			default:
				break prepare_batch
			}
		}

		if len(items) > 0 {
			fn(items)
		}
	}
}

// BatchTimeout reads items from in channel or tmeout, and then prepares a batch of items to the handler.
// Not like the Batch function, it waits until timeout if there is no sufficient items for batch.
func BatchTimeout[T any](in chan T, batchSize int, timeout time.Duration, fn func(items []T)) error {
	if batchSize <= 1 {
		return errors.New("batch is unnecessary")
	}
	var items = make([]T, 0, batchSize)

	exitLoop := false
	
	for {
		items = items[:0]
		timer := time.NewTimer(timeout)
	prepare_batch:
		for {
			select {
			case item, ok := <-in:
				if !ok {
					exitLoop = true
					break prepare_batch
				}
				items = append(items, item)
				if len(items) == batchSize {
					break prepare_batch
				}
			case <-timer.C:
				break prepare_batch
			}
		}
		timer.Stop()

		if len(items) > 0 {
			fn(items)
		}

		if exitLoop {
			return nil
		}
	}
}
