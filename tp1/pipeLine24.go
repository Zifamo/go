package main

import (
	"fmt"
)

const finLigne int = -1
const finImage int = -2

func main() {
	c1 := make (chan int)
	c2 := make (chan int)
	canSynch := make (chan bool)
	go func (){
		source(c1)
		canSynch <- true
	}()

	go func (){
		seuillage(20,c1,c2)
		canSynch <- true
	}()

	go func (){
		affichage (c2)
		canSynch <- true
	}()

	<- canSynch
	<- canSynch
	<- canSynch
}

func source(c1 chan int) {
	for {
		image := [12]int{100, 200, 150,finLigne, 32, 250, 18, finLigne, 47, 242, 99, finImage}
		for _,e := range image {
			c1 <- e
		}
	}
	 
}

func seuillage(seuil int ,c1 chan int, c2 chan int) {
	var pixel int
	for {
		pixel =<- c1
		if(pixel < seuil){
			c2 <- 0
		}
		if (pixel >= seuil){
			c2 <- 1
		}
	}
	
}

func affichage(c2 chan int) {
	var entier int
	for {
			entier =<- c2
			switch{
			case entier == -1 :
				fmt.Println()
			case entier == -2 :
				fmt.Println()
				fmt.Println()
			default :
				fmt.Println(entier)
			}
	}
}