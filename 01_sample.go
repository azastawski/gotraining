// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const p string = "death and taxes"
const q = 42

const (
	A = iota
	B = iota
)

const (
	C = iota //0
	D
)

func main() {
	fmt.Println("Hello, world")
	PrintVariables("Jan Michael Smith")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
}

//this is exported?
func PrintVariables(name string) {
	var b string

	fmt.Println(b)

	b = name
	fmt.Println(b)

	fmt.Println(reverse(b))
}

//this is not exported
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
