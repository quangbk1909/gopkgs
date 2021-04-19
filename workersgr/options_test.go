package workersgr

import (
	"errors"
	"sync"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	errsMap := sync.Map{}
	newError := func() error {
		err := errors.New("some error")
		errsMap.Store(err, true)
		return err
	}

	wp := New(
		OptionErrHandler(func(err error) {
			_, ok := errsMap.Load(err)
			if !ok {
				t.Fatalf("error handler received an error that does not exist, error: %v", err)
			}
			errsMap.Store(err, false)
		}),
	)
	for i := 0; i < 100; i++ {
		wp.Run(func() error {
			return newError()
		})
	}
	wp.Wait()
	errsMap.Range(func(key, value interface{}) bool {
		mustExist := value.(bool)
		if mustExist {
			t.Fatal("an error is not received by error handler")
		}
		return true
	})
}

func TestMaxGoroutines(t *testing.T) {
	tests := []struct {
		name          string
		maxGoroutines int
	}{
		{
			name:          "max goroutines = 10000",
			maxGoroutines: 10000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wg := New(OptionMaxGoroutines(test.maxGoroutines))
			done := make(chan struct{})
			cStarted := make(chan struct{}, test.maxGoroutines)
			cntStarted := 0
			for i := 0; i < test.maxGoroutines; i++ {
				wg.Run(func() error {
					cStarted <- struct{}{}
					<-done
					return nil
				})
			}
			for range cStarted {
				cntStarted++
				if cntStarted == test.maxGoroutines {
					break
				}
			}
			if wg.tryAcquireOneSlot() {
				t.Fatal("still got a slot after reach max goroutines limit")
			}
			close(done)
			// check that a new go routine could be run if available
			wg.Run(func() error {
				return nil
			})
			wg.Wait()
		})
	}
}
