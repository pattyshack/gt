package main

import (
  "fmt"

  "github.com/pattyshack/bt/go/argparse"
)

type autoCompleteSubCmd struct {
  *mainCmd
  *argparse.Command

  verboseLevel int
  fdPath string
}

func (cmd *autoCompleteSubCmd) SetupCommand() {
  cmd.IntEnumVar(
    &cmd.verboseLevel,
    "verbose-level",
    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
    0,
    "Log verbose level")

  // Suggest values without value validation.
  cmd.StringVar(
    &cmd.fdPath,
    "file-descriptor-path",
    "stdout",
    "file descriptor path",
    nil,
    argparse.NewStringEnum("stdin", "stdout", "stderr"))

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
  fmt.Println("  file descriptor path:", cmd.fdPath)
  return nil
}
