package filesystem

import (
  "context"
  "io/fs"
  "os"
  "path/filepath"
)

var Local FileSystem = NewLocalFileSystem()

const (
  defaultLocalFilePerm = 0644
  defaultLocalDirPerm = 0755
)

type localFileSystemParams struct {
  openFileFlag int

  // file and dir perm need to be separate if we want to support Copy.
  filePerm FileMode
  dirPerm FileMode
}

func (params *localFileSystemParams) SetContext(ctx context.Context) {
  // Do nothing.
}

func (params *localFileSystemParams) SetFilePerm(perm FileMode) {
  params.filePerm = perm
}

func (params *localFileSystemParams) SetDirPerm(perm FileMode) {
  params.dirPerm = perm
}

// Applicable to Create Append, and WriteFile.
func WithLocalOpenFileFlag(flag int) Option {
  return func(params CommonOptions) {
    options, ok := params.(*localFileSystemParams)
    if ok {
      options.openFileFlag = flag
    }
  }
}

// This separation is mostly for testing extended file system.
type minLocalFileSystem struct {}

func (minLocalFileSystem) Abs(filePath string) (string, error) {
  if filePath == "" {
    return "", NewInvalidArgumentError("", "", "empty path")
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
  op string,
  filePath string,
  flag int,
  options []Option,
) (
  FileWriter,
  error,
) {
  absPath, err := local.Abs(filePath)
  if err != nil {
    return nil, WrapError(op, filePath, err)
  }

  params := &localFileSystemParams{
    openFileFlag: flag,
    filePerm: defaultLocalFilePerm,
  }

  for _, update := range options {
    update(params)
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
  return local.openFile(
    "Create",
    filePath,
    os.O_CREATE | os.O_TRUNC | os.O_WRONLY,
    options)
}

func (local minLocalFileSystem) Append(
  filePath string,
  options ...Option,
) (
  FileWriter,
  error,
) {
  return local.openFile(
    "Append",
    filePath,
    os.O_CREATE | os.O_APPEND | os.O_WRONLY,
    options)
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
