// struct creation
package pets

import (
	"fmt"
	"strings"
	"time"
)

type Dog struct {
	Name      string
	Color     string
	Breed     string
	lastslept time.Time
	Animal
}

func (d Dog) needsSleep() bool {
	return time.Now().Sub(d.lastslept) > 4*time.Hour
}

func (d *Dog) sleep() {
	d.lastslept = time.Now()
}

// func (d Dog) Feed(food string) string {
// 	return fmt.Sprintf("%s is eating %s", d.Name, food)
// }

func (d Dog) GiveAttention(activity string) string {
	if d.needsSleep() {
		d.sleep()
		return "your dog is tired and need sleep"
	}
	response := ""
	switch strings.ToUpper(activity) {
	case "PET":
		response = "wags tail"
	case "Playing Fetch":
		response = "return the ball and jump waiting"
	default:
		response = "barks"
	}
	return fmt.Sprintf("%s loves attention,%s will cause him to %s", d.Name, activity, response)
}

// Encapsulation
func NewDog(name, colour, breed string, lastslept time.Time) Dog {
	return Dog{
		Name:      name,
		Color:     colour,
		Breed:     breed,
		lastslept: lastslept,
		Animal:    Animal{lastAte: time.Now()},
	}
}
