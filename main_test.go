package main

import (
	"testing"
)

func TestLimbo(t *testing.T) {
	clientSeed := ""
	serverSeed := ""

	result := calculateLimbo(clientSeed, serverSeed, 1)
	result = float64(int(result*100)) / 100
	expected := 15.07
	if result != expected {
		t.Errorf("Result not what was expected. Got %f. Expected %f", result, expected)
	}

	result = calculateLimbo(clientSeed, serverSeed, 2)
	result = float64(int(result*100)) / 100
	expected = 1.82

	if result != expected {
		t.Errorf("Result not what was expected. Got %f. Expected %f", result, expected)
	}
}

func TestDice(t *testing.T) {
	clientSeed := ""
	serverSeed := ""

	result := calculateDice(clientSeed, serverSeed, 1)
	result = float64(int(result*100)) / 100

	expected := 6.56
	if result != expected {
		t.Errorf("Result not what was expected. Got %f. Expected %f", result, expected)
	}
}
