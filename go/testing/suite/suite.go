package suite

import (
  "reflect"
  "strings"
  "testing"
)

var (
  // lower case -> camel case
  specialFuncNames = map[string]string {
    "setupsuite": "SetupSuite",
    "setuptest": "SetupTest",
    "setupbenchmark": "SetupBenchmark",
    "cleanupsuite": "CleanupSuite",
    "cleanuptest": "CleanupTest",
    "cleanupbenchmark": "CleanupBenchmark",
  }
)

// If this interface is defined for the suite, then RunTests/RunBenchmarks
// will call this setup method before any of the suite's tests/benchmarks.
// This is called once per RunTests/RunBenchmarks.
type SuiteSetup interface {
  // CAUTION: This expects testing.TB rather than *testing.T or *testing.B.
  SetupSuite(tb testing.TB)
}

// If this interface is defined for the suite, then RunTests/RunBenchmarks
// will call this cleanup method after tests/benchmark completed.  This is
// called once per RunTests/RunBenchmarks.
type SuiteCleanup interface {
  // CAUTION: This expects testing.TB rather than *testing.T or *testing.B.
  CleanupSuite(tb testing.TB)
}

// If this interface is defined for the suite, then RunTests will call this
// setup method before each test in the suite.
type TestSetup interface {
  SetupTest(t *testing.T)
}

// If this interface is defined for the suite, then RunTests will call this
// cleanup method after each test in the suite.
type TestCleanup interface {
  CleanupTest(t *testing.T)
}

// If this interface is defined for the suite, then RunBenchmarks will call
// this setup method before each benchmark test run in the suite.
type BenchmarkSetup interface {
  SetupBenchmark(b *testing.B)
}

// If this interface is defined for the suite, then RunBenchmarks will call
// this cleanup method after each benchmark test run in the suite.
type BenchmarkCleanup interface {
  CleanupBenchmark(b *testing.B)
}

func checkSuiteSignatures[SuiteT any](tb testing.TB, suite SuiteT) {
  tb.Helper()

  suiteValue := reflect.ValueOf(suite)
  suiteType := reflect.TypeOf(suite)

  for i := 0; i < suiteValue.NumMethod(); i++ {
    methodName := suiteType.Method(i).Name

    // suiteValue.Method instead of suiteType.Method to get the receiver-binded
    // method.
    methodValue := suiteValue.Method(i)

    if strings.HasPrefix(methodName, "Test") {
      _, ok := methodValue.Interface().(func(*testing.T))
      if !ok {
        tb.Fatalf("Found unexpected method signature for %s", methodName)
      }

    } else if strings.HasPrefix(methodName, "Benchmark") {
      _, ok := methodValue.Interface().(func(*testing.B))
      if !ok {
        tb.Fatalf("Found unexpected method signature for %s", methodName)
      }

    } else {
      lowerMethodName := strings.ToLower(methodName)

      specialFuncName, ok := specialFuncNames[lowerMethodName]
      if !ok {
        continue
      }

      if methodName != specialFuncName {
        // This is probably a typo.
        tb.Errorf(
          "Found %s method defined for suite (expected %s)",
          methodName,
          specialFuncName)
      }

      if strings.HasSuffix(methodName, "Suite") {
        _, ok := methodValue.Interface().(func(testing.TB))
        if !ok {
          tb.Fatalf("Found unexpected method signature for %s", methodName)
        }

      } else if strings.HasSuffix(methodName, "Test") {
        _, ok := methodValue.Interface().(func(*testing.T))
        if !ok {
          tb.Fatalf("Found unexpected method signature for %s", methodName)
        }
      } else if strings.HasSuffix(methodName, "Benchmark") {
        _, ok := methodValue.Interface().(func(*testing.B))
        if !ok {
          tb.Fatalf("Found unexpected method signature for %s", methodName)
        }
      }
    }
  }
}

// Suite's tests are of the form
//    `func (testSuite) Test<Name>(*testing.T)`
// RunTests will make use of SuiteSetup, TestSetup, TestCleanup, and
// SuiteCleanup if they are defined.
//
// Note that the suite's name is not used for constructing the test cases'
// names.
func RunTests[SuiteT any](outerT *testing.T, suite SuiteT) {
  outerT.Helper()

  checkSuiteSignatures(outerT, suite)

  var iSuite any = suite  // We can't type cast suite directly

  cleanupSuite, ok := iSuite.(SuiteCleanup)
  if ok {
    outerT.Cleanup(func() { cleanupSuite.CleanupSuite(outerT) })
  }

  setupSuite, ok := iSuite.(SuiteSetup)
  if ok {
    setupSuite.SetupSuite(outerT)
  }

  setupTest, ok := iSuite.(TestSetup)
  if !ok {
    setupTest = nil
  }

  cleanupTest, ok := iSuite.(TestCleanup)
  if !ok {
    cleanupTest = nil
  }

  suiteValue := reflect.ValueOf(suite)
  suiteType := reflect.TypeOf(suite)

  for i := 0; i < suiteValue.NumMethod(); i++ {
    methodName := suiteType.Method(i).Name

    if !strings.HasPrefix(methodName, "Test") {
      continue
    }

    // suiteValue.Method instead of suiteType.Method to get the receiver-binded
    // method.
    methodValue := suiteValue.Method(i)

    // We've already verify the signature
    testFunc := methodValue.Interface().(func(*testing.T))

    outerT.Run(
      methodName[4:],  // trim "Test" prefix
      func(innerT *testing.T) {
        if cleanupTest != nil {
          innerT.Cleanup(func() { cleanupTest.CleanupTest(innerT) })
        }

        if setupTest != nil {
          setupTest.SetupTest(innerT)
        }

        testFunc(innerT)
      })
  }
}

// Suite's benchmarks are of the form
//    `func (testSuite) Benchmark<Name>(*testing.B)`
// RunBenchmarks will make use of SuiteSetup, BenchmarkSetup,
// BenchmarkCleanup, and SuiteCleanup if they are defined.
//
// Note that the suite's name is not used for constructing the benchmark cases'
// names.
func RunBenchmarks[SuiteT any](outerB *testing.B, suite SuiteT) {
  outerB.Helper()

  outerB.StopTimer()

  checkSuiteSignatures(outerB, suite)

  var iSuite any = suite  // We can't type cast suite directly

  cleanupSuite, ok := iSuite.(SuiteCleanup)
  if ok {
    outerB.Cleanup(func() { cleanupSuite.CleanupSuite(outerB) })
  }

  setupSuite, ok := iSuite.(SuiteSetup)
  if ok {
    setupSuite.SetupSuite(outerB)
  }

  setupBenchmark, ok := iSuite.(BenchmarkSetup)
  if !ok {
    setupBenchmark = nil
  }

  cleanupBenchmark, ok := iSuite.(BenchmarkCleanup)
  if !ok {
    cleanupBenchmark = nil
  }

  suiteValue := reflect.ValueOf(suite)
  suiteType := reflect.TypeOf(suite)

  for i := 0; i < suiteValue.NumMethod(); i++ {
    methodName := suiteType.Method(i).Name

    if !strings.HasPrefix(methodName, "Benchmark") {
      continue
    }

    // suiteValue.Method instead of suiteType.Method to get the receiver-binded
    // method.
    methodValue := suiteValue.Method(i)

    // We've already verify the signature
    benchmarkFunc := methodValue.Interface().(func(*testing.B))

    outerB.Run(
      methodName[9:],  // trim "Benchmark" prefix
      func(innerB *testing.B) {
        innerB.StopTimer()

        if cleanupBenchmark != nil {
          innerB.Cleanup(func() { cleanupBenchmark.CleanupBenchmark(innerB) })
        }


        if setupBenchmark != nil {
          setupBenchmark.SetupBenchmark(innerB)
        }

        innerB.StartTimer()
        benchmarkFunc(innerB)
        innerB.StopTimer()
      })
  }
}
