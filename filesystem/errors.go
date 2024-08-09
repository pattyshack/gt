package filesystem

import (
  "errors"
  "fmt"
  "io/fs"
)

var (
  ErrInvalid = fs.ErrInvalid
  ErrPermission = fs.ErrPermission
  ErrExist = fs.ErrExist
  ErrNotExist = fs.ErrNotExist
  ErrClosed = fs.ErrClosed
)

func WrapError(
  op string,
  path string,
  err error,
) error {
  return &PathError{
    Op: op,
    Path: path,
    Err: err,
  }
}

func WrapErrorf(
  op string,
  path string,
  err error,
  messageTemplate string,
  args ...any,
) error {
  return WrapError(
    op,
    path,
    fmt.Errorf(messageTemplate + ": %w", append(args, err)...))
}

func IsNotExistError(err error) bool {
  return errors.Is(err, ErrNotExist)
}

func NewNotExistError(
  op string,
  path string,
  messageTemplate string,
  args ...any,
) error {
  return WrapErrorf(op, path, ErrNotExist, messageTemplate, args...)
}

func IsExistError(err error) bool {
  return errors.Is(err, ErrExist)
}

func NewExistError(
  op string,
  path string,
  messageTemplate string,
  args ...any,
) error {
  return WrapErrorf(op, path, ErrExist, messageTemplate, args...)
}

func IsInvalidArgumentError(err error) bool {
  return errors.Is(err, ErrInvalid)
}

func NewInvalidArgumentError(
  op string,
  path string,
  messageTemplate string,
  args ...any,
) error {
  return WrapErrorf(op, path, ErrInvalid, messageTemplate, args...)
}

func IsPermissionError(err error) bool {
  return errors.Is(err, ErrPermission)
}

func NewPermissionError(
  op string,
  path string,
  messageTemplate string,
  args ...any,
) error {
  return WrapErrorf(op, path, ErrPermission, messageTemplate, args...)
}
