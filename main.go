//package gogame
// Haven't built an http server yet. Need to work out basics of requests in and out

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"math/rand"
	"net/http"
)








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
