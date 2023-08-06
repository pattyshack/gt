package argparse

import (
  "fmt"
  "sort"
  "strings"
  "time"
)

// EnumType is both a ValueValidator and a ValueSuggestor.
type EnumType[T any] struct{
  ValueType[T]

  // Note: we're sorting / comparing the value strings rather than values
  // directly since strings are always comparable.  This is computationally
  // inefficient, but reduces the amount of boilerplate code.
  valueStrings []string
}

func NewEnumType[T any](
  marshaler ValueType[T],
  values ...T,
) *EnumType[T] {
  valueStrings := make([]string, 0, len(values))
  for _, value := range values {
    valueStrings = append(valueStrings, marshaler.Marshal(value))
  }
  sort.Strings(valueStrings)

  return &EnumType[T]{
    ValueType: marshaler,
    valueStrings: valueStrings,
  }
}

func (enum *EnumType[T]) Validate(value T) error {
  valueString := enum.Marshal(value)
  for _, enumValueString := range enum.valueStrings {
    if valueString == enumValueString {
      return nil
    }
  }

  return fmt.Errorf("invalid enum value")
}

func (enum *EnumType[T]) TypeDescription() string {
  values := make([]string, 0, len(enum.valueStrings))
  for _, value := range enum.valueStrings {
    values = append(values, fmt.Sprintf("%#v", value))
  }

  return "one of {" + strings.Join(values, ", ") + "}"
}

func (enum *EnumType[T]) Suggest(valuePrefix string) []Suggestion {
  suggestions := make([]Suggestion, 0, len(enum.valueStrings))
  for _, value := range enum.valueStrings {
    suggestions = append(
      suggestions,
      Suggestion{
        Value: value,
        IsPrefix: false,
      })
  }
  return suggestions
}

func NewStringEnumType(enumValues ...string) *EnumType[string] {
  return NewEnumType[string](StringType{}, enumValues...)
}

func NewIntEnumType(enumValues ...int) *EnumType[int] {
  return NewEnumType[int](IntType{}, enumValues...)
}

func NewInt64EnumType(enumValues ...int64) *EnumType[int64] {
  return NewEnumType[int64](Int64Type{}, enumValues...)
}

func NewUintEnumType(enumValues ...uint) *EnumType[uint] {
  return NewEnumType[uint](UintType{}, enumValues...)
}

func NewUint64EnumType(enumValues ...uint64) *EnumType[uint64] {
  return NewEnumType[uint64](Uint64Type{}, enumValues...)
}

func NewDurationEnumType(enumValues ...time.Duration) *EnumType[time.Duration] {
  return NewEnumType[time.Duration](DurationType{}, enumValues...)
}
