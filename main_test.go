package main

import (
	"fmt"
	"testing"
)

func TestLimbo(t *testing.T) {
	clientSeed := ""
	serverSeed := ""

	result := calculateLimbo(clientSeed, serverSeed, 1)
	resultFormatted := fmt.Sprintf("%.2f", result)
	expected := "15.07"
	if resultFormatted != "15.07" {
		t.Errorf("Result not what was expected. Got %s. Expected %s", resultFormatted, expected)
	}

	result = calculateLimbo(clientSeed, serverSeed, 2)
	resultFormatted = fmt.Sprintf("%.2f", result)
	expected = "1.82"

	if resultFormatted != "1.82" {
		t.Errorf("Result not what was expected. Got %s. Expected %s", resultFormatted, expected)
	}
}
