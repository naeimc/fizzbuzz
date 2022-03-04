package fizzbuzz

import "fmt"

func Example() {
	start := 1.0
	increment := 1.0
	pattern := "%0.f"

	fb := NewFizzBuzzer(start, increment, pattern)

	fb.AddCheck(3.0, "Fizz")
	fb.AddCheck(5.0, "Buzz")

	for i := 1; i <= 15; i++ {
		fmt.Printf("%s ", fb.Next())
	}
	// Output:
	// 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz
}
