package main

import (
	"fmt"
)

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
	image := [9]int{100, 200, 150, 32, 250, 18, 47, 242, 99}
	for i := 0; i < 8; i++ {
		c1 <- image[i]
	}
	 
}

func seuillage(seuil int ,c1 chan int, c2 chan int) {
	var pixel int
	for i := 0; i < 8; i++ {
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
	for i := 0; i < 8; i++ {
		entier =<- c2
		fmt.Println(entier)
	}
}