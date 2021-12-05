package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// opens the required text file and checks for errors
	ascii, err := os.Open("standard.txt")
	check(err)

	// fileScanner := bufio.NewScanner(ascii)
	scanner := bufio.NewScanner(ascii)
	// scans the file and calls it "scanner"
	fileMap := make(map[rune][]string)
	// create a map with key: value to rune: []string
	firstline := true
	// checks for first line which is empty
	firstcharvalue := 32
	// the 1st character is rune 32
	inputScanner := bufio.NewScanner(os.Stdin)
	// stdin is input on command line
	fmt.Println("please input what you would like to convert:  ")

	inputScanner.Scan() // this scans user input
	// scans user input for word to convert
	userInput := inputScanner.Text() // this creates the variable = to user input
	var Toconvert []rune
	// creates a slice of rune to later store individual letters
	for scanner.Scan() {
		// scans the file named scanner
		if scanner.Text() == "" && !firstline {
			// if empty string and is not first line
			firstcharvalue++
			// increase ascii value
		} else if !firstline {
			// if not first line append map with ascii char and its values from the scanner
			fileMap[rune(firstcharvalue)] = append(fileMap[rune(firstcharvalue)], scanner.Text())
		}
		firstline = false
		// sets first line to false and is placed outside loop to stop ending it early
	}
	for _, convert := range userInput {
		// creates a for range loop based on user input
		Toconvert = append(Toconvert, convert)
		// appends individual characters to a slice of runes
	}
	for index, character := range fileMap {
		// creates a for range loop the lenght of the map
		// for _, letter:= range Toconvert{
		for i := 0; i < len(Toconvert); i++ {
			// runs a for range loop, in the lenght of how many characters to convert
			if Toconvert[i] == index {
				// if the character to convert == the individual key in the map
				for i := 0; i < 8; i++ {
					// creates a for loop to print out each line, less than 8 as there are only 8 lines per character
					fmt.Println(character[i])
				}
			}
			//	}
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
