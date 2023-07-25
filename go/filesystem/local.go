package filesystem

import (
  "context"
  "fmt"
  "io/fs"
  "os"
  "path/filepath"
  "strconv"
)

var Local FileSystem = NewLocalFileSystem()

const (
  localFileSystemName = "Local"
  localOpenFileFlagOption = "OpenFileFlag"

  defaultLocalFilePerm = 0644
  defaultLocalDirPerm = 0755
)

type localFileSystemParams struct {
  openFileFlag int

  // file and dir perm need to be separate if we want to support Copy.
  filePerm FileMode
  dirPerm FileMode

  err error
}

func (options *localFileSystemParams) SetContext(ctx context.Context) {
  // Do nothing.
}

func (options *localFileSystemParams) SetFilePerm(perm FileMode) {
  options.filePerm = perm
}

func (options *localFileSystemParams) SetDirPerm(perm FileMode) {
  options.dirPerm = perm
}

func (options *localFileSystemParams) SetOption(
  fsImplName string,
  optionName string,
  optionValue string,
) {
  if fsImplName != localFileSystemName {
    return
  }

  if optionName == localOpenFileFlagOption {
    val, err := strconv.ParseInt(optionValue, 0, strconv.IntSize)
    if err != nil {
      options.err = fmt.Errorf(
        "invalid open file flag option (%s): %w",
        optionValue,
        err)
    } else {
      options.openFileFlag = int(val)
    }
  }
}

// Applicable to Create Append, and WriteFile.
func WithLocalOpenFileFlag(flag int) Option {

  return func(options Options) {
    options.SetOption(
      localFileSystemName,
      localOpenFileFlagOption,
      strconv.Itoa(flag))
  }
}

// This separation is mostly for testing extended file system.
type minLocalFileSystem struct {}

func (minLocalFileSystem) Abs(filePath string) (string, error) {
  if filePath == "" {
    return "", NewInvalidArgumentError("Abs", "", "empty path")
  }

  return filepath.Abs(filePath)
}

func (local minLocalFileSystem) Open(
  filePath string,
  options ...Option,
) (
  FileReader,
  error,
) {
  absPath, err := local.Abs(filePath)
  if err != nil {
    return nil, WrapError("Open", filePath, err)
  }

  return os.Open(absPath)
}

func (local minLocalFileSystem) openFile(
  filePath string,
  flag int,
  options []Option,
) (
  FileWriter,
  error,
) {
  absPath, err := local.Abs(filePath)
  if err != nil {
    return nil, err
  }

  params := &localFileSystemParams{
    openFileFlag: flag,
    filePerm: defaultLocalFilePerm,
  }

  for _, update := range options {
    update(params)
  }

  if params.err != nil {
    return nil, params.err
  }

  return os.OpenFile(absPath, params.openFileFlag, params.filePerm)
}

func (local minLocalFileSystem) Create(
  filePath string,
  options ...Option,
) (
  FileWriter,
  error,
) {
  writer, err := local.openFile(
    filePath,
    os.O_CREATE | os.O_TRUNC | os.O_WRONLY,
    options)
  if err != nil {
    return nil, WrapError("Create", filePath, err)
  }

  return writer, nil
}

func (local minLocalFileSystem) Append(
  filePath string,
  options ...Option,
) (
  FileWriter,
  error,
) {
  writer, err := local.openFile(
    filePath,
    os.O_CREATE | os.O_APPEND | os.O_WRONLY,
    options)
  if err != nil {
    return nil, WrapError("Append", filePath, err)
  }

  return writer, nil
}

func (local minLocalFileSystem) Mkdir(
  dirPath string,
  options ...Option,
) error {
  absPath, err := local.Abs(dirPath)
  if err != nil {
    return WrapError("Mkdir", dirPath, err)
  }

  params := &localFileSystemParams{
    dirPerm: defaultLocalDirPerm,
  }

  for _, update := range options {
    update(params)
  }

  if params.err != nil {
    return WrapError("Mkdir", dirPath, params.err)
  }

  return os.Mkdir(absPath, params.dirPerm)
}

func (local minLocalFileSystem) Rename(
  srcPath string,
  destPath string,
  options ...Option,
) error {
  absSrcPath, err := local.Abs(srcPath)
  if err != nil {
    return WrapError("Rename", srcPath, err)
  }

  absDestPath, err := local.Abs(destPath)
  if err != nil {
    return WrapError("Rename", destPath, err)
  }

  return os.Rename(absSrcPath, absDestPath)
}

func (local minLocalFileSystem) Remove(
  filePath string,
  options ...Option,
) error {
  absPath, err := local.Abs(filePath)
  if err != nil {
    return WrapError("Remove", filePath, err)
  }

  return os.Remove(absPath)
}

func (local minLocalFileSystem) Chmod(
  filePath string,
  perm FileMode,
  options ...Option,
) error {
  absPath, err := local.Abs(filePath)
  if err != nil {
    return WrapError("Chmod", filePath, err)
  }

  return os.Chmod(absPath, perm & PermissionBits)
}

type localFileSystem struct {
  minLocalFileSystem
}

func NewLocalFileSystem() FileSystem {
  return localFileSystem{}
}

func (local localFileSystem) Stat(
  filePath string,
  options ...Option,
) (
  FileInfo,
  error,
) {
  absPath, err := local.Abs(filePath)
  if err != nil {
    return nil, WrapError("Stat", filePath, err)
  }

  return os.Stat(absPath)
}

func (local localFileSystem) ReadDir(
  dirPath string,
  options ...Option,
) (
  []DirEntry,
  error,
) {
  absPath, err := local.Abs(dirPath)
  if err != nil {
    return nil, WrapError("ReadDir", dirPath, err)
  }

  return os.ReadDir(absPath)
}

func (local localFileSystem) WalkDir(
  dirPath string,
  fn WalkDirFunc,
  options ...Option,
) error {
  absPath, err := local.Abs(dirPath)
  if err != nil {
    return WrapError("WalkDir", dirPath, err)
  }

  return fs.WalkDir(ToFS(local, options), absPath, fn)
}

func (local localFileSystem) Glob(
  pattern string,
  options ...Option,
) (
  []string,
  error,
) {
  absPattern, err := local.Abs(pattern)
  if err != nil {
    return nil, WrapError("Glob", pattern, err)
  }

  return fs.Glob(ToFS(local, options), absPattern)
}

func (local localFileSystem) ReadFile(
  filePath string,
  options ...Option,
) (
  []byte,
  error,
) {
  absPath, err := local.Abs(filePath)
  if err != nil {
    return nil, WrapError("ReadFile", filePath, err)
  }

  return os.ReadFile(absPath)
}

func (local localFileSystem) WriteFile(
  filePath string,
  data []byte,
  options ...Option,
) error {
  absPath, err := local.Abs(filePath)
  if err != nil {
    return WrapError("WriteFile", filePath, err)
  }

  params := &localFileSystemParams{
    filePerm: defaultLocalFilePerm,
  }

  for _, update := range options {
    update(params)
  }

  if params.err != nil {
    return WrapError("WriteFile", filePath, params.err)
  }

  return os.WriteFile(absPath, data, params.filePerm)
}

func (local localFileSystem) CopyFile(
  srcPath string,
  destPath string,
  options ...Option,
) error {
  return CopyFile(local, srcPath, local, destPath, options...)
}

func (local localFileSystem) CopyAll(
  srcPath string,
  destPath string,
  options ...Option,
) error {
  return CopyAll(local, srcPath, local, destPath, options...)
}

func (local localFileSystem) MkdirAll(
  dirPath string,
  options ...Option,
) error {
  absPath, err := local.Abs(dirPath)
  if err != nil {
    return WrapError("MkdirAll", dirPath, err)
  }

  params := &localFileSystemParams{
    dirPerm: defaultLocalDirPerm,
  }

  for _, update := range options {
    update(params)
  }

  if params.err != nil {
    return WrapError("MkdirAll", dirPath, params.err)
  }

  return os.MkdirAll(absPath, params.dirPerm)
}

func (local localFileSystem) RemoveAll(
  filePath string,
  options ...Option,
) error {
  absPath, err := local.Abs(filePath)
  if err != nil {
    return WrapError("RemoveAll", filePath, err)
  }

  return os.RemoveAll(absPath)
}
