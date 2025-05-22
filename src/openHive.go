package main

import (
	"fmt"
	"openHive/persons"
)

func main() {

	var choice, exit int
	fmt.Println("Welcome to Open Hive, an open source service desk platform")

	fmt.Println("What module do you want to access?")
	fmt.Println("1 - Persons")
	fmt.Println("2 - Tickets")
	fmt.Scanln(&choice)

	for exit != 2 {
		switch choice {
		case 1:
			persons.Persons()
		default:
			fmt.Println("Error")
		}

		fmt.Println("Wanna exit? 1 - yes | 2 - no")
		fmt.Scanln(&exit)

		fmt.Println("What module do you want to access?")
		fmt.Println("1 - Persons")
		fmt.Println("2 - Tickets")
		fmt.Scanln(&choice)
	}

}
