package concurrency

import (
	"reflect"
	"strings"
	"sync"
	"time"
)

type Data struct {
	sync.Mutex
}

func (d *Data) Test(s string) {
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}

func PrintPrimeB() {
	origin := make(chan int)
	wait := make(chan struct{})
	go func(origin chan int) {
		for num := 2; num < 100; num++ {
			origin <- num
		}
		close(origin)
	}(origin)

	go processor(origin, wait)
	<-wait
}

func processor(seq chan int, wait chan struct{}) {
	prime, ok := <-seq
	if !ok {
		close(wait)
		return
	}
	out := make(chan int)
	go processor(out, wait)
	for num := range seq {
		if num%prime != 0 {
			out <- num
		}
	}
	close(out)

}

func PrintPrimeA() {
	origin := make(chan int)

	go func(origin chan int) {
		for num := 2; num < 100; num++ {
			origin <- num
		}
		close(origin)
	}(origin)

	for {
		primes := make(chan int)
		prime, ok := <-origin
		if !ok {
			return
		}
		//fmt.Println(prime)
		go func(prime int, recv chan int, send chan int) {
			for i := range recv {
				if i%prime != 0 {
					send <- i
				}
			}
			close(send)
		}(prime, origin, primes)
		origin = primes
	}
}

func add(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}

	var ret reflect.Value
	switch args[0].Kind() {
	case reflect.Int:
		n := 0
		for _, a := range args {
			n += int(a.Int())
		}
		ret = reflect.ValueOf(n)
	case reflect.String:
		ss := make([]string, 0, len(args))
		for _, s := range args {
			ss = append(ss, s.String())
		}
		ret = reflect.ValueOf(strings.Join(ss, ""))
	}
	results = append(results, ret)

	return
}

func MakeAdd(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), add)
	fn.Set(v)
}
