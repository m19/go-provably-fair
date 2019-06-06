package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"math"
)

const (
	houseEdge = .99
)

// create a hmacSha256 with the serverSeed as the secret and clientSeed:nonce:round as the data
// returns the hmacSha256 as an array of bytes
func hmacSha256(clientSeed string, serverSeed string, nonce int, round int) []byte {
	h := hmac.New(sha256.New, []byte(serverSeed))
	seed := fmt.Sprintf("%s:%d:%d", clientSeed, nonce, round)
	h.Write([]byte(seed))

	return h.Sum(nil)
}

func chunkArray(arr []byte, chunkSize int) [][]byte {
	var groups [][]byte

	for i := 0; i < len(arr); i = i + chunkSize {
		value := arr[i : i+chunkSize]
		groups = append(groups, value)
	}

	return groups
}

func bytesToFloat64(arr []byte) float64 {
	var total float64

	for i := 0; i < len(arr); i++ {
		value := float64(arr[i]) / math.Pow(256, float64(i+1))
		total = total + value
	}

	return total
}

func calculateRolls(clientSeed string, serverSeed string, nonce int, round int, chunk int, limit int) []float64 {
	bytes := hmacSha256(clientSeed, serverSeed, nonce, round)
	bytesArray := chunkArray(bytes, chunk)

	if chunk*limit > 32 {
		panic("hash is only 32 bytes big")
	}

	if limit < 8 {
		bytesArray = bytesArray[0:limit]
	}

	var totals []float64

	for i := 0; i < len(bytesArray); i++ {
		totals = append(totals, bytesToFloat64(bytesArray[i]))
	}

	return totals
}

func calculateLimbo(clientSeed string, serverSeed string, nonce int) float64 {
	chunk := 4
	limit := 1
	round := 0
	total := calculateRolls(clientSeed, serverSeed, nonce, round, chunk, limit)[0]
	total = total * 100000000

	result := 1000000 / (math.Floor(total) + 1) * houseEdge

	return result * 100
}

func calculateDice(clientSeed string, serverSeed string, nonce int) float64 {
	chunk := 4
	limit := 1
	round := 0
	total := calculateRolls(clientSeed, serverSeed, nonce, round, chunk, limit)[0]
	total = total * 10001

	return total / 100
}

func calculateCards(clientSeed string, serverSeed string, nonce int) []float64 {
	chunk := 4
	limit := 8

	var cardIndexes []float64

	// this will get the first 6*8 = 48 cards
	for i := 0; i < 6; i++ {
		rolls := calculateRolls(clientSeed, serverSeed, nonce, i, chunk, limit)
		for j := 0; j < len(rolls); j++ {
			value := math.Floor(rolls[j] * 52)
			cardIndexes = append(cardIndexes, value)
		}
	}

	// we only need 4 more bytes to get to 52 cards
	rolls := calculateRolls(clientSeed, serverSeed, nonce, 6, chunk, 4)
	for j := 0; j < len(rolls); j++ {
		value := math.Floor(rolls[j] * 52)
		cardIndexes = append(cardIndexes, value)
	}

	return cardIndexes
}

func main() {
	clientSeed := ""
	serverSeed := ""
	nonce := 1

	fmt.Println("limbo:", calculateLimbo(clientSeed, serverSeed, nonce))
	fmt.Println("dice:", calculateDice(clientSeed, serverSeed, nonce))
	fmt.Println("cards:", calculateCards(clientSeed, serverSeed, nonce))
}
