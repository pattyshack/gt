package main

import (
  "fmt"

  "github.com/pattyshack/bt/argparse"
)

// Suggest values without validation.
type fileDescriptorType struct {
  *argparse.EnumType[string]
}

func (fdt fileDescriptorType) TypeDescription() string {
  return fdt.EnumType.TypeDescription() + " or a file descriptor"
}

func (fileDescriptorType) Validate(value string) error {
  return nil
}

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

  cmd.Var(
    argparse.NewValue[string](
      &cmd.fdPath,
      "stdout",
      &fileDescriptorType{
        EnumType: argparse.NewStringEnumType("stdin", "stdout", "stderr"),
      }),
    "file-descriptor-path",
    "file descriptor path")

  pets := argparse.NewStringEnumType(
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
      ValueType: pets,
    })
}

func (cmd *autoCompleteSubCmd) Execute(args []string) error {
  fmt.Println("Flags:")
  fmt.Println("  verbose level:", cmd.verboseLevel)
  fmt.Println("  file descriptor path:", cmd.fdPath)
  fmt.Println("Args:", args)
  return nil
}
