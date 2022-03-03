/*
A FizzBuzz utility.
*/
package fizzbuzz

import (
	"fmt"
	"math"
)

// A Checker is used to perform the division in a FizzBuzzer.
type Checker struct {
	Divisor float64 // What to divide by.
	Text    string  // The test to return in modulus == 0.
}

// Create a new Checker.
func NewCheck(divisor float64, text string) *Checker {
	return &Checker{float64(divisor), text}
}

// Run the check on a number.
//
// Checks the remaineder of dividend / *Checker.Divisor.
// If remainder == 0 return *Checker.Text,
// otherwise return an empty string.
func (c *Checker) Check(dividend float64) string {
	if math.Mod(dividend, c.Divisor) == 0.0 {
		return c.Text
	}
	return ""
}

// A FizzBuzser
type FizzBuzzer struct {
	Value     float64   // The current value to fizzbuzz.
	Increment float64   // What to increment the current value by after a successful FizzBuzz.
	Checks    []Checker // The checks to run while FizzBuzzing.
}

// Create a new FizzBuzz utility.
func NewFizzBuzzer(start float64, increment float64) *FizzBuzzer {
	return &FizzBuzzer{Value: start, Increment: increment}
}

// Add a check to the fizzbuzz utility.
func (fb *FizzBuzzer) AddCheck(divisor float64, text string) {
	fb.Checks = append(fb.Checks, *NewCheck(divisor, text))
}

// FizzBuzz the next value.
//
// This will run each check on the current value, adding its result to the output string.
// If the output string is empty after all strings it is set to the current value.
// The current value is then incremented before the output is returned.
func (fb *FizzBuzzer) Next() string {
	var out string
	for _, c := range fb.Checks {
		out += c.Check(fb.Value)
	}
	if out == "" {
		out = fmt.Sprintf("%0.f", fb.Value)
	}
	fb.Value += fb.Increment
	return out
}
