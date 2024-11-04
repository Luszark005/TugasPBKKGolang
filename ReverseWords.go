package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kalimat (minimal 3 kata): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(input)

	if len(words) < 3 {
		fmt.Println("3 Kata minimal bro -_-")
		return
	}

	for i, word := range words {
		words[i] = reverseWord(word)
	}

	result := strings.Join(words, " ")
	fmt.Printf("Hasil Reversenya : %s\n", result)
}
