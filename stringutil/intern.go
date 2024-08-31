package stringutil

import (
	"reflect"
	"unsafe"
)

func UnsafeBytesToString(bytes []byte) string {
	var unsafeString string
	bytesPtr := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	stringPtr := (*reflect.StringHeader)(unsafe.Pointer(&unsafeString))
	stringPtr.Data = bytesPtr.Data
	stringPtr.Len = bytesPtr.Len
	return unsafeString
}

type InternPool struct {
	pool map[string]string
}

func NewInternPool() *InternPool {
	return &InternPool{
		pool: map[string]string{},
	}
}

func (pool *InternPool) Intern(str string) string {
	interned, ok := pool.pool[str]
	if ok {
		return interned
	}

	pool.pool[str] = str
	return str
}

func (pool *InternPool) InternBytes(bytes []byte) string {
	interned, ok := pool.pool[UnsafeBytesToString(bytes)]
	if ok {
		return interned
	}

	str := string(bytes)
	pool.pool[str] = str
	return str
}

func (pool *InternPool) GetIntern(str string) (string, bool) {
	interned, ok := pool.pool[str]
	return interned, ok
}

func (pool *InternPool) GetInternBytes(bytes []byte) (string, bool) {
	interned, ok := pool.pool[UnsafeBytesToString(bytes)]
	return interned, ok
}
