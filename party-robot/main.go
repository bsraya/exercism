package main

import (
	"fmt"
	"strconv"
)

func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	first := fmt.Sprintf("Welcome to my party, %s!", name)
	second := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	third := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return first + "\n" + second + "\n" + third
}

func main() {
	welcome := Welcome("Christine")
	happyBirthday := HappyBirthday("Christine", 18)
	assignTable := AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)
	fmt.Println(welcome)
	fmt.Println(happyBirthday)
	fmt.Println(assignTable)
}
