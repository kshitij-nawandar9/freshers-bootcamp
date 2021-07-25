package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var totalRating uint64
	for i:=0;i<200;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			waitingTime := rand.Intn(10)
			time.Sleep(time.Duration(waitingTime) * time.Millisecond)

			rating  := uint64(rand.Intn(10))
			atomic.AddUint64(&totalRating, rating)
		}()
	}
	//time.Sleep(5*time.Second)
	wg.Wait()
	fmt.Println("Total student rating",totalRating)
	average :=float64(totalRating)/200
	fmt.Println("Average student rating : ",average)

}
