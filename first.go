package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is demo for Git")
	fmt.Println("When's saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
	fmt.Println(time.Now())
}
