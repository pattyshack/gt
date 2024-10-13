package lexutil

import (
	"sort"
	"sync"
)

type ErrorEmitter struct {
	mutex sync.Mutex
	errs  []error // guarded by mutex
}

func (emitter *ErrorEmitter) MergeFrom(other *ErrorEmitter) {
	emitter.EmitErrors(other.errors()...)
}

func (emitter *ErrorEmitter) errors() []error {
	emitter.mutex.Lock()
	defer emitter.mutex.Unlock()

	return emitter.errs
}

func (emitter *ErrorEmitter) Errors() []error {
	errs := emitter.errors()

	sorted := make([]error, len(errs), len(errs))
	copy(sorted, errs)
	sort.Sort(ErrorsByLocation(sorted))

	return sorted
}

func (emitter *ErrorEmitter) HasErrors() bool {
	emitter.mutex.Lock()
	defer emitter.mutex.Unlock()

	return len(emitter.errs) > 0
}

func (emitter *ErrorEmitter) Emit(
	loc Location,
	format string,
	args ...interface{}) {

	emitter.EmitErrors(NewLocationError(loc, format, args...))
}

func (emitter *ErrorEmitter) EmitErrors(errs ...error) {
	emitter.mutex.Lock()
	defer emitter.mutex.Unlock()

	emitter.errs = append(emitter.errs, errs...)
}

// Analysis or transform pass
type Pass[T any] interface {
	Process(T)
}

func Process[T any](
	node T,
	passes [][]Pass[T], // sequence of parallelizable passes
	shouldEarlyExit func() bool, // optional
) {
	for _, parallelPasses := range passes {
		wg := sync.WaitGroup{}
		wg.Add(len(parallelPasses))
		for _, pass := range parallelPasses {
			go func(pass Pass[T]) {
				pass.Process(node)
				wg.Done()
			}(pass)
		}

		wg.Wait()

		if shouldEarlyExit != nil && shouldEarlyExit() {
			return
		}
	}
}
