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

func bytesToFloat64(arr [][]byte) float64 {
	var total float64

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			value := float64(arr[i][j]) / math.Pow(256, float64(j+1))
			total = total + value
		}
	}

	return total
}

func calculateLimbo() float64 {
	clientSeed := ""
	serverSeed := ""
	nonce := 1
	round := 0

	limit := 1

	bytes := hmacSha256(clientSeed, serverSeed, nonce, round)
	bytesArray := chunkArray(bytes, 4)

	bytesArray = bytesArray[0:limit]

	total := bytesToFloat64(bytesArray)
	total = total * 100000000

	fmt.Println(bytes)
	fmt.Println(bytesArray)
	fmt.Println(total)

	result := 1000000 / (math.Floor(total) + 1) * houseEdge

	return result * 100
}

func main() {
	fmt.Println(calculateLimbo())
}
