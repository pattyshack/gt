package main

import (
  "fmt"

  "github.com/pattyshack/bt/go/argparse"
)

type noPosArgsSubCmd struct {
  *mainCmd
  *argparse.Command

  subFlag bool
}

func (cmd *noPosArgsSubCmd) SetupCommand() {
  cmd.BoolVar(&cmd.subFlag, "sub-flag", false, "subcommand 1's bool flag")
  cmd.SetCommandFunc(cmd.Execute)
}

func (cmd *noPosArgsSubCmd) Execute(args []string) error {
  cmd.mainCmd.PrintFlags()
  fmt.Println()
  fmt.Println("sub1:")
  fmt.Println("=====")
  fmt.Println("sub1's flag:", cmd.subFlag)
  fmt.Println("sub1's args:", args)
  return nil
}

type posArgsSubCmd struct {
  *mainCmd
  *argparse.Command

  subFlag string
}

func (cmd *posArgsSubCmd) SetupCommand(varArgs bool) {
  cmd.StringVar(
    &cmd.subFlag,
    "sub-flag",
    "sub2",
    "subcommand 2's string flag",
    nil,
    nil)
  xEnum := argparse.NewStringEnum("x", "xx", "xxx", "xxxx")
  zEnum := argparse.NewStringEnum("z", "Z")
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
