package gophermark

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func BenchmarkWithoutGopherMark(b *testing.B) {
	str := strings.Repeat("I am a cat I have a hat I have a house I have a mouse!", 100)

	var count int
	for i := 0; i < b.N; i++ {
		count = 1
		b.StartTimer()
		count += strings.Count(str, "mouse!")
		b.StopTimer()
		if count != 101 {
			panic("FAIL")
		}

	}
}

func BenchmarkAThing(b *testing.B) {
	Benchmark(b, func() {
		str := strings.Repeat("I am a cat I have a hat I have a house I have a mouse!", 100)
		var count int

		Setup(func() {
			count = 1
		})

		Run(func() {
			count += strings.Count(str, "mouse!")
		})

		SanityCheck(func() {
			ct := 100 + 1
			Verify(str, ShouldContainSubstring, "hat")
			Verify(count, ShouldEqual, ct)
		})
	})

}
