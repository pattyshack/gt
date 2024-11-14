package lexutil

import (
	"fmt"
	"io"
)

type Reader[T any] interface {
	// Use io.EOF to indicate end of reader stream.
	Read(buffer []T) (int, error)
}

type eofReader[T any] struct{}

func (eofReader[T]) Read(output []T) (int, error) {
	return 0, io.EOF
}

type StatsCollector[T any] interface {
	// Collect stats from items about to be discard/read from the BufferedReader.
	CollectStats([]T, error)
}

type BufferedReader[T any] struct {
	base Reader[T]

	buffer []T

	// buffer[startIdx:startIdx+numBuffered] holds the unconsumed buffered data
	startIdx    int
	numBuffered int

	statsCollector StatsCollector[T]
}

func NewBufferedReader[T any](
	base Reader[T],
	initialBufferSize int,
) *BufferedReader[T] {
	return NewBufferedReaderWithStatsCollector[T](base, initialBufferSize, nil)
}

func NewBufferedReaderWithStatsCollector[T any](
	base Reader[T],
	initialBufferSize int,
	statsCollector StatsCollector[T],
) *BufferedReader[T] {
	if initialBufferSize < 0 {
		initialBufferSize = 0
	}
	return newBufferedReader[T](
		base,
		make([]T, initialBufferSize),
		0,
		statsCollector)
}

func NewBufferedReaderFromSlice[T any](
	content []T,
) *BufferedReader[T] {
	return NewBufferedReaderWithStatsCollectorFromSlice[T](content, nil)
}

func NewBufferedReaderWithStatsCollectorFromSlice[T any](
	content []T,
	statsCollector StatsCollector[T],
) *BufferedReader[T] {
	return newBufferedReader[T](
		eofReader[T]{},
		content,
		len(content),
		statsCollector)
}

func newBufferedReader[T any](
	base Reader[T],
	content []T,
	numBuffered int,
	statsCollector StatsCollector[T],
) *BufferedReader[T] {
	return &BufferedReader[T]{
		base:           base,
		buffer:         content,
		startIdx:       0,
		numBuffered:    numBuffered,
		statsCollector: statsCollector,
	}
}

func (reader *BufferedReader[T]) fill(num int) (int, error) {
	if num < 0 {
		return reader.startIdx, fmt.Errorf("BufferedReader: negative fill number")
	}

	if reader.numBuffered >= num {
		return reader.startIdx + num, nil
	}

	_, ok := reader.base.(eofReader[T])
	if ok {
		return reader.startIdx + reader.numBuffered, io.EOF
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
	peeked, err := reader.Peek(num)

	if reader.statsCollector != nil {
		reader.statsCollector.CollectStats(peeked, err)
	}

	numDiscarded := reader.discard(num)
	return numDiscarded, err
}

func (reader *BufferedReader[T]) Read(output []T) (int, error) {
	peeked, err := reader.Peek(len(output))
	copy(output, peeked)

	if reader.statsCollector != nil {
		reader.statsCollector.CollectStats(peeked, err)
	}

	numDiscarded := reader.discard(len(peeked))
	return numDiscarded, err
}

func (reader *BufferedReader[T]) Next() (T, error) {
	var t T
	peeked, err := reader.Peek(1)
	if err != nil {
		return t, err
	}

	if len(peeked) != 1 {
		panic("should never happen")
	}

	t = peeked[0]
	_, err = reader.Discard(1)
	if err != nil {
		panic("should never happen")
	}

	return t, nil
}

type LocationStatsCollector struct {
	Location
}

func NewLocationStatsCollector(fileName string) *LocationStatsCollector {
	return &LocationStatsCollector{
		Location: Location{
			FileName: fileName,
			Line:     1,
			Column:   0,
		},
	}
}

func (collector *LocationStatsCollector) CollectStats(bytes []byte, err error) {
	for _, b := range bytes {
		if b == '\n' {
			collector.Line += 1
			collector.Column = 0
		} else {
			collector.Column += 1
		}
	}
}

type BufferedByteLocationReader struct {
	*BufferedReader[byte]
	*LocationStatsCollector
}

func NewBufferedByteLocationReader(
	fileName string,
	reader io.Reader,
	initialBufferSize int,
) BufferedByteLocationReader {
	collector := NewLocationStatsCollector(fileName)

	return BufferedByteLocationReader{
		BufferedReader: NewBufferedReaderWithStatsCollector[byte](
			reader,
			initialBufferSize,
			collector),
		LocationStatsCollector: collector,
	}
}

func NewBufferedByteLocationReaderFromSlice(
	fileName string,
	content []byte,
) BufferedByteLocationReader {
	collector := NewLocationStatsCollector(fileName)

	return BufferedByteLocationReader{
		BufferedReader: NewBufferedReaderWithStatsCollectorFromSlice[byte](
			content,
			collector),
		LocationStatsCollector: collector,
	}
}
