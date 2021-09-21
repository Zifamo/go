package main

import (
	"fmt"
)

const nbX int = 10

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	s := make(chan bool)
	
	go func(){
		source(c1)
		s <- true
	}()
	
	go func(){
		traitement(c1, c2, 5)
		s <- true
	}()

	go func(){
		affichage(c2)
		s <- true
	}()
		
	<- s
	<- s
	<- s
		
}

func source(out chan int) {
	for i := 0; i < nbX; i++ {
		out <- i+1
	}
}

func traitement(in chan int, out chan int, k int) {
	var valIn, valOut int
	for i := 0; i < nbX; i++ {
		valIn = <-in
		valOut = valIn + k
		out <- valOut
	}
}

func affichage(in chan int) {
	var resIn int
	for i := 0; i < nbX; i++ {
		resIn = <-in
		fmt.Println(resIn)
	}
}
