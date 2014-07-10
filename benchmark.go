package gophermark

import (
	"testing"
)

type GoconveyAssertFunc func(actual interface{}, expected ...interface{}) string

type benchScope struct {
	Setup       func()
	Run         func()
	SanityCheck func()
	VerifyList  []func() string
}

var curBench *benchScope

func Benchmark(b *testing.B, main func()) {
	// stop the timer and initialize scope
	b.StopTimer()
	if curBench != nil {
		b.Fatalf("GopherMark: calls to Benchmark() cannot be nested.")
	}
	curBench = &benchScope{}

	// run main, which should set scope vars
	main()

	// sanity checks
	if curBench.Run == nil {
		b.Fatalf("GopherMark: no benchmarks to run! Please set Run() " +
			"definitions within this benchmark.")
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

		}

		if curBench.SanityCheck != nil {
			curBench.SanityCheck()
			for _, verify := range curBench.VerifyList {
				out := verify()
				if out != "" {
					b.Fatalf("GopherMark Verify Failure: \n%s", out)
				}
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

func SanityCheck(sanityFunc func()) {
	if curBench == nil {
		panic("GopherMark: can only call SanityCheck() inside of Benchmark()")
	}
	curBench.SanityCheck = sanityFunc

}

// for more, see
//  https://github.com/smartystreets/goconvey/blob/master/convey/assertions/
func Verify(actual interface{}, assert GoconveyAssertFunc, expected ...interface{}) {
	verifyFunc := func() string {
		return assert(actual, expected...)
	}

	curBench.VerifyList = append(curBench.VerifyList, verifyFunc)
}
