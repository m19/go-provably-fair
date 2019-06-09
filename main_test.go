package main

import (
	"fmt"
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
	expected = 1.82

	if result != expected {
		t.Errorf("Result not what was expected. Got %f. Expected %f", result, expected)
	}
}

func BenchmarkLimbo(b *testing.B) {
	bustCount := 0
	clientSeed := ""
	serverSeed := ""

	for n := 0; n < b.N; n++ {
		result := calculateLimbo(clientSeed, serverSeed, n)

		if result < 1 {
			bustCount = bustCount + 1
		}
	}

	fmt.Println("Bust count: ", bustCount)
}

func TestDice(t *testing.T) {
	clientSeed := ""
	serverSeed := ""

	result := calculateDice(clientSeed, serverSeed, 1)

	expected := 6.56
	if result != expected {
		t.Errorf("Result not what was expected. Got %f. Expected %f", result, expected)
	}
}

func TestCards(t *testing.T) {
	clientSeed := ""
	serverSeed := ""

	result := shuffleCards(clientSeed, serverSeed, 1)

	expected := []float64{3, 49, 27, 46, 10, 10, 50, 32, 31, 47, 23, 23, 17, 11, 2, 50, 49, 40, 12, 47, 29, 41, 22, 38, 1, 13, 21, 25, 36, 13, 27, 7, 3, 16, 21, 17, 34, 42, 40, 20, 14, 47, 6, 38, 0, 16, 9, 27, 33, 36, 36, 17}

	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("Result not what was expected. Got %f. Expected %f", result[i], expected[i])
		}
	}
}

func TestKeno(t *testing.T) {
	clientSeed := ""
	serverSeed := ""

	result := shuffleKeno(clientSeed, serverSeed, 1)
	expected := []int{3, 39, 21, 35, 9, 8, 38, 25, 24, 36}

	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Errorf("Result not what was expected. Got %d Expected %d", result[i], expected[i])
		}
	}
}
