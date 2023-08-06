package argparse

import (
  "flag"
  "path/filepath"
  "os"
  "time"
)

var CommandLine = newCommand(
  filepath.Base(os.Args[0]),
  "",
  flag.CommandLine)

func AddSubcommand(name string, description string) *Command {
  return CommandLine.AddSubcommand(name, description)
}

func Var(
  value flag.Getter,
  name string,
  description string,
) {
  CommandLine.Var(value, name, description)
}

func BoolVar(
  ptr *bool,
  name string,
  value bool,
  description string,
) {
  CommandLine.BoolVar(ptr, name, value, description)
}

func Bool(
  name string,
  value bool,
  description string,
) *bool {
  return CommandLine.Bool(name, value, description)
}


func StringVar(
  ptr *string,
  name string,
  value string,
  description string,
) {
  CommandLine.StringVar(ptr, name, value, description)
}

func String(
  name string,
  value string,
  description string,
) *string {
  return CommandLine.String(name, value, description)
}

func StringEnumVar(
  ptr *string,
  name string,
  enumValues []string,
  value string,
  description string,
) {
  CommandLine.StringEnumVar(ptr, name, enumValues, value, description)
}

func StringEnum(
  name string,
  enumValues []string,
  value string,
  description string,
) *string {
  return CommandLine.StringEnum(name, enumValues, value, description)
}

func IntVar(
  ptr *int,
  name string,
  value int,
  description string,
) {
  CommandLine.IntVar(ptr, name, value, description)
}

func Int(
  name string,
  value int,
  description string,
) *int {
  return CommandLine.Int(name, value, description)
}

func IntEnumVar(
  ptr *int,
  name string,
  enumValues []int,
  value int,
  description string,
) {
  CommandLine.IntEnumVar(ptr, name, enumValues, value, description)
}

func IntEnum(
  name string,
  enumValues []int,
  value int,
  description string,
) *int {
  return CommandLine.IntEnum(name, enumValues, value, description)
}

func Int64Var(
  ptr *int64,
  name string,
  value int64,
  description string,
) {
  CommandLine.Int64Var(ptr, name, value, description)
}

func Int64(
  name string,
  value int64,
  description string,
) *int64 {
  return CommandLine.Int64(name, value, description)
}

func Int64EnumVar(
  ptr *int64,
  name string,
  enumValues []int64,
  value int64,
  description string,
) {
  CommandLine.Int64EnumVar(ptr, name, enumValues, value, description)
}

func Int64Enum(
  name string,
  enumValues []int64,
  value int64,
  description string,
) *int64 {
  return CommandLine.Int64Enum(name, enumValues, value, description)
}

func UintVar(
  ptr *uint,
  name string,
  value uint,
  description string,
) {
  CommandLine.UintVar(ptr, name, value, description)
}

func Uint(
  name string,
  value uint,
  description string,
) *uint {
  return CommandLine.Uint(name, value, description)
}

func UintEnumVar(
  ptr *uint,
  name string,
  enumValues []uint,
  value uint,
  description string,
) {
  CommandLine.UintEnumVar(ptr, name, enumValues, value, description)
}

func UintEnum(
  name string,
  enumValues []uint,
  value uint,
  description string,
) *uint {
  return CommandLine.UintEnum(name, enumValues, value, description)
}

func Uint64Var(
  ptr *uint64,
  name string,
  value uint64,
  description string,
) {
  CommandLine.Uint64Var(ptr, name, value, description)
}

func Uint64(
  name string,
  value uint64,
  description string,
) *uint64 {
  return CommandLine.Uint64(name, value, description)
}

func Uint64EnumVar(
  ptr *uint64,
  name string,
  enumValues []uint64,
  value uint64,
  description string,
) {
  CommandLine.Uint64EnumVar(ptr, name, enumValues, value, description)
}

func Uint64Enum(
  name string,
  enumValues []uint64,
  value uint64,
  description string,
) *uint64 {
  return CommandLine.Uint64Enum(name, enumValues, value, description)
}

func Float64Var(
  ptr *float64,
  name string,
  value float64,
  description string,
) {
  CommandLine.Float64Var(ptr, name, value, description)
}

func Float64(
  name string,
  value float64,
  description string,
) *float64 {
  return CommandLine.Float64(name, value, description)
}

func DurationVar(
  ptr *time.Duration,
  name string,
  value time.Duration,
  description string,
) {
  CommandLine.DurationVar(ptr, name, value, description)
}

func Duration(
  name string,
  value time.Duration,
  description string,
) *time.Duration {
  return CommandLine.Duration(name, value, description)
}

func DurationEnumVar(
  ptr *time.Duration,
  name string,
  enumValues []time.Duration,
  value time.Duration,
  description string,
) {
  CommandLine.DurationEnumVar(ptr, name, enumValues, value, description)
}

func DurationEnum(
  name string,
  enumValues []time.Duration,
  value time.Duration,
  description string,
) *time.Duration {
  return CommandLine.DurationEnum(name, enumValues, value, description)
}

func Execute() error {
  return CommandLine.Execute(os.Args[1:])
}

