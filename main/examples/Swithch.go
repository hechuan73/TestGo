package main

import "fmt"
import "time"

func main()  {

	for i := 1; i <= 5; i++ {
		switch i {
		case 1: fmt.Println(i)
		case 2: fmt.Println(i)
		case 3: fmt.Println(i)
		default:
			fmt.Println("default")
		}
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

}
