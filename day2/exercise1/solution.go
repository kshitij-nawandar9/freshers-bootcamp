package main

import (
	"fmt"
	"sync"
)

var state = make(map[string]int)

func parseWord(wordChannel <-chan string, wg *sync.WaitGroup) {
	word:=<-wordChannel
	fmt.Println("in routine for ",word)
	var mutex = &sync.Mutex{}
	defer wg.Done()
	for _,char := range word{
		mutex.Lock()
		state[string(char)]++
		mutex.Unlock()
	}
}

func main() {
	words := [5]string {"quick","brown","fox","lazy","dog"}

	var wg sync.WaitGroup
	wordChannel := make(chan string, 1)

	for _,word := range words{
		wordChannel<-word
		wg.Add(1)
		go parseWord(wordChannel,&wg)
	}
	//time.Sleep(5*time.Second)
	wg.Wait()
	fmt.Println(state)
}
