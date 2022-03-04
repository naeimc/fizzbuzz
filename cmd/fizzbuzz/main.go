package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/naeimc/fizzbuzz"
)

var first float64
var last float64
var increment float64
var separator string
var pattern string
var help bool

const (
	USAGE             = "%s [-help] [-f first_number] [-l last_number] [-s separator] [-i increment_amount] [-p pattern]\n"
	HELP_NOTIFICATION = "Run \"fizzbuzz -help\" for more information"

	FIRST_FLAGS   = "f first"
	FIRST_DEFAULT = 1.0
	FIRST_USAGE   = "the first number <integer>"

	LAST_FLAGS   = "l last"
	LAST_DEFAULT = 100.0
	LAST_USAGE   = "the last number <integer>"

	INCREMENT_FLAGS   = "i increment"
	INCREMENT_DEFAULT = 1.0
	INCREMENT_USAGE   = "how much to increment the value by <integer>"

	SEPARATOR_FLAGS   = "s separator"
	SEPARATOR_DEFAULT = "\n"
	SEPARATOR_USAGE   = "the separator between each string <string>"

	PATTERN_FLAGS   = "p pattern"
	PATTERN_DEFAULT = "%0.f"
	PATTERN_USAGE   = "the pattern used to format the number"

	HELP_FLAGS   = "h help"
	HELP_DEFAULT = false
	HELP_USAGE   = "show a help message"
)

func init() {

	for _, s := range strings.Split(FIRST_FLAGS, " ") {
		flag.Float64Var(&first, s, FIRST_DEFAULT, FIRST_USAGE)
	}

	for _, s := range strings.Split(LAST_FLAGS, " ") {
		flag.Float64Var(&last, s, LAST_DEFAULT, LAST_USAGE)
	}

	for _, s := range strings.Split(INCREMENT_FLAGS, " ") {
		flag.Float64Var(&increment, s, INCREMENT_DEFAULT, INCREMENT_USAGE)
	}

	for _, s := range strings.Split(SEPARATOR_FLAGS, " ") {
		flag.StringVar(&separator, s, SEPARATOR_DEFAULT, SEPARATOR_USAGE)
	}

	for _, s := range strings.Split(PATTERN_FLAGS, " ") {
		flag.StringVar(&pattern, s, PATTERN_DEFAULT, PATTERN_USAGE)
	}

	for _, s := range strings.Split(HELP_FLAGS, " ") {
		flag.BoolVar(&help, s, HELP_DEFAULT, HELP_USAGE)
	}

	flag.Usage = usage
}

func main() {
	flag.Parse()

	if help {
		showHelp()
		os.Exit(0)
	}

	command := strings.ToLower(flag.Arg(0))

	switch command {
	case "":
		fb := fizzbuzz.NewFizzBuzzer(first, increment, pattern)
		fb.AddCheck(3.0, "Fizz")
		fb.AddCheck(5.0, "Buzz")
		if first > last {
			last = first
		}
		for i := first; i <= last; i += increment {
			fmt.Printf("%s%s", fb.Next(), separator)
		}
	default:
		fmt.Printf("unknown command: %s\n", flag.Arg(0))
		usage()
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "usage: "+USAGE, os.Args[0])
	fmt.Println(HELP_NOTIFICATION)
}

func showHelp() {

	args := make([][3]string, 0)
	args = append(args, [3]string{FIRST_FLAGS, FIRST_USAGE, fmt.Sprint(FIRST_DEFAULT)})
	args = append(args, [3]string{LAST_FLAGS, LAST_USAGE, fmt.Sprint(LAST_DEFAULT)})
	args = append(args, [3]string{INCREMENT_FLAGS, INCREMENT_USAGE, fmt.Sprint(INCREMENT_DEFAULT)})
	args = append(args, [3]string{SEPARATOR_FLAGS, SEPARATOR_USAGE, SEPARATOR_DEFAULT})
	args = append(args, [3]string{PATTERN_FLAGS, PATTERN_USAGE, PATTERN_DEFAULT})
	args = append(args, [3]string{HELP_FLAGS, HELP_USAGE, fmt.Sprint(HELP_DEFAULT)})

	fmt.Print("A FizzBuzz utility.\n\n")
	fmt.Print("Usage:\n\n")
	fmt.Printf("    %s\n\n", fmt.Sprintf(USAGE, os.Args[0]))
	fmt.Print("The arguments are:\n\n")

	for _, a := range args {
		fmt.Print("\t")
		for _, s := range strings.Split(a[0], " ") {
			fmt.Printf("-%s ", s)
		}
		fmt.Print("\n")

		if a[2] == "\n" {
			fmt.Printf("\t\t%s (default: \"\\n\")\n", a[1])
		} else if a[0] == "h help" {
			fmt.Printf("\t\t%s\n", a[1])
		} else {
			fmt.Printf("\t\t%s (default: %s)\n", a[1], fmt.Sprint(a[2]))
		}
	}
}
