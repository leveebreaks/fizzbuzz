package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	Fizz     = "Fizz"
	Buzz     = "Buzz"
	FizzBuzz = "Fizz Buzz"
)

func main() {
	length := GetLength()

	PrintSequence(length)
}

func GetLength() int {
	length := 100

	if len(os.Args) > 1 {
		arg, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Argument must be an integer. Length is set to 100 by default.")
		} else {
			length = arg
		}
	}

	return length
}

func PrintSequence(length int) {
	for i := 1; i <= length; i++ {
		switch {
		case i%15 == 0:
			fmt.Println(FizzBuzz)
		case i%5 == 0:
			fmt.Println(Buzz)
		case i%3 == 0:
			fmt.Println(Fizz)
		default:
			fmt.Println(i)
		}
	}
}
