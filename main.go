package main

import (
	"fmt"
	"os"
	"bufio"
)

func ReadFile() map[int][]string { // We want to return a map for simplicity
	file, err := os.Open(os.Args[1]) // go run main.go standard.txt; standard.txt would be 1st argument in this case; alternatively we can use file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Invalid input. The named file does not exist")
	}
	defer file.Close() // To close file is good practice; defer allows for all operations to be carried out before closing file
	scanner := bufio.NewScanner(file) // Scans the file and calls it "scanner"
	fileInput := make(map[int][]string) // Alternative syntax: var fileInput map[int][]string
	firstLine := true // We want to ignore the first line as it's empty in standard.txt file; checks for first line; see line 24
	count := 32 // In ascii manual (man ascii in terminal) integer 32 (decimal value) represents space (line 2 to 9 in standard.txt file)
	// The first character is rune 32; (maybe use rune instead of int?)
	inputScanner := bufio.NewScanner(os.Stdin) // Stdin is input on command line; standard input
	fmt.Println("Please input what you would like to convert: ")
	inputScanner.Scan() // This scans user input for word to convert
	for scanner.Scan() { // Scans the variable named scanner
		if scanner.Text() == "" && !firstLine { // If empty string and is not first line
			count++		// Increase ascii value
		} else if !firstLine { // If not first line append map with ascii char and its values from the scanner
			fileInput[count] = append(fileInput[count], scanner.Text()) // fileInput = append(fileInput, scanner.Text()) - we cannot append to a map but we can append to slices; fileInput is of type map[int][]string
		} 
	firstLine = false // ???Sets first line to false and is placed outside loop to stop ending the loop early
	} // ask Karolis line 29 and 30 of his code
	return fileInput
}


/*
func ReadFile() []string {
	var fileInput []string
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input. The named file does not exist")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileInput = append(fileInput, scanner.Text())
	}
	return fileInput
}
*/
func main() {
	ReadFile()
	// m := make(map[int][]string)
	// m[1] = []string{"      ", "      ", "      ", "      ", "      ", "      ", "      ", "      "}
	// m[2] = []string{" _  ", "| | ", "| | ", "| | ", "|_| ", "(_) ", "    ", "    "}
	// m[3] = []string{"", "", "", "", "", "", "", ""}
	// m[4] = []string{"", "", "", "", "", "", "", ""}
	// m[34] = []string{"           ","    /\\     ", "   /  \\    ", "  / /\\ \\   ", " / ____ \\  ", "/_/    \\_\\ ", "           ", "           "} // Alternative to for loop: fmt.Println(m[34][0]) fmt.Println(m[34][1]) fmt.Println(m[34][2])
	
	// for i := 0; i < 8; i++ {
	// 	fmt.Println(m[34][i])
	// }

	// for i := range map[int][]string {
	//  	fmt.Println(map[int][]string[i]) 
	// }

	// d := []string{"Hel", "S", "An", "Mam"}
// fmt.Println(a[0])
// fmt.Println(a[1])
// fmt.Println(a[2])

// for i := range m[1] {
// 	fmt.Println(m[1][i]) 
// }

// for i := range m[2] {
// 	fmt.Println(m[2][i]) 
// }

// for i := range m[34] {
// 	fmt.Println(m[34][i]) 
// }
	// m[1] = s
	// m[6] = d


	// fmt.Println(m[6][2])
	

	
}


/*
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
fmt.Println(m["Amy"])
delete(m, "James")
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