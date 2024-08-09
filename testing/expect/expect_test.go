package expect_test

import (
  "fmt"
  "testing"

  "github.com/pattyshack/gt/testing/expect"
)

type testStruct struct {
  X int
  y string
}

func (testStruct) Foo() {}

type testInterface interface {
  Foo()
}

func TestSame(t *testing.T) {
  expect.Same(t, 1, 1)
  expect.Same(t, 1.0, 1.0)

  expect.Same(t, "hello", "hello")

  s1 := testStruct{14, "foo"}
  s2 := testStruct{14, "foo"}
  expect.Same(t, s1, s2)

  x := new(int)
  y := x
  expect.Same(t, x, y)
}

func TestSameFail(t *testing.T) {
  // Compiler error
  // expect.Same(t, 1, 1.0)

  skip := true

  // Runtime test failures
  t.Run("int", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Same(t, 1, 3)
  })

  t.Run("float64", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Same(t, 1.0, 3.0)
  })

  t.Run("string", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Same(t, "foo", "bar")
  })

  t.Run("struct", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Same(t, testStruct{14, "foo"}, testStruct{14, "abc"})
  })

  t.Run("struct pointer", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    x := testStruct{14, "foo"}
    y := x
    expect.Same(t, &x, &y)
  })
}

func TestEqual(t *testing.T) {
  expect.Equal(t, 1, 1)
  expect.Equal(t, 1.0, 1.0)

  expect.Equal(t, "hello", "hello")
  expect.Equal(t, []byte("hi"), []byte{'h', 'i'})

  expect.Equal(t, []int{1, 3, 5, 7}, []int{1, 3, 5, 7})

  x := 1
  y := 1
  expect.Equal(t, &x, &y)

  s1 := testStruct{14, "foo"}
  s2 := testStruct{14, "foo"}
  expect.Equal(t, s1, s2)
  expect.Equal(t, &s1, &s2)
}

func TestEqualFail(t *testing.T) {
  // Compiler error
  // expect.Equal(t, 1, 1.0)

  skip := true

  // Runtime test failures
  t.Run("int", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Equal(t, 1, 3, "different ints %v != %d", 1, 3)
  })

  t.Run("float64", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Equal(t, 1.0, 3.0)
  })

  t.Run("string", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Equal(t, "foo", "bar")
  })

  t.Run("empty vs nil slice", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Equal(t, nil, []byte{})
  })

  t.Run("slice", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Equal(t, []int{1, 2}, []int{1, 2, 3})
  })

  t.Run("struct", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Equal(t, testStruct{14, "foo"}, testStruct{14, "abc"})
  })

  t.Run("struct pointer", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Equal(t, &testStruct{14, "foo"}, &testStruct{14, "abc"})
  })
}

func TestNotEqual(t *testing.T) {
  expect.NotEqual(t, 1, 3)
  expect.NotEqual(t, 1.0, 3.0)

  expect.NotEqual(t, "foo", "bar")
  expect.NotEqual(t, nil, []byte{})

  expect.NotEqual(t, []int{1, 2}, []int{1, 2, 3})

  expect.NotEqual(t, testStruct{14, "foo"}, testStruct{14, "abc"})

  expect.NotEqual(t, &testStruct{14, "foo"}, &testStruct{14, "abc"})
}

func TestNotEqualFail(t *testing.T) {
  // Compiler error
  // expect.NotEqual(t, 1, 1.0)

  skip := true

  t.Run("int", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.NotEqual(t, 1, 1)
  })

  t.Run("float64", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.NotEqual(t, 1.0, 1.0)
  })

  t.Run("string", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.NotEqual(t, "hello", "hello")
  })

  t.Run("byte slice", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.NotEqual(t, []byte("hi"), []byte{'h', 'i'})
  })

  t.Run("int slice", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.NotEqual(t, []int{1, 3, 5, 7}, []int{1, 3, 5, 7})
  })

  t.Run("int pointer", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    x := 1
    y := 1
    expect.NotEqual(t, &x, &y)
  })

  t.Run("struct", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    s1 := testStruct{14, "foo"}
    s2 := testStruct{14, "foo"}
    expect.NotEqual(t, s1, s2)
  })

  t.Run("struct pointer", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    s1 := testStruct{14, "foo"}
    s2 := testStruct{14, "foo"}
    expect.NotEqual(t, &s1, &s2)
  })
}

func TestNil(t *testing.T) {
  expect.Nil(t, nil, "untyped nil")

  var intSlice []int
  expect.Nil(t, intSlice)

  var intPtr *int
  expect.Nil(t, intPtr)

  var structPtr *testStruct
  expect.Nil(t, structPtr)

  var testIface testInterface
  expect.Nil(t, testIface)

  expect.Nil(t, func() error { return nil }())

  var testChan chan float64
  expect.Nil(t, testChan)

  var testMap map[uint]bool
  expect.Nil(t, testMap)
}

func TestNilFail(t *testing.T) {
  skip := true

  t.Run("int", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Nil(t, 1)
  })

  t.Run("byte slice", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Nil(t, []byte{})
  })

  t.Run("string slice", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Nil(t, []string{"a"})
  })

  t.Run("int ptr", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    intPtr := new(int)
    expect.Nil(t, intPtr)
  })

  t.Run("struct ptr", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    structPtr := &testStruct{}
    expect.Nil(t, structPtr)
  })

  t.Run("interface", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    var testIface testInterface = &testStruct{}
    expect.Nil(t, testIface)
  })

  t.Run("error", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Nil(t, func() error { return fmt.Errorf("error") }())
  })

  t.Run("chan", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    testChan := make(chan float64)
    expect.Nil(t, testChan)
  })

  t.Run("map", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    testMap := map[uint]bool{}
    expect.Nil(t, testMap)
  })
}

func TestNoNil(t *testing.T) {
  expect.NotNil(t, "string")
  expect.NotNil(t, new(int))
}

func TestNoNilFail(t *testing.T) {
  t.Skip()
  expect.NotNil(t, nil)
}

func TestError(t *testing.T) {
  expect.Error(t, fmt.Errorf("bad"), "")
  expect.Error(
    t,
    func() error { return fmt.Errorf("... snippet ...") }(),
    "snippet")
}

func TestErrorFail(t *testing.T) {
  skip := true

  t.Run("nil", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Error(t, func() error { return nil }(), "")
  })

  t.Run("snippet", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.Error(
      t,
      func() error { return fmt.Errorf("... snippet ...") }(),
      "other")
  })
}

func TestTrueFalse(t *testing.T) {
  expect.True(t, true)
  expect.False(t, false)
}

func TestTrueFalseFail(t *testing.T) {
  skip := true

  t.Run("true", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.True(t, false, "something went wrong")
  })

  t.Run("false", func(t *testing.T) {
    if skip {
      t.Skip()
    }
    expect.False(t, true)
  })
}
