package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var noOfArgs int = 0
	var arguments []string
	var inputStr string
	
	outputAscii()

	argument := os.Args[1]
	fmt.Println(argument)
	
	arguments = os.Args
	noOfArgs = len(arguments)

	switch noOfArgs {
	case 0: // case 0 may be used in IPC (inter-process communication) and Unix pipes e.g. go run . "Hello\n" | cat -e; | is used to combine two or more commands and allows them to operate simultaneously and permits data to be transferred between them; with cat -e "$" shows at the end of the line and useful to squeeze multiple lines into a single line
		fmt.Println("Error: No arguments.")
		return // return so that switch statement doesn't continue to the next case
	case 1:
		fmt.Println("Error: No arguments.")
		return
	case 2:
		/*Correct number of arguments, so we can continue.*/
		inputStr = arguments[1] // This reads what you input on the terminal as first argument	
		fmt.Println("Input string:", inputStr+";", "Number of arguments:", noOfArgs)
		fmt.Println("Correct number of arguments.")
		ReadFile()
		return
	case 3:
		fmt.Println("Error: There is only one argument allowed - \"a string inside quotes\"")
		return
	case 4:
		fmt.Println("Error: There is only one argument allowed - \"a string inside quotes\"")
		return
	default:
		/*Too many arguments*/
		fmt.Println("Error: Too many arguments.")
		return
	} /*switch*/
} /*main*/

func ReadFile() map[int][]string { // We want to return a map for simplicity
	var emptySliceToFill []string        // Declaring an empty slice of string to use; we can then append to it later
	fileInput := make(map[int][]string)  // Alternative syntax: var fileInput map[int][]string
	file, err := os.Open("standard.txt") // file, err := os.Open("os.Args[1]") => e.g. go run main.go standard.txt; standard.txt would be 1st argument in this case
	if err != nil {
		fmt.Println("Invalid input. The named file does not exist.")
	}
	defer file.Close()                // To close file is good practice; defer allows for all operations to be carried out before closing file
	scanner := bufio.NewScanner(file) // Scans the file and calls it "scanner"
	count := 32                       // In ascii manual (man ascii in terminal) integer 32 (decimal value) represents space (line 2 to 9 in standard.txt file); the first character is rune 32 (decimal value 32 - space); we could use fileInput := make(map[rune][]string) with slight changes to code elsewhere; there are 95 characters to print (dec value 32 - 126 in ascii manual)
	countLines := 0
	for scanner.Scan() { // Scans the variable named scanner
		emptySliceToFill = append(emptySliceToFill, scanner.Text())
		countLines++
		if countLines == 9 { // If nine lines have been read
			fileInput[count] = emptySliceToFill // Assigning this emptySliceToFill to map[32]
			count++                             // count = count + 1; 33
			// fmt.Println(emptySliceToFill)
			emptySliceToFill = []string{}       // We can also use emptySliceToFill = nil; reset emptySliceToFill to zero after nine lines have been counted
			countLines = 0
		}
	}
	for i, v := range fileInput[count] {
		if i < 9 {
		fmt.Println(i, v)
		}
	}
	return fileInput
}

func outputAscii() {
	var showAscii map[int][]string = ReadFile()
	fmt.Println("Show ascii map:", showAscii)
	fmt.Println(showAscii[104][0])
	fmt.Println(showAscii[104][1])
	fmt.Println(showAscii[104][2])
	fmt.Println(showAscii[104][3])
	fmt.Println(showAscii[104][4])
	fmt.Println(showAscii[104][5])
	fmt.Println(showAscii[104][6])
	fmt.Println(showAscii[104][7])
	fmt.Println(showAscii[104][8])

	for i := 0; i < 9 ; i++ {
		fmt.Println(showAscii[104][i] + showAscii[33][i])
	}

	
	// See below: why are we casting to a string?
	
	// fmt.Println(string(showAscii[104][0]) + showAscii[101][0] + showAscii[108][0] + showAscii[108][0] + showAscii[111][0])
	// fmt.Println(string(showAscii[104][1]) + showAscii[101][1] + showAscii[108][1] + showAscii[108][1] + showAscii[111][1])
	// fmt.Println(string(showAscii[104][2]) + showAscii[101][2] + showAscii[108][2] + showAscii[108][2] + showAscii[111][2])
	// fmt.Println(string(showAscii[104][3]) + showAscii[101][3] + showAscii[108][3] + showAscii[108][3] + showAscii[111][3])
	// fmt.Println(string(showAscii[104][4]) + showAscii[101][4] + showAscii[108][4] + showAscii[108][4] + showAscii[111][4])
	// fmt.Println(string(showAscii[104][5]) + showAscii[101][5] + showAscii[108][5] + showAscii[108][5] + showAscii[111][5])
	// fmt.Println(string(showAscii[104][6]) + showAscii[101][6] + showAscii[108][6] + showAscii[108][6] + showAscii[111][6])
	// fmt.Println(string(showAscii[104][7]) + showAscii[101][7] + showAscii[108][7] + showAscii[108][7] + showAscii[111][7])
	// fmt.Println(string(showAscii[104][8]) + showAscii[101][8] + showAscii[108][8] + showAscii[108][8] + showAscii[111][8])

	// fmt.Println(showAscii[33])
	// fmt.Println(len(showAscii[33]))

	// for i, el := range showAscii[33] {
	// 	fmt.Println(i, el)
	// }

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
