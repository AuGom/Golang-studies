package main

import (
	"errors"
	"fmt"
)

// errors are the last return value and have type error
func f1(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42") // errors.New constructs a basic error value with the given message

	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// possible to use custom types as errors by implementing the Error() method
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
