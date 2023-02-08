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

func getRandomWord() string {
	rand.Seed(time.Now().Unix())
	return wordArray[rand.Intn(len(wordArray))]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	randomWord = getRandomWord()

	for {
		showDisplay()

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

func showDisplay() {
	fmt.Println(hangmanArray[len(wrongGuesses)])
	fmt.Println("Secret Word:", getCoveredWord())
	fmt.Println("Incorrect Guesses:", strings.Join(wrongGuesses, ", "))
}
