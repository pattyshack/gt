package filesystem

import (
  "testing"

  "github.com/pattyshack/gt/testing/suite"
)

func TestExtendedMinimalLocalFileSystem(t *testing.T) {
  extended := &LocalSuite{
    fsImpl: ExtendMinimalFileSystem(minLocalFileSystem{}),
    fsRootDir: "/",
  }
  suite.RunTests(t, extended)
}
