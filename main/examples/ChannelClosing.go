package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// 二值形式，前者为值，后者表示通道是否关闭
			j, hasNext := <- jobs
			if hasNext {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 通道同步等待任务结束
	<- done
}
