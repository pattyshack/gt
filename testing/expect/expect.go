package expect

import (
  "bytes"
  "fmt"
  "path"
  "reflect"
  "runtime"
  "strings"
  "testing"
)

func formatValue(value any) string {
  if value == nil {
    return "nil"
  }

  switch reflect.ValueOf(value).Type().Kind() {
  case reflect.Uintptr, reflect.Pointer, reflect.UnsafePointer:
    return fmt.Sprintf("%p %#v (%v)", value, value, value)
  }
  return fmt.Sprintf("%#v (%v)", value, value)
}

func writeTrace(buffer *bytes.Buffer) {
  currentPC, _, _, recoverable := runtime.Caller(0)
  if !recoverable {
    return
  }

  currentFunc := runtime.FuncForPC(currentPC)
  if currentFunc == nil {
    return
  }

  if path.Ext(currentFunc.Name()) != ".writeTrace" {
    panic("This should never occur")
  }

  testingParentPkg := path.Dir(currentFunc.Name())  // exclude "expect"

  if path.Base(testingParentPkg) != "testing" {
    panic("This should never occur")
  }

  skip := 1
  first := true
  for {
    pc, file, line, recoverable := runtime.Caller(skip)
    if !recoverable {
      break
    }

    skip++

    funcEntry := runtime.FuncForPC(pc)
    if funcEntry == nil {
      break
    }

    fullName := funcEntry.Name()

    parentPkg := path.Dir(fullName)
    name := path.Base(fullName)

    // Skip printing out golang's test runner
    if parentPkg == "." && strings.HasPrefix(name, "testing.") {
      break
    }

    if parentPkg == testingParentPkg {
      pkgName, _, _ := strings.Cut(name, ".")
      if !strings.HasSuffix(pkgName, "_test") {
        continue
      }
    }

    if first {
      buffer.WriteString("\nTrace:")
      first = false
    }

    buffer.WriteString("\n\t")
    buffer.WriteString(funcEntry.Name())
    buffer.WriteString("\n\t\t")
    buffer.WriteString(file)
    buffer.WriteString(":")
    buffer.WriteString(fmt.Sprintf("%d", line))
  }
}

func logFailure(
  tb testing.TB,
  additionalMsg []any,
  template string,
  args ...any,
) {
  tb.Helper()

  buffer := bytes.NewBuffer(nil)
  buffer.WriteString("Error: ")
  buffer.WriteString(fmt.Sprintf(template, args...))

  if len(additionalMsg) > 0 {
    buffer.WriteString("\nMessage: ")
    buffer.WriteString(
      fmt.Sprintf(additionalMsg[0].(string), additionalMsg[1:]...))
  }

  writeTrace(buffer)

  tb.Fatalf(string(buffer.Bytes()))
}

func Same[T comparable](
  tb testing.TB,
  expected T,
  actual T,
  additionalMsg ...any,
) {
  if expected != actual {
    logFailure(
      tb,
      additionalMsg,
      "Expected same (==)\nExpected: %s\nActual:   %s",
      formatValue(expected),
      formatValue(actual))
  }
}

func Equal[T any](
  tb testing.TB,
  expected T,
  actual T,
  additionalMsg ...any,
) {
  tb.Helper()

  if !reflect.DeepEqual(expected, actual) {
    logFailure(
      tb,
      additionalMsg,
      "Expected deep equal\nExpected: %s\nActual:   %s",
      formatValue(expected),
      formatValue(actual))
  }
}

func NotEqual[T any](
  tb testing.TB,
  expected T,
  actual T,
  additionalMsg ...any,
) {
  tb.Helper()

  if reflect.DeepEqual(expected, actual) {
    logFailure(
      tb,
      additionalMsg,
      "Expected not deep equal\nActual: %s",
      formatValue(actual))
  }
}

func isNil(value any) bool {
  if value == nil {  // deal with untyped nil
    return true
  }

  rValue := reflect.ValueOf(value)
  switch rValue.Type().Kind() {
  case reflect.Uintptr, reflect.Chan, reflect.Func, reflect.Interface,
      reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:

    return rValue.IsNil()
  default:
    return false
  }
}

func Nil(
  tb testing.TB,
  value any,
  additionalMsg ...any,
) {
  tb.Helper()

  if !isNil(value) {
    logFailure(
      tb,
      additionalMsg,
      "Expected nil\nActual: %s",
      formatValue(value))
  }
}

func NotNil(
  tb testing.TB,
  value any,
  additionalMsg ...any,
) {
  tb.Helper()

  if isNil(value) {
    logFailure(tb, additionalMsg, "Expected not nil")
  }
}

func Error(
  tb testing.TB,
  err error,
  expectedSnippet string,
  additionalMsg ...any,
) {
  tb.Helper()

  if err == nil {
    logFailure(tb, additionalMsg, "Expected error\nActual: nil")
  }

  if expectedSnippet != "" && !strings.Contains(err.Error(), expectedSnippet) {
    logFailure(
      tb,
      additionalMsg,
      "Expected error snippet\nExpected snippet: %s\nActual error msg: %s",
      expectedSnippet,
      err.Error())
  }
}

func True(tb testing.TB, value bool, additionalMsg ...any) {
  tb.Helper()

  if !value {
    logFailure(tb, additionalMsg, "Expected true")
  }
}

func False(tb testing.TB, value bool, additionalMsg ...any) {
  tb.Helper()

  if value {
    logFailure(tb, additionalMsg, "Expected false")
  }
}
