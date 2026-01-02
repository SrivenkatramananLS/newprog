package main

import ("fmt"
		"strings"
	)


func main() {
	//loginvalidation()
	//greet()
	//mathematical()
	tmp()
}

func loginvalidation() {
	var username string
	var password string

	fmt.Print("Enter Username: ")
	fmt.Scan(&username)

	fmt.Print("Enter Password: ")
	fmt.Scan(&password)

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)


	if username == "Srivenkat" && password == "Sri@6799" {
		fmt.Println("Login success")
		runProgram()
	} else {
		fmt.Println("Enter valid input")
	}
}

func runProgram() {
	var name string = "ranjini"
	var age int = 25
	height := 5.8

	fmt.Println(name, age, height)

	state()
	stentence()
	peopleage()
	mark()
}

func state() {
	city := "karur"
	temperature := 43
	israining := false

	fmt.Printf("City: %s, Temperature: %d°C, israining: %t\n",
		city, temperature, israining)
}

func stentence() {
	name := "ranjini"
	age := 25

	fmt.Printf("My name is %s and I am %d years old\n", name, age)
}

func peopleage() {
	age := 18

	if age >= 18 {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not eligible to vote")
	}
}

func mark() {
	mark := 75

	if mark >= 90 {
		fmt.Println("Grade A")
	} else if mark >= 60 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}
}
