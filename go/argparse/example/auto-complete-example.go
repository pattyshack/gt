package main

import (
  "fmt"

  "github.com/pattyshack/bt/go/argparse"
)

type autoCompleteSubCmd struct {
  *mainCmd
  *argparse.Command

  verboseLevel int
}

func (cmd *autoCompleteSubCmd) SetupCommand() {
  cmd.IntEnumVar(
    &cmd.verboseLevel,
    "verbose-level",
    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
    0,
    "Log verbose level")

  pets := argparse.NewStringEnum(
    "cat",
    "dog",
    "bunny",
    "hamster",
    "bird",
    "fish")
  cmd.SetCommandFunc(
    cmd.Execute,
    argparse.PositionalArgument{
      Name: "pet-type",
      Description: "A pet type",
      NumExpected: 1,
      ValueValidator: pets,
      ValueSuggestor: pets,
    })
}

func (cmd *autoCompleteSubCmd) Execute(args []string) error {
  fmt.Println("Flags:")
  fmt.Println("  verbose level:", cmd.verboseLevel)
  return nil
}
