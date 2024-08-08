package suite_test

import (
	"testing"

	"github.com/pattyshack/bt/go/testing/expect"
	"github.com/pattyshack/bt/go/testing/suite"
)

// Note: suite's struct name is not included in the test's full name.
type Suite struct {
	setupSuite   map[string]bool
	cleanupSuite map[string]bool

	setupTests   map[string]bool
	tests        map[string]bool
	cleanupTests map[string]bool

	failTestSetup   bool
	failTest        bool
	failTestCleanup bool

	setupBenchmarks   map[string]int
	benchmarks        map[string]int
	cleanupBenchmarks map[string]int

	failBenchmarkSetup   bool
	failBenchmark        bool
	failBenchmarkCleanup bool
}

func newSuite(
	failTestSetup bool,
	failTest bool,
	failTestCleanup bool,
	failBenchmarkSetup bool,
	failBenchmark bool,
	failBenchmarkCleanup bool,
) *Suite {
	return &Suite{
		setupSuite:   map[string]bool{},
		cleanupSuite: map[string]bool{},

		setupTests:   map[string]bool{},
		tests:        map[string]bool{},
		cleanupTests: map[string]bool{},

		failTestSetup:   failTestSetup,
		failTest:        failTest,
		failTestCleanup: failTestCleanup,

		setupBenchmarks:   map[string]int{},
		benchmarks:        map[string]int{},
		cleanupBenchmarks: map[string]int{},

		failBenchmarkSetup:   failBenchmarkSetup,
		failBenchmark:        failBenchmark,
		failBenchmarkCleanup: failBenchmarkCleanup,
	}
}

func (s *Suite) SetupSuite(tb testing.TB) {
	expect.Equal(tb, 0, len(s.setupSuite))
	expect.Equal(tb, 0, len(s.setupTests))
	expect.Equal(tb, 0, len(s.setupBenchmarks))

	expect.Equal(tb, 0, len(s.tests))
	expect.Equal(tb, 0, len(s.benchmarks))

	expect.Equal(tb, 0, len(s.cleanupSuite))
	expect.Equal(tb, 0, len(s.cleanupTests))
	expect.Equal(tb, 0, len(s.cleanupBenchmarks))

	s.setupSuite[tb.Name()] = true
}

func (s *Suite) CleanupSuite(tb testing.TB) {
	expect.Equal(tb, 1, len(s.setupSuite))
	expect.Equal(tb, 0, len(s.cleanupSuite))

	expect.True(tb, s.setupSuite[tb.Name()])
	expect.False(tb, s.cleanupSuite[tb.Name()])

	s.cleanupSuite[tb.Name()] = true

	// Note: we can't verify if the test suite had cleanup while still inside the
	// test.  The best we can do is panic in CleanupSuite to see if it ran or not.
	// panic("GOT HERE")
}

func (s *Suite) SetupTest(t *testing.T) {
	expect.Equal(t, 1, len(s.setupSuite))
	expect.Equal(t, 0, len(s.cleanupSuite))

	expect.False(t, s.setupTests[t.Name()])
	expect.False(t, s.tests[t.Name()])
	expect.False(t, s.cleanupTests[t.Name()])

	s.setupTests[t.Name()] = true

	expect.False(t, s.failTestSetup)
}

func (s *Suite) CleanupTest(t *testing.T) {
	expect.Equal(t, 1, len(s.setupSuite))
	expect.Equal(t, 0, len(s.cleanupSuite))

	expect.True(t, s.setupTests[t.Name()])
	expect.True(t, s.tests[t.Name()])
	expect.False(t, s.cleanupTests[t.Name()])

	s.cleanupTests[t.Name()] = true

	expect.False(t, s.failTestCleanup)
}

func (s *Suite) test(t *testing.T) {
	expect.Equal(t, 1, len(s.setupSuite))
	expect.Equal(t, 0, len(s.cleanupSuite))

	expect.True(t, s.setupTests[t.Name()])
	expect.False(t, s.tests[t.Name()])
	expect.False(t, s.cleanupTests[t.Name()])

	s.tests[t.Name()] = true

	expect.False(t, s.failTest)
}

func (s *Suite) SetupBenchmark(b *testing.B) {
	expect.Equal(b, 1, len(s.setupSuite))
	expect.Equal(b, 0, len(s.cleanupSuite))

	s.setupBenchmarks[b.Name()]++

	currentRun := s.setupBenchmarks[b.Name()]

	expect.Equal(b, currentRun-1, s.benchmarks[b.Name()])
	expect.Equal(b, currentRun-1, s.cleanupBenchmarks[b.Name()])

	expect.False(b, s.failBenchmarkSetup)
}

func (s *Suite) CleanupBenchmark(b *testing.B) {
	expect.Equal(b, 1, len(s.setupSuite))
	expect.Equal(b, 0, len(s.cleanupSuite))

	currentRun := s.setupBenchmarks[b.Name()]

	s.cleanupBenchmarks[b.Name()]++

	expect.Equal(b, currentRun, s.benchmarks[b.Name()])
	expect.Equal(b, currentRun, s.cleanupBenchmarks[b.Name()])

	expect.False(b, s.failBenchmarkCleanup)
}

func (s *Suite) benchmark(b *testing.B) {
	expect.Equal(b, 1, len(s.setupSuite))
	expect.Equal(b, 0, len(s.cleanupSuite))

	currentRun := s.setupBenchmarks[b.Name()]

	s.benchmarks[b.Name()]++

	expect.Equal(b, currentRun, s.benchmarks[b.Name()])
	expect.Equal(b, currentRun-1, s.cleanupBenchmarks[b.Name()])

	expect.False(b, s.failBenchmark)
}

func (s *Suite) TestUnittestCase1(t *testing.T) {
	s.test(t)
}

func (s *Suite) TestUnittestCase2(t *testing.T) {
	s.test(t)
}

func (s *Suite) TestUnittestCase3(t *testing.T) {
	s.test(t)
}

/*
// Bad test signature results in runtime testing failure.
func (s *Suite) TestBadSignature(t *testing.T, _ int) {
  s.test(t)
}
*/

func (s *Suite) BenchmarkBenchCase1(b *testing.B) {
	s.benchmark(b)
}

func (s *Suite) BenchmarkBenchCase2(b *testing.B) {
	s.benchmark(b)
}

func (s *Suite) BenchmarkBenchCase3(b *testing.B) {
	s.benchmark(b)
}

/*
// Bad benchmark signature results in runtime testing failure.
func (s *Suite) BenchmarkBadSignature(b *testing.B, _ int) {
  s.benchmark(b)
}
*/

func TestSuite(t *testing.T) {
	// In general, tests should just create the suite and call suite.RunTests().
	tests := newSuite(false, false, false, false, false, false)
	suite.RunTests(t, tests)

	// The following are additional checks to verifies suite internal
	// implementation, and should not be parts of real tests.

	expectedTestCases := map[string]bool{
		"TestSuite/UnittestCase1": true,
		"TestSuite/UnittestCase2": true,
		"TestSuite/UnittestCase3": true,
	}

	expect.Equal(t, expectedTestCases, tests.setupTests)
	expect.Equal(t, expectedTestCases, tests.tests)
	expect.Equal(t, expectedTestCases, tests.cleanupTests)

	expectedSuite := map[string]bool{"TestSuite": true}

	expect.Equal(t, expectedSuite, tests.setupSuite)

	// Note: we can't verify if the test suite had cleanup while still inside the
	// test.  The best we can do is panic in CleanupSuite to see if it ran or not.
}

func BenchmarkSuite(b *testing.B) {
	// In general, benchmarks should just create the suite and call
	// suite.RunBenchmarks().
	benchmarks := newSuite(false, false, false, false, false, false)
	suite.RunBenchmarks(b, benchmarks)

	// The following are additional checks to verifies suite internal
	// implementation, and should not be parts of real benchmarks.

	expect.Equal(b, 3, len(benchmarks.setupBenchmarks))
	expect.Equal(b, benchmarks.setupBenchmarks, benchmarks.benchmarks)
	expect.Equal(b, benchmarks.setupBenchmarks, benchmarks.cleanupBenchmarks)

	expectedSuite := map[string]bool{"BenchmarkSuite": true}

	expect.Equal(b, expectedSuite, benchmarks.setupSuite)

	// Note: we can't verify if the test suite had cleanup while still inside the
	// test.  The best we can do is panic in CleanupSuite to see if it ran or not.
}
