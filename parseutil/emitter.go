package parseutil

import (
	"sort"
	"sync"
)

type Emitter struct {
	mutex sync.Mutex
	errs  []error // guarded by mutex
}

func (emitter *Emitter) errors() []error {
	emitter.mutex.Lock()
	defer emitter.mutex.Unlock()

	return emitter.errs
}

func (emitter *Emitter) Errors() []error {
	errs := emitter.errors()

	sorted := make([]error, len(errs), len(errs))
	copy(sorted, errs)
	sort.Sort(ErrorsByLocation(sorted))

	return sorted
}

func (emitter *Emitter) HasErrors() bool {
	emitter.mutex.Lock()
	defer emitter.mutex.Unlock()

	return len(emitter.errs) > 0
}

func (emitter *Emitter) Emit(
	loc Location,
	format string,
	args ...interface{},
) error {
	err := NewLocationError(loc, format, args...)
	emitter.EmitErrors(err)
	return err
}

func (emitter *Emitter) EmitErrors(errs ...error) {
	emitter.mutex.Lock()
	defer emitter.mutex.Unlock()

	emitter.errs = append(emitter.errs, errs...)
}
