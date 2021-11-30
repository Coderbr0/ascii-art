package main

import (
	"fmt"
	"os"
	"bufio"
)

func ReadFile() map[int][]string {
	var fileInput map[int][]string
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input. The named file does not exist")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fileInput = append(fileInput, scanner.Text())
		fileInput[count] = append(fileInput[count], scanner.Text())
	}
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
	//strArr := ReadFile()
	m := make(map[int][]string)
	m[1] = []string{"      ", "      ", "      ", "      ", "      ", "      ", "      ", "      "}
	m[2] = []string{" _  ", "| | ", "| | ", "| | ", "|_| ", "(_) ", "    ", "    "}
	m[3] = []string{"", "", "", "", "", "", "", ""}
	m[4] = []string{"", "", "", "", "", "", "", ""}
	m[34] = []string{"           ","    /\\     ", "   /  \\    ", "  / /\\ \\   ", " / ____ \\  ", "/_/    \\_\\ ", "           ", "           "} // Alternative to for loop: fmt.Println(m[34][0]) fmt.Println(m[34][1]) fmt.Println(m[34][2])
	
	// d := []string{"Hel", "S", "An", "Mam"}
// fmt.Println(a[0])
// fmt.Println(a[1])
// fmt.Println(a[2])

for i := range m[1] {
	fmt.Println(m[1][i]) 
}

for i := range m[2] {
	fmt.Println(m[2][i]) 
}

for i := range m[34] {
	fmt.Println(m[34][i]) 
}
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