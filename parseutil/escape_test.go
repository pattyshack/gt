package parseutil

import (
	"testing"

	"github.com/pattyshack/gt/testing/expect"
	"github.com/pattyshack/gt/testing/suite"
)

type UnescapeStringSuite struct{}

func (UnescapeStringSuite) TestNoEscaped(t *testing.T) {
	expect.Equal(t, "asdf", Unescape("asdf"))
	expect.Equal(t, "世界", Unescape("世界"))
}

func (UnescapeStringSuite) TestBasicEscaped(t *testing.T) {
	expect.Equal(
		t,
		"a\as\bd\ff\nz\rx\tc\vv\\q'w\"e`r",
		Unescape("a\\as\\bd\\ff\\nz\\rx\\tc\\vv\\\\q\\'w\\\"e\\`r"))

	expect.Equal(t, "ab", Unescape("a\\\nb"))
}

func (UnescapeStringSuite) TestOctalEscaped(t *testing.T) {
	expect.Equal(t, "abcDEF", Unescape("a\\142c\\104E\\106"))
}

func (UnescapeStringSuite) Test2BytesHexEscaped(t *testing.T) {
	expect.Equal(t, "abcDEF", Unescape("a\\x62c\\x44E\\x46"))
}

func (UnescapeStringSuite) Test4BytesHexEscaped(t *testing.T) {
	expect.Equal(t, "a世b界c", Unescape("a\\u4e16b\\u754cc"))
}

func (UnescapeStringSuite) Test8BytesHexEscaped(t *testing.T) {
	expect.Equal(t, "a世b界c", Unescape("a\\U00004e16b\\U0000754cc"))
}

func TestUnescapeString(t *testing.T) {
	suite.RunTests(t, UnescapeStringSuite{})
}
