package argparse

import (
  "flag"
  "fmt"
  "reflect"
  "strconv"
  "time"
)


type Suggestion struct {
  Value string

  // Suggestions of the form --<flag>=<value suggestion>
  IsFlagValueAssignment bool

  // When true, don't add a trailing space to start a new argument
  IsPrefix bool
}

type ValueTypeInfo interface {
  TypeName() string
  TypeDescription() string

  Suggest(valuePrefix string) []Suggestion
}

type ValueSuggestor interface {
  TypeDescription() string

  Suggest(valuePrefix string) []Suggestion
}

type ValueType[T any] interface {
  ValueTypeInfo

  // NOTE: The marshaled value should be deterministic, e.g., serialized map
  // entries should be sorted.
  Marshal(value T) string
  Unmarshal(unparsedValue string) (T, error)

  Validate(value T) error
}

type TypedGetter[T any] interface {
  ValueType[T]
  flag.Getter
}

type value[T any] struct {
  ptr *T

  ValueType[T]
}

func (val *value[T]) Set(str string) error {
  value, err := val.Unmarshal(str)
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

  err = val.Validate(value)
  if err != nil {
    return err
  }

  *val.ptr = value
  return nil
}

func (val *value[T]) String() string {
  return val.Marshal(*val.ptr)
}

func (val *value[T]) TypeName() string {
  var t T
  return reflect.TypeOf(t).Name()
}

func (val *value[T]) Get() any {
  return *val.ptr
}

func NewValue[T any](
  ptr *T,
  defaultValue T,
  marshaler ValueType[T],
) flag.Getter {
  err := marshaler.Validate(defaultValue)
  if err != nil {
    panic(fmt.Sprintf(
      "invalid default value (%s): %s",
      marshaler.Marshal(defaultValue),
      err))
  }

  *ptr = defaultValue
  return &value[T]{
    ptr: ptr,
    ValueType: marshaler,
  }
}

type boolType struct {}

func (boolType) TypeName() string {
  return "bool"
}

func (boolType) TypeDescription() string {
  return ""
}

func (boolType) Marshal(value bool) string {
  return strconv.FormatBool(value)
}

func (boolType) Unmarshal(value string) (bool, error) {
  return strconv.ParseBool(value)
}

func (boolType) Validate(value bool) error {
  return nil
}

func (boolType) Suggest(valuePrefix string) []Suggestion {
  return nil
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
) flag.Getter {
  *ptr = defaultValue
  return &boolValue{
    value: value[bool]{
      ptr: ptr,
      ValueType: boolType{},
    },
  }
}

func (boolValue) IsBoolFlag() bool {
  return true
}

type StringType struct {}

func (StringType) TypeName() string {
  return "string"
}

func (StringType) TypeDescription() string {
  return ""
}

func (StringType) Marshal(value string) string {
  return value
}

func (StringType) Unmarshal(value string) (string, error) {
  return value, nil
}

func (StringType) Validate(value string) error {
  return nil
}

func (StringType) Suggest(valuePrefix string) []Suggestion {
  return nil
}

type IntType struct {}

func (IntType) TypeName() string {
  return "int"
}

func (IntType) TypeDescription() string {
  return ""
}

func (IntType) Marshal(value int) string {
  return strconv.Itoa(value)
}

func (IntType) Unmarshal(value string) (int, error) {
  val, err := strconv.ParseInt(value, 0, strconv.IntSize)
  return int(val), err
}

func (IntType) Validate(value int) error {
  return nil
}

func (IntType) Suggest(valuePrefix string) []Suggestion {
  return nil
}

type Int64Type struct {}

func (Int64Type) TypeName() string {
  return "int64"
}

func (Int64Type) TypeDescription() string {
  return ""
}

func (Int64Type) Marshal(value int64) string {
  return strconv.FormatInt(value, 10)
}

func (Int64Type) Unmarshal(value string) (int64, error) {
  return strconv.ParseInt(value, 0, 64)
}

func (Int64Type) Validate(value int64) error {
  return nil
}

func (Int64Type) Suggest(valuePrefix string) []Suggestion {
  return nil
}

type UintType struct {}

func (UintType) TypeName() string {
  return "uint"
}

func (UintType) TypeDescription() string {
  return ""
}

func (UintType) Marshal(value uint) string {
  return strconv.FormatUint(uint64(value), 10)
}

func (UintType) Unmarshal(value string) (uint, error) {
  val, err := strconv.ParseUint(value, 0, strconv.IntSize)
  return uint(val), err
}

func (UintType) Validate(value uint) error {
  return nil
}

func (UintType) Suggest(valuePrefix string) []Suggestion {
  return nil
}

type Uint64Type struct {}

func (Uint64Type) TypeName() string {
  return "uint64"
}

func (Uint64Type) TypeDescription() string {
  return ""
}

func (Uint64Type) Marshal(value uint64) string {
  return strconv.FormatUint(value, 10)
}

func (Uint64Type) Unmarshal(value string) (uint64, error) {
  return strconv.ParseUint(value, 0, strconv.IntSize)
}

func (Uint64Type) Validate(value uint64) error {
  return nil
}

func (Uint64Type) Suggest(valuePrefix string) []Suggestion {
  return nil
}

type Float64Type struct {}

func (Float64Type) TypeName() string {
  return "float64"
}

func (Float64Type) TypeDescription() string {
  return ""
}

func (Float64Type) Marshal(value float64) string {
  return strconv.FormatFloat(value,'g', -1, 64)
}

func (Float64Type) Unmarshal(value string) (float64, error) {
  return strconv.ParseFloat(value, 64)
}

func (Float64Type) Validate(value float64) error {
  return nil
}

func (Float64Type) Suggest(valuePrefix string) []Suggestion {
  return nil
}

type DurationType struct {}

func (DurationType) TypeName() string {
  return "time.Duration"
}

func (DurationType) TypeDescription() string {
  return ""
}

func (DurationType) Marshal(value time.Duration) string {
  return value.String()
}

func (DurationType) Unmarshal(value string) (time.Duration, error) {
  return time.ParseDuration(value)
}

func (DurationType) Validate(value time.Duration) error {
  return nil
}

func (DurationType) Suggest(valuePrefix string) []Suggestion {
  return nil
}

