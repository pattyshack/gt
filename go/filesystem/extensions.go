package filesystem

import (
  "io"
  "io/fs"
  "path"
)

type fsWithOptions struct {
  fileSystem MinimalFileSystem
  options []Option
}

func (fso fsWithOptions) Open(filePath string) (fs.File, error) {
  return fso.fileSystem.Open(filePath, fso.options...)
}

// This returns a fs.FS object usable by io/fs's functions that uses fs.FS as
// input.  In general, this should not be used directly.
func ToFS(fileSystem MinimalFileSystem, options []Option) fs.FS {
  return fsWithOptions{
    fileSystem: fileSystem,
    options: options,
  }
}

// A Stat implementation using only MinimalFileSystem.  File system implementor
// may use this to implement the Stat method.
func Stat(
  minFS MinimalFileSystem,
  filePath string,
  options ...Option,
) (
  FileInfo,
  error,
) {
  info, err := fs.Stat(ToFS(minFS, options), filePath)
  if err != nil {
    return nil, WrapError("Stat", filePath, err)
  }

  return info, nil
}

// A ReadDir implementation using only MinimalFileSystem.  File system
// implementor may use this to implement the ReadDir method.
func ReadDir(
  minFS MinimalFileSystem,
  dirPath string,
  options ...Option,
) (
  []DirEntry,
  error,
) {
  entries, err := fs.ReadDir(ToFS(minFS, options), dirPath)
  if err != nil {
    return nil, WrapError("ReadDir", dirPath, err)
  }

  return entries, nil
}

// A WalkDir implementation using only MinimalFileSystem.  File system
// implementor may use this to implement the WalkDir method.
func WalkDir(
  minFS MinimalFileSystem,
  dirPath string,
  fn WalkDirFunc,
  options ...Option,
) error {
  err := fs.WalkDir(ToFS(minFS, options), dirPath, fn)
  if err != nil {
    return WrapError("WalkDir", dirPath, err)
  }

  return nil
}

// A Glob implementation using only MinimalFileSystem.  File system implementor
// may use this to implement the Glob method.
func Glob(
  minFS MinimalFileSystem,
  pattern string,
  options ...Option,
) (
  []string,
  error,
) {
  matches, err := fs.Glob(ToFS(minFS, options), pattern)
  if err != nil {
    return nil, WrapError("Glob", pattern, err)
  }

  return matches, nil
}

// A ReadFile implementation using only MinimalFileSystem.  File system
// implementor may use this to implement the ReadFile method.
func ReadFile(
  minFS MinimalFileSystem,
  filePath string,
  options ...Option,
) (
  []byte,
  error,
) {
  content, err := fs.ReadFile(ToFS(minFS, options), filePath)
  if err != nil {
    return nil, WrapError("ReadFile", filePath, err)
  }

  return content, nil
}

// A WriteFile implementation using only MinimalFileSystem.  File system
// implementor may use this to implement the WriteFile method.
func WriteFile(
  minFS MinimalFileSystem,
  filePath string,
  data []byte,
  options ...Option,
) error {
  err := writeFile(minFS, filePath, data, options)
  if err != nil {
    return WrapError("WriteFile", filePath, err)
  }

  return nil
}

func writeFile(
  minFS MinimalFileSystem,
  filePath string,
  data []byte,
  options []Option,
) error {
  file, err := minFS.Create(filePath, options...)
  if err != nil {
    return err
  }

  _, err = file.Write(data)
  closeErr := file.Close()

  if err != nil {
    return err
  }

  return closeErr
}

// A MkdirAll implementation using only MinimalFileSystem.
// File system implementor may use this to implement the MkdirAll method.
func MkdirAll(
  minFS MinimalFileSystem,
  dirPath string,
  options ...Option,
) error {
  err := mkdirAll(minFS, path.Clean(dirPath), options)
  if err != nil {
    return WrapError("MkdirAll", dirPath, err)
  }

  return nil
}

func mkdirAll(
  minFS MinimalFileSystem,
  dirPath string,
  options []Option,
) error {
  stat, err := Stat(minFS, dirPath, options...)
  if err == nil {
    if stat.IsDir() {
      return nil
    }
    return NewExistError("MkdirAll", dirPath, "cannot overwrite file")
  } else if !IsNotExistError(err) {
    return err
  }

  err = mkdirAll(minFS, path.Dir(dirPath), options)
  if err != nil {
    return err
  }

  return minFS.Mkdir(dirPath, options...)
}

// A RemoveAll implementation using only MinimalFileSystem.
// File system implementor may use this to implement the RemoveAll method.
func RemoveAll(
  minFS MinimalFileSystem,
  filePath string,
  options ...Option,
) error {
  err := removeAll(minFS, filePath, options)
  if err != nil {
    return WrapError("RemoveAll", filePath, err)
  }

  return nil
}

func removeAll(
  minFS MinimalFileSystem,
  filePath string,
  options []Option,
) error {
  info, err := Stat(minFS, filePath, options...)
  if err != nil {
    if IsNotExistError(err) {
      return nil
    }
    return err
  }

  if !info.IsDir() {
    return minFS.Remove(filePath, options...)
  }

  return removeDir(minFS, filePath, options)
}

func removeDir(
  minFS MinimalFileSystem,
  dirPath string,
  options []Option,
) error {
  entries, err := ReadDir(minFS, dirPath, options...)
  if err != nil {
    return err
  }

  for _, entry := range entries {
    childPath := path.Join(dirPath, entry.Name())

    if entry.IsDir() {
      err := removeDir(minFS, childPath, options)
      if err != nil {
        return err
      }
    } else {
      err := minFS.Remove(childPath, options...)
      if err != nil {
        return err
      }
    }
  }

  return minFS.Remove(dirPath, options...)
}

// This Copies a src file to dest.  If dest is a file, the file content is
// overwritten.  if dest is a directory, this returns an error.
//
// This implementation only uses only MinimalFileSystem.  File system
// implementor may use this to implement the CopyFile method.
func CopyFile(
  srcFS MinimalFileSystem,
  srcPath string,
  destFS MinimalFileSystem,
  destPath string,
  options ...Option,
) error {
  srcInfo, err := Stat(srcFS, srcPath, options...)
  if err != nil {
    return WrapError("CopyFile", srcPath, err)
  }

  if srcInfo.IsDir() {
    return NewInvalidArgumentError("CopyFile", srcPath, "cannot copy directory")
  }

  destInfo, err := Stat(destFS, destPath, options...)
  if err == nil {
    if destInfo.IsDir() {
      return NewInvalidArgumentError(
        "CopyFile",
        destPath,
        "cannot overwrite directory")
    }
  } else if !IsNotExistError(err) {
    return WrapError("CopyFile", destPath, err)
  }

  err = copyFile(srcFS, srcPath, srcInfo.Mode(), destFS, destPath, options)
  if err != nil {
    return WrapError("CopyFile", srcPath, err)
  }

  return nil
}

func copyFile(
  srcFS MinimalFileSystem,
  srcPath string,
  srcPerm FileMode,
  destFS MinimalFileSystem,
  destPath string,
  options []Option,
) (err error) {
  var reader FileReader
  reader, err = srcFS.Open(srcPath, options...)
  if err != nil {
    return err
  }
  defer func() {
    closeErr := reader.Close()
    if err == nil {
      err = closeErr
    }
  }()

  // Try to copy over permission as default, but allow user to override it.
  // XXX: maybe switch to chmod instead of depending on file creation
  // permission.
  writerOptions := append([]Option{WithFilePerm(srcPerm)}, options...)

  var writer FileWriter
  writer, err = destFS.Create(destPath, writerOptions...)
  if err != nil {
    return err
  }
  defer func() {
    closeErr := writer.Close()
    if err == nil {
      err = closeErr
    }
  }()

  _, err = io.Copy(writer, reader)
  return err
}

// Recursively copy src to dest.
//
// Behavior:
//  1. src is a file:
//    a. dest does not exist: copy src file to dest (try to create using src's
//       permission if caller didn't specify file perm option)
//    b. dest is a file: overwrite dest file with src (use dest's original
//       permission)
//    c. dest is a directory: error
//  2. src is a directory:
//    a. dest does not exist: recursively copy src directory to dest (try to
//       create using src's permissions if caller didn't specify file/dir
//       perm options)
//    b. dest exist: error
//
// (If we ignore file/directory metadata, e.g., permission, fs.Rename(src, dest)
// should be equivalent to fs.CopyAll(src, dest) follow by fs.RemoveAll(src).
//
// This implementation only uses only MinimalFileSystem.  File system
// implementor may use this to implement the CopyAll method.
func CopyAll(
  srcFS MinimalFileSystem,
  srcPath string,
  destFS MinimalFileSystem,
  destPath string,
  options ...Option,
) error {
  srcInfo, err := Stat(srcFS, srcPath, options...)
  if err != nil {
    return WrapError("CopyAll", srcPath, err)
  }

  destInfo, err := Stat(destFS, destPath, options...)
  if err == nil {
    if destInfo.IsDir() {
      return NewInvalidArgumentError(
        "CopyAll",
        destPath,
        "cannot overwrite directory")
    } else if srcInfo.IsDir() {
      return NewInvalidArgumentError(
        "CopyAll",
        destPath,
        "cannot overwrite file")
    }
  } else if !IsNotExistError(err) {
    return WrapError("CopyAll", destPath, err)
  }

  if srcInfo.IsDir() {
    err = copyDir(srcFS, srcPath, srcInfo.Mode(), destFS, destPath, options)
  } else {
    err = copyFile(srcFS, srcPath, srcInfo.Mode(), destFS, destPath, options)
  }

  if err != nil {
    return WrapError("CopyAll", srcPath, err)
  }

  return nil
}

func copyDir(
  srcFS MinimalFileSystem,
  srcPath string,
  srcPerm FileMode,
  destFS MinimalFileSystem,
  destPath string,
  options []Option,
) error {
  // Try to copy over permission as default, but allow user to override it.
  // XXX: maybe switch to chmod instead of depending on dir creation
  // permission.
  mkdirOptions := append([]Option{WithDirPerm(srcPerm)}, options...)

  err := destFS.Mkdir(destPath, mkdirOptions...)
  if err != nil {
    return err
  }

  entries, err := ReadDir(srcFS, srcPath, options...)
  if err != nil {
    return err
  }

  for _, entry := range entries {
    childSrcPath := path.Join(srcPath, entry.Name())
    childDestPath := path.Join(destPath, entry.Name())

    childSrcInfo, err := Stat(srcFS, childSrcPath)
    if err != nil {
      return err
    }

    if childSrcInfo.IsDir() {
      err = copyDir(
        srcFS,
        childSrcPath,
        childSrcInfo.Mode(),
        destFS,
        childDestPath,
        options)
      if err != nil {
        return WrapError("CopyAll", childSrcPath, err)
      }
    } else {
      err = copyFile(
        srcFS,
        childSrcPath,
        childSrcInfo.Mode(),
        destFS,
        childDestPath,
        options)
      if err != nil {
        return WrapError("CopyAll", childSrcPath, err)
      }
    }
  }

  return nil
}

type extendedFileSystem struct {
  MinimalFileSystem
}

// Given a MinimalFileSystem, which may or may not have fully implemented the
// full FileSystem interface, return a full FileSystem implementation.
func ExtendMinimalFileSystem(minFS MinimalFileSystem) FileSystem {
  fileSystem, ok := minFS.(FileSystem)
  if ok {
    return fileSystem
  }

  return &extendedFileSystem{
    MinimalFileSystem: minFS,
  }
}

func (fileSystem *extendedFileSystem) Stat(
  filePath string,
  options ...Option,
) (
  FileInfo,
  error,
) {
  type FSStat interface {
    Stat(string, ...Option) (FileInfo, error)
  }

  stat, ok := fileSystem.MinimalFileSystem.(FSStat)
  if ok {
    return stat.Stat(filePath, options...)
  }

  return Stat(fileSystem.MinimalFileSystem, filePath, options...)
}

func (fileSystem *extendedFileSystem) ReadDir(
  dirPath string,
  options ...Option,
) (
  []DirEntry,
  error,
) {
  type FSReadDir interface {
    ReadDir(string, ...Option) ([]DirEntry, error)
  }

  readDir, ok := fileSystem.MinimalFileSystem.(FSReadDir)
  if ok {
    return readDir.ReadDir(dirPath, options...)
  }

  return ReadDir(fileSystem.MinimalFileSystem, dirPath, options...)
}

func (fileSystem *extendedFileSystem) WalkDir(
  dirPath string,
  fn WalkDirFunc,
  options ...Option,
) error {
  type FSWalkDir interface {
    WalkDir(string, WalkDirFunc, ...Option) error
  }

  walkDir, ok := fileSystem.MinimalFileSystem.(FSWalkDir)
  if ok {
    return walkDir.WalkDir(dirPath, fn, options...)
  }

  return WalkDir(fileSystem.MinimalFileSystem, dirPath, fn, options...)
}

func (fileSystem *extendedFileSystem) Glob(
  pattern string,
  options ...Option,
) (
  []string,
  error,
) {
  type FSGlob interface {
    Glob(string, ...Option) ([]string, error)
  }

  glob, ok := fileSystem.MinimalFileSystem.(FSGlob)
  if ok {
    return glob.Glob(pattern, options...)
  }

  return Glob(fileSystem.MinimalFileSystem, pattern, options...)
}

func (fileSystem *extendedFileSystem) ReadFile(
  filePath string,
  options ...Option,
) (
  []byte,
  error,
) {
  type FSReadFile interface {
    ReadFile(string, ...Option) ([]byte, error)
  }

  readFile, ok := fileSystem.MinimalFileSystem.(FSReadFile)
  if ok {
    return readFile.ReadFile(filePath, options...)
  }

  return ReadFile(fileSystem.MinimalFileSystem, filePath, options...)
}

func (fileSystem *extendedFileSystem) WriteFile(
  filePath string,
  data []byte,
  options ...Option,
) error {
  type FSWriteFile interface {
    WriteFile(string, []byte, ...Option) error
  }

  writeFile, ok := fileSystem.MinimalFileSystem.(FSWriteFile)
  if ok {
    return writeFile.WriteFile(filePath, data, options...)
  }

  return WriteFile(
    fileSystem.MinimalFileSystem,
    filePath,
    data,
    options...)
}

func (fileSystem *extendedFileSystem) CopyFile(
  srcFilePath string,
  destFilePath string,
  options ...Option,
) error {
  type FSCopyFile interface {
    CopyFile(string, string, ...Option) error
  }

  copyFile, ok := fileSystem.MinimalFileSystem.(FSCopyFile)
  if ok {
    return copyFile.CopyFile(srcFilePath, destFilePath, options...)
  }

  return CopyFile(
    fileSystem.MinimalFileSystem,
    srcFilePath,
    fileSystem.MinimalFileSystem,
    destFilePath,
    options...)
}

func (fileSystem *extendedFileSystem) CopyAll(
  srcPath string,
  destPath string,
  options ...Option,
) error {
  type FSCopyAll interface {
    CopyAll(string, string, ...Option) error
  }

  copyFile, ok := fileSystem.MinimalFileSystem.(FSCopyAll)
  if ok {
    return copyFile.CopyAll(srcPath, destPath, options...)
  }

  return CopyAll(
    fileSystem.MinimalFileSystem,
    srcPath,
    fileSystem.MinimalFileSystem,
    destPath,
    options...)
}

func (fileSystem *extendedFileSystem) MkdirAll(
  dirPath string,
  options ...Option,
) error {
  type FSMkdirAll interface {
    MkdirAll(string, ...Option) error
  }

  mkdirAll, ok := fileSystem.MinimalFileSystem.(FSMkdirAll)
  if ok {
    return mkdirAll.MkdirAll(dirPath, options...)
  }

  return MkdirAll(fileSystem.MinimalFileSystem, dirPath, options...)
}

func (fileSystem *extendedFileSystem) RemoveAll(
  filePath string,
  options ...Option,
) error {
  type FSRemoveAll interface {
    RemoveAll(string, ...Option) error
  }

  removeAll, ok := fileSystem.MinimalFileSystem.(FSRemoveAll)
  if ok {
    return removeAll.RemoveAll(filePath, options...)
  }

  return RemoveAll(fileSystem.MinimalFileSystem, filePath, options...)
}
