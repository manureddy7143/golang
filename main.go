package main

import (
	"fmt"
	"time"

	"github.com/golang/pets"
)

// This is a hello world prgram
func main() {
	sleeptime := time.Now()
	var animals []pets.Pet
	//sleetime := time.Now().Add(time.Duration(-5) * time.Hour)
	animals = append(animals, pets.NewDog("oreo", "black", "mixed", sleeptime))
	animals = append(animals, pets.NewCat("bomba", "white", "mixed"))

	for _, pet := range animals {
		if pet.IsHungry() {
			fmt.Println(pet.Feed("steak", "animal"))
		} else {
			fmt.Println("the animal isn't hungray,waiting")
			time.Sleep(5 * time.Second)
			fmt.Println(pet.Feed("kibele", "animal"))
		}
		//fmt.Println(pet.Feed("steak"))
		fmt.Println(pet.GiveAttention("playing fetch"))
	}
}
