package pipeline

import (
	"io"
	"log"
	"sort"
)

func ArraySource(data ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range data {
			out <- n
			log.Println("write data to chan", n)
			//time.Sleep(time.Second)
		}
		log.Println("end write data to chan")
		close(out)
	}()
	return out
}

func SortInMem(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var data []int
		for n := range ch {
			data = append(data, n)
			log.Println("read data from chan", n)
		}
		log.Println("before sort data:", data)
		sort.Ints(data)
		for _, d := range data {
			out <- d
		}
		close(out)
	}()
	return out
}

func SourceFrom(reader io.Reader) <-chan int {
	return nil
}

//func RandomSink()
