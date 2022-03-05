# FizzBuzz
An implementation of FizzBuzz.

## Contents
 - [Getting Started](#getting-started)
    - [Command Line](#command-line)
    - [Code](#code)
 - [Installation](#installation)
 - [Usage](#usage)
    - [API](#api)
    - [Command Line Arguments](#command-line-arguments)
 - [License](#license)

## Getting Started
### Command Line
A command line utilty is provided that can be used as follows.
```bash
$ ./fizzbuzz -s " "
```
```bash
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz 31 32 Fizz 34 Buzz Fizz 37 38 Fizz Buzz 41 Fizz 43 44 FizzBuzz 46 47 Fizz 49 Buzz Fizz 52 53 Fizz Buzz 56 Fizz 58 59 FizzBuzz 61 62 Fizz 64 Buzz Fizz 67 68 Fizz Buzz 71 Fizz 73 74 FizzBuzz 76 77 Fizz 79 Buzz Fizz 82 83 Fizz Buzz 86 Fizz 88 89 FizzBuzz 91 92 Fizz 94 Buzz Fizz 97 98 Fizz Buzz
```

### Code
The FizzBuzzer can also be imported using:
```go
import "github.com/naeimc/fizzbuzz"

func main() {
    fb := NewFizzBuzzer(1.0, 1.0, "%0.f")

    fb.AddCheck(3.0, "Fizz")
    fb.AddCheck(5.0, "Buzz")

    for i := 1; i <= 100; i++ {
        fmt.Printf("%s ", fb.Next())
    }
}
```
```bash
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz 31 32 Fizz 34 Buzz Fizz 37 38 Fizz Buzz 41 Fizz 43 44 FizzBuzz 46 47 Fizz 49 Buzz Fizz 52 53 Fizz Buzz 56 Fizz 58 59 FizzBuzz 61 62 Fizz 64 Buzz Fizz 67 68 Fizz Buzz 71 Fizz 73 74 FizzBuzz 76 77 Fizz 79 Buzz Fizz 82 83 Fizz Buzz 86 Fizz 88 89 FizzBuzz 91 92 Fizz 94 Buzz Fizz 97 98 Fizz Buzz
```

## Installation
```bash
$ git clone github.com/naeimc/fizzbuzz
$ cd ./fizzbuzz
$ go build ./cmd/fizzbuzz
```

## Usage
### API
Import the package as follows:
```go
import "github.com/naeimc/fizzbuzz"
```

Then create a FizzBuzzer:
```go
func NewFizzBuzzer(start float64, increment float64, pattern string) *FizzBuzzer
```
```go
fb := fizzbuzz.NewFizzBuzzer(1.0, 1.0, "%0.f")
```
This sets up a FizzBuzzer starting at '1.0', increasing the sequence by '1.0' and formatting the output using '%0.f'.

Setting up the FizzBuzzer then requires adding 'Checks'. Checks run on each number in the sequence and returns the Text if it passes.
```go
func (fb *FizzBuzzer) AddCheck(divisor float64, text string)
```
```go
fb.AddCheck(3.0, "Fizz")
fb.AddCheck(5.0, "Buzz")
```
A 'Check' divides the number in the sequence by the Divisor. If the remainder is 0 it returns the Text. Otherwise an empty string is returned.

To run through FizzBuzzer, one can repeatedly call FizzBuzzer.Next.
```go
func (fb *FizzBuzzer) Next() string
```
```go
for i := 1; i <= 100; i++ {
	fmt.Printf("%s ", fb.Next())
}
```
Next applies each check to a number adding its text to the output. If the output is empty after all checks it is set to the number formatted by the pattern. After the output is set the value is increased by the increment.

```bash
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz 31 32 Fizz 34 Buzz Fizz 37 38 Fizz Buzz 41 Fizz 43 44 FizzBuzz 46 47 Fizz 49 Buzz Fizz 52 53 Fizz Buzz 56 Fizz 58 59 FizzBuzz 61 62 Fizz 64 Buzz Fizz 67 68 Fizz Buzz 71 Fizz 73 74 FizzBuzz 76 77 Fizz 79 Buzz Fizz 82 83 Fizz Buzz 86 Fizz 88 89 FizzBuzz 91 92 Fizz 94 Buzz Fizz 97 98 Fizz Buzz
```
Next can be called repeatedly until the sequence is completed.

### Command Line Arguments
```bash
$ # sets the separator between output, defaults to "/n"
$ ./fizzbuzz -s " "
$ ./fizzbuzz -separator " "
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz ...

$ # sets the first value to start the FizzBuzz sequence
$ ./fizzbuzz -s " " -f 95
$ ./fizzbuzz -s " " -first 95
Buzz Fizz 97 98 Fizz Buzz

$ # sets the value to end the FizzBuzz sequence, if -first > -last the last value will be set to first
$ ./fizzbuzz -s " " -l 5
$ ./fizzbuzz -s " " -last 5
1 2 Fizz 4 Buzz

$ # sets the amount to increment each value by
$ ./fizzbuzz -s " " -i 14
$ ./fizzbuzz -s " " -increment 14
1 FizzBuzz 29 43 Fizz 71 Buzz Fizz

$ # the pattern used to format the number output
$ ./fizzbuzz -s " " -p "%f"
$ ./fizzbuzz -s " " -pattern "%f"
1.000000 2.000000 Fizz 4.000000 Buzz Fizz 7.000000 8.000000 Fizz Buzz 11.000000 Fizz 13.000000 14.000000 FizzBuzz ...
```

A help message is also available:
```bash
$ ./fizzbuzz -h
$ ./fizzbuzz -help
```

## License
[MIT](LICENSE)