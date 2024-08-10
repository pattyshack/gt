package lexer

import (
	"bytes"
	"io"
	"testing"

	"github.com/pattyshack/gt/testing/expect"
	"github.com/pattyshack/gt/testing/suite"
)

type BufferedReaderSuite struct {
	content  string
	base     *bytes.Buffer
	buffered *BufferedReader[byte]
}

func (s *BufferedReaderSuite) SetupTest(t *testing.T) {
	s.content = "0123456789"
	s.base = bytes.NewBufferString(s.content)
	s.buffered = NewBufferedReader(s.base, 4)
}

func (s *BufferedReaderSuite) TestPeekNegativeNum(t *testing.T) {
	bytes, err := s.buffered.Peek(-1)
	expect.Equal(t, "", string(bytes))
	expect.Error(t, err, "negative")
	expect.NotEqual(t, err, io.EOF)
	expect.Equal(t, []byte{0, 0, 0, 0}, s.buffered.buffer)
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(1)
	expect.Equal(t, "0", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 4, s.buffered.numBuffered)

	numDiscarded, err := s.buffered.Discard(3)
	expect.Equal(t, 3, numDiscarded)
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 3, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(-1)
	expect.Equal(t, "", string(bytes))
	expect.Error(t, err, "negative")
	expect.NotEqual(t, err, io.EOF)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 3, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(1)
	expect.Equal(t, "3", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 3, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	expect.Equal(t, "456789", string(s.base.Bytes()))
}

func (s *BufferedReaderSuite) TestPeekZeroNum(t *testing.T) {
	bytes, err := s.buffered.Peek(0)
	expect.Equal(t, "", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, []byte{0, 0, 0, 0}, s.buffered.buffer)
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(2)
	expect.Equal(t, "01", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 4, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(0)
	expect.Equal(t, "", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 4, s.buffered.numBuffered)

	readBytes := make([]byte, 2)
	numRead, err := s.buffered.Read(readBytes)
	expect.Nil(t, err)
	expect.Equal(t, 2, numRead)
	expect.Equal(t, "01", string(readBytes))
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 2, s.buffered.startIdx)
	expect.Equal(t, 2, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(0)
	expect.Equal(t, "", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 2, s.buffered.startIdx)
	expect.Equal(t, 2, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(readBytes[:1])
	expect.Nil(t, err)
	expect.Equal(t, 1, numRead)
	expect.Equal(t, "2", string(readBytes[:1]))
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 3, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(0)
	expect.Equal(t, "", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 3, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	expect.Equal(t, "456789", string(s.base.Bytes()))
}

func (s *BufferedReaderSuite) TestPeekNoRead(t *testing.T) {
	bytes, err := s.buffered.Peek(2)
	expect.Equal(t, "01", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 4, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(1)
	expect.Equal(t, "0", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 4, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(4)
	expect.Equal(t, "0123", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 4, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(3)
	expect.Equal(t, "012", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 4, s.buffered.numBuffered)

	numDiscarded, err := s.buffered.Discard(1)
	expect.Nil(t, err)
	expect.Equal(t, 1, numDiscarded)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(2)
	expect.Equal(t, "12", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(1)
	expect.Equal(t, "1", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(3)
	expect.Equal(t, "123", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)

	expect.Equal(t, "456789", string(s.base.Bytes()))
}

func (s *BufferedReaderSuite) TestPeekResizeAndShiftBuffer(t *testing.T) {
	s.buffered = NewBufferedReader(s.base, 2)

	bytes, err := s.buffered.Peek(1)
	expect.Equal(t, "0", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "01", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 2, s.buffered.numBuffered)

	numDiscarded, err := s.buffered.Discard(1)
	expect.Equal(t, 1, numDiscarded)
	expect.Nil(t, err)
	expect.Equal(t, "01", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(3)
	expect.Equal(t, "123", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "123", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)

	numDiscarded, err = s.buffered.Discard(2)
	expect.Equal(t, 2, numDiscarded)
	expect.Nil(t, err)
	expect.Equal(t, "123", string(s.buffered.buffer))
	expect.Equal(t, 2, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(4)
	expect.Equal(t, "3456", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "3456", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 4, s.buffered.numBuffered)

	numDiscarded, err = s.buffered.Discard(1)
	expect.Equal(t, 1, numDiscarded)
	expect.Nil(t, err)
	expect.Equal(t, "3456", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)

	bytes, err = s.buffered.Peek(10)
	expect.Equal(t, io.EOF, err)
	expect.Equal(t, "456789", string(bytes))
	expect.Equal(
		t,
		[]byte{'4', '5', '6', '7', '8', '9', 0, 0, 0, 0},
		s.buffered.buffer)
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 6, s.buffered.numBuffered)

	expect.Equal(t, "", string(s.base.Bytes()))
}

func (s *BufferedReaderSuite) TestPeekReadWithoutShift(t *testing.T) {
	bytes, err := s.buffered.Peek(1)
	expect.Equal(t, "0", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 4, s.buffered.numBuffered)

	numDiscarded, err := s.buffered.Discard(1)
	expect.Equal(t, 1, numDiscarded)
	expect.Nil(t, err)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)

	// Simulate previous Read did not completely fill the entire buffer, e.g.,
	// read from network socket.
	s.buffered.buffer = []byte{'0', '1', '2', '3', 0, 0, 0, 0}

	bytes, err = s.buffered.Peek(5)
	expect.Equal(t, "12345", string(bytes))
	expect.Nil(t, err)
	expect.Equal(t, "01234567", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 7, s.buffered.numBuffered)

	expect.Equal(t, "89", string(s.base.Bytes()))
}

func (s *BufferedReaderSuite) TestReadShiftBuffer(t *testing.T) {
	bytes := make([]byte, 3)
	numRead, err := s.buffered.Read(bytes)
	expect.Nil(t, err)
	expect.Equal(t, 3, numRead)
	expect.Equal(t, "012", string(bytes))
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 3, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(bytes[:2])
	expect.Nil(t, err)
	expect.Equal(t, 2, numRead)
	expect.Equal(t, "34", string(bytes[:2]))
	expect.Equal(t, "3456", string(s.buffered.buffer))
	expect.Equal(t, 2, s.buffered.startIdx)
	expect.Equal(t, 2, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(bytes)
	expect.Nil(t, err)
	expect.Equal(t, 3, numRead)
	expect.Equal(t, "567", string(bytes))
	expect.Equal(t, "5678", string(s.buffered.buffer))
	expect.Equal(t, 3, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(bytes)
	expect.Equal(t, io.EOF, err)
	expect.Equal(t, 2, numRead)
	expect.Equal(t, "89", string(bytes[:2]))
	expect.Equal(t, "8978", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(bytes)
	expect.Equal(t, io.EOF, err)
	expect.Equal(t, 0, numRead)
	expect.Equal(t, "8978", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)
}

func (s *BufferedReaderSuite) TestDiscardShiftBuffer(t *testing.T) {
	numDiscard, err := s.buffered.Discard(2)
	expect.Nil(t, err)
	expect.Equal(t, 2, numDiscard)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 2, s.buffered.startIdx)
	expect.Equal(t, 2, s.buffered.numBuffered)

	numDiscard, err = s.buffered.Discard(3)
	expect.Nil(t, err)
	expect.Equal(t, 3, numDiscard)
	expect.Equal(t, "2345", string(s.buffered.buffer))
	expect.Equal(t, 3, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	numDiscard, err = s.buffered.Discard(1)
	expect.Nil(t, err)
	expect.Equal(t, 1, numDiscard)
	expect.Equal(t, "2345", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)

	numDiscard, err = s.buffered.Discard(3)
	expect.Nil(t, err)
	expect.Equal(t, 3, numDiscard)
	expect.Equal(t, "6789", string(s.buffered.buffer))
	expect.Equal(t, 3, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	numDiscard, err = s.buffered.Discard(3)
	expect.Equal(t, io.EOF, err)
	expect.Equal(t, 1, numDiscard)
	expect.Equal(t, "9789", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)

	numDiscard, err = s.buffered.Discard(1)
	expect.Equal(t, io.EOF, err)
	expect.Equal(t, 0, numDiscard)
	expect.Equal(t, "9789", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)
}

func (s *BufferedReaderSuite) TestReadExtraNoEofErr(t *testing.T) {
	bytes := make([]byte, 8)
	numRead, err := s.buffered.Read(bytes)
	expect.Nil(t, err)
	expect.Equal(t, 8, numRead)
	expect.Equal(t, "01234567", string(bytes))
	expect.Equal(t, "01234567", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(bytes[:1])
	expect.Nil(t, err)
	expect.Equal(t, 1, numRead)
	expect.Equal(t, "8", string(bytes[:1]))
	expect.Equal(t, "89234567", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 1, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(bytes[:1])
	expect.Nil(t, err)
	expect.Equal(t, 1, numRead)
	expect.Equal(t, "9", string(bytes[:1]))
	expect.Equal(t, "89234567", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(bytes[:1])
	expect.Equal(t, io.EOF, err)
	expect.Equal(t, 0, numRead)
	expect.Equal(t, "89234567", string(s.buffered.buffer))
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)
}

func (s *BufferedReaderSuite) TestDiscardNegativeNum(t *testing.T) {
	numDiscarded, err := s.buffered.Discard(-1)
	expect.Error(t, err, "negative")
	expect.Equal(t, 0, numDiscarded)
	expect.Equal(t, []byte{0, 0, 0, 0}, s.buffered.buffer)
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)

	bytes := make([]byte, 1)
	numRead, err := s.buffered.Read(bytes)
	expect.Nil(t, err)
	expect.Equal(t, 1, numRead)
	expect.Equal(t, "0", string(bytes))
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)

	numDiscarded, err = s.buffered.Discard(-1)
	expect.Error(t, err, "negative")
	expect.Equal(t, 0, numDiscarded)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)
}

func (s *BufferedReaderSuite) TestReadZeroNum(t *testing.T) {
	bytes := make([]byte, 1)

	numRead, err := s.buffered.Read(bytes[:0])
	expect.Nil(t, err)
	expect.Equal(t, 0, numRead)
	expect.Equal(t, []byte{0, 0, 0, 0}, s.buffered.buffer)
	expect.Equal(t, 0, s.buffered.startIdx)
	expect.Equal(t, 0, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(bytes)
	expect.Nil(t, err)
	expect.Equal(t, 1, numRead)
	expect.Equal(t, "0", string(bytes))
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)

	numRead, err = s.buffered.Read(bytes[:0])
	expect.Nil(t, err)
	expect.Equal(t, 0, numRead)
	expect.Equal(t, "0123", string(s.buffered.buffer))
	expect.Equal(t, 1, s.buffered.startIdx)
	expect.Equal(t, 3, s.buffered.numBuffered)
}

func TestBufferedReader(t *testing.T) {
	suite.RunTests(t, &BufferedReaderSuite{})
}
