package filesystem

import (
  "os"
  "path"
  "testing"

  "github.com/pattyshack/gt/testing/expect"
  "github.com/pattyshack/gt/testing/suite"
)

type SubLocalSuite struct {
  LocalSuite
}

func (s *SubLocalSuite) SetupTest(t *testing.T) {
  s.LocalSuite.SetupTest(t)

  // sub file system rooted at the tmp parent dir
  s.fsRootDir = path.Dir(s.testDir)
  s.fsImpl = NewSubFileSystem(Local, s.fsRootDir)
}

func (s *SubLocalSuite) TestAbs(t *testing.T) {
  t.Skip()  // ignore LocalSuite's TestAbs
}

func (s *SubLocalSuite) TestAbsWithWorkingDirectoryNotAtRoot(t *testing.T) {
  origWd, err := os.Getwd()
  expect.Nil(t, err)

  t.Cleanup(func() {
    err := os.Chdir(origWd)
    expect.Nil(t, err)
  })

  // Set working directory to a non-"root" sub directory
  err = os.Chdir(path.Join(s.testDir, "foo"))
  expect.Nil(t, err)

  // absolute paths are always ok / cleaned.

  filePath, err := s.fsImpl.Abs("/foo/abc/../bar")
  expect.Nil(t, err)
  expect.Equal(t, "/foo/bar", filePath)

  filePath, err = s.fsImpl.Abs("/../../..")
  expect.Nil(t, err)
  expect.Equal(t, "/", filePath)

  // relative paths are ok

  filePath, err = s.fsImpl.Abs(".")
  expect.Nil(t, err)
  expect.Equal(t, "/" + path.Base(s.testDir) + "/foo", filePath)

  filePath, err = s.fsImpl.Abs("bar")
  expect.Nil(t, err)
  expect.Equal(t, "/" + path.Base(s.testDir) + "/foo/bar", filePath)

  filePath, err = s.fsImpl.Abs("..")
  expect.Nil(t, err)
  expect.Equal(t, "/" + path.Base(s.testDir), filePath)

  filePath, err = s.fsImpl.Abs("../..")
  expect.Nil(t, err)
  expect.Equal(t, "/", filePath)

  filePath, err = s.fsImpl.Abs("../../..")
  expect.Nil(t, err)
  expect.Equal(t, "/", filePath)

  filePath, err = s.fsImpl.Abs("../../../foo/bar")
  expect.Nil(t, err)
  expect.Equal(t, "/foo/bar", filePath)
}

func (s *SubLocalSuite) TestAbsWithWorkingDirectoryAtRoot(t *testing.T) {
  origWd, err := os.Getwd()
  expect.Nil(t, err)

  t.Cleanup(func() {
    err := os.Chdir(origWd)
    expect.Nil(t, err)
  })

  // Set working directory to "root" directory
  err = os.Chdir(s.fsRootDir)
  expect.Nil(t, err)

  // absolute paths are always ok / cleaned.

  filePath, err := s.fsImpl.Abs("/foo/abc/../bar")
  expect.Nil(t, err)
  expect.Equal(t, "/foo/bar", filePath)

  filePath, err = s.fsImpl.Abs("/../../..")
  expect.Nil(t, err)
  expect.Equal(t, "/", filePath)

  // relative paths are ok

  filePath, err = s.fsImpl.Abs(".")
  expect.Nil(t, err)
  expect.Equal(t, "/", filePath)

  filePath, err = s.fsImpl.Abs("..")
  expect.Nil(t, err)
  expect.Equal(t, "/", filePath)

  filePath, err = s.fsImpl.Abs("foo")
  expect.Nil(t, err)
  expect.Equal(t, "/foo", filePath)

  filePath, err = s.fsImpl.Abs("../foo/bar")
  expect.Nil(t, err)
  expect.Equal(t, "/foo/bar", filePath)
}

func (s *SubLocalSuite) TestAbsWithoutWorkingDirectory(t *testing.T) {
  origWd, err := os.Getwd()
  expect.Nil(t, err)
  expect.NotEqual(t, "/", origWd)

  t.Cleanup(func() {
    err := os.Chdir(origWd)
    expect.Nil(t, err)
  })

  // underlying file system's working directory is inaccessible from sub file
  // system.
  err = os.Chdir("/")
  expect.Nil(t, err)

  // absolute paths are always ok / cleaned.

  filePath, err := s.fsImpl.Abs("/foo/abc/../bar")
  expect.Nil(t, err)
  expect.Equal(t, "/foo/bar", filePath)

  filePath, err = s.fsImpl.Abs("/../../..")
  expect.Nil(t, err)
  expect.Equal(t, "/", filePath)

  // relative paths are invalid

  _, err = s.fsImpl.Abs(".")
  expect.Error(t, err, "working directory unavailable")
}

func TestSubLocalFileSystem(t *testing.T) {
  sub := &SubLocalSuite{}
  suite.RunTests(t, sub)
}
