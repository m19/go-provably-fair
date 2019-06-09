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

func bytesToNumber(arr []byte) float64 {
	var number float64

	for i := 0; i < len(arr); i++ {
		value := float64(arr[i]) / math.Pow(256, float64(i+1))
		number = number + value
	}

	return number
}

func chunkArray(arr []byte, chunkSize int) [][]byte {
	var groups [][]byte

	for i := 0; i < len(arr); i = i + chunkSize {
		value := arr[i : i+chunkSize]
		groups = append(groups, value)
	}

	return groups
}

func calculateRolls(clientSeed string, serverSeed string, nonce int, round int, limit int) []float64 {
	bytes := hmacSha256(clientSeed, serverSeed, nonce, round)

	if limit > 8 {
		panic("hash is only 8*4 = 32 bytes big")
	}

	bytesArray := chunkArray(bytes, 4)

	if limit < 8 {
		bytesArray = bytesArray[0:limit]
	}

	var numbers []float64

	for i := 0; i < len(bytesArray); i++ {
		numbers = append(numbers, bytesToNumber(bytesArray[i]))
	}

	return numbers
}

func calculateLimbo(clientSeed string, serverSeed string, nonce int) float64 {
	limit := 1
	round := 0
	total := calculateRolls(clientSeed, serverSeed, nonce, round, limit)[0]
	total = total * 1000000

	result := 1000000 / (math.Floor(total) + 1) * houseEdge

	rounded := float64(int(result*100)) / 100

	return rounded
}

// this calculates a dice roll between 0 and 100.00
func calculateDice(clientSeed string, serverSeed string, nonce int) float64 {
	limit := 1
	round := 0
	total := calculateRolls(clientSeed, serverSeed, nonce, round, limit)[0]
	total = total * 10001 / 100

	rounded := float64(int(total*100)) / 100

	return rounded
}

// this will get 52 random cards from a deck of cards. A deck consists of 52 cards. 4 suits and 13 ranks per suit
// a card can appear more than once, this is to make sure the player doesn't have a huge advantage
func shuffleCards(clientSeed string, serverSeed string, nonce int) []float64 {
	limit := 8

	var cardIndexes []float64

	// this will get the first 6*8 = 48 cards
	for round := 0; round < 6; round++ {
		rolls := calculateRolls(clientSeed, serverSeed, nonce, round, limit)
		for j := 0; j < len(rolls); j++ {
			value := math.Floor(rolls[j] * 52)
			cardIndexes = append(cardIndexes, value)
		}
	}

	round := 6
	limit = 4

	// we only need 4 more bytes to get to 52 cards
	rolls := calculateRolls(clientSeed, serverSeed, nonce, round, 4)
	for i := 0; i < len(rolls); i++ {
		value := math.Floor(rolls[i] * 52)
		cardIndexes = append(cardIndexes, value)
	}

	return cardIndexes
}

func shuffleKeno(clientSeed string, serverSeed string, nonce int) []int {
	var startingNumbers []int
	var chosenNumbers []int

	for i := 0; i < 40; i++ {
		startingNumbers = append(startingNumbers, i+1)
	}

	// we want 10 numbers total
	var rolls []float64
	rolls = append(rolls, calculateRolls(clientSeed, serverSeed, nonce, 0, 8)...)
	rolls = append(rolls, calculateRolls(clientSeed, serverSeed, nonce, 1, 2)...)

	for i := 0; i < len(rolls); i++ {
		// which position?
		position := int(math.Floor(rolls[i] * (40 - float64(i))))
		// that means the chosen number is this
		chosenNumber := startingNumbers[position]
		// remove the chosenNumber from the board, this way we're sure we can't pick the same element twice
		startingNumbers = append(startingNumbers[:position], startingNumbers[position+1:]...)
		chosenNumbers = append(chosenNumbers, chosenNumber)
	}

	return chosenNumbers
}

func displayCards(cards []float64) []string {
	deck := [...]string{"♦2", "♥2", "♠2", "♣2", "♦3", "♥3", "♠3", "♣3", "♦4", "♥4", "♠4", "♣4", "♦5", "♥5", "♠5", "♣5", "♦6", "♥6", "♠6", "♣6", "♦7", "♥7", "♠7", "♣7", "♦8", "♥8", "♠8", "♣8", "♦9", "♥9", "♠9", "♣9", "♦10", "♥10", "♠10", "♣10", "♦J", "♥J", "♠J", "♣J", "♦Q", "♥Q", "♠Q", "♣Q", "♦K", "♥K", "♠K", "♣K", "♦A", "♥A", "♠A", "♣A"}

	var shuffledDeck []string

	for i := 0; i < len(cards); i++ {
		cardIndex := int(cards[i])
		shuffledDeck = append(shuffledDeck, deck[cardIndex])
	}

	return shuffledDeck
}

func main() {
	clientSeed := ""
	serverSeed := ""
	nonce := 1

	fmt.Println("limbo:", calculateLimbo(clientSeed, serverSeed, nonce))
	fmt.Println("dice:", calculateDice(clientSeed, serverSeed, nonce))
	fmt.Println("cards:", shuffleCards(clientSeed, serverSeed, nonce))
	fmt.Println("cards:", displayCards(shuffleCards(clientSeed, serverSeed, nonce)))
	fmt.Println("keno:", shuffleKeno(clientSeed, serverSeed, nonce))
}
