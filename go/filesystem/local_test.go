package filesystem

import (
  "io"
  "os"
  "path"
  "sort"
  "strings"
  "testing"

  "github.com/pattyshack/bt/go/testing/expect"
  "github.com/pattyshack/bt/go/testing/suite"
)

type testDirEntry struct {
  name string
  isDir bool
  perm FileMode
}

type LocalSuite struct {
  testDir string

  fsImpl FileSystem

  // Which os directory fsImpl thinks is the fs' root.
  fsRootDir string
}

func (s *LocalSuite) toFsPath(t *testing.T, osPath string) string {
  osPath = path.Clean(osPath)

  if osPath == s.fsRootDir {
    return "/"
  }

  root := s.fsRootDir
  if root != "/" {
    root += "/"
  }

  expect.True(
    t,
    strings.HasPrefix(osPath, root),
    "osPath (%s) is not accessible by fs (%s)",
    osPath,
    s.fsRootDir)

  return "/" + osPath[len(root):]
}

// Create the following in the temp directory:
// abc                  0644 abc
// dir1/                0755
// dir1/blah            0644 blah blah blah
// dir1/dir2/           0755
// dir1/dir2/dir3/      0755
// dir1/dir2/dir3/junk  0600 junk
// dir1/dir2/other.file 0604 other file
// dir1/dir2/some.file  0640 some file
// foo/                 0750
// foo/bar/             0750
// foo/bar/abc/         0755
// foo/bar/abc/def/     0700
// foo/bar/abc/def/ghi  0600 alphabets
// foo/file1            0644 file 1's content
// foo/file2            0640 file 2's content
// hello.txt            0644 hello
// world.txt            0604 world
func (s *LocalSuite) SetupTest(t *testing.T) {
  s.testDir = t.TempDir()

  err := os.WriteFile(
    path.Join(s.testDir, "abc"),
    []byte("abc"),
    0666)
  expect.Nil(t, err)

  err = os.MkdirAll(
    path.Join(s.testDir, "dir1/dir2/dir3"),
    0755)
  expect.Nil(t, err)

  err = os.WriteFile(
    path.Join(s.testDir, "dir1/blah"),
    []byte("blah blah blah"),
    0644)
  expect.Nil(t, err)

  err = os.WriteFile(
    path.Join(s.testDir, "dir1/dir2/some.file"),
    []byte("some file"),
    0640)
  expect.Nil(t, err)

  err = os.WriteFile(
    path.Join(s.testDir, "dir1/dir2/other.file"),
    []byte("other file"),
    0604)
  expect.Nil(t, err)

  err = os.WriteFile(
    path.Join(s.testDir, "dir1/dir2/dir3/junk"),
    []byte("junk"),
    0600)
  expect.Nil(t, err)

  err = os.MkdirAll(
    path.Join(s.testDir, "foo/bar"),
    0750)
  expect.Nil(t, err)

  err = os.MkdirAll(
    path.Join(s.testDir, "foo/bar/abc"),
    0755)
  expect.Nil(t, err)

  err = os.MkdirAll(
    path.Join(s.testDir, "foo/bar/abc/def"),
    0700)
  expect.Nil(t, err)

  err = os.WriteFile(
    path.Join(s.testDir, "foo/bar/abc/def/ghi"),
    []byte("alphabets"),
    0600)
  expect.Nil(t, err)

  err = os.WriteFile(
    path.Join(s.testDir, "foo/file1"),
    []byte("file 1's content"),
    0666)
  expect.Nil(t, err)

  err = os.WriteFile(
    path.Join(s.testDir, "foo/file2"),
    []byte("file 2's content"),
    0660)
  expect.Nil(t, err)

  err = os.WriteFile(
    path.Join(s.testDir, "hello.txt"),
    []byte("hello"),
    0664)
  expect.Nil(t, err)

  err = os.WriteFile(
    path.Join(s.testDir, "world.txt"),
    []byte("world"),
    0604)
  expect.Nil(t, err)
}

func (s *LocalSuite) TestAbs(t *testing.T) {
  _, err := s.fsImpl.Abs("")
  expect.Error(t, err, "empty")

  currentDir, err := s.fsImpl.Abs(".")
  expect.Nil(t, err)

  info, err := s.fsImpl.Stat(currentDir)
  expect.Nil(t, err)
  expect.True(t, info.IsDir())
}

func (s *LocalSuite) TestOpenStat(t *testing.T) {
  reader, err := s.fsImpl.Open(s.toFsPath(t, path.Join(s.testDir, "dir1")))
  expect.Nil(t, err)

  info, err := reader.Stat()
  expect.Nil(t, err)

  expect.True(t, info.IsDir())
  expect.Equal(t, 0755, info.Mode() & PermissionBits)
}

func (s *LocalSuite) TestOpenRead(t *testing.T) {
  reader, err := s.fsImpl.Open(
    s.toFsPath(t, path.Join(s.testDir, "world.txt")))
  expect.Nil(t, err)

  content, err := io.ReadAll(reader)
  expect.Nil(t, err)

  expect.Equal(t, "world", string(content))

  err = reader.Close()
  expect.Nil(t, err)
}

func (s *LocalSuite) TestOpenReadDir(t *testing.T) {
  reader, err := s.fsImpl.Open(s.toFsPath(t, s.testDir))
  expect.Nil(t, err)

  entries, err := reader.ReadDir(0)
  expect.Nil(t, err)

  names := []string{}
  for _, entry := range entries {
    names = append(names, entry.Name())
  }

  // ReadDir entries are unsorted.
  sort.Strings(names)

  expect.Equal(
    t,
    []string{"abc", "dir1", "foo", "hello.txt", "world.txt"},
    names)
}

func (s *LocalSuite) TestCreateNewFileWithDefaultPerm(t *testing.T) {
  newOsFilePath := path.Join(s.testDir, "create-new-file")

  _, err := os.Stat(newOsFilePath)
  expect.True(t, IsNotExistError(err))

  writer, err := s.fsImpl.Create(s.toFsPath(t, newOsFilePath))
  expect.Nil(t, err)

  fileContent := "new file's content"

  _, err = writer.Write([]byte(fileContent))
  expect.Nil(t, err)

  err = writer.Close()
  expect.Nil(t, err)

  content, err := os.ReadFile(newOsFilePath)
  expect.Nil(t, err)

  expect.Equal(t, fileContent, string(content))

  info, err := os.Stat(newOsFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0644, info.Mode() & PermissionBits)
}

func (s *LocalSuite) TestCreateNewFileWithCustomPerm(t *testing.T) {
  newOsFilePath := path.Join(s.testDir, "create-new-file")

  _, err := os.Stat(newOsFilePath)
  expect.True(t, IsNotExistError(err))

  writer, err := s.fsImpl.Create(
    s.toFsPath(t, newOsFilePath),
    WithFilePerm(0604))
  expect.Nil(t, err)

  fileContent := "new file with custom perm's content"

  _, err = writer.Write([]byte(fileContent))
  expect.Nil(t, err)

  err = writer.Close()
  expect.Nil(t, err)

  content, err := os.ReadFile(newOsFilePath)
  expect.Nil(t, err)

  expect.Equal(t, fileContent, string(content))

  info, err := os.Stat(newOsFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0604, info.Mode() & PermissionBits)
}

func (s *LocalSuite) TestCreateTruncateExistingFile(t *testing.T) {
  testOsFilePath := path.Join(s.testDir, "dir1/dir2/some.file")

  // Sanity checks
  info, err := os.Stat(testOsFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0640, info.Mode() & PermissionBits)

  content, err := os.ReadFile(testOsFilePath)
  expect.Nil(t, err)
  expect.Equal(t, "some file", string(content))

  // Actual test

  writer, err := s.fsImpl.Create(
    s.toFsPath(t, testOsFilePath),
    WithFilePerm(0604))
  expect.Nil(t, err)

  newContent := "new file content"

  _, err = writer.Write([]byte(newContent))
  expect.Nil(t, err)

  err = writer.Close()
  expect.Nil(t, err)

  info, err = os.Stat(testOsFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())

  // Original file's permission is perserved.
  expect.Equal(t, 0640, info.Mode() & PermissionBits)

  content, err = os.ReadFile(testOsFilePath)
  expect.Nil(t, err)
  expect.Equal(t, newContent, string(content))
}

func (s *LocalSuite) TestAppendToNewFileWithDefaultPerm(t *testing.T) {
  newOsFilePath := path.Join(s.testDir, "create-new-file")

  _, err := os.Stat(newOsFilePath)
  expect.True(t, IsNotExistError(err))

  writer, err := s.fsImpl.Append(s.toFsPath(t, newOsFilePath))
  expect.Nil(t, err)

  fileContent := "new file's content"

  _, err = writer.Write([]byte(fileContent))
  expect.Nil(t, err)

  err = writer.Close()
  expect.Nil(t, err)

  content, err := os.ReadFile(newOsFilePath)
  expect.Nil(t, err)

  expect.Equal(t, fileContent, string(content))

  info, err := os.Stat(newOsFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0644, info.Mode() & PermissionBits)
}

func (s *LocalSuite) TestAppendToNewFileWithCustomPerm(t *testing.T) {
  newFilePath := path.Join(s.testDir, "create-new-file")

  _, err := os.Stat(newFilePath)
  expect.True(t, IsNotExistError(err))

  writer, err := s.fsImpl.Append(s.toFsPath(t, newFilePath), WithFilePerm(0604))
  expect.Nil(t, err)

  fileContent := "new file with custom perm's content"

  _, err = writer.Write([]byte(fileContent))
  expect.Nil(t, err)

  err = writer.Close()
  expect.Nil(t, err)

  content, err := os.ReadFile(newFilePath)
  expect.Nil(t, err)

  expect.Equal(t, fileContent, string(content))

  info, err := os.Stat(newFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0604, info.Mode() & PermissionBits)
}

func (s *LocalSuite) TestAppendToExistingFile(t *testing.T) {
  testFilePath := path.Join(s.testDir, "dir1/dir2/some.file")

  // Sanity checks
  info, err := os.Stat(testFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0640, info.Mode() & PermissionBits)

  content, err := os.ReadFile(testFilePath)
  expect.Nil(t, err)
  expect.Equal(t, "some file", string(content))

  // Actual test

  writer, err := s.fsImpl.Append(
    s.toFsPath(t, testFilePath),
    WithFilePerm(0604))
  expect.Nil(t, err)

  _, err = writer.Write([]byte(" more stuff"))
  expect.Nil(t, err)

  err = writer.Close()
  expect.Nil(t, err)

  info, err = os.Stat(testFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())

  // Original file's permission is perserved.
  expect.Equal(t, 0640, info.Mode() & PermissionBits)

  content, err = os.ReadFile(testFilePath)
  expect.Nil(t, err)
  expect.Equal(t, "some file more stuff", string(content))
}

func (s *LocalSuite) TestMkdirWithDefaultPerm(t *testing.T) {
  testFilePath := path.Join(s.testDir, "new-dir")

  // Sanity checks
  _, err := os.Stat(testFilePath)
  expect.True(t, IsNotExistError(err))

  err = s.fsImpl.Mkdir(s.toFsPath(t, testFilePath))
  expect.Nil(t, err)

  info, err := os.Stat(testFilePath)
  expect.Nil(t, err)

  expect.True(t, info.IsDir())
  expect.Equal(t, 0755, info.Mode() & PermissionBits)
}

func (s *LocalSuite) TestMkdirWithCustomPerm(t *testing.T) {
  testFilePath := path.Join(s.testDir, "new-dir")

  // Sanity checks
  _, err := os.Stat(testFilePath)
  expect.True(t, IsNotExistError(err))

  err = s.fsImpl.Mkdir(s.toFsPath(t, testFilePath), WithDirPerm(0705))
  expect.Nil(t, err)

  info, err := os.Stat(testFilePath)
  expect.Nil(t, err)

  expect.True(t, info.IsDir())
  expect.Equal(t, 0705, info.Mode() & PermissionBits)
}

func (s *LocalSuite) TestRename(t *testing.T) {
  destFilePath := path.Join(s.testDir, "renamed-file")

  _, err := os.Stat(destFilePath)
  expect.True(t, IsNotExistError(err))

  err = s.fsImpl.Rename(
    s.toFsPath(t, path.Join(s.testDir, "world.txt")),
    s.toFsPath(t, destFilePath))
  expect.Nil(t, err)

  info, err := os.Stat(destFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0604, info.Mode() & PermissionBits)

  content, err := os.ReadFile(destFilePath)
  expect.Nil(t, err)

  expect.Equal(t, "world", string(content))
}

func (s *LocalSuite) TestRemove(t *testing.T) {
  targetFilePath := path.Join(s.testDir, "hello.txt")

  _, err := os.Stat(targetFilePath)
  expect.Nil(t, err)

  err = s.fsImpl.Remove(s.toFsPath(t, targetFilePath))
  expect.Nil(t, err)

  _, err = os.Stat(targetFilePath)
  expect.True(t, IsNotExistError(err))
}

func (s *LocalSuite) TestChmod(t *testing.T) {
  testFilePath := path.Join(s.testDir, "abc")

  info, err := os.Stat(testFilePath)
  expect.Nil(t, err)
  expect.Equal(t, 0644, info.Mode() & PermissionBits)

  err = s.fsImpl.Chmod(s.toFsPath(t, testFilePath), 0600)
  expect.Nil(t, err)

  info, err = os.Stat(testFilePath)
  expect.Nil(t, err)
  expect.Equal(t, 0600, info.Mode() & PermissionBits)
}

func (s *LocalSuite) TestStat(t *testing.T) {
  info, err := s.fsImpl.Stat(s.toFsPath(t, path.Join(s.testDir, "dir1")))
  expect.Nil(t, err)

  expect.True(t, info.IsDir())
  expect.Equal(t, 0755, info.Mode() & PermissionBits)

  info, err = s.fsImpl.Stat(
    s.toFsPath(t, path.Join(s.testDir, "dir1/dir2/other.file")))
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0604, info.Mode() & PermissionBits)
}

func (s *LocalSuite) TestReadDir(t *testing.T) {
  result, err := s.fsImpl.ReadDir(s.toFsPath(t, s.testDir))
  expect.Nil(t, err)

  entries := make([]testDirEntry, 0, len(result))
  for _, entry := range result {
    entries = append(
      entries,
      testDirEntry{
        name: entry.Name(),
        isDir: entry.IsDir(),
      })
  }

  expect.Equal(
    t,
    []testDirEntry{
      {
        name: "abc",
      },
      {
        name: "dir1",
        isDir: true,
      },
      {
        name: "foo",
        isDir: true,
      },
      {
        name: "hello.txt",
      },
      {
        name: "world.txt",
      },
    },
    entries)
}

func (s *LocalSuite) TestWalkDir(t *testing.T) {
  paths := []string{}

  walkRoot := s.toFsPath(t, s.testDir)
  err := s.fsImpl.WalkDir(
    walkRoot,
    func(path string, _ DirEntry, err error) error {
      if err != nil {
        return err
      }

      if !strings.HasPrefix(path, walkRoot) {
        panic("walking unexpected path: " + path)
      }

      paths = append(paths, path[len(walkRoot):])
      return nil
    })
  expect.Nil(t, err)

  expect.Equal(
    t,
    []string{
      "",  // root
      "/abc",
      "/dir1",
      "/dir1/blah",
      "/dir1/dir2",
      "/dir1/dir2/dir3",
      "/dir1/dir2/dir3/junk",
      "/dir1/dir2/other.file",
      "/dir1/dir2/some.file",
      "/foo",
      "/foo/bar",
      "/foo/bar/abc",
      "/foo/bar/abc/def",
      "/foo/bar/abc/def/ghi",
      "/foo/file1",
      "/foo/file2",
      "/hello.txt",
      "/world.txt",
    },
    paths)
}

func (s *LocalSuite) TestGlob(t *testing.T) {
  matches, err := s.fsImpl.Glob(
    s.toFsPath(t, path.Join(s.testDir, "*.txt")))
  expect.Nil(t, err)

  expect.Equal(
    t,
    []string{
      s.toFsPath(t, path.Join(s.testDir, "hello.txt")),
      s.toFsPath(t, path.Join(s.testDir, "world.txt")),
    },
    matches)
}

func (s *LocalSuite) TestReadFile(t *testing.T) {
  content, err := s.fsImpl.ReadFile(
    s.toFsPath(t, path.Join(s.testDir, "hello.txt")))
  expect.Nil(t, err)

  expect.Equal(t, "hello", string(content))
}

func (s *LocalSuite) TestWriteFileNewFileWithDefaultPerm(t *testing.T) {
  testOsFilePath := path.Join(s.testDir, "some-new-file.txt")

  _, err := os.Stat(testOsFilePath)
  expect.True(t, IsNotExistError(err))

  err = s.fsImpl.WriteFile(s.toFsPath(t, testOsFilePath), []byte("3.14159"))
  expect.Nil(t, err)

  info, err := os.Stat(testOsFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0644, info.Mode() & PermissionBits)

  content, err := os.ReadFile(testOsFilePath)
  expect.Nil(t, err)
  expect.Equal(t, "3.14159", string(content))
}

func (s *LocalSuite) TestWriteFileNewFileWithCustomPerm(t *testing.T) {
  testOsFilePath := path.Join(s.testDir, "some-new-file.txt")

  _, err := os.Stat(testOsFilePath)
  expect.True(t, IsNotExistError(err))

  err = s.fsImpl.WriteFile(
    s.toFsPath(t, testOsFilePath),
    []byte("3.14159"),
    WithFilePerm(0604))
  expect.Nil(t, err)

  info, err := os.Stat(testOsFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0604, info.Mode() & PermissionBits)

  content, err := os.ReadFile(testOsFilePath)
  expect.Nil(t, err)
  expect.Equal(t, "3.14159", string(content))
}

func (s *LocalSuite) TestWriteFileOverwriteFile(t *testing.T) {
  testOsFilePath := path.Join(s.testDir, "abc")

  info, err := os.Stat(testOsFilePath)
  expect.Nil(t, err)
  expect.False(t, info.IsDir())

  err = s.fsImpl.WriteFile(
    s.toFsPath(t, testOsFilePath),
    []byte("overwrite"),
    WithFilePerm(0604))
  expect.Nil(t, err)

  info, err = os.Stat(testOsFilePath)
  expect.Nil(t, err)

  expect.False(t, info.IsDir())
  expect.Equal(t, 0644, info.Mode() & PermissionBits)

  content, err := os.ReadFile(testOsFilePath)
  expect.Nil(t, err)
  expect.Equal(t, "overwrite", string(content))
}

func (s *LocalSuite) TestCopyFileWithSourcePerm(t *testing.T) {
  destOsPath := path.Join(s.testDir, "foo/duplicate-junk")

  err := s.fsImpl.CopyFile(
    s.toFsPath(t, path.Join(s.testDir, "dir1/dir2/dir3/junk")),
    s.toFsPath(t, destOsPath))
  expect.Nil(t, err)

  info, err := os.Stat(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, 0600, info.Mode() & PermissionBits)

  content, err := os.ReadFile(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, "junk", string(content))
}

func (s *LocalSuite) TestCopyFileWithCustomPerm(t *testing.T) {
  destOsPath := path.Join(s.testDir, "foo/duplicate-junk")

  err := s.fsImpl.CopyFile(
    s.toFsPath(t, path.Join(s.testDir, "dir1/dir2/dir3/junk")),
    s.toFsPath(t, destOsPath),
    WithFilePerm(0640))
  expect.Nil(t, err)

  info, err := os.Stat(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, 0640, info.Mode() & PermissionBits)

  content, err := os.ReadFile(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, "junk", string(content))
}

func (s *LocalSuite) TestCopyFileOverwriteFile(t *testing.T) {
  destOsPath := path.Join(s.testDir, "world.txt")

  err := s.fsImpl.CopyFile(
    s.toFsPath(t, path.Join(s.testDir, "dir1/dir2/dir3/junk")),
    s.toFsPath(t, destOsPath),
    WithFilePerm(0640))
  expect.Nil(t, err)

  info, err := os.Stat(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, 0604, info.Mode() & PermissionBits)

  content, err := os.ReadFile(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, "junk", string(content))
}

func (s *LocalSuite) TestCopyFileCannotOverwriteDir(t *testing.T) {
  err := s.fsImpl.CopyFile(
    s.toFsPath(t, path.Join(s.testDir, "dir1/dir2/dir3/junk")),
    s.toFsPath(t, path.Join(s.testDir, "foo")),
    WithFilePerm(0640))
  expect.Error(t, err, "cannot overwrite directory")
}

func (s *LocalSuite) TestCopyAllFileWithSourcePerm(t *testing.T) {
  destOsPath := path.Join(s.testDir, "world2.txt")

  _, err := os.Stat(destOsPath)
  expect.True(t, IsNotExistError(err))

  err = s.fsImpl.CopyAll(
    s.toFsPath(t, path.Join(s.testDir, "world.txt")),
    s.toFsPath(t, destOsPath))
  expect.Nil(t, err)

  info, err := os.Stat(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, 0604, info.Mode() & PermissionBits)

  content, err := os.ReadFile(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, "world", string(content))
}

func (s *LocalSuite) TestCopyAllRecursiveDirCopy(t *testing.T) {
  destOsPath := path.Join(s.testDir, "foo2")

  err := s.fsImpl.CopyAll(
    s.toFsPath(t, path.Join(s.testDir, "foo")),
    s.toFsPath(t, destOsPath))
  expect.Nil(t, err)

  expectedContents := map[string]string{
    "foo2/bar/abc/def/ghi": "alphabets",
    "foo2/file1": "file 1's content",
    "foo2/file2": "file 2's content",
  }

  expectedPerms := map[string]FileMode {
    "foo2": 0750,
    "foo2/bar": 0750,
    "foo2/bar/abc": 0755,
    "foo2/bar/abc/def": 0700,
    "foo2/bar/abc/def/ghi": 0600,
    "foo2/file1": 0644,
    "foo2/file2": 0640,
  }

  for filePath, perm := range expectedPerms {
    fullPath := path.Join(s.testDir, filePath)
    info, err := os.Stat(fullPath)
    expect.Nil(t, err)
    expect.Equal(t, perm, info.Mode() & PermissionBits, filePath)

    expectedContent, ok := expectedContents[filePath]
    if ok {
      expect.False(t, info.IsDir(), filePath)

      content, err := os.ReadFile(fullPath)
      expect.Nil(t, err)
      expect.Equal(t, expectedContent, string(content))
    } else {
      expect.True(t, info.IsDir(), filePath)
    }
  }

  paths := []string{}
  walkRoot := s.toFsPath(t, destOsPath)
  err = s.fsImpl.WalkDir(
    walkRoot,
    func(path string, _ DirEntry, err error) error {
      if err != nil {
        return err
      }

      if !strings.HasPrefix(path, walkRoot) {
        panic("walking unexpected path: " + path)
      }

      paths = append(paths, path[len(walkRoot):])
      return nil
    })
  expect.Nil(t, err)

  expect.Equal(
    t,
    []string{
      "",
      "/bar",
      "/bar/abc",
      "/bar/abc/def",
      "/bar/abc/def/ghi",
      "/file1",
      "/file2",
    },
    paths)
}

func (s *LocalSuite) TestCopyAllDirWithSourcePerm(t *testing.T) {
  destOsPath := path.Join(s.testDir, "dir1/dir2/dir3-copy")

  err := s.fsImpl.CopyAll(
    s.toFsPath(t, path.Join(s.testDir, "dir1/dir2/dir3")),
    s.toFsPath(t, destOsPath))
  expect.Nil(t, err)

  info, err := os.Stat(destOsPath)
  expect.Nil(t, err)
  expect.True(t, info.IsDir())
  expect.Equal(t, 0755, info.Mode() & PermissionBits)

  entries, err := os.ReadDir(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, 1, len(entries))
  expect.Equal(t, "junk", entries[0].Name())

  junkCopyOsPath := path.Join(destOsPath, "junk")

  info, err = os.Stat(junkCopyOsPath)
  expect.Nil(t, err)
  expect.False(t, info.IsDir())
  expect.Equal(t, 0600, info.Mode() & PermissionBits)

  content, err := os.ReadFile(junkCopyOsPath)
  expect.Nil(t, err)
  expect.Equal(t, "junk", string(content))
}

func (s *LocalSuite) TestCopyAllFileWithCustomPerm(t *testing.T) {
  destOsPath := path.Join(s.testDir, "world2.txt")

  _, err := os.Stat(destOsPath)
  expect.True(t, IsNotExistError(err))

  err = s.fsImpl.CopyAll(
    s.toFsPath(t, path.Join(s.testDir, "world.txt")),
    s.toFsPath(t, destOsPath),
    WithFilePerm(0640))
  expect.Nil(t, err)

  info, err := os.Stat(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, 0640, info.Mode() & PermissionBits)

  content, err := os.ReadFile(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, "world", string(content))
}

func (s *LocalSuite) TestCopyAllDirWithCustomPerm(t *testing.T) {
  destOsPath := path.Join(s.testDir, "dir1/dir2/dir3-copy")

  err := s.fsImpl.CopyAll(
    s.toFsPath(t, path.Join(s.testDir, "dir1/dir2/dir3")),
    s.toFsPath(t, destOsPath),
    WithDirPerm(0700),
    WithFilePerm(0640))
  expect.Nil(t, err)

  info, err := os.Stat(destOsPath)
  expect.Nil(t, err)
  expect.True(t, info.IsDir())
  expect.Equal(t, 0700, info.Mode() & PermissionBits)

  entries, err := os.ReadDir(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, 1, len(entries))
  expect.Equal(t, "junk", entries[0].Name())

  junkCopyOsPath := path.Join(destOsPath, "junk")

  info, err = os.Stat(junkCopyOsPath)
  expect.Nil(t, err)
  expect.False(t, info.IsDir())
  expect.Equal(t, 0640, info.Mode() & PermissionBits)

  content, err := os.ReadFile(junkCopyOsPath)
  expect.Nil(t, err)
  expect.Equal(t, "junk", string(content))
}

func (s *LocalSuite) TestCopyAllFileOverwriteFile(t *testing.T) {
  destOsPath := path.Join(s.testDir, "world.txt")
  err := s.fsImpl.CopyAll(
    s.toFsPath(t, path.Join(s.testDir, "hello.txt")),
    s.toFsPath(t, destOsPath))
  expect.Nil(t, err)

  // world's original perm is preserved.
  info, err := os.Stat(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, 0604, info.Mode() & PermissionBits)

  // world's content is overwritten.
  content, err := os.ReadFile(destOsPath)
  expect.Nil(t, err)
  expect.Equal(t, "hello", string(content))
}

func (s *LocalSuite) TestCopyAllFileCannotOverwriteDir(t *testing.T) {
  err := s.fsImpl.CopyAll(
    s.toFsPath(t, path.Join(s.testDir, "hello.txt")),
    s.toFsPath(t, path.Join(s.testDir, "dir1")))
  expect.Error(t, err, "cannot overwrite directory")
}

func (s *LocalSuite) TestCopyAllDirCannotOverwriteFile(t *testing.T) {
  err := s.fsImpl.CopyAll(
    s.toFsPath(t, path.Join(s.testDir, "dir1")),
    s.toFsPath(t, path.Join(s.testDir, "hello.txt")))
  expect.Error(t, err, "cannot overwrite file")
}

func (s *LocalSuite) TestCopyAllDirCannotOverwriteDir(t *testing.T) {
  err := s.fsImpl.CopyAll(
    s.toFsPath(t, path.Join(s.testDir, "dir1")),
    s.toFsPath(t, path.Join(s.testDir, "foo")))
  expect.Error(t, err, "cannot overwrite directory")
}

func (s *LocalSuite) TestMkdirAllWithDefaultPerm(t *testing.T) {
  err := s.fsImpl.MkdirAll(
    s.toFsPath(t, path.Join(s.testDir, "foo/bar/a/b/c/d/e")))
  expect.Nil(t, err)

  paths := []string{
    "foo",
    "foo/bar",
    "foo/bar/a",
    "foo/bar/a/b",
    "foo/bar/a/b/c",
    "foo/bar/a/b/c/d",
    "foo/bar/a/b/c/d/e",
  }

  for _, filePath := range paths {
    info, err := os.Stat(path.Join(s.testDir, filePath))
    expect.Nil(t, err)
    expect.True(t, info.IsDir())

    if filePath == "foo" || filePath == "foo/bar" {
      expect.Equal(t, 0750, info.Mode() & PermissionBits)
    } else {
      expect.Equal(t, 0755, info.Mode() & PermissionBits)
    }
  }
}

func (s *LocalSuite) TestMkdirAllWithCustomPerm(t *testing.T) {
  err := s.fsImpl.MkdirAll(
    s.toFsPath(t, path.Join(s.testDir, "foo/bar/a/b/c/d/e")),
    WithDirPerm(0705))
  expect.Nil(t, err)

  paths := []string{
    "foo",
    "foo/bar",
    "foo/bar/a",
    "foo/bar/a/b",
    "foo/bar/a/b/c",
    "foo/bar/a/b/c/d",
    "foo/bar/a/b/c/d/e",
  }

  for _, filePath := range paths {
    info, err := os.Stat(path.Join(s.testDir, filePath))
    expect.Nil(t, err)
    expect.True(t, info.IsDir())

    if filePath == "foo" || filePath == "foo/bar" {
      expect.Equal(t, 0750, info.Mode() & PermissionBits)
    } else {
      expect.Equal(t, 0705, info.Mode() & PermissionBits)
    }
  }
}

func (s *LocalSuite) TestRemoveAll(t *testing.T) {
  walkRoot := s.toFsPath(t, s.testDir)
  listAll := func() []string {
    paths := []string{}
    err := s.fsImpl.WalkDir(
      walkRoot,
      func(path string, _ DirEntry, err error) error {
        if err != nil {
          return err
        }

        if !strings.HasPrefix(path, walkRoot) {
          panic("walking unexpected path: " + path)
        }

        paths = append(paths, path[len(walkRoot):])
        return nil
      })
    expect.Nil(t, err)

    return paths
  }

  expect.Equal(
    t,
    []string{
      "",  // root
      "/abc",
      "/dir1",
      "/dir1/blah",
      "/dir1/dir2",
      "/dir1/dir2/dir3",
      "/dir1/dir2/dir3/junk",
      "/dir1/dir2/other.file",
      "/dir1/dir2/some.file",
      "/foo",
      "/foo/bar",
      "/foo/bar/abc",
      "/foo/bar/abc/def",
      "/foo/bar/abc/def/ghi",
      "/foo/file1",
      "/foo/file2",
      "/hello.txt",
      "/world.txt",
    },
    listAll())

  err := s.fsImpl.RemoveAll(s.toFsPath(t, path.Join(s.testDir, "dir1")))
  expect.Nil(t, err)

  expect.Equal(
    t,
    []string{
      "",  // root
      "/abc",
      "/foo",
      "/foo/bar",
      "/foo/bar/abc",
      "/foo/bar/abc/def",
      "/foo/bar/abc/def/ghi",
      "/foo/file1",
      "/foo/file2",
      "/hello.txt",
      "/world.txt",
    },
    listAll())
}

func TestLocalFileSystem(t *testing.T) {
  local := &LocalSuite{
    fsImpl: Local,
    fsRootDir: "/",
  }
  suite.RunTests(t, local)
}
