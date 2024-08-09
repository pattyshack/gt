package filesystem

import (
  "io"
  "path"
  "testing"

  "github.com/pattyshack/gt/testing/expect"
  "github.com/pattyshack/gt/testing/suite"
)

type readOnlySuite struct {
  fs FileSystem
  testDir string
}

func (s *readOnlySuite) SetupTest(t *testing.T) {
  s.fs = NewReadOnlyFileSystem(Local)

  s.testDir = setupLocalTest(t)
}

func (s *readOnlySuite) TestAbs(t *testing.T) {
  absPath, err := s.fs.Abs(".")
  expect.Nil(t, err)
  expect.True(t, path.IsAbs(absPath))
}

func (s *readOnlySuite) TestOpen(t *testing.T) {
  reader, err := s.fs.Open(path.Join(s.testDir, "world.txt"))
  expect.Nil(t, err)

  content, err := io.ReadAll(reader)
  expect.Nil(t, err)

  expect.Equal(t, "world", string(content))

  err = reader.Close()
  expect.Nil(t, err)
}

func (s *readOnlySuite) TestCreate(t *testing.T) {
  _, err := s.fs.Create(path.Join(s.testDir, "new-file"))
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestAppend(t *testing.T) {
  _, err := s.fs.Append(path.Join(s.testDir, "new-file"))
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestMkdir(t *testing.T) {
  err := s.fs.Mkdir(path.Join(s.testDir, "new-dir"))
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestRename(t *testing.T) {
  err := s.fs.Rename(
    path.Join(s.testDir, "hello.txt"),
    path.Join(s.testDir, "hello2"))
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestRemove(t *testing.T) {
  err := s.fs.Remove(path.Join(s.testDir, "hello.txt"))
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestChmod(t *testing.T) {
  err := s.fs.Chmod(path.Join(s.testDir, "world.txt"), 0666)
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestStat(t *testing.T) {
  info, err := s.fs.Stat(path.Join(s.testDir, "world.txt"))
  expect.Nil(t, err)
  expect.False(t, info.IsDir())
  expect.Equal(t, 0604, info.Mode() & PermissionBits)
}

func (s *readOnlySuite) TestReadDir(t *testing.T) {
  entries, err := s.fs.ReadDir(path.Join(s.testDir, "foo/bar/abc/def"))
  expect.Nil(t, err)

  expect.Equal(t, 1, len(entries))
  expect.Equal(t, "ghi", entries[0].Name())
}

func (s *readOnlySuite) TestWalkDir(t *testing.T) {
  result := []string{}

  root := path.Join(s.testDir, "foo/bar")
  err := s.fs.WalkDir(
    root,
    func(filePath string, _ DirEntry, err error) error {
      if err != nil {
        return err
      }

      result = append(result, filePath[len(root):])
      return nil
    })
  expect.Nil(t, err)

  expect.Equal(
    t,
    []string{
      "",  // root
      "/abc",
      "/abc/def",
      "/abc/def/ghi",
    },
    result)
}

func (s *readOnlySuite) TestGlob(t *testing.T) {
  matches, err := s.fs.Glob(path.Join(s.testDir, "world*"))
  expect.Nil(t, err)
  expect.Equal(t, []string{path.Join(s.testDir,"world.txt")}, matches)
}

func (s *readOnlySuite) TestReadFile(t *testing.T) {
  content, err := s.fs.ReadFile(path.Join(s.testDir, "world.txt"))
  expect.Nil(t, err)
  expect.Equal(t, "world", string(content))
}

func (s *readOnlySuite) TestWriteFile(t *testing.T) {
  err := s.fs.WriteFile(path.Join(s.testDir, "new-file"), []byte("data"))
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestCopyFile(t *testing.T) {
  err := s.fs.CopyFile(
    path.Join(s.testDir, "hello.txt"),
    path.Join(s.testDir, "hello2"))
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestCopyAll(t *testing.T) {
  err := s.fs.CopyAll(path.Join(s.testDir, "foo"), path.Join(s.testDir, "zzz"))
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestMkdirAll(t *testing.T) {
  err := s.fs.MkdirAll(path.Join(s.testDir, "1/2/3/4"))
  expect.Error(t, err, "read-only mode")
}

func (s *readOnlySuite) TestRemoveAll(t *testing.T) {
  err := s.fs.RemoveAll(s.testDir)
  expect.Error(t, err, "read-only mode")
}

func TestReadOnlyFileSystem(t *testing.T) {
  suite.RunTests(t, &readOnlySuite{})
}
