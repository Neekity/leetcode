package main

import (
	"fmt"
	"math/rand"
	"neekity.com/leetcode/src/middle"
	"sync"
	"time"
)

type X int

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

var wg sync.WaitGroup
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"json",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	fmt.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	//lists1 := common.TransferNodes([]int{1, 2, 3, 4, 5})
	fmt.Println(middle.PermuteUnique([]int{1, 1, 3}))

}

func counter(origin chan int) {
	for num := 2; num < 100; num++ {
		origin <- num
	}
	close(origin)
}

func filter(prime int, recv chan int, send chan int) {
	for i := range recv {
		if i%prime != 0 {
			send <- i
		}
	}
	close(send)
}

func processor(seq chan int, wait chan struct{}) {

	prime, ok := <-seq
	if !ok {
		close(wait)
		return
	}
	fmt.Println(prime, "是素数")
	out := make(chan int)
	go processor(out, wait)
	for num := range seq {
		if num%prime != 0 {
			out <- num
		}
	}
	close(out)

}

func player(name string, cout chan int) {
	defer wg.Done()
	for {
		fmt.Printf("选手 %s 等待对方的来球！\n", name)
		ball, ok := <-cout
		if !ok {
			fmt.Printf("选手 %s 赢得了比赛！\n", name)
			return
		}
		n := rand.Intn(100)
		if n%57 == 0 {
			fmt.Printf("选手 %s 击球失败！\n", name)
			close(cout)
			return
		}
		fmt.Printf("选手 %s 成功击球，第%d次！\n", name, ball)
		ball++
		cout <- ball
	}
}

func testSlice(slice []int) []int {
	return append(slice, 1, 2, 3)
}
