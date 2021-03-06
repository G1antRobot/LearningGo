package main

import (
	"fmt"
	"github.com/G1antRobot/LearningGo/models"
)

const (
	pi   = 3.4
	name = "swapnil"
)

func testPointers(x *string) {
	fmt.Println("received:", *x)
}

func main() {
	fmt.Println("Starting GO ! Hello World !")
	u := models.User{
		ID: 2,
		FirstName: "Swapnil",
		LastName: "Karyekar",
		Address: "1900 S EADs St, Arlington VA.",

	}
	fmt.Println(u)

}
