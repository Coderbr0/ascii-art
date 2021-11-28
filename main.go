package main

import "fmt"

/*func ReadFile() []string {
	var fileInput []string
	file, err := os.Open(os.Args[1]) // We do this first to open the file
	if err != nil {
		fmt.Println("Invalid Input") // Alternative with fmt.Println("Invalid input. The named file does not exist"); a descriptive message may make it easier for the user to understand the error; go run main.go samp.txt => Invalid Input
	}
*/
func main() {
	// m := make(map[int][]string)

	a := []string{"           ","    /\\     ", "   /  \\    ", "  / /\\ \\   ", " / ____ \\  ", "/_/    \\_\\ ", "           ", "           "}
	// d := []string{"Hel", "S", "An", "Mam"}
// fmt.Println(a[0])
// fmt.Println(a[1])
// fmt.Println(a[2])

for i := range a {
	fmt.Println(a[i] + a[i]) 
}
	// m[1] = s
	// m[6] = d


	// fmt.Println(m[6][2])
	

	//ascii-art
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