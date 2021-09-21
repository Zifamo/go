package main

import (
	"fmt"
)

const nbX int = 10
const ordre int = 3

func main() {

	s := make(chan bool)
	
	var tabCan [ordre+1] chan float64
	for i:= range tabCan {
		tabCan[i] = make (chan float64)
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

func source(out chan float64) {
	var x int = 0
	var P int = 1
	var S int = 0
	out <- float64(x)
	out <- float64(P)
	out <- float64(U)

}

func traitement(in chan float64, out chan float64, k int) {
	var valIn, valOut int
	for i := 0; i < nbX; i++ {
		valIn = <-in
		valOut = valIn + k
		out <- valOut
	}
}

func affichage(in chan float64) {
	var resIn int
	for i := 0; i < nbX; i++ {
		resIn = <-in
		fmt.Println(resIn)
	}
}
