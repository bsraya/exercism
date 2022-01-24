package main

import (
	"strconv"
	"testing"
)

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

type person struct {
	name string
	age  int
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

}
