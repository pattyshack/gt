package lexutil

import (
	"errors"
	"fmt"
)

type Location struct {
	FileName string

	Line int // 1 based

	// Note: We'll use byte position within the line instead of unicode symbol
	// position since some unicode symbols are composed of multiple unicode
	// runes.  It's too much work to figure out all the cases.
	Column int // 0 based
}

func (loc Location) String() string {
	return fmt.Sprintf("%s:%v:%v", loc.FileName, loc.Line, loc.Column)
}

func (loc Location) ShortString() string {
	return fmt.Sprintf("%v:%v", loc.Line, loc.Column)
}

type LocationError struct {
	Loc Location
	Err error
}

func NewLocationError(
	loc Location,
	format string,
	args ...interface{},
) LocationError {
	return LocationError{
		Loc: loc,
		Err: fmt.Errorf(format, args...),
	}
}

func (le LocationError) Error() string {
	return le.Loc.String() + ": " + le.Err.Error()
}

func (le LocationError) Unwrap() error {
	return le.Err
}

// Sort non-location errors first, followed by location sorted errors
type ErrorsByLocation []error

func (s ErrorsByLocation) Len() int { return len(s) }

func (s ErrorsByLocation) Swap(i int, j int) { s[i], s[j] = s[j], s[i] }

func (s ErrorsByLocation) Less(i int, j int) bool {
	var err1 LocationError
	ok1 := errors.As(s[i], &err1)

	var err2 LocationError
	ok2 := errors.As(s[j], &err2)

	if !ok1 {
		return ok2
	}

	if !ok2 {
		return false
	}

	if err1.Loc.FileName != err2.Loc.FileName {
		return err1.Loc.FileName < err2.Loc.FileName
	}

	if err1.Loc.Line != err2.Loc.Line {
		return err1.Loc.Line < err2.Loc.Line
	}

	return err1.Loc.Column < err2.Loc.Column
}
