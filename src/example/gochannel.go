package example

import (
	"fmt"
)

func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
		fmt.Println("producer:", i)
	}
	close(data)
}

func consumer(data chan int, done chan bool) {
	for x := range data {
		fmt.Println("consumer:", x)
	}
	done <- true
}

func Test() {
	done := make(chan bool)
	data := make(chan int)
	go consumer(data, done)
	go producer(data)
	<-done
}

func RangeSlice() {
	data := []int{10, 20, 30}
	for i, x := range data[:] {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		fmt.Println(x, data[i])
		fmt.Println(&data)
	}
	fmt.Println(&data)
}
