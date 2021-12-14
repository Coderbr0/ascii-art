package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ReadFile()
}

func ReadFile() map[int][]string { // We want to return a map for simplicity
	var emptySliceToFill []string // Declaring an empty slice of string to use; we can then append to it later
	fileInput := make(map[int][]string) // Alternative syntax: var fileInput map[int][]string
	file, err := os.Open(os.Args[1])    // go run main.go standard.txt; standard.txt would be 1st argument in this case; alternatively we can use file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Invalid input. The named file does not exist.")
	}
	defer file.Close()                // To close file is good practice; defer allows for all operations to be carried out before closing file
	scanner := bufio.NewScanner(file) // Scans the file and calls it "scanner"
	count := 32          // In ascii manual (man ascii in terminal) integer 32 (decimal value) represents space (line 2 to 9 in standard.txt file); the first character is rune 32 (decimal value 32 - space); we could use fileInput := make(map[rune][]string) with slight changes to code elsewhere; there are 95 characters to print (dec value 32 - 126 in ascii manual)
	countLines := 0
	for scanner.Scan() { // Scans the variable named scanner
		emptySliceToFill = append(emptySliceToFill, scanner.Text())
		countLines++
		if countLines == 9 { // If nine lines have been read
			fileInput[count] = emptySliceToFill
			count++
			emptySliceToFill = []string{} // We can also use emptySliceToFill = nil; reset emptySliceToFill to zero after nine lines have been counted
			countLines = 0
		}
	}
	
	
// 	fmt.Println(string(fileInput[104][0]) + fileInput[101][0] + fileInput[108][0] + fileInput[108][0] + fileInput[111][0])
// 	fmt.Println(string(fileInput[104][1]) + fileInput[101][1] + fileInput[108][1] + fileInput[108][1] + fileInput[111][1])
// 	fmt.Println(string(fileInput[104][2]) + fileInput[101][2] + fileInput[108][2] + fileInput[108][2] + fileInput[111][2])
// 	fmt.Println(string(fileInput[104][3]) + fileInput[101][3] + fileInput[108][3] + fileInput[108][3] + fileInput[111][3])
// 	fmt.Println(string(fileInput[104][4]) + fileInput[101][4] + fileInput[108][4] + fileInput[108][4] + fileInput[111][4])
// 	fmt.Println(string(fileInput[104][5]) + fileInput[101][5] + fileInput[108][5] + fileInput[108][5] + fileInput[111][5])
// 	fmt.Println(string(fileInput[104][6]) + fileInput[101][6] + fileInput[108][6] + fileInput[108][6] + fileInput[111][6])
// 	fmt.Println(string(fileInput[104][7]) + fileInput[101][7] + fileInput[108][7] + fileInput[108][7] + fileInput[111][7])
// 	fmt.Println(string(fileInput[104][8]) + fileInput[101][8] + fileInput[108][8] + fileInput[108][8] + fileInput[111][8])

// 	// fmt.Println(fileInput[33])
// 	// fmt.Println(len(fileInput[34]))

// // fmt.Println(string(fileInput[34][2]) + fileInput[32][2])
// // fmt.Println(string(fileInput[34][3]) + fileInput[32][3])
// for i := 0; i < 9 ; i++ {
// 	fmt.Println(string(fileInput[34][i]) + fileInput[33][i])
// }


	// for i, el := range fileInput[33] {
	// 	fmt.Println(i, el)
	// }
	//}
	// for i, v := range fileInput[count+1] {
	// 	if i < 9 {
	// 	fmt.Println(i, v)
	// 	}
	
	return fileInput
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
