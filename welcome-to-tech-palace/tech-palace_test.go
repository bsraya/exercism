package main

import "testing"

func TestWelcomeMessage(t *testing.T) {
	if got, want := WelcomeMessage("Bijon"), "Welcome to the Tech Palace, BIJON"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestAddBorders(t *testing.T) {
	if got, want := AddBorder("Welcome!", 10), "**********\nWelcome!\n**********"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCleanupMessage(t *testing.T) {
	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`
	if got, want := CleanupMessage(message), "BUY NOW, SAVE 10%"; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
