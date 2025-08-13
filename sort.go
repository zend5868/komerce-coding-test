package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input one line of words (S) : ")
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	vowels := "aeiou"
	var vowelChars, consonantChars strings.Builder

	for _, ch := range input {
		if ch == ' ' {
			continue // abaikan spasi
		}
		if strings.ContainsRune(vowels, ch) {
			vowelChars.WriteRune(ch)
		} else if ch >= 'a' && ch <= 'z' {
			consonantChars.WriteRune(ch)
		}
	}

	fmt.Println("Vowel Characters     :", vowelChars.String())
	fmt.Println("Consonant Characters :", consonantChars.String())
}
