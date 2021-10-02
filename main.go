package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const lowerLetters string = "abcdefghijklmnopqrstuvwxyz"
const upperLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZS"
const numbers string = "0123456789"
const symbols string = "`~-_=+[{]}|;:\\,<.>/?!@#$%^&*()\"'"

var useLetters bool
var useNumbers bool
var useSymbols bool
var useUpperCase bool
var useAll bool
var passLength int

func generateArray() (result string) {
	if useAll == true {
		result += lowerLetters + upperLetters + numbers + symbols
	} else {
		if useLetters == true {
			if useUpperCase == true {
				result += upperLetters + lowerLetters
			} else {
				result += lowerLetters
			}
		}
		if useNumbers == true {
			result += numbers
		}
		if useSymbols == true {
			result += symbols
		}
	}
	return result
}

func printUsage() {
	fmt.Println("SYNOPSIS")
	fmt.Print("	gopass [OPTIONS]\n\n")
	fmt.Println("DESCRIPTION")
	fmt.Print("	gopass is a small password generation utility for the commandline. It generates a password and prints to standard output\n\n")
	fmt.Println("OPIONS")
	fmt.Println("--all 					Use all character types in password")
	fmt.Println("--length				The character length of the password")
	fmt.Println("-l						Use letters in the password")
	fmt.Println("-n")
	fmt.Println("-s")
	fmt.Println("--use-upper-case")

}

func main() {

	// flag.BoolVar(&useLetters, "letters", false, "Use letters in generation")
	flag.BoolVar(&useLetters, "l", false, "Use letters in generation")
	// flag.BoolVar(&useNumbers, "numbers", false, "Use numbers in generation")
	flag.BoolVar(&useNumbers, "n", false, "Use numbers in generation")
	// flag.BoolVar(&useSymbols, "symbols", false, "Use symbols in generation")
	flag.BoolVar(&useSymbols, "s", false, "Use symbols in generation")

	flag.BoolVar(&useUpperCase, "use-upper-case", false, "Use uppercase letters in generation")
	flag.BoolVar(&useAll, "all", false, "Use all character types in a password")
	flag.IntVar(&passLength, "length", 8, "Password Length")


	flag.Parse()

	var pass string
	templateString := generateArray()
	rand.Seed(time.Now().UnixNano())

	if useAll || useLetters || useNumbers || useSymbols == true {
		for i := 1; i < passLength; i++ {
			index := rand.Intn(len(templateString))
			pick := string(templateString[index])
			pass += pick
		}
	} else {
		fmt.Println("No options specified!")
		os.Exit(1)
	}
	fmt.Println(pass)
	// printUsage()
}
