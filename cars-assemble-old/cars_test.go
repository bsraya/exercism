package main

import (
	"testing"
)

func TestSuccessRate(t *testing.T) {
	if got, want := SuccessRate(0), 0.0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(1), 1.0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(2), 1.0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(3), 1.0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(4), 1.0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(5), 0.9; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(6), 0.9; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(7), 0.9; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(8), 0.9; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(9), 0.77; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := SuccessRate(10), 0.77; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalculateProductionRatePerHour(t *testing.T) {
	if got, want := CalculateProductionRatePerHour(7), 1392.3; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalculateProductionRatePerMinute(t *testing.T) {
	if got, want := CalculateProductionRatePerMinute(5), 16; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalculateLimitedProductionRatePerHour(t *testing.T) {
	if got, want := CalculateLimitedProductionRatePerHour(2, 1000.0), 442.0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	if got, want := CalculateLimitedProductionRatePerHour(7, 1000.0), 1000.0; got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
