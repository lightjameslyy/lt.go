package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
)

var (
	globalCnt = int64(0)
)

func main()  {
	singleFlight := singleflight.Group{}

	wg := sync.WaitGroup{}

	fn := func() (interface{}, error) {
		fmt.Println("call fn")
		//time.Sleep(time.Second)
		return atomic.AddInt64(&globalCnt, 1), nil
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(rank int) {
			defer wg.Done()
			val, err, shared := singleFlight.Do("key", fn)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("rank: %v, val: %v, shared: %v\n", rank, val, shared)
		}(i)
	}

	wg.Wait()
}
