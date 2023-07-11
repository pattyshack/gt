package argparse

import (
  "flag"
  "fmt"
  "reflect"
  "strconv"
  "time"
)

type WithValidValueDescription interface {
  ValidValueDescription() string
}

type TypedGetter interface {
  TypeName() string

  flag.Getter
}

type TypeGetterWithValidValueDescription interface {
  TypedGetter

  WithValidValueDescription
}

type ValueMarshaler[T any] interface {
  // NOTE: The marshaled value should be deterministic, e.g., serialized map
  // entries should be sorted.
  Marshal(T) string
  Unmarshal(string) (T, error)
}

type ValueValidator[T any] interface {
  Validate(T) error

  WithValidValueDescription
}

type value[T any] struct {
  ptr *T

  marshaler ValueMarshaler[T]
  validator ValueValidator[T]
}

func (val *value[T]) Set(str string) error {
  value, err := val.marshaler.Unmarshal(str)
  if err != nil {
    // Mask strconv error detail to match flag lib's error messaging.
    numErr, ok := err.(*strconv.NumError)
    if ok {
      if numErr.Err == strconv.ErrSyntax {
        return fmt.Errorf("parse error")
      }

      if numErr.Err == strconv.ErrRange {
        return fmt.Errorf("value out of range")
      }
    }

    return err
  }

  if val.validator != nil {
    err := val.validator.Validate(value)
    if err != nil {
      return err
    }
  }

  *val.ptr = value
  return nil
}

func (val *value[T]) String() string {
  return val.marshaler.Marshal(*val.ptr)
}

func (val *value[T]) TypeName() string {
  var t T
  return reflect.TypeOf(t).Name()
}

func (val *value[T]) Get() any {
  return *val.ptr
}

type valueWithValidValueDescription[T any] struct {
  value[T]
}

func (val *valueWithValidValueDescription[T]) ValidValueDescription() string {
  return val.validator.ValidValueDescription()
}

// This return a TypedGetter, which could specialize into
// TypeGetterWithValidValueDescription if an validator was provided.
func NewValue[T any](
  ptr *T,
  defaultValue T,
  marshaler ValueMarshaler[T],
  validator ValueValidator[T],  // optional
) TypedGetter {
  *ptr = defaultValue

  if validator == nil {
    return &value[T]{
      ptr: ptr,
      marshaler: marshaler,
    }
  }

  err := validator.Validate(defaultValue)
  if err != nil {
    panic(fmt.Sprintf(
      "invalid default value (%s): %s",
      marshaler.Marshal(defaultValue),
      err))
  }

  return &valueWithValidValueDescription[T]{
    value: value[T]{
      ptr: ptr,
      marshaler: marshaler,
      validator: validator,
    },
  }
}

type boolMarshaler struct {}

func (boolMarshaler) Marshal(value bool) string {
  return strconv.FormatBool(value)
}

func (boolMarshaler) Unmarshal(value string) (bool, error) {
  return strconv.ParseBool(value)
}

// NOTE: we need to special case bool since the flag library parses bool
// differently.
type boolValue struct {
  value[bool]
}

func NewBoolValue(
  ptr *bool,
  defaultValue bool,
  description string,
) TypedGetter {
  *ptr = defaultValue
  return &boolValue{
    value: value[bool]{
      ptr: ptr,
      marshaler: boolMarshaler{},
    },
  }
}

func (boolValue) IsBoolFlag() bool {
  return true
}

type stringMarshaler struct {}

func (stringMarshaler) Marshal(value string) string {
  return value
}

func (stringMarshaler) Unmarshal(value string) (string, error) {
  return value, nil
}

type intMarshaler struct {}

func (intMarshaler) Marshal(value int) string {
  return strconv.Itoa(value)
}

func (intMarshaler) Unmarshal(value string) (int, error) {
  val, err := strconv.ParseInt(value, 0, strconv.IntSize)
  return int(val), err
}

type int64Marshaler struct {}

func (int64Marshaler) Marshal(value int64) string {
  return strconv.FormatInt(value, 10)
}

func (int64Marshaler) Unmarshal(value string) (int64, error) {
  return strconv.ParseInt(value, 0, 64)
}

type uintMarshaler struct {}

func (uintMarshaler) Marshal(value uint) string {
  return strconv.FormatUint(uint64(value), 10)
}

func (uintMarshaler) Unmarshal(value string) (uint, error) {
  val, err := strconv.ParseUint(value, 0, strconv.IntSize)
  return uint(val), err
}

type uint64Marshaler struct {}

func (uint64Marshaler) Marshal(value uint64) string {
  return strconv.FormatUint(value, 10)
}

func (uint64Marshaler) Unmarshal(value string) (uint64, error) {
  return strconv.ParseUint(value, 0, strconv.IntSize)
}

type float64Marshaler struct {}

func (float64Marshaler) Marshal(value float64) string {
  return strconv.FormatFloat(value,'g', -1, 64)
}

func (float64Marshaler) Unmarshal(value string) (float64, error) {
  return strconv.ParseFloat(value, 64)
}

type durationMarshaler struct {}

func (durationMarshaler) Marshal(value time.Duration) string {
  return value.String()
}

func (durationMarshaler) Unmarshal(value string) (time.Duration, error) {
  return time.ParseDuration(value)
}
