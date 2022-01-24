package main

import (
	"fmt"
	"strconv"
	"testing"
)

type person struct {
	name string
	age  int
}

type description struct {
	name      string
	table     int
	neighbor  string
	direction string
	distance  float64
}

func TestWelcome(t *testing.T) {
	names := []string{
		"John",
		"Christiane",
		"Frank",
	}

	for _, name := range names {
		if got, want := Welcome(name), "Welcome to my party, "+name+"!"; got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestHappyBirthday(t *testing.T) {
	people := []person{
		{
			name: "John",
			age:  18,
		},
		{
			name: "Christiane",
			age:  20,
		},
		{
			name: "Frank",
			age:  25,
		},
	}

	for _, person := range people {
		want := "Happy birthday " + person.name + "! You are now " + strconv.Itoa(person.age) + " years old!"
		if got := HappyBirthday(person.name, person.age); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestAssignTable(t *testing.T) {
	descriptions := []description{
		{"Christiane", 27, "Frank", "on the left", 23.7834298},
		{"Xuân Jing", 4, "Renée", "by the façade", 23.470103},
		{"Chihiro", 22, "Akachi Chikondi", "straight ahead", 9.2394381},
	}

	for _, i := range descriptions {
		first := fmt.Sprintf("Welcome to my party, %s!", i.name)
		second := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", i.table, i.direction, i.distance)
		third := fmt.Sprintf("You will be sitting next to %s.", i.neighbor)
		want := first + "\n" + second + "\n" + third

		if got := AssignTable(i.name, i.table, i.neighbor, i.direction, i.distance); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}
