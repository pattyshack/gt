package filesystem

const readOnlyErrMsg = "file system is in read-only mode"

type readOnlyFileSystem struct {
  FileSystem
}

func NewReadOnlyFileSystem(fs FileSystem) FileSystem {
  return readOnlyFileSystem{
    FileSystem: fs,
  }
}

func (readOnlyFileSystem) Create(
  filePath string,
  options ...Option,
) (
  FileWriter,
  error,
) {
  return nil, NewPermissionError("Create", filePath, readOnlyErrMsg)
}

func (readOnlyFileSystem) Append(
  filePath string,
  options ...Option,
) (
  FileWriter,
  error,
) {
  return nil, NewPermissionError("Append", filePath, readOnlyErrMsg)
}

func (readOnlyFileSystem) Mkdir(
  dirPath string,
  options ...Option,
) error {
  return NewPermissionError("Mkdir", dirPath, readOnlyErrMsg)
}

func (readOnlyFileSystem) Rename(
  srcPath string,
  destPath string,
  options ...Option,
) error {
  return NewPermissionError("Rename", srcPath, readOnlyErrMsg)
}

func (readOnlyFileSystem) Remove(
  filePath string,
  options ...Option,
) error {
  return NewPermissionError("Remove", filePath, readOnlyErrMsg)
}

func (readOnlyFileSystem) Chmod(
  filePath string,
  perm FileMode,
  options ...Option,
) error {
  return NewPermissionError("Chmod", filePath, readOnlyErrMsg)
}

func (readOnlyFileSystem) WriteFile(
  filePath string,
  data []byte,
  options ...Option,
) error {
  return NewPermissionError("WriteFile", filePath, readOnlyErrMsg)
}

func (readOnlyFileSystem) CopyFile(
  srcPath string,
  destPath string,
  options ...Option,
) error {
  return NewPermissionError("CopyFile", srcPath, readOnlyErrMsg)
}

func (readOnlyFileSystem) CopyAll(
  srcPath string,
  destPath string,
  options ...Option,
) error {
  return NewPermissionError("CopyAll", srcPath, readOnlyErrMsg)
}

func (readOnlyFileSystem) MkdirAll(
  dirPath string,
  options ...Option,
) error {
  return NewPermissionError("MkdirAll", dirPath, readOnlyErrMsg)
}

func (readOnlyFileSystem) RemoveAll(
  filePath string,
  options ...Option,
) error {
  return NewPermissionError("RemoveAll", filePath, readOnlyErrMsg)
}
