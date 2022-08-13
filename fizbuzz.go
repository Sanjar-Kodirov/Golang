package main

import "fmt"

func FizzBuzz() string{
	fmt.Println("Please enter number...")
	var myNumber int64
	fmt.Scan(&myNumber)
	if myNumber % 15 == 0 {
		return "Fizz buzzz"
	}else if myNumber % 5 == 0 {
		return "Fizz"
	}else if myNumber % 3 == 0 {
		return "Buzz"
	}
	return "You entered wrong number"
}

