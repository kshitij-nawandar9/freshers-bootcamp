package main

import (
	"fmt"
)

func main() {
	words := [5]string {"quick","brown","fox","lazy","dog"}
	var state = make(map[string]int)

	charChannel := make(chan string)

	go func() {
		for{
			char,done:=<-charChannel
			if done{
				state[char]++
			} else{
				return
			}

		}
	}()

	for _,word := range words{
		for _,char := range word{
			charChannel<-string(char)
		}
	}

	//time.Sleep(5*time.Second)
	fmt.Println(state)
}
