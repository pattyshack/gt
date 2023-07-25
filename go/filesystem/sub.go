package filesystem

import (
  "path"
  "strings"
)

type subFileSystem struct {
  baseFS FileSystem
  baseRootDir string
}

// Return a new sub file system rooted at baseRootDir.  Files outside of
// baseRootDir are inaccessible from the sub file system.
func NewSubFileSystem(baseFS FileSystem, baseRootDir string) FileSystem {
  return subFileSystem{
    baseFS: baseFS,
    baseRootDir: baseRootDir,
  }
}

func (sub subFileSystem) Abs(filePath string) (string, error) {
  if filePath == "" {
    return "", NewInvalidArgumentError("Abs", "", "empty path")
  }

  // Note: non-abs filePath must be clean prior to joining with sub file
  // system's working directory to ensure the path is nested inside the sub
  // file system's root.
  filePath = path.Clean(filePath)
  if path.IsAbs(filePath) {
    return filePath, nil
  }

  baseWorkingDir, err := sub.baseFS.Abs(".")
  if err != nil {
    return "", err
  }

  subWorkingDir := baseWorkingDir
  if sub.baseRootDir != "/" {
    if baseWorkingDir == sub.baseRootDir {
      subWorkingDir = "/"
    } else {
      prefix := sub.baseRootDir + "/"
      if strings.HasPrefix(baseWorkingDir, prefix) {
        subWorkingDir = "/" + baseWorkingDir[len(prefix):]
      } else {
        return "", NewInvalidArgumentError(
          "Abs",
          "",
          "working directory unavailable")
      }
    }
  }

  return path.Clean(path.Join(subWorkingDir, filePath)), nil
}

func (sub subFileSystem) toBasePath(filePath string) (string, error) {
  filePath, err := sub.Abs(filePath)
  if err != nil {
    return "", err
  }

  return path.Join(sub.baseRootDir, filePath), nil
}

func (sub subFileSystem) Open(
  filePath string,
  options ...Option,
) (
  FileReader,
  error,
) {
  baseFilePath, err := sub.toBasePath(filePath)
  if err != nil {
    return nil, WrapError("Open", filePath, err)
  }

  return sub.baseFS.Open(baseFilePath, options...)
}

func (sub subFileSystem) Create(
  filePath string,
  options ...Option,
) (
  FileWriter,
  error,
) {
  baseFilePath, err := sub.toBasePath(filePath)
  if err != nil {
    return nil, WrapError("Create", filePath, err)
  }

  return sub.baseFS.Create(baseFilePath, options...)
}

func (sub subFileSystem) Append(
  filePath string,
  options ...Option,
) (
  FileWriter,
  error,
) {
  baseFilePath, err := sub.toBasePath(filePath)
  if err != nil {
    return nil, WrapError("Append", filePath, err)
  }

  return sub.baseFS.Append(baseFilePath, options...)
}

func (sub subFileSystem) Mkdir(dirPath string, options ...Option) error {
  baseDirPath, err := sub.toBasePath(dirPath)
  if err != nil {
    return WrapError("Mkdir", dirPath, err)
  }

  return sub.baseFS.Mkdir(baseDirPath, options...)
}

func (sub subFileSystem) Rename(
  srcPath string,
  destPath string,
  options ...Option,
) error {
  baseSrcPath, err := sub.toBasePath(srcPath)
  if err != nil {
    return WrapError("Rename", srcPath, err)
  }

  baseDestPath, err := sub.toBasePath(destPath)
  if err != nil {
    return WrapError("Rename", destPath, err)
  }

  return sub.baseFS.Rename(baseSrcPath, baseDestPath, options...)
}

func (sub subFileSystem) Remove(filePath string, options ... Option) error {
  baseFilePath, err := sub.toBasePath(filePath)
  if err != nil {
    return WrapError("Remove", filePath, err)
  }

  return sub.baseFS.Remove(baseFilePath, options...)
}

func (sub subFileSystem) Chmod(
  filePath string,
  perm FileMode,
  options ...Option,
) error {
  baseFilePath, err := sub.toBasePath(filePath)
  if err != nil {
    return WrapError("Chmod", filePath, err)
  }

  return sub.baseFS.Chmod(baseFilePath, perm, options...)
}

func (sub subFileSystem) Stat(
  filePath string,
  options ...Option,
) (
  FileInfo,
  error,
) {
  baseFilePath, err := sub.toBasePath(filePath)
  if err != nil {
    return nil, WrapError("Stat", filePath, err)
  }

  return sub.baseFS.Stat(baseFilePath, options...)
}

func (sub subFileSystem) ReadDir(
  dirPath string,
  options ...Option,
) (
  []DirEntry,
  error,
) {
  baseDirPath, err := sub.toBasePath(dirPath)
  if err != nil {
    return nil, WrapError("ReadDir", dirPath, err)
  }

  return sub.baseFS.ReadDir(baseDirPath, options...)
}

func (sub subFileSystem) WalkDir(
  dirPath string,
  fn WalkDirFunc,
  options ...Option,
) error {
  // Note: If we use base file system's WalkDir rather than the generic
  // WalkDir implementation, we would need to intercept the fn input and
  // convert the walked base paths back to sub paths.
  return WalkDir(sub, dirPath, fn, options...)
}

func (sub subFileSystem) Glob(
  pattern string,
  options ...Option,
) (
  []string,
  error,
) {
  // Note: If we use base file system's Glob rather than the generic
  // Glob implementation, we would need to convert the matched base paths
  // back to sub paths.
  return Glob(sub, pattern, options...)
}

func (sub subFileSystem) ReadFile(
  filePath string,
  options ...Option,
) (
  []byte,
  error,
) {
  baseFilePath, err := sub.toBasePath(filePath)
  if err != nil {
    return nil, WrapError("ReadFile", filePath, err)
  }

  return sub.baseFS.ReadFile(baseFilePath, options...)
}

func (sub subFileSystem) WriteFile(
  filePath string,
  data []byte,
  options ...Option,
) error {
  baseFilePath, err := sub.toBasePath(filePath)
  if err != nil {
    return WrapError("WriteFile", filePath, err)
  }

  return sub.baseFS.WriteFile(baseFilePath, data, options...)
}

func (sub subFileSystem) CopyFile(
  srcFilePath string,
  destFilePath string,
  options ...Option,
) error {
  baseSrcFilePath, err := sub.toBasePath(srcFilePath)
  if err != nil {
    return WrapError("CopyFile", srcFilePath, err)
  }

  baseDestFilePath, err := sub.toBasePath(destFilePath)
  if err != nil {
    return WrapError("CopyFile", destFilePath, err)
  }

  return sub.baseFS.CopyFile(baseSrcFilePath, baseDestFilePath, options...)
}

func (sub subFileSystem) CopyAll(
  srcPath string,
  destPath string,
  options ...Option,
) error {
  baseSrcPath, err := sub.toBasePath(srcPath)
  if err != nil {
    return WrapError("CopyAll", srcPath, err)
  }

  baseDestPath, err := sub.toBasePath(destPath)
  if err != nil {
    return WrapError("CopyAll", destPath, err)
  }

  return sub.baseFS.CopyAll(baseSrcPath, baseDestPath, options...)
}

func (sub subFileSystem) MkdirAll(dirPath string, options ...Option) error {
  baseDirPath, err := sub.toBasePath(dirPath)
  if err != nil {
    return WrapError("MkdirAll", dirPath, err)
  }

  return sub.baseFS.MkdirAll(baseDirPath, options...)
}

func (sub subFileSystem) RemoveAll(filePath string, options ...Option) error {
  baseFilePath, err := sub.toBasePath(filePath)
  if err != nil {
    return WrapError("RemoveAll", filePath, err)
  }

  return sub.baseFS.RemoveAll(baseFilePath, options...)
}
