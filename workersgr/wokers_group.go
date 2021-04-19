package workersgr

import (
	"sync"
)

type settings struct {
	errHandler    func(error)
	maxGoroutines int
}

type WorkersGroup struct {
	nGoroutines             int
	condGoroutinesAvailable *sync.Cond
	wg                      sync.WaitGroup

	errors         chan<- error
	errHandlerChan chan struct{}

	settings
}

func New(opts ...Option) *WorkersGroup {
	w := &WorkersGroup{}
	for _, opt := range opts {
		opt.apply(&w.settings)
	}

	w.condGoroutinesAvailable = sync.NewCond(&sync.Mutex{})

	if w.errHandler != nil {
		w.errHandlerChan = make(chan struct{})
		c := make(chan error)
		w.errors = c
		go func() {
			// signal to sender: more error is consumed from error channel
			defer close(w.errHandlerChan)

			for err := range c {
				w.errHandler(err)
			}
		}()
	}
	return w
}

func (w *WorkersGroup) Run(f func() error) {
	w.acquireGoroutineSlot()
	go func() {
		defer w.releaseOneSlot()
		if err := f(); err != nil {
			if w.errors != nil {
				select {
				case w.errors <- err:
				case <-w.errHandlerChan:
				}
			}
		}
	}()
}

func (w *WorkersGroup) Wait() {
	w.wg.Wait()
	if w.errors != nil {
		close(w.errors)
		<-w.errHandlerChan
	}
}

func (w *WorkersGroup) acquireGoroutineSlot() {
	w.condGoroutinesAvailable.L.Lock()
	defer w.condGoroutinesAvailable.L.Unlock()
	for !w.tryAcquireOneSlot() {
		w.condGoroutinesAvailable.Wait()
	}
}

func (w *WorkersGroup) tryAcquireOneSlot() bool {
	if w.reachGoroutinesLimit() {
		return false
	}
	w.nGoroutines++
	w.wg.Add(1)
	return true
}

func (w *WorkersGroup) releaseOneSlot() {
	w.condGoroutinesAvailable.L.Lock()
	defer w.condGoroutinesAvailable.L.Unlock()
	w.nGoroutines--
	w.wg.Done()
	w.condGoroutinesAvailable.Signal()
}

func (w *WorkersGroup) reachGoroutinesLimit() bool {
	// max goroutines <= 0 mean no limit
	if w.maxGoroutines <= 0 {
		return false
	}
	return w.nGoroutines >= w.maxGoroutines
}
