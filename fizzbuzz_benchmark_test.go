package fizzbuzz

import "testing"

func BenchmarkFizzBuzzer_Next(b *testing.B) {
	fb := NewFizzBuzzer(1.0, 1.0)
	fb.AddCheck(3.0, "Fizz")
	fb.AddCheck(5.0, "Buzz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 1.0; i <= 100; i++ {
			fb.Next()
		}
	}
}
