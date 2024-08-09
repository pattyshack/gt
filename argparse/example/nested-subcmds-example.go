package main

import (
  "fmt"
  "time"

  "github.com/pattyshack/bt/argparse"
)

type nestedSubCmd struct {
  *mainCmd
  *argparse.Command

  subFlag int
  otherFlag string
}

func (cmd *nestedSubCmd) SetupCommand() {
  cmd.IntVar(
    &cmd.subFlag,
    "sub-flag",
    123,
    "subcommand 3's int flag")
  cmd.StringVar(
    &cmd.otherFlag,
    "other-flag",
    "foo",
    "subcommand 3's other flag")
}

func (cmd *nestedSubCmd) PrintFlags() {
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

type nestedSubSubCmd struct {
  parentCmd
  *argparse.Command

  name string

  nestedSubFlag time.Duration
}

func (cmd *nestedSubSubCmd) SetupCommand() {
  cmd.DurationVar(
    &cmd.nestedSubFlag,
    "nested-sub-flag",
    time.Second,
    "nested-sub-command 3's duration flag")

  cmd.SetCommandFunc(
    cmd.Execute,
    argparse.PositionalArgument{
      Name: "varargs",
      Description: "any string",
      NumExpected: 0,
      VarArgs: true,
    })
}

func (cmd *nestedSubSubCmd) Execute(args []string) error {
  cmd.parentCmd.PrintFlags()
  fmt.Println()
  fmt.Println("nested subcmd:", cmd.name)
  fmt.Println("==============")
  fmt.Println("nested subcmd's flag:", cmd.nestedSubFlag)
  fmt.Println("nested subcmd's args:", args)
  return nil
}
