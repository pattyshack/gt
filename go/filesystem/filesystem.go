package filesystem

import (
  "context"
  "io"
  "io/fs"
)

const (
  PermissionBits = fs.ModePerm
)

// Note: DirEntry's Type()/Info() usage should generally be avoided since
// type/info may not be fully populated.  User should use Stat the actual file
// for accurate information.
type DirEntry = fs.DirEntry

type FileInfo = fs.FileInfo

// Note: In general, only the PermissionBits in FileMode can be specified as
// argument; the file system implementator should ignore the non-permission
// bits.  FileInfo/DirEntry's Mode() may include non-permission bits.
type FileMode = fs.FileMode

type PathError = fs.PathError
type WalkDirFunc = fs.WalkDirFunc

// XXX: Maybe support Seek
//
// Note: Unlike FileSystem's ReadDir, FileReader's ReadDir is unsorted.
//
// Note: DirEntry's IsDir/Type/Info maybe not be fully populated by file system
// implementation.  User must Stat the actual file for accurate information.
type FileReader = fs.ReadDirFile

// XXX: Maybe support Sync
type FileWriter = io.WriteCloser

// Standard options across all implementations.  (The implementation may
// choose to ignore these options, but must expose the api anyway).
type Options interface {
  SetContext(ctx context.Context)

  // File and dir permissions need to be separate in order to support Copy.
  SetFilePerm(perm FileMode)
  SetDirPerm(perm FileMode)

  // All other file system implementation specific options should be set via
  // this method (Standardizing the option setting interface simplifies
  // file system api forwarding/proxying).  File system implementation should
  // ignore unrelated options.
  SetOption(fsImplName string, optionName string, optionValue string)
}

// Note that the user may pass in options not associated to the file system's
// method implementation, in which case, the implementation should ignore the
// option.
type Option func(Options)

// Applicable to all file system methods.
func WithContext(ctx context.Context) Option {
  return func(options Options) {
    options.SetContext(ctx)
  }
}

// Applicable to file creation operations (Create, Append, WriteFile, CopyFile,
// and CopyAll).  File system implementation may place additional (umask)
// restriction on the given permission.
func WithFilePerm(perm FileMode) Option {
  return func(options Options) {
    options.SetFilePerm(perm & PermissionBits)
  }
}

// Applicable to directory creation operations (Mkdir, MkdirAll, and CopyAll).
// File system implementation may place additional (umask) restriction on the
// given permission.
func WithDirPerm(perm FileMode) Option {
  return func(options Options) {
    options.SetDirPerm(perm & PermissionBits)
  }
}

// Similar to fs.FS, but extended to support write operations.  All methods
// should return error on empty path string inputs.
//
// Note: Windows-style backslash (and volume named) paths are not supported.
// User should use golang's "path" (instead of "path/filepath") for path
// manipulation, and use the file system's Abs method to resolve non-absolute
// paths.
//
// Note: This is the minimal set of operations that the file system
// implementor must provide (The implementation should return error for
// unsupported functionalities).  If the file system implementation does not
// implement the full set of FileSystem operations, the implementator should
// use ExtendMinimalFileSystem() to provide implementation for the remaining
// operations.
type MinimalFileSystem interface {
  // Equivalent to path/filepath.Abs.  If the file system implementation
  // does not support working directory, this should return error when the
  // path is not absolute.  This should always return error when filePath
  // is an empty string.
  Abs(filePath string) (string, error)

  // This returns a FileReader that behaves as if it was opened with O_RDONLY.
  Open(filePath string, options ...Option) (FileReader, error)

  // This returns a FileWriter that behaves as if it was opened with
  // O_CREATE | O_WRONLY | O_TRUNC.  The default permission should be 0664 or
  // tighter if the file system supports unix style permission.
  //
  // If a file perm option is provided, the file system implementation may
  // place additional (umask) restriction on the provided permission.
  //
  // This returns an error if Create is not supported.
  Create(filePath string, options ...Option) (FileWriter, error)

  // This returns a FileWriter that behaves as if it was opened with
  // O_CREATE |O_WRONLY | O_APPEND.  The default permission should be 0664
  // or tighter if the file system supports unix style permission.
  //
  // If a file perm option is provided, the file system implementation may
  // place additional (umask) restriction on the provided permission.
  //
  // This returns an error if Append is not supported.
  Append(filePath string, options ...Option) (FileWriter, error)

  // This creates a directory.  The default permission should be 0775 or
  // tighter.
  //
  // If a dir perm option is provided, the file system implementation may
  // place additional (umask) restriction on the provided permission.
  //
  // This returns an error if Mkdir is not supported.
  Mkdir(dirPath string, options ...Option) error

  // This renames a file.
  //
  // This returns an error if Mkdir is not supported.
  // Note that Rename may not be an atomic operation (the implementation
  // could for example copy then delete).
  Rename(srcPath string, descPath string, options ...Option) error

  // This remove a file or an empty directory.
  //
  // This returns an error if Remove is not supported.
  Remove(filePath string, options ... Option) error

  // Intended for changing permission bits.  The file system implementor
  // should ignore non-permission file mode bits.
  Chmod(filePath string, perm FileMode, options ...Option) error

  // Note: It doesn't make sense to support unix style Chown since uid/gid are
  // unix / localhost specific.  Local uid/gid are meaningless in windows /
  // distributed file systems.  Maybe support Chown by user/group name instead?
}

// FileSystem specify additional operations on top of MinimalFileSystem for
// ease of use.
//
// Note: all path input must be absolute (i.e., not relative to local working
// directory) since the working directory is meaningless to all but the local
// file system implementation rooted at '/'.  File system implementator must
// guard against non-absolute paths.
type FileSystem interface {
  MinimalFileSystem

  // This returns the path's FileInfo.
  Stat(filePath string, options ...Option) (FileInfo, error)

  // This returns a directory's content, sorted by filename.
  ReadDir(dirPath string, options ...Option) ([]DirEntry, error)

  // This walks a directory.
  WalkDir(dirPath string, fn WalkDirFunc, options ...Option) error

  // This globs a file pattern.
  Glob(pattern string, options ...Option) ([]string, error)

  // This returns a file's content.
  ReadFile(filePath string, options ...Option) ([]byte, error)

  // This writes a file.
  WriteFile(
    filePath string,
    data []byte,
    options ...Option,
  ) error

  // Copy the file at srcFilePath to destFilePath.  If destFilePath is a file,
  // the file content is overwritten.  If destFilePath is a directory, this
  // returns error.
  CopyFile(
    srcFilePath string,
    destFilePath string,
    options ...Option,
  ) error

  // Recursively copy src to dest.
  CopyAll(
    srcePath string,
    destPath string,
    options ...Option,
  ) error

  // This recursively creates the directory dirPath.
  MkdirAll(dirPath string, options ...Option) error

  // This removes all files under filePath.
  RemoveAll(filePath string, options ...Option) error
}

