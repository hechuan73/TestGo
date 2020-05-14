package main

import "fmt"
import "errors"

func main() {
	for _, num := range []int {7, 42} {
		if r, e := f1(num); e != nil{
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, num := range []int {7, 42} {
		if r, e := f2(num); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	// ok 用于判断e是否是argError这个结构体的一个实例
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	
}

type argError struct {
	arg int
	prob string
}

func (err *argError) Error() string {
	return fmt.Sprintf("%d -> %s", err.arg, err.prob)
}

func f2(arg int) (int, error)  {
	if arg == 42 {
		return -1, &argError{arg, "arg is invalid"}
	}
	return arg+3, nil
}

func f1(arg int) (int, error)  {
	if arg == 42 {
		return -1, errors.New("arg is invalid")
	}

	return arg+3, nil
}