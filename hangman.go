package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
 +---+
 0   |
/|\  |
/ \  |
    ===

Secret Word: M_N___
Incorrect Guesses: N

Sorry, you're dead. The word is MONKEY.
Correct, the secret word is MONKEY.

Please enter only one letter.
Please enter a letter.
You have already guessed that letter.
*/

var hangmanArray = [7]string{
	// 0
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===",
	// 1
	" +---+\n" +
		" 0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===",
	// 2
	" +---+\n" +
		" 0   |\n" +
		" |   |\n" +
		"     |\n" +
		"    ===",
	// 3
	" +---+\n" +
		" 0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===",
	// 4
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"     |\n" +
		"    ===",
	// 5
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/    |\n" +
		"    ===",
	// 6
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/ \\  |\n" +
		"    ===",
}
var wordArray = [7]string{
	"JAZZ", "ZIGZAG", "ZINC", "ZIPPER",
	"ZODIAC", "ZOMBIE", "FLUFF",
}

var randomWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

const letterPool string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func readLetter(reader *bufio.Reader) string {
	var letter string
	var err error
	for len(letter) != 1 {
		fmt.Print("Enter a letter: ")
		letter, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		letter = strings.Trim(letter, "\n")
		letter = strings.Trim(letter, "\r")
		if len(letter) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}
		if !strings.Contains(letterPool, strings.ToUpper(letter)) {
			fmt.Println("Please enter a letter.")
			letter = ""
			continue
		}
		if strings.Contains(guessedLetters, strings.ToUpper(letter)) {
			fmt.Println("You have already guessed that letter.")
			letter = ""
		}
	}
	return letter
}

func getCoveredWord() string {
	var hiddenWord = ""
	for _, letter := range randomWord {
		if strings.Contains(strings.Join(correctLetters, ""), string(letter)) {
			hiddenWord += string(letter)
		} else {
			hiddenWord += "_"
		}
	}
	return hiddenWord
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().Unix())
	randomWord = wordArray[rand.Intn(len(wordArray))]
	// Show game board
	// Get a letter from the user
	//
	// A. If the guess is correct
	//   1. Are there more letters to guess?
	//   2. If not, player wins
	//   3. Otherwise, add letter to guessedLetters and correctLetters
	// B. If the guess is incorrect
	// 1. Add new letter to guessedLetters and wrongGuesses
	// 2. Check if player died
	for {
		fmt.Println(hangmanArray[len(wrongGuesses)])
		fmt.Println("Secret Word:", getCoveredWord())
		fmt.Println("Incorrect Guesses:", strings.Join(wrongGuesses, ", "))
		var letter = strings.ToUpper(readLetter(reader))
		if strings.Contains(randomWord, letter) {
			correctLetters = append(correctLetters, letter)
			if randomWord == getCoveredWord() {
				break
			}
		} else {
			wrongGuesses = append(wrongGuesses, letter)
			if len(wrongGuesses) >= 6 {
				break
			}
		}
		guessedLetters += letter
	}
	if getCoveredWord() == randomWord {
		fmt.Println("A winrar is you! The word is indeed " + randomWord)
	} else {
		fmt.Println(hangmanArray[6])
		fmt.Println("You lost, better luck next time! The word was " + randomWord)
	}
}
