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
	var pixel int = 6
	c1 <- pixel 
}

func seuillage(seuil int ,c1 chan int, c2 chan int) {
	var pixel int
	pixel =<- c1
	if(pixel < seuil){
		c2 <- 0
	}
	if (pixel >= seuil){
		c2 <- 1
	}
}

func affichage(c2 chan int) {
	var entier int
	entier =<- c2
	fmt.Println(entier)
}
