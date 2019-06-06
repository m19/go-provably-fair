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
	if resultFormatted != "15.07" {
		t.Error("Result not what was expected")
	}

	result = calculateLimbo(clientSeed, serverSeed, 2)
	resultFormatted = fmt.Sprintf("%.2f", result)
	if resultFormatted != "1.83" {
		t.Error("Result not what was expected")
	}
}
