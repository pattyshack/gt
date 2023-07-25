package main

import (
  "flag"
  "fmt"
  "time"
  "net/url"

  "github.com/pattyshack/bt/go/argparse"
)

type urlMarshaler struct {}

func (urlMarshaler) Marshal(url url.URL) string {
  return url.String()
}

func (urlMarshaler) Unmarshal(value string) (url.URL, error) {
  u, err := url.Parse(value)
  if err != nil {
    return url.URL{}, err
  }

  return *u, nil
}

type mainCmd struct {
  *argparse.Command

  externalLibFlag string

  boolVarFlag bool
  stringVarFlag string
  stringEnumVarFlag string
  intVarFlag int
  intEnumVarFlag int
  int64VarFlag int64
  int64EnumVarFlag int64
  uintVarFlag uint
  uintEnumVarFlag uint
  uint64VarFlag uint64
  uint64EnumVarFlag uint64
  float64VarFlag float64
  durationVarFlag time.Duration
  durationEnumVarFlag time.Duration

  boolFlag *bool
  stringFlag *string
  stringEnumFlag *string
  intFlag *int
  intEnumFlag *int
  int64Flag *int64
  int64EnumFlag *int64
  uintFlag *uint
  uintEnumFlag *uint
  uint64Flag *uint64
  uint64EnumFlag *uint64
  float64Flag *float64
  durationFlag *time.Duration
  durationEnumFlag *time.Duration

  urlFlag url.URL
}

func (cmd *mainCmd) PrintFlags() {
  fmt.Println("main flags:")
  fmt.Println("===========")
  fmt.Println("bool flags:", cmd.boolVarFlag, *cmd.boolFlag)
  fmt.Println("string flags:", cmd.stringVarFlag, *cmd.stringFlag)
  fmt.Println("string enum flags:", cmd.stringEnumVarFlag, *cmd.stringEnumFlag)
  fmt.Println("int flags:", cmd.intVarFlag, *cmd.intFlag)
  fmt.Println("int enum flags:", cmd.intEnumVarFlag, *cmd.intEnumFlag)
  fmt.Println("int64 flags:", cmd.int64VarFlag, *cmd.int64Flag)
  fmt.Println("int64 enum flags:", cmd.int64EnumVarFlag, *cmd.int64EnumFlag)
  fmt.Println("uint flags:", cmd.uintVarFlag, *cmd.uintFlag)
  fmt.Println("uint enum flags:", cmd.uintEnumVarFlag, *cmd.uintEnumFlag)
  fmt.Println("uint64 flags:", cmd.uint64VarFlag, *cmd.uint64Flag)
  fmt.Println("uint64 enum flags:", cmd.uint64EnumVarFlag, *cmd.uint64EnumFlag)
  fmt.Println("float64 flags:", cmd.float64VarFlag, *cmd.float64Flag)
  fmt.Println("duration flags:", cmd.durationVarFlag, *cmd.durationFlag)
  fmt.Println(
    "duration enum flags:",
    cmd.durationEnumVarFlag,
    *cmd.durationEnumFlag)

  fmt.Println("url flag:", cmd.urlFlag)
}

func (cmd *mainCmd) SetupCommand() {
  argparse.BoolVar(&cmd.boolVarFlag, "bool.var-flag",  true, "BoolVar()")
  cmd.boolFlag = argparse.Bool("bool.flag",  false, "Bool()")

  argparse.StringVar(
    &cmd.stringVarFlag,
    "string.var-flag",
    "string var",
    "StringVar()",
    nil,
    nil)
  cmd.stringFlag = argparse.String(
    "string.flag",
    "string",
    "String()",
    nil,
    nil)

  argparse.StringEnumVar(
    &cmd.stringEnumVarFlag,
    "string.enum.var-flag",
    []string{"value1", "value2", "value3"},
    "value1",
    "StringEnumVar()")
  cmd.stringEnumFlag = argparse.StringEnum(
    "string.enum.flag",
    []string{"ccc", "bbb", "aaa", "ddd"},
    "bbb",
    "StringEnum()")

  argparse.IntVar(
    &cmd.intVarFlag,
    "int.var-flag",
    10,
    "IntVar()",
    nil,
    nil)
  cmd.intFlag = argparse.Int("int.flag", 20, "Int()", nil, nil)

  argparse.IntEnumVar(
    &cmd.intEnumVarFlag,
    "int.enum.var-flag",
    []int{1, 2, 3, 4, 5},
    3,
    "IntEnumVar()")
  cmd.intEnumFlag = argparse.IntEnum(
    "int.enum.flag",
    []int{10, 20, 30, 40},
    20,
    "IntEnum()")

  argparse.Int64Var(
    &cmd.int64VarFlag,
    "int64.var-flag",
    300,
    "Int64Var()",
    nil,
    nil)
  cmd.int64Flag = argparse.Int64("int64.flag", 400, "Int64()", nil, nil)

  argparse.Int64EnumVar(
    &cmd.int64EnumVarFlag,
    "int64.enum.var-flag",
    []int64{100, 200, 300, 400, 500},
    300,
    "Int64EnumVar()")
  cmd.int64EnumFlag = argparse.Int64Enum(
    "int64.enum.flag",
    []int64{101, 201, 301, 401},
    201,
    "Int64Enum()")

  argparse.UintVar(
    &cmd.uintVarFlag,
    "uint.var-flag",
    4000,
    "UintVar()",
    nil,
    nil)
  cmd.uintFlag = argparse.Uint("uint.flag", 20, "Uint()", nil, nil)

  argparse.UintEnumVar(
    &cmd.uintEnumVarFlag,
    "uint.enum.var-flag",
    []uint{1000, 2000, 3000, 4000, 5000},
    3000,
    "UintEnumVar()")
  cmd.uintEnumFlag = argparse.UintEnum(
    "uint.enum.flag",
    []uint{1001, 2001, 3001, 4001},
    2001,
    "UintEnum()")

  argparse.Uint64Var(
    &cmd.uint64VarFlag,
    "uint64.var-flag",
    50000,
    "Uint64Var()",
    nil,
    nil)
  cmd.uint64Flag = argparse.Uint64("uint64.flag", 400, "Uint64()", nil, nil)

  argparse.Uint64EnumVar(
    &cmd.uint64EnumVarFlag,
    "uint64.enum.var-flag",
    []uint64{10000, 20000, 30000, 40000, 50000},
    30000,
    "Uint64EnumVar()")
  cmd.uint64EnumFlag = argparse.Uint64Enum(
    "uint64.enum.flag",
    []uint64{10001, 20001, 30001, 40001},
    20001,
    "Uint64Enum()")

  argparse.Float64Var(
    &cmd.float64VarFlag,
    "float64.var-flag",
    3.14,
    "Float64Var()",
    nil,
    nil)
  cmd.float64Flag = argparse.Float64(
    "float64.flag",
    2.15,
    "Float64()",
    nil,
    nil)

  argparse.DurationVar(
    &cmd.durationVarFlag,
    "duration.var-flag",
    time.Minute,
    "Duration()",
    nil,
    nil)
  cmd.durationFlag = argparse.Duration(
    "duration.flag",
    time.Hour,
    "Duration()",
    nil,
    nil)

  argparse.DurationEnumVar(
    &cmd.durationEnumVarFlag,
    "duration.enum.var-flag",
    []time.Duration{time.Second, time.Minute, time.Hour},
    time.Second,
    "DurationEnumVar()")
  cmd.durationEnumFlag = argparse.DurationEnum(
    "duration.enum.flag",
    []time.Duration{10 *time.Second, 10 *time.Minute, 10 * time.Hour},
    10 *time.Second,
    "DurationEnum()")

  defaultUrl, err := url.Parse("https://github.com/pattyshack/bt")
  if err != nil {
    panic(err)
  }
  argparse.Var(
    argparse.NewValue[url.URL](
      &cmd.urlFlag,
      *defaultUrl,
      urlMarshaler{},
      nil),
    "url.var-flag",
    "test custom flag var",
    nil)
}

func main() {
  cmd := &mainCmd{
    Command: argparse.CommandLine,
  }

  // simulate flag defined by thirdparty library.
  flag.StringVar(&cmd.externalLibFlag, "external-lib-flag", "", "")

  argparse.CommandLine.Description = "Example usage of the argparse library.\n"+
    "Different subcommands demo different features of library."

  cmd.SetupCommand()

  noPosArgs := &noPosArgsSubCmd{
    mainCmd: cmd,
    Command: cmd.AddSubcommand(
      "noargs",
      "no subcommands, cmd func with no positional args"),
  }
  noPosArgs.SetupCommand()

  fixedPosArgs := &posArgsSubCmd{
    mainCmd: cmd,
    Command: cmd.AddSubcommand(
      "fixed-args",
      "no subcommands, cmd func with positional args with no varargs"),
  }
  fixedPosArgs.SetupCommand(false)

  varPosArgs := &posArgsSubCmd{
    mainCmd: cmd,
    Command: cmd.AddSubcommand(
      "varargs",
      "no subcommands, cmd func with varargs"),
  }
  varPosArgs.SetupCommand(true)

  nested := &nestedSubCmd{
    mainCmd: cmd,
    Command: cmd.AddSubcommand(
      "nested-subcommands",
      "nested subcommands, no cmd func"),
  }
  nested.SetupCommand()

  nestedSub1 := &nestedSubSubCmd{
    parentCmd: nested,
    Command: nested.AddSubcommand("sub1", "some nested subcommand"),
    name: "sub1",
  }
  nestedSub1.SetupCommand()

  nestedSub2 := &nestedSubSubCmd{
    parentCmd: nested,
    Command: nested.AddSubcommand("sub2", "some other nested subcommand"),
    name: "sub2",
  }
  nestedSub2.SetupCommand()

  autoComplete := &autoCompleteSubCmd{
    mainCmd: cmd,
    Command: cmd.AddSubcommand(
      "auto-complete-example",
      "Auto complete example. All flag/positional args are auto-completeable"),
  }
  autoComplete.SetupCommand()

  err := argparse.Execute()
  if err != nil {
    panic(err)
  }
}
