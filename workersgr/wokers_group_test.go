package workersgr_test

import (
	"sync"
	"testing"

	"gitlab.id.vin/platform/gopkgs/workersgr"
)

func TestWorkersPool(t *testing.T) {
	const numIterations = 100
	wp := workersgr.New()
	executed := sync.Map{}
	for i := 0; i < numIterations; i++ {
		j := i
		wp.Run(func() error {
			executed.Store(j, struct {}{})
			return nil
		})
	}
	wp.Wait()
	for i := 0; i < numIterations; i++ {
		_, ok := executed.Load(i)
		if !ok {
			t.Fatalf("iteration %d has not been executed yet", i)
		}
	}
}
