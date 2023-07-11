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
  argparse.BoolVar(&cmd.boolVarFlag, "bool-var",  true, "BoolVar()")
  cmd.boolFlag = argparse.Bool("bool",  false, "Bool()")

  argparse.StringVar(
    &cmd.stringVarFlag,
    "string-var",
    "string var",
    "StringVar()",
    nil,
    nil)
  cmd.stringFlag = argparse.String("string", "string", "String()", nil, nil)

  argparse.StringEnumVar(
    &cmd.stringEnumVarFlag,
    "string-enum-var",
    []string{"value1", "value2", "value3"},
    "value1",
    "StringEnumVar()")
  cmd.stringEnumFlag = argparse.StringEnum(
    "string-enum",
    []string{"ccc", "bbb", "aaa", "ddd"},
    "bbb",
    "StringEnum()")

  argparse.IntVar(
    &cmd.intVarFlag,
    "int-var",
    10,
    "IntVar()",
    nil,
    nil)
  cmd.intFlag = argparse.Int("int", 20, "Int()", nil, nil)

  argparse.IntEnumVar(
    &cmd.intEnumVarFlag,
    "int-enum-var",
    []int{1, 2, 3, 4, 5},
    3,
    "IntEnumVar()")
  cmd.intEnumFlag = argparse.IntEnum(
    "int-enum",
    []int{10, 20, 30, 40},
    20,
    "IntEnum()")

  argparse.Int64Var(
    &cmd.int64VarFlag,
    "int64-var",
    300,
    "Int64Var()",
    nil,
    nil)
  cmd.int64Flag = argparse.Int64("int64", 400, "Int64()", nil, nil)

  argparse.Int64EnumVar(
    &cmd.int64EnumVarFlag,
    "int64-enum-var",
    []int64{100, 200, 300, 400, 500},
    300,
    "Int64EnumVar()")
  cmd.int64EnumFlag = argparse.Int64Enum(
    "int64-enum",
    []int64{101, 201, 301, 401},
    201,
    "Int64Enum()")

  argparse.UintVar(
    &cmd.uintVarFlag,
    "uint-var",
    4000,
    "UintVar()",
    nil,
    nil)
  cmd.uintFlag = argparse.Uint("uint", 20, "Uint()", nil, nil)

  argparse.UintEnumVar(
    &cmd.uintEnumVarFlag,
    "uint-enum-var",
    []uint{1000, 2000, 3000, 4000, 5000},
    3000,
    "UintEnumVar()")
  cmd.uintEnumFlag = argparse.UintEnum(
    "uint-enum",
    []uint{1001, 2001, 3001, 4001},
    2001,
    "UintEnum()")

  argparse.Uint64Var(
    &cmd.uint64VarFlag,
    "uint64-var",
    50000,
    "Uint64Var()",
    nil,
    nil)
  cmd.uint64Flag = argparse.Uint64("uint64", 400, "Uint64()", nil, nil)

  argparse.Uint64EnumVar(
    &cmd.uint64EnumVarFlag,
    "uint64-enum-var",
    []uint64{10000, 20000, 30000, 40000, 50000},
    30000,
    "Uint64EnumVar()")
  cmd.uint64EnumFlag = argparse.Uint64Enum(
    "uint64-enum",
    []uint64{10001, 20001, 30001, 40001},
    20001,
    "Uint64Enum()")

  argparse.Float64Var(
    &cmd.float64VarFlag,
    "float64-var",
    3.14,
    "Float64Var()",
    nil,
    nil)
  cmd.float64Flag = argparse.Float64("float64", 2.15, "Float64()", nil, nil)

  argparse.DurationVar(
    &cmd.durationVarFlag,
    "duration-var",
    time.Minute,
    "Duration()",
    nil,
    nil)
  cmd.durationFlag = argparse.Duration(
    "duration",
    time.Hour,
    "Duration()",
    nil,
    nil)

  argparse.DurationEnumVar(
    &cmd.durationEnumVarFlag,
    "duration-enum-var",
    []time.Duration{time.Second, time.Minute, time.Hour},
    time.Second,
    "DurationEnumVar()")
  cmd.durationEnumFlag = argparse.DurationEnum(
    "duration-enum",
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
    "url-var",
    "test custom flag var",
    nil)
}

type subCmd1 struct {
  *mainCmd
  *argparse.Command

  subFlag bool
}

func (cmd *subCmd1) SetupCommand() {
  cmd.BoolVar(&cmd.subFlag, "sub-flag", false, "subcommand 1's bool flag")
  cmd.SetCommandFunc(cmd.Execute)
}

func (cmd *subCmd1) Execute(args []string) error {
  cmd.mainCmd.PrintFlags()
  fmt.Println()
  fmt.Println("sub1:")
  fmt.Println("=====")
  fmt.Println("sub1's flag:", cmd.subFlag)
  fmt.Println("sub1's args:", args)
  return nil
}

type subCmd2 struct {
  *mainCmd
  *argparse.Command

  subFlag string
}

func (cmd *subCmd2) SetupCommand(varArgs bool) {
  cmd.StringVar(
    &cmd.subFlag,
    "sub-flag",
    "sub2",
    "subcommand 2's string flag",
    nil,
    nil)
  xEnum := argparse.NewStringEnum([]string{"x", "xx", "xxx", "xxxx"})
  zEnum := argparse.NewStringEnum([]string{"z", "Z"})
  cmd.SetCommandFunc(
    cmd.Execute,
    argparse.PositionalArgument{
      Name: "x",
      Description: "first arg (3x)",
      NumExpected: 3,
      ValueValidator: xEnum,
      ValueSuggestor: xEnum,
    },
    argparse.PositionalArgument{
      Name: "y",
      Description: "second arg (2x)",
      NumExpected: 2,
    },
    argparse.PositionalArgument{
      Name: "z",
      Description: "3rd arg (2x)",
      NumExpected: 2,
      VarArgs: varArgs,
      ValueValidator: zEnum,
      ValueSuggestor: zEnum,
    })
}

func (cmd *subCmd2) Execute(args []string) error {
  cmd.mainCmd.PrintFlags()
  fmt.Println()
  fmt.Println("sub2:")
  fmt.Println("=====")
  fmt.Println("sub2's flag:", cmd.subFlag)
  fmt.Println("sub2's args:", args)
  return nil
}

type subCmd3 struct {
  *mainCmd
  *argparse.Command

  subFlag int
  otherFlag string
}

func (cmd *subCmd3) SetupCommand() {
  cmd.IntVar(
    &cmd.subFlag,
    "sub-flag",
    123,
    "subcommand 3's int flag",
    nil,
    nil)
  cmd.StringVar(
    &cmd.otherFlag,
    "other-flag",
    "foo",
    "subcommand 3's other flag",
    nil,
    nil)
}

func (cmd *subCmd3) PrintFlags() {
  cmd.mainCmd.PrintFlags()
  fmt.Println()
  fmt.Println("sub3:")
  fmt.Println("=====")
  fmt.Println("sub3's flag:", cmd.subFlag)
  fmt.Println("sub3's other flag:", cmd.otherFlag)
}

type parentCmd interface {
  PrintFlags()
}

type nestedSubCmd struct {
  parentCmd
  *argparse.Command

  name string

  nestedSubFlag time.Duration
}

func (cmd *nestedSubCmd) SetupCommand() {
  cmd.DurationVar(
    &cmd.nestedSubFlag,
    "nested-sub-flag",
    time.Second,
    "nested-sub-command 3's duration flag",
    nil,
    nil)

  cmd.SetCommandFunc(
    cmd.Execute,
    argparse.PositionalArgument{
      Name: "varargs",
      Description: "any string",
      NumExpected: 0,
      VarArgs: true,
    })
}

func (cmd *nestedSubCmd) Execute(args []string) error {
  cmd.parentCmd.PrintFlags()
  fmt.Println()
  fmt.Println("nested subcmd:", cmd.name)
  fmt.Println("==============")
  fmt.Println("nested subcmd's flag:", cmd.nestedSubFlag)
  fmt.Println("nested subcmd's args:", args)
  return nil
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

  subcmd1 := &subCmd1{
    mainCmd: cmd,
    Command: cmd.AddSubcommand(
      "noargs",
      "no subcommands, cmd func with no positional args"),
  }
  subcmd1.SetupCommand()

  subcmd2a := &subCmd2{
    mainCmd: cmd,
    Command: cmd.AddSubcommand(
      "fixed-args",
      "no subcommands, cmd func with positional args with no varargs"),
  }
  subcmd2a.SetupCommand(false)

  subcmd2b := &subCmd2{
    mainCmd: cmd,
    Command: cmd.AddSubcommand(
      "varargs",
      "no subcommands, cmd func with varargs"),
  }
  subcmd2b.SetupCommand(true)

  subcmd3 := &subCmd3{
    mainCmd: cmd,
    Command: cmd.AddSubcommand(
      "nested-subcommands",
      "nested subcommands, no cmd func"),
  }
  subcmd3.SetupCommand()

  subcmd3a := &nestedSubCmd{
    parentCmd: subcmd3,
    Command: subcmd3.AddSubcommand("3a", "some nested subcommand"),
    name: "3a",
  }
  subcmd3a.SetupCommand()

  subcmd3b := &nestedSubCmd{
    parentCmd: subcmd3,
    Command: subcmd3.AddSubcommand("3b", "some other nested subcommand"),
    name: "3b",
  }
  subcmd3b.SetupCommand()

  err := argparse.Execute()
  if err != nil {
    panic(err)
  }
}
