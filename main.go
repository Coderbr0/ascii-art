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
	var emptySlice []string // Declaring an empty slice of string to use; we can then append to it later
	fileInput := make(map[int][]string) // Alternative syntax: var fileInput map[int][]string
	file, err := os.Open(os.Args[1])    // go run main.go standard.txt; standard.txt would be 1st argument in this case; alternatively we can use file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Invalid input. The named file does not exist")
	}
	defer file.Close()                // To close file is good practice; defer allows for all operations to be carried out before closing file
	scanner := bufio.NewScanner(file) // Scans the file and calls it "scanner"
	// inputScanner := bufio.NewScanner(os.Stdin) // Stdin is input on command line; standard input
	// fmt.Println("Please input what you would like to convert: ")
	// inputScanner.Scan() // This scans user input for word to convert
	// userInput := inputScanner.Text() // This creates the variable equal to user input; returns a string
	count := 32          // In ascii manual (man ascii in terminal) integer 32 (decimal value) represents space (line 2 to 9 in standard.txt file); the first character is rune 32 (decimal value 32 - space); we could use fileInput := make(map[rune][]string) with slight changes to code elsewhere; there are 95 characters to print (dec value 32 - 126 in ascii manual)
	countLines := 0
	for scanner.Scan() { // Scans the variable named scanner
		emptySlice = append(emptySlice, scanner.Text())
		countLines++
		if countLines == 9 { // If nine lines have been read
			fileInput[count] = emptySlice
			count++
			emptySlice = []string{} // We can also use emptySlice = nil
			countLines = 0
		}	
	}
	// fmt.Println(len(fileInput[34]))

// fmt.Println(string(fileInput[34][2]) + fileInput[32][2])
// fmt.Println(string(fileInput[34][3]) + fileInput[32][3])

for i := 0; i <9 ; i++{
	fmt.Println(string(fileInput[34][i]) + fileInput[32][i])
}


	// for _, el := range fileInput[34] {
	// 	fmt.Println(el)
	// }
		
	
	
		
		
		
		
	
	
	
	
	
	
	
	
	
	
	
	

	
	//}
	// for i, v := range fileInput[count+1] {
	// 	if i < 9 {
	// 	fmt.Println(i, v)
	// 	}
	// fmt.Println(fileInput[count][0])
	// fmt.Println(fileInput[count][1])
	// fmt.Println(fileInput[count][2])
	return fileInput
}

// 	for _, convertRune := range userInput { // Creates a for range loop based on user input; by default when iterating over a string, we obtain value as a rune
// 		emptySlice = append(emptySlice, string(convertRune)) // Casting individual characters (runes) to a string then appending to a slice of string
// 	}
// 	for index := range fileInput { // Creates a for range loop the length of the map; in this case, index is key of map and character is value of map
// 		fmt.Println(index)
// 		for _, character := range emptySlice {
// 			for i := 0; i < len(emptySlice); i++ { // runs a for range loop, in the lenght of how many characters to convert //	fmt.Println("number", number, "index", index)
// 			fmt.Println(emptySlice)
// 			if int(emptySlice[i][0]) == index { // if the character to convert == the individual key in the map
// 					for i := 0; i < 8; i++ { // creates a for loop to print out each line, less than 8 as there are only 8 lines per character
// 						fmt.Println(character[i])
// 					}
// 				}

// 			}
// 		}
// 	}
//
// }

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
*/
