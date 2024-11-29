package sync

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert.Equal(t, 3, counter.Value())
	})

	t.Run("Runs concurrently", func(t *testing.T) {
		counter := NewCounter()
		workerCount := 1000

		var waitGroup sync.WaitGroup
		waitGroup.Add(workerCount)

		for i := 0; i < workerCount; i++ {
			go func() {
				counter.Inc()
				waitGroup.Done()
			}()
		}

		waitGroup.Wait()

		assert.Equal(t, workerCount, counter.Value())
	})
}
