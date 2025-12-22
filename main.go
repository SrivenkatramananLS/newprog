package main

import "fmt"

func main() {
	var name string = "ranjini"
	var age int = 25
	height :=  5.8

	fmt.Println(name, age, height)
	state()
	stentence()
}

func state(){
	var city string = "karur"
	var temperature int = 43
	
	israining := false

fmt.Printf("City: %s, Temperature: %d°C, israining: %t\n", city, temperature, israining)

}

func stentence(){
      
	name := "ranjini"
	age := 25

	fmt.Printf("My name is %s and I am %d years old\n",name,age)

}