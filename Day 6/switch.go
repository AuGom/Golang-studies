package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	switch i { //ordinary switch
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("none")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //its possible to use more than one statement
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

}