package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func convertToBinary(input int) string {
	quotient := input / 2
	remainder := input % 2

	if quotient == 0 {
		return strconv.Itoa(remainder)
	}

	return convertToBinary(quotient) + strconv.Itoa(remainder)
}

func formatToBinary(input int64) string {
	return strconv.FormatInt(input, 2)
}

func main() {
	fmt.Printf("Int Size = %d\n", strconv.IntSize)

	rand.Seed(time.Now().UnixNano()) // Without this, the rand will always produce the same number.
	testInput := rand.Intn(255)

	fmt.Printf("Converting %d to binary...\n", testInput)

	fmt.Println(convertToBinary(testInput))
	fmt.Println(formatToBinary(int64(testInput)))
}
