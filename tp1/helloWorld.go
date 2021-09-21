package main

import (
	"fmt"
)
/*
question 1.

func main() {
	fmt.Println("Hello")
	fmt.Println("world")
}

question 2.

func main() {
	
	go func() {
		fmt.Println("Hello")
	}()

	fmt.Println("world")
}
absence de synchronisation donc arrêt du processus père sans affichage de Hello
*/

func main() {
	c:= make (chan bool)
	go func() {
		fmt.Println("Hello")
		c <- true
	}()
	<- c

	fmt.Println("world")
}