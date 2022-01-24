package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := ""
	for i := 0; i < numStarsPerLine; i++ {
		stars += "*"
	}
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func main() {
	string := "      My name is Bijon         "
	// strings.TrimSpace() -> removes the whitespaces at the beginning and the end of a sentence
	// strings.Field() -> splits a sentences into words
	// strings.Join(["...", "..."], " ") join seperated words with a single whitespace
	fmt.Println(strings.Join(strings.Fields(strings.TrimSpace(string)), " "))
	fmt.Println(strings.ReplaceAll("*****************     My name is Bijon", "*", ""))
}
