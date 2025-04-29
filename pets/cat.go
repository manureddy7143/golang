package pets

import (
	"fmt"
	"time"
)

type Cat struct {
	Name  string
	Color string
	Breed string
	Animal
}

func (c Cat) GiveAttention(activity string) string {

	return fmt.Sprintf("the cat is ignoring attempts to %s", activity)
}

func NewCat(name, colour, breed string) Pet {

	return Cat{
		Name:   name,
		Color:  colour,
		Breed:  breed,
		Animal: Animal{lastAte: time.Now()},
	}
}
