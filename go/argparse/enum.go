package argparse

import (
  "fmt"
  "sort"
  "strings"
  "time"
)

// Enum is both a ValueValidator and a ValueSuggestor.
type Enum[T any] struct{
  marshaler ValueMarshaler[T]

  // Note: we're sorting / comparing the value strings rather than values
  // directly since strings are always comparable.  This is computationally
  // inefficient, but reduces the amount of boilerplate code.
  valueStrings []string
}

func NewEnum[T any](
  marshaler ValueMarshaler[T],
  values ...T,
) *Enum[T] {
  valueStrings := make([]string, 0, len(values))
  for _, value := range values {
    valueStrings = append(valueStrings, marshaler.Marshal(value))
  }
  sort.Strings(valueStrings)

  return &Enum[T]{
    marshaler: marshaler,
    valueStrings: valueStrings,
  }
}

func (enum *Enum[T]) Validate(value T) error {
  valueString := enum.marshaler.Marshal(value)
  for _, enumValueString := range enum.valueStrings {
    if valueString == enumValueString {
      return nil
    }
  }

  return fmt.Errorf("invalid enum value")
}

func (enum *Enum[T]) ValidValueDescription() string {
  values := make([]string, 0, len(enum.valueStrings))
  for _, value := range enum.valueStrings {
    values = append(values, fmt.Sprintf("%#v", value))
  }

  return "{" + strings.Join(values, ", ") + "}"
}

func (enum *Enum[T]) Suggest(valuePrefix string) []Suggestion {
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

func NewStringEnum(enumValues ...string) *Enum[string] {
  return NewEnum[string](stringMarshaler{}, enumValues...)
}

func NewIntEnum(enumValues ...int) *Enum[int] {
  return NewEnum[int](intMarshaler{}, enumValues...)
}

func NewInt64Enum(enumValues ...int64) *Enum[int64] {
  return NewEnum[int64](int64Marshaler{}, enumValues...)
}

func NewUintEnum(enumValues ...uint) *Enum[uint] {
  return NewEnum[uint](uintMarshaler{}, enumValues...)
}

func NewUint64Enum(enumValues ...uint64) *Enum[uint64] {
  return NewEnum[uint64](uint64Marshaler{}, enumValues...)
}

func NewDurationEnum(enumValues ...time.Duration) *Enum[time.Duration] {
  return NewEnum[time.Duration](durationMarshaler{}, enumValues...)
}
