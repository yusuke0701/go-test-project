package main

import "fmt"

const (
	fizz     = "Fizz"
	buzz     = "Buzz"
	fizzBuzz = "Fizz Buzz"

	start = 1
	end   = 101
)

func main() {
	for i := start; i <= end; i++ {
		if i % 15 == 0 {
			fmt.Printf("%s", fizzBuzz)
		} else if i % 3 == 0 {
			fmt.Printf("%s", fizz)
		} else if i % 5 == 0 {
			fmt.Printf("%s", buzz)
		} else {
			fmt.Printf("%d", i)
		}
		fmt.Print(",")
	}
}
