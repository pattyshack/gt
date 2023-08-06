package main

import (
  "fmt"

  "github.com/pattyshack/bt/go/argparse"
  "github.com/pattyshack/bt/go/filesystem"
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
    "subcommand 2's string flag")
  xEnum := argparse.NewStringEnumType("x", "xx", "xxx", "xxxx")
  cmd.SetCommandFunc(
    cmd.Execute,
    argparse.PositionalArgument{
      Name: "x",
      Description: "first arg (3x)",
      NumExpected: 3,
      ValueType: xEnum,
    },
    argparse.PositionalArgument{
      Name: "y",
      Description: "second arg (2x)",
      NumExpected: 2,
      ValueType: argparse.NewFilePathType(filesystem.Local, true),
    },
    argparse.PositionalArgument{
      Name: "z",
      Description: "3rd arg (2x)",
      NumExpected: 2,
      VarArgs: varArgs,
      ValueType: argparse.NewFilePathType(filesystem.Local, false),
    })
}

func (cmd *posArgsSubCmd) Execute(args []string) error {
  cmd.mainCmd.PrintFlags()
  fmt.Println()
  fmt.Println("sub1:")
  fmt.Println("=====")
  fmt.Println("sub1's flag:", cmd.subFlag)
  fmt.Println("sub1's args:", args)
  return nil
}

