package sort

import (
	"sync"
	"time"
)

func Sleepsort(n int,a *[]int ,wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(n))
	*a = append(*a,n)
}

func Checkerr(err error) {
	if (err != nil) {
		panic(err)
	}
}