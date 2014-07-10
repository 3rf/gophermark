package gophermark

import (
	"testing"
)

type benchScope struct {
	B           *testing.B
	Setup       func()
	Run         func()
	RunParallel func()
	Verify      func()
}

var curBench *benchScope

func Benchmark(b *testing.B, main func()) {
	// stop the timer and initialize scope
	b.StopTimer()
	if curBench != nil {
		b.Fatalf("GopherMark: calls to Benchmark() cannot be nested.")
	}
	curBench = &benchScope{B: b}

	// run main, which should set scope vars
	main()

	// sanity checks
	if curBench.Run == nil && curBench.RunParallel == nil {
		b.Fatalf("GopherMark: no benchmarks to run! Please set Run() or" +
			"RunParallel() definitions within this benchmark.")
	}

	// TODO both set

	// regular Run
	if curBench.Run != nil {
		for i := 0; i < b.N; i++ {

			if curBench.Setup != nil {
				curBench.Setup()
			}

			b.StartTimer()
			curBench.Run()
			b.StopTimer()

			if curBench.Verify != nil {
				curBench.Verify()
			}
		}
	}

	curBench = nil
}

func Setup(setupFunc func()) {
	if curBench == nil {
		panic("GopherMark: can only call Setup() inside of Benchmark()")
	}
	curBench.Setup = setupFunc
}

func Run(benchFunc func()) {
	if curBench == nil {
		panic("GopherMark: can only call Run() inside of Benchmark()")
	}
	curBench.Run = benchFunc
}

func Verify(verifyFunc func()) {
	if curBench == nil {
		panic("GopherMark: can only call Verify() inside of Benchmark()")
	}
	curBench.Verify = verifyFunc
}
