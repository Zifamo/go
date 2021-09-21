package main

import (
	"fmt"
)

const nbX int = 10
const nbProc int = 3

func main() {

	s := make(chan bool)
	
	var tabCan [nbProc+1] chan int
	for i:= range tabCan {
		tabCan[i] = make (chan int)
	}

	tabK := [3]int{2,3,4}

	go func(){
		source(tabCan[0])
		s <- true
	}()
	
	for i := 0; i < nbProc; i++ {
		go func(j int){
		traitement(tabCan[j], tabCan[j+1], tabK[j])
		s <- true
		}(i)
	}

	go func(){
		affichage(tabCan[nbProc])
		s <- true
	}()
		
	for i := 0; i < nbProc+2; i++{
		<- s
	}
		
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
