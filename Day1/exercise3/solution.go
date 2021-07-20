//package exercise3

package main

import (
"fmt"
)

type Salary interface {
	compute() int
}

type fullTime struct {
	duration int
}

type contractor struct {
	duration int
}

type freelancer struct {
	duration int
}


func (fte fullTime) compute() int {
	return fte.duration*500
}

func (cont contractor) compute() int {
	return cont.duration*100
}

func (fl freelancer) compute() int {
	return fl.duration*10
}


func main() {
	fte:=fullTime{30}
	cont:=contractor{20}
	fl:=freelancer{25}
	fmt.Println(fte.compute())
	fmt.Println(cont.compute())
	fmt.Println(fl.compute())
}


