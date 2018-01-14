package reloader

import (
	"sync"

	"github.com/dc0d/config"
	"github.com/dc0d/retry"
	"github.com/fsnotify/fsnotify"
)

// Reloader .
type Reloader struct {
	w        *fsnotify.Watcher
	l        *sync.Mutex
	stop     chan struct{}
	stopOnce *sync.Once
	events   chan fsnotify.Event
	onError  func(error)
	onEvent  func(fsnotify.Event)
}

// New .
func New(onError func(error)) *Reloader {
	if onError == nil {
		panic(ErrNilOnError)
	}
	w, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	res := &Reloader{
		w:        w,
		l:        &sync.Mutex{},
		stop:     make(chan struct{}),
		stopOnce: &sync.Once{},
		onError:  onError,
	}
	res.events = res.w.Events
	go retry.Retry(
		res.loop,
		-1,
		onError)
	return res
}

// Stop .
func (r *Reloader) Stop() {
	r.l.Lock()
	defer r.l.Unlock()
	r.stopOnce.Do(func() {
		close(r.stop)
		r.w.Close()
		r.w = nil
	})
}

// SetOnEvent .
func (r *Reloader) SetOnEvent(onEvent func(fsnotify.Event)) {
	if onEvent == nil {
		return
	}
	r.l.Lock()
	defer r.l.Unlock()
	r.onEvent = onEvent
}

// Add .
func (r *Reloader) Add(
	fileName string,
	dirs ...string) error {
	r.l.Lock()
	defer r.l.Unlock()
	if r.w == nil {
		return ErrStopped
	}
	fp, err := config.RelativeResource(fileName, dirs...)
	if err != nil {
		return err
	}
	return r.w.Add(fp)
}

func (r *Reloader) loop() error {
	for {
		select {
		case <-r.stop:
			return nil
		case e, ok := <-r.events:
			if !ok {
				return nil
			}
			r.handleEvent(e)
		}
	}
}

func (r *Reloader) handleEvent(ev fsnotify.Event) {
	r.l.Lock()
	defer r.l.Unlock()
	if r.w == nil {
		return
	}
	if r.onEvent == nil {
		return
	}
	go r.onEvent(ev)
}

// errors
var (
	ErrNilOnError = errorf("NIL onError")
	ErrStopped    = errorf("STOPPED")
)
