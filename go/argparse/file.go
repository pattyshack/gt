package argparse

import (
  "fmt"
  "strings"

  "github.com/pattyshack/bt/go/filesystem"
)

type FilePathType struct {
  StringType

  filesystem.FileSystem
  directoryOnly bool
}

func NewFilePathType(
  fileSystem filesystem.FileSystem,
  directoryOnly bool,
) FilePathType {
  return FilePathType{
    FileSystem: fileSystem,
    directoryOnly: directoryOnly,
  }
}

func (suggestor FilePathType) TypeDescription() string {
  if suggestor.directoryOnly {
    return "a directory"
  }

  return "a file"
}

func (suggestor FilePathType) Validate(path string) error {
  info, err := suggestor.Stat(path)
  if err != nil {
    return err
  }

  if suggestor.directoryOnly && !info.IsDir() {
    return fmt.Errorf("not a directory")
  }

  return nil
}

func (filePath FilePathType) Suggest(prefix string) []Suggestion {
  suggestions := []Suggestion{}

  parentDir := ""
  realParentPrefix := ""
  filePrefix := ""

  components := strings.Split(prefix, "/")
  if len(components) == 1 {  // this must be relative to working directory.
    parentDir = "."
    realParentPrefix = ""
    filePrefix = prefix
  } else {
    parentDir = strings.Join(components[:len(components) - 1], "/")
    realParentPrefix = parentDir + "/"

    if parentDir == "" {  // i.e., joining a single empty component string
      parentDir = "/"
    }

    filePrefix = components[len(components) - 1]
  }

  if filePrefix == "." || filePrefix == ".." {
    suggestions = append(
      suggestions,
      Suggestion{
        Value: realParentPrefix + "./",
        IsPrefix: true,
      },
      Suggestion{
        Value: realParentPrefix + "../",
        IsPrefix: true,
      })
  }

  entries, err := filePath.ReadDir(parentDir)
  if err != nil {
    return suggestions
  }

  matched := false
  for _, entry := range entries {
    if !strings.HasPrefix(entry.Name(), filePrefix) {
      continue
    }

    fileName := realParentPrefix + entry.Name()
    info, err := filePath.Stat(fileName)
    if err != nil || !info.IsDir() {
      if !filePath.directoryOnly {
        suggestions = append(
          suggestions,
          Suggestion{
            Value: fileName,
            IsPrefix: false,
          })
        matched = true
      }
    } else {
      suggestions = append(
        suggestions,
        Suggestion{
          Value: fileName + "/",
          IsPrefix: true,
        })
      matched = true
    }
  }

  if filePrefix == "" && !matched {
    // suggest adding <space> to complete the path
    return append(
      suggestions,
      Suggestion{
        Value: prefix,
        IsPrefix: false,
      })
  }

  return suggestions
}
