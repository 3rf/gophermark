package gophermark

import (
	"strings"
	"testing"
)

func BenchmarkWithoutGopherMark(b *testing.B) {
	str := strings.Repeat("I am a cat I have a hat I have a house I have a mouse!", 100)
	for i := 0; i < b.N; i++ {
		count := strings.Count(str, "mouse!")

		b.StopTimer()
		if count != 100 {
			panic("FAIL")
		}
		b.StartTimer()

	}
}

func BenchmarkAThing(b *testing.B) {
	str := strings.Repeat("I am a cat I have a hat I have a house I have a mouse!", 100)

	Benchmark(b, func() {

		Run(func() {
			count := strings.Count(str, "mouse!")

			Verify(func() {
				if count != 100 {
					panic("FAIL")
				}
			})
		})
	})

}
