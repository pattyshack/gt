package argparse

import (
  "flag"
  "fmt"
  "io"
  "os"
  "sort"
  "strings"
  "time"
)

const (
  autoSuggestFlagName = "auto-suggest"
)

type boolFlag interface {
  IsBoolFlag() bool
}

type ValueSuggestor interface {
  Suggest(valuePrefix string) []string
}

type PositionalArgument struct {
  Name string
  Description string

  // Expected number of arguments of this type, or in VarArgs' case, the
  // expected minimum number.  Maybe zero when VarArgs is true; otherwise,
  // must be positive.
  NumExpected int

  // Can only be true for the last positional argument.
  // XXX: maybe relax this to at most one varargs per positional arg list.  We
  // can still easily validate the list on actual cmd execution, but this
  // causes combinatorial explosion for auto suggestion.
  VarArgs bool

  ValueValidator[string]  // optional
  ValueSuggestor  // optional
}

type Command struct {
  Name string
  Description string

  flagSet *flag.FlagSet

  autoSuggest bool  // predefined flag

  parentCommand *Command  // May be nil if it's a top level command.

  flagValueSuggestors map[string]ValueSuggestor

  subcommands map[string]*Command

  cmdFunc func([]string) error
  positionalArgs []PositionalArgument
}

func newCommand(
  name string,
  description string,
  flagSet *flag.FlagSet,
) *Command {
  if name == "" {
    panic("command name cannot be empty")
  }

  cmd := &Command{
    Name:name,
    Description: description,
    flagSet: flagSet,
    subcommands: map[string]*Command{},
  }

  flagSet.Usage = func() {
    fmt.Fprintf(flagSet.Output(), cmd.UsageString())
  }

  return cmd
}

func NewCommand(
  name string,
  description string,
  flagOutput io.Writer,
  errorHandling flag.ErrorHandling,
) *Command {
  flagSet := flag.NewFlagSet(name, errorHandling)
  flagSet.SetOutput(flagOutput)

  return newCommand(name, description, flagSet)
}

func (cmd *Command) UsageString() string {
  baseCmdForm := cmd.Name
  parentCmd := cmd.parentCommand
  for parentCmd != nil {
    baseCmdForm = parentCmd.Name + " " + baseCmdForm
    parentCmd = parentCmd.parentCommand
  }

  options := ""
  cmd.flagSet.VisitAll(func (flagDef *flag.Flag) {
    if flagDef.Name == autoSuggestFlagName {
      return
    }

    options += "  -" + flagDef.Name

    typeName := ""
    validValueDescription := ""
    typed, ok := flagDef.Value.(TypedGetter)
    if ok {
      typeName = typed.TypeName()

      descriptor, ok := typed.(WithValidValueDescription)
      if ok {
        validValueDescription = descriptor.ValidValueDescription()
      }
    } else {
      // This flag was defined using the flag lib directly (probably by a
      // thirdparty library).  Fallback to flag lib's guestimation.
      typeName, _ = flag.UnquoteUsage(flagDef)
      if typeName == "" {
        _, ok := flagDef.Value.(boolFlag)
        if ok {
          typeName = "bool"
        }
      }
    }

    if typeName != "" {
      options += " " + typeName
    }

    if validValueDescription != "" {
      options += fmt.Sprintf(
        " (default: %#v; valid values: %s)\n",
        flagDef.DefValue,
        validValueDescription)
    } else {
      options += fmt.Sprintf(" (default: %#v)\n", flagDef.DefValue)
    }

    if flagDef.Usage != "" {
      options += "\t" + strings.ReplaceAll(flagDef.Usage, "\n", "\n\t") + "\n"
    }
  })

  if len(options) > 0 {
    baseCmdForm += " [options]"
    options = "\nOptions:\n" + options
  }

  subcmds := ""
  subcommandForm := ""
  if len(cmd.subcommands) > 0 {
    subcommandForm = baseCmdForm + " <subcommand> ..."

    sortedSubcommands := make([]*Command, 0, len(cmd.subcommands))
    for _, sub := range cmd.subcommands {
      sortedSubcommands = append(sortedSubcommands, sub)
    }

    sort.Slice(
      sortedSubcommands,
      func (i int, j int) bool {
        return sortedSubcommands[i].Name < sortedSubcommands[j].Name
      })

    subcmds = "\nSubcommands:\n"
    for _, sub := range sortedSubcommands {
      subcmds += "  " + sub.Name + "\n"
      if sub.Description != "" {
        desc := strings.ReplaceAll(sub.Description, "\n", "\n\t") 
        subcmds += "\t" + desc + "\n"
      }
    }
  }

  posArgs := ""
  posArgsForm := ""
  if cmd.cmdFunc != nil {
    posArgsForm = baseCmdForm
    posArgs = "\nPositional arguments:\n"

    for _, arg := range cmd.positionalArgs {
      for i := 0; i < arg.NumExpected; i++ {
        posArgsForm += " <" + arg.Name + ">"
      }

      if arg.VarArgs {
        posArgsForm += " [<" + arg.Name + "> ...]"
      }

      if arg.ValueValidator == nil {
        posArgs += "  " + arg.Name + "\n"
      } else {
        posArgs += fmt.Sprintf(
          "  %s (valid values: %s)\n",
          arg.Name,
          arg.ValueValidator.ValidValueDescription())
      }

      if arg.Description != "" {
        desc := strings.ReplaceAll(arg.Description, "\n", "\n\t")
        posArgs += "\t" + desc + "\n"
      }
    }
  }

  usage := "USAGE: "
  if subcommandForm == "" && posArgsForm == "" {
    // subcommands / cmdFunc not specified.  This will error on actual
    // execution.
    usage += baseCmdForm + "\n"
  }

  if subcommandForm != "" {
    usage += subcommandForm + "\n"

    if posArgsForm != "" {
      usage += "  or\n"
    }
  }

  if posArgsForm != "" {
    usage += posArgsForm + "\n"
  }

  if cmd.Description != "" {
    usage += "\n" + cmd.Description + "\n"
  }

  usage += subcmds + posArgs + options

  return usage
}

// AddSubcommand associates a new subcommand to the command.
func (cmd *Command) AddSubcommand(name string, description string) *Command {
  _, ok := cmd.subcommands[name]
  if ok {
    panic("duplicate subcommand name specified: " + name)
  }

  subcommand := NewCommand(
    name,
    description,
    cmd.flagSet.Output(),
    cmd.flagSet.ErrorHandling())
  subcommand.parentCommand = cmd

  cmd.subcommands[name] = subcommand

  return subcommand
}

// SetCommandFunc specifies the function associated with the command and the
// command's expected positional arguments.  This should be called at most once.
//
// In general, command func should not be specified when subcommands are
// specified.  If both command func and subcommands are specified, subcommands
// take higher priority than command func.
func (cmd *Command) SetCommandFunc(
  cmdFunc func([]string) error,
  positionalArgs ...PositionalArgument,
) {
  if cmdFunc == nil {
    panic("cmdFunc cannot be nil")
  }

  for idx, arg := range positionalArgs {
    minNumExpected := 1
    if arg.VarArgs {
      if idx + 1 != len(positionalArgs) {
        panic(fmt.Sprintf("invalid VarArgs: %s (%d)", arg.Name, idx))
      }

      minNumExpected = 0
    }

    if arg.NumExpected < minNumExpected {
      panic(fmt.Sprintf("invalid NumExpected: %s (%d)", arg.Name, idx))
    }
  }

  cmd.cmdFunc = cmdFunc
  cmd.positionalArgs = positionalArgs
}

func (cmd *Command) Var(
  value TypedGetter,
  name string,
  description string,
  suggestor ValueSuggestor,  // optional
) {
  cmd.flagSet.Var(value, name, description)

  // TODO suggestor
}

func (cmd *Command) BoolVar(
  ptr *bool,
  name string,
  value bool,
  description string,
) {
  // TODO bool suggestor
  cmd.Var(
    NewBoolValue(ptr, value, description),
    name,
    description,
    nil)
}

func (cmd *Command) Bool(
  name string,
  value bool,
  description string,
) *bool {
  ptr := new(bool)
  cmd.BoolVar(ptr, name, value, description)
  return ptr
}

func (cmd *Command) StringVar(
  ptr *string,
  name string,
  value string,
  description string,
  validator ValueValidator[string],  // optional
  suggestor ValueSuggestor,  // optional
) {
  cmd.Var(
    NewValue[string](ptr, value, stringMarshaler{}, validator),
    name,
    description,
    suggestor)
}

func (cmd *Command) String(
  name string,
  value string,
  description string,
  validator ValueValidator[string],  // optional
  suggestor ValueSuggestor,  // optional
) *string {
  ptr := new(string)
  cmd.StringVar(ptr, name, value, description, validator, suggestor)
  return ptr
}

func (cmd *Command) StringEnumVar(
  ptr *string,
  name string,
  enumValues []string,
  value string,
  description string,
) {
  enum := NewStringEnum(enumValues)
  cmd.Var(
    NewValue[string](ptr, value, stringMarshaler{}, enum),
    name,
    description,
    enum)
}

func (cmd *Command) StringEnum(
  name string,
  enumValues []string,
  value string,
  description string,
) *string {
  ptr := new(string)
  cmd.StringEnumVar(ptr, name, enumValues, value, description)
  return ptr
}

func (cmd *Command) IntVar(
  ptr *int,
  name string,
  value int,
  description string,
  validator ValueValidator[int],  // optional
  suggestor ValueSuggestor,  // optional
) {
  cmd.Var(
    NewValue[int](ptr, value, intMarshaler{}, validator),
    name,
    description,
    suggestor)
}

func (cmd *Command) Int(
  name string,
  value int,
  description string,
  validator ValueValidator[int],  // optional
  suggestor ValueSuggestor,  // optional
) *int {
  ptr := new(int)
  cmd.IntVar(ptr, name, value, description, validator, suggestor)
  return ptr
}

func (cmd *Command) IntEnumVar(
  ptr *int,
  name string,
  enumValues []int,
  value int,
  description string,
) {
  enum := NewIntEnum(enumValues)
  cmd.Var(
    NewValue[int](ptr, value, intMarshaler{}, enum),
    name,
    description,
    enum)
}

func (cmd *Command) IntEnum(
  name string,
  enumValues []int,
  value int,
  description string,
) *int {
  ptr := new(int)
  cmd.IntEnumVar(ptr, name, enumValues, value, description)
  return ptr
}

func (cmd *Command) Int64Var(
  ptr *int64,
  name string,
  value int64,
  description string,
  validator ValueValidator[int64],  // optional
  suggestor ValueSuggestor,  // optional
) {
  cmd.Var(
    NewValue[int64](ptr, value, int64Marshaler{}, validator),
    name,
    description,
    suggestor)
}

func (cmd *Command) Int64(
  name string,
  value int64,
  description string,
  validator ValueValidator[int64],  // optional
  suggestor ValueSuggestor,  // optional
) *int64 {
  ptr := new(int64)
  cmd.Int64Var(ptr, name, value, description, validator, suggestor)
  return ptr
}

func (cmd *Command) Int64EnumVar(
  ptr *int64,
  name string,
  enumValues []int64,
  value int64,
  description string,
) {
  enum := NewInt64Enum(enumValues)
  cmd.Var(
    NewValue[int64](ptr, value, int64Marshaler{}, enum),
    name,
    description,
    enum)
}

func (cmd *Command) Int64Enum(
  name string,
  enumValues []int64,
  value int64,
  description string,
) *int64 {
  ptr := new(int64)
  cmd.Int64EnumVar(ptr, name, enumValues, value, description)
  return ptr
}

func (cmd *Command) UintVar(
  ptr *uint,
  name string,
  value uint,
  description string,
  validator ValueValidator[uint],  // optional
  suggestor ValueSuggestor,  // optional
) {
  cmd.Var(
    NewValue[uint](ptr, value, uintMarshaler{}, validator),
    name,
    description,
    suggestor)
}

func (cmd *Command) Uint(
  name string,
  value uint,
  description string,
  validator ValueValidator[uint],  // optional
  suggestor ValueSuggestor,  // optional
) *uint {
  ptr := new(uint)
  cmd.UintVar(ptr, name, value, description, validator, suggestor)
  return ptr
}

func (cmd *Command) UintEnumVar(
  ptr *uint,
  name string,
  enumValues []uint,
  value uint,
  description string,
) {
  enum := NewUintEnum(enumValues)
  cmd.Var(
    NewValue[uint](ptr, value, uintMarshaler{}, enum),
    name,
    description,
    enum)
}

func (cmd *Command) UintEnum(
  name string,
  enumValues []uint,
  value uint,
  description string,
) *uint {
  ptr := new(uint)
  cmd.UintEnumVar(ptr, name, enumValues, value, description)
  return ptr
}

func (cmd *Command) Uint64Var(
  ptr *uint64,
  name string,
  value uint64,
  description string,
  validator ValueValidator[uint64],  // optional
  suggestor ValueSuggestor,  // optional
) {
  cmd.Var(
    NewValue[uint64](ptr, value, uint64Marshaler{}, validator),
    name,
    description,
    suggestor)
}

func (cmd *Command) Uint64(
  name string,
  value uint64,
  description string,
  validator ValueValidator[uint64],  // optional
  suggestor ValueSuggestor,  // optional
) *uint64 {
  ptr := new(uint64)
  cmd.Uint64Var(ptr, name, value, description, validator, suggestor)
  return ptr
}

func (cmd *Command) Uint64EnumVar(
  ptr *uint64,
  name string,
  enumValues []uint64,
  value uint64,
  description string,
) {
  enum := NewUint64Enum(enumValues)
  cmd.Var(
    NewValue[uint64](ptr, value, uint64Marshaler{}, enum),
    name,
    description,
    enum)
}

func (cmd *Command) Uint64Enum(
  name string,
  enumValues []uint64,
  value uint64,
  description string,
) *uint64 {
  ptr := new(uint64)
  cmd.Uint64EnumVar(ptr, name, enumValues, value, description)
  return ptr
}

func (cmd *Command) Float64Var(
  ptr *float64,
  name string,
  value float64,
  description string,
  validator ValueValidator[float64],  // optional
  suggestor ValueSuggestor,  // optional
) {
  cmd.Var(
    NewValue[float64](ptr, value, float64Marshaler{}, validator),
    name,
    description,
    suggestor)
}

func (cmd *Command) Float64(
  name string,
  value float64,
  description string,
  validator ValueValidator[float64],  // optional
  suggestor ValueSuggestor,  // optional
) *float64 {
  ptr := new(float64)
  cmd.Float64Var(ptr, name, value, description, validator, suggestor)
  return ptr
}

func (cmd *Command) DurationVar(
  ptr *time.Duration,
  name string,
  value time.Duration,
  description string,
  validator ValueValidator[time.Duration],  // optional
  suggestor ValueSuggestor,  // optional
) {
  cmd.Var(
    NewValue[time.Duration](ptr, value, durationMarshaler{}, validator),
    name,
    description,
    suggestor)
}

func (cmd *Command) Duration(
  name string,
  value time.Duration,
  description string,
  validator ValueValidator[time.Duration],  // optional
  suggestor ValueSuggestor,  // optional
) *time.Duration {
  ptr := new(time.Duration)
  cmd.DurationVar(ptr, name, value, description, validator, suggestor)
  return ptr
}

func (cmd *Command) DurationEnumVar(
  ptr *time.Duration,
  name string,
  enumValues []time.Duration,
  value time.Duration,
  description string,
) {
  enum := NewDurationEnum(enumValues)
  cmd.Var(
    NewValue[time.Duration](ptr, value, durationMarshaler{}, enum),
    name,
    description,
    enum)
}

func (cmd *Command) DurationEnum(
  name string,
  enumValues []time.Duration,
  value time.Duration,
  description string,
) *time.Duration {
  ptr := new(time.Duration)
  cmd.DurationEnumVar(ptr, name, enumValues, value, description)
  return ptr
}

func (cmd *Command) Suggest(args []string) {
  // TODO
}

func (cmd *Command) Execute(args []string) error {
  err := cmd.flagSet.Parse(args)
  if err != nil {
    return fmt.Errorf("flag parse error: %w", err)
  }

  parsedArgs := cmd.flagSet.Args()

  if cmd.autoSuggest {
    cmd.Suggest(parsedArgs)
    return nil
  }

  actionFunc, parsedArgs, err := cmd.parseAction(parsedArgs)
  if err != nil {
    // Make action parsing error behave the same way as flag parsing error.
    fmt.Fprintln(cmd.flagSet.Output(), err.Error())
    cmd.flagSet.Usage()
    switch cmd.flagSet.ErrorHandling() {
    case flag.ExitOnError:
      os.Exit(2)
    case flag.PanicOnError:
      panic(err)
    }
    return err
  }

  return actionFunc(parsedArgs)
}

func (cmd *Command) parseAction(args []string) (
  func([]string) error,
  []string,
  error,
) {
  if len(cmd.subcommands) == 0 && cmd.cmdFunc == nil {
    // This is a programming error
    panic("subcommand/command function not defined for command:" + cmd.Name)
  }

  if len(args) > 0 {
    subcmd, ok := cmd.subcommands[args[0]]
    if ok {
      return subcmd.Execute, args[1:], nil
    }
  }

  if cmd.cmdFunc == nil {
    return nil, nil, fmt.Errorf("no subcommand specified")
  }

  err := cmd.validateCommandFunc(args)
  if err != nil {
    return nil, nil, err
  }

  return cmd.cmdFunc, args, nil
}

func (cmd *Command) validateCommandFunc(args []string) error {
  totalExpected := 0
  hasVarArgs := false
  for _, param := range cmd.positionalArgs {
    totalExpected += param.NumExpected
    if param.VarArgs {
      hasVarArgs = true
    }
  }

  if len(args) < totalExpected || (!hasVarArgs && len(args) > totalExpected) {
    return fmt.Errorf("unexpected number of positional arguments.")
  }

  idx := 0
  for _, param := range cmd.positionalArgs {
    if param.ValueValidator == nil {
      if param.VarArgs {  // i.e., last arg
        break  // no need to check the remaining arguments
      }

      idx += param.NumExpected
      continue
    }

    numExpected := param.NumExpected
    if param.VarArgs {
      numExpected = len(args) - idx
    }

    for i := 0; i < numExpected; i++ {
      err := param.ValueValidator.Validate(args[idx])
      if err != nil {
        return fmt.Errorf(
          "invalid value for argument <%s> at positional index %d (%#v): %w",
          param.Name,
          idx,
          args[idx],
          err)
      }
      idx++
    }
  }

  return nil
}
