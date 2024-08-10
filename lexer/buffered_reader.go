package lexer

import (
	"fmt"
	"io"
)

type Reader[T any] interface {
	// Use io.EOF to indicate end of reader stream.
	Read(T []T) (int, error)
}

type BufferedReader[T any] struct {
	base Reader[T]

	buffer []T

	// buffer[startIdx:startIdx+numBuffered] holds the unconsumed buffered data
	startIdx    int
	numBuffered int
}

func NewBufferedReader[T any](
	base Reader[T],
	initialBufferSize int,
) *BufferedReader[T] {
	return &BufferedReader[T]{
		base:        base,
		buffer:      make([]T, initialBufferSize),
		startIdx:    0,
		numBuffered: 0,
	}
}

func (reader *BufferedReader[T]) fill(num int) (int, error) {
	if num < 0 {
		return reader.startIdx, fmt.Errorf("BufferedReader: negative fill number")
	}

	if reader.numBuffered >= num {
		return reader.startIdx + num, nil
	}

	if len(reader.buffer) < num {
		resizedBuffer := make([]T, num)
		copy(
			resizedBuffer,
			reader.buffer[reader.startIdx:reader.startIdx+reader.numBuffered])
		reader.buffer = resizedBuffer

		reader.startIdx = 0
	} else if len(reader.buffer) < reader.startIdx+num {
		copy(
			reader.buffer,
			reader.buffer[reader.startIdx:reader.startIdx+reader.numBuffered])

		reader.startIdx = 0
	}

	remaining := num - reader.numBuffered
	for remaining > 0 {
		numRead, err := reader.base.Read(
			reader.buffer[reader.startIdx+reader.numBuffered:])
		reader.numBuffered += numRead
		remaining -= numRead

		if err == io.ErrUnexpectedEOF {
			// unexpected eof is never checked by users ...
			err = io.EOF
		}

		if err != nil {
			return reader.startIdx + reader.numBuffered, err
		}
	}

	return reader.startIdx + num, nil
}

func (reader *BufferedReader[T]) Peek(num int) ([]T, error) {
	endIdx, err := reader.fill(num)
	return reader.buffer[reader.startIdx:endIdx], err
}

func (reader *BufferedReader[T]) discard(num int) int {
	if num <= 0 {
		return 0
	} else if reader.numBuffered > num {
		reader.startIdx += num
		reader.numBuffered -= num
	} else {
		num = reader.numBuffered
		reader.startIdx = 0
		reader.numBuffered = 0
	}

	return num
}

func (reader *BufferedReader[T]) Discard(num int) (int, error) {
	_, err := reader.fill(num)
	numDiscarded := reader.discard(num)
	return numDiscarded, err
}

func (reader *BufferedReader[T]) Read(output []T) (int, error) {
	peeked, err := reader.Peek(len(output))
	copy(output, peeked)

	numDiscarded := reader.discard(len(peeked))
	return numDiscarded, err
}
