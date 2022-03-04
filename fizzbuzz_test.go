package fizzbuzz

import (
	"testing"
)

func TestNewFizzBuzzer(t *testing.T) {
	start := 1.0
	step := 1.0
	pattern := "%0.f"

	fb := NewFizzBuzzer(start, step, pattern)

	if fb.Value != start {
		t.Errorf("start: expected %f, got %f", start, fb.Value)
	}

	if fb.Increment != step {
		t.Errorf("step: expected %f, got %f", step, fb.Increment)
	}
}

func TestFizzBuzzer_AddCheck(t *testing.T) {
	divisor := 3.0
	text := "Fizz"

	fb := NewFizzBuzzer(1.0, 1.0, "%0.f")
	fb.AddCheck(divisor, text)

	if result := fb.Checks[0].Divisor; result != divisor {
		t.Errorf("divisor: expected %f, got %f", divisor, result)
	}

	if result := fb.Checks[0].Text; result != text {
		t.Errorf("text: expected  %s, got %s", text, result)
	}
}

func TestFizzBuzzer_Next(t *testing.T) {
	strs := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

	fb := NewFizzBuzzer(1.0, 1.0, "%0.f")
	fb.AddCheck(3.0, "Fizz")
	fb.AddCheck(5.0, "Buzz")

	for index, text := range strs {
		next := fb.Next()
		if text != next {
			t.Errorf("expected %s for %d, got %s", text, index+1, next)
		}
	}
}
