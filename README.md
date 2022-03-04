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
./fizzbuzz -s " "
```
Output:
```bash
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz 31 32 Fizz 34 Buzz Fizz 37 38 Fizz Buzz 41 Fizz 43 44 FizzBuzz 46 47 Fizz 49 Buzz Fizz 52 53 Fizz Buzz 56 Fizz 58 59 FizzBuzz 61 62 Fizz 64 Buzz Fizz 67 68 Fizz Buzz 71 Fizz 73 74 FizzBuzz 76 77 Fizz 79 Buzz Fizz 82 83 Fizz Buzz 86 Fizz 88 89 FizzBuzz 91 92 Fizz 94 Buzz Fizz 97 98 Fizz Buzz
```

### Code
The FizzBuzzer can also be imported using:
```go
import "github.com/naeimc/fizzbuzz"
```
And implemented with:
```go
fb := NewFizzBuzzer(1.0, 1.0, "%0.f")

fb.AddCheck(3.0, "Fizz")
fb.AddCheck(5.0, "Buzz")

for i := 1; i <= 100; i++ {
	fmt.Printf("%s ", fb.Next())
}
```
Output:
```bash
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz 31 32 Fizz 34 Buzz Fizz 37 38 Fizz Buzz 41 Fizz 43 44 FizzBuzz 46 47 Fizz 49 Buzz Fizz 52 53 Fizz Buzz 56 Fizz 58 59 FizzBuzz 61 62 Fizz 64 Buzz Fizz 67 68 Fizz Buzz 71 Fizz 73 74 FizzBuzz 76 77 Fizz 79 Buzz Fizz 82 83 Fizz Buzz 86 Fizz 88 89 FizzBuzz 91 92 Fizz 94 Buzz Fizz 97 98 Fizz Buzz
```

## Installation
```bash
git clone github.com/naeimc/fizzbuzz
cd ./fizzbuzz
go build ./cmd/fizzbuzz
```
## Usage
### API
Import the package as follows:
```go
import "github.com/naeimc/fizzbuzz"
```
The FizzBuzz utility operates through the FizzBuzzer struct:
```go
func NewFizzBuzzer(start float64, increment float64) *FizzBuzzer
```
The following sets a FizzBuzzer starting at "1", increasing the sequence by "1" and formatting the output using "%0.f"
```go
start := 1.0
increment := 1.0
pattern := "%0.f"
fb := fizzbuzz.NewFizzBuzzer(start, increment, pattern)
```
Setting up the FizzBuzzer then requires adding 'Checks'. Checks run on each number in the sequence and returns the text if it passes.
```go
fizzDivisor := 3.0
fizzText := "Fizz"
fb.AddCheck(fizzDivisor, fizzText)
buzzDivisor := 5.0
buzzText := "Buzz"
fb.AddCheck(buzzDivisor, buzzText)
```
A Check divides the number in the sequence by the divisor. If the remainder is 0 it returns the text. Otherwise nothing is returned.

To run through FizzBuzzer, one can repeatedly call FizzBuzzer.Next().
```go
func (fb *FizzBuzzer) Next() string
```
```go
for i := 1; i <= 100; i++ {
	fmt.Printf("%s ", fb.Next())
}
```
Next() applies each check to a number adding its text to the output. If the output is empty it is set to the number formatted by the pattern. After the output is set the value is increased by the increment.

Output:
```bash
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz 31 32 Fizz 34 Buzz Fizz 37 38 Fizz Buzz 41 Fizz 43 44 FizzBuzz 46 47 Fizz 49 Buzz Fizz 52 53 Fizz Buzz 56 Fizz 58 59 FizzBuzz 61 62 Fizz 64 Buzz Fizz 67 68 Fizz Buzz 71 Fizz 73 74 FizzBuzz 76 77 Fizz 79 Buzz Fizz 82 83 Fizz Buzz 86 Fizz 88 89 FizzBuzz 91 92 Fizz 94 Buzz Fizz 97 98 Fizz Buzz
```
Next can be called repeatedly until the sequence is completedly.
### Command Line Arguments
The command line utility includes a number of arguments:
#### s, separator
Sets the separator between output. Defaults to "/n".
```bash
./fizzbuzz -s " "
./fizzbuzz -separator " "
```
Output:
```bash
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz 31 32 Fizz 34 Buzz Fizz 37 38 Fizz Buzz 41 Fizz 43 44 FizzBuzz 46 47 Fizz 49 Buzz Fizz 52 53 Fizz Buzz 56 Fizz 58 59 FizzBuzz 61 62 Fizz 64 Buzz Fizz 67 68 Fizz Buzz 71 Fizz 73 74 FizzBuzz 76 77 Fizz 79 Buzz Fizz 82 83 Fizz Buzz 86 Fizz 88 89 FizzBuzz 91 92 Fizz 94 Buzz Fizz 97 98 Fizz Buzz
```
#### f, first
Sets the first value to start the FizzBuzz sequence.
```bash
./fizzbuzz -f 95
./fizzbuzz -first 95
```
Output:
```bash
Buzz
Fizz
97
98
Fizz
Buzz
```
#### l, last
Sets the value to end the FizzBuzz sequence.
If -first > -last the last value will be set to first.
```bash
./fizzbuzz -l 5
./fizzbuzz -last 5
```
Output:
```bash
1
2
Fizz
4
Buzz
```
#### i, increment
Sets the amount to increment each value by.
```bash
./fizzbuzz -i 14
./fizzbuzz -increment 14
```
Output:
```bash
1
FizzBuzz
29
43
Fizz
71
Buzz
Fizz
```
#### p, pattern
The pattern used to format the number output.
```bash
./fizzbuzz -p "%f"
./fizzbuzz -pattern "%f"
```
Output:
```bash
1.000000
2.000000
Fizz
4.000000
Buzz
Fizz
7.000000
8.000000
Fizz
Buzz
11.000000
...
```
#### h, help
Show a help message.
```bash
./fizzbuzz -h
./fizzbuzz -help
```
Output:
```bash
A FizzBuzz utility.

Usage:

    C:\Users\Naeim\Projects\fizzbuzz\fizzbuzz.exe [-help] [-f first_number] [-l last_number] [-s separator] [-i increment_amount] [-p pattern]


The arguments are:

        -f -first
                the first number <integer> (default: 1)
        -l -last
                the last number <integer> (default: 100)
        -i -increment
                how much to increment the value by <integer> (default: 1)
        -s -separator
                the separator between each string <string> (default: "\n")
        -p -pattern
                the pattern used to format the number (default: %0.f)
        -h -help
                show a help message

```

## License

[MIT](LICENSE)