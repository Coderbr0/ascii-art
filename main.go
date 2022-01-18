package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		outputAscii(ReadFile(), os.Args[1]) // ReadFile() is passed into asciiMap map[int][]string and os.Args[1] is passed into input string in parameter of outputAscii function
	} else {
		fmt.Println("Error: Wrong amount of arguments.")
	}
}

func ReadFile() map[int][]string { // We want to return a map for simplicity
	fileInput := make(map[int][]string)  // Alternative syntax: var fileInput map[int][]string
	file, err := os.Open("standard.txt") // file, err := os.Open("os.Args[1]") => e.g. go run main.go standard.txt; standard.txt would be 1st argument in this case
	if err != nil {
		fmt.Println("Invalid input. The named file does not exist.")
		os.Exit(1) // Exits program with "exit status 1" if file is not found; alternative: fmt.Println(err.Error()); "open standard.txt: no such file or directory"
	}
	defer file.Close()                // To close file is good practice; defer allows for all operations to be carried out before closing file
	scanner := bufio.NewScanner(file) // Scans the file and calls it "scanner"
	count := 31                       // Ignoring first line which is blank; in ascii manual (man ascii in terminal) integer 32 (decimal value) represents space (line 2 to 9 in standard.txt file); the first character is rune 32 (decimal value 32 - space); we could use fileInput := make(map[rune][]string) with slight changes to code elsewhere; there are 95 characters to print (dec value 32 - 126 in ascii manual)
	for scanner.Scan() {              // Scans the variable named scanner
		if scanner.Text() == "" {
			count++
		} else {
			fileInput[count] = append(fileInput[count], scanner.Text())
		}
	}
	return fileInput
}

func outputAscii(asciiMap map[int][]string, input string) {
	inputSlice := strings.Split(input, "\\n") // Takes input e.g. "Hello\nWorld" and converts to slice. In this case, we have two elements (words) in slice (splitting by new line)
	for _, input = range inputSlice {         // Range over inputSlice (Hello World) word by word
		for i := 0; i < 8; i++ { // Iterates through 8 lines as that's how many lines there are for each ascii character
			for _, inputChar := range input { // Range over input (H e l l o W o r l d) character by character
				fmt.Print(asciiMap[int(inputChar)][i]) // Casting inputChar (rune) to int as our map key is int; [i] print by index
			}
			fmt.Println() // Alternative: fmt.Println(""); print new line is required so that characters are printed properly as opposed to on one line e.g. go run . "a" =>  / _` | | (_| |  \__,_|
		}
	}
}

/*
package main

import "fmt"

func main() {
m := make(map[int][]string)
m[1] = []string{"      ", "      ", "      ", "      ", "      ", "      ", "      ", "      "}
m[2] = []string{" _  ", "| | ", "| | ", "| | ", "|_| ", "(_) ", "    ", "    "}
m[3] = []string{"", "", "", "", "", "", "", ""}
m[4] = []string{"", "", "", "", "", "", "", ""}
m[34] = []string{"           ","    /\\     ", "   /  \\    ", "  / /\\ \\   ", " / ____ \\  ", "/_/    \\_\\ ", "           ", "           "} // Alternative to for loop: fmt.Println(m[34][0]) fmt.Println(m[34][1]) fmt.Println(m[34][2]) up to fmt.Println(m[34][7])

	for i := 0; i < 8; i++ {
	fmt.Println(m[2][i])
	}
-------------------------------
Alternative syntax:

	for i := range m[34] {
	fmt.Println(m[34][i])
	}
}
-------------------------------
package main

import "fmt"

func main() {
m := make(map[string]int)
m["James"] = 42
m["Amy"] = 16
fmt.Println(m["Amy"], m["James"])
delete(m, "James")
_, ok := m["James"] // Checking if item exists in a map
fmt.Println(ok)
fmt.Println(m["James"])
}
-------------------------------
Alternative syntax:

package main

import "fmt"

func main() {
m := map[string]int{
"James": 42,
"Amy":   16}
fmt.Println(m["Amy"], m["James"])
delete(m, "James")
_, ok := m["James"] // Checking if item exists in a map
fmt.Println(ok)
fmt.Println(m["James"])
}
-------------------------------
package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[72] = "H"
	m[101] = "e"
	m[108] = "l"
	m[108] = "l"
	m[111] = "o"

	fmt.Println(m[72], m[101], m[108], m[108], m[111])
	fmt.Println(m)
}
-------------------------------
scanner := bufio.NewScanner(file) // Scans the file and calls it "scanner"
inputScanner := bufio.NewScanner(os.Stdin) // Stdin is input on command line; standard input
fmt.Println("Please input what you would like to convert: ")
inputScanner.Scan() // This scans user input for word to convert
userInput := inputScanner.Text() // This creates the variable equal to user input; returns a string
*/
