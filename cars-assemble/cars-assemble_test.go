package main

import (
	"testing"
)

func TestCalculateWorkingCarsPerHour(t *testing.T) {
	got := CalculateWorkingCarsPerHour(1547, 90)
	want := 1392.3
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalculateWorkingCarsPerMinute(t *testing.T) {
	got := CalculateWorkingCarsPerMinute(1105, 90)
	want := 16
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCalculateCost(t *testing.T) {
	if got, want := CalculateCost(37), uint(355000); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	if got, want := CalculateCost(21), uint(200000); got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
