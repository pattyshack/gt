package lexutil

import (
	"sync"
)

type ErrorEmitter struct {
	mutex sync.Mutex
	errs  []error // guarded by mutex
}

func (emitter *ErrorEmitter) Errors() []error {
	emitter.mutex.Lock()
	defer emitter.mutex.Unlock()

	errs := make([]error, len(emitter.errs), len(emitter.errs))
	copy(errs, emitter.errs)
	return errs
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

	// Accumulated errors from Process
	Errors() []error
}

func Process[T any](
	node T,
	passes [][]Pass[T], // sequence of parallelizable passes
	earlyExitErrThreshold int, // 0 or less = unlimited
) []error {
	errors := []error{}
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

		for _, pass := range parallelPasses {
			errors = append(errors, pass.Errors()...)
		}

		if earlyExitErrThreshold > 0 && len(errors) >= earlyExitErrThreshold {
			return errors
		}
	}

	return errors
}
