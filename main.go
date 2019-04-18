//package gogame

package main

import (
	"fmt"
	"time"
	"math/rand"
)


var r = rand.New(rand.NewSource(time.Now().Unix())

type person struct {
	first string
	last  string
}

type hero struct {
	person
	class string
	strength   int
	resilience int
	agility    int
}

func (s hero) speak() {
	fmt.Println("I am", s.first, s.last, " a strong ", s.class)
}

func (i hero) attack() {
	luck := roll()
	//fmt.Println(luck)
	fmt.Println(i.first, i.strength)
	fmt.Println(i.strength, luck, luck + i.strength)
	x := luck + i.strength
	fmt.Println(x)
	//return "cool"
}


func roll() int {
	//x := r.Intn(10)
	//fmt.Println("you rolled a ", x)
	//return x
	fmt.Println("roll called")
	return 3
}

func main() {
	fmt.Println(currenttime)
	warrior := hero{
		person: person{
			first: "James",
			last:  "Bond",
		},
		class: "Warrior"
		strength:   3,
		resilience: 2,
		agility:    1,
	}

	//warrior.speak()
	warrior.attack()
	fmt.Println(warrior.strength)

}
