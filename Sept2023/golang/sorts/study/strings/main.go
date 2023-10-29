package main

import "fmt"

func main() {
	s := "aåå\u2318\u00E0\u0061\u0300"
	fmt.Println("bytes")
	for i := 0; i < len(s); i++ {
		fmt.Printf("% x\t", s[i])
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\t", string(s[i]))
	}
	fmt.Printf("\n----\n")
	fmt.Println("runes")
	for _, r := range s {
		fmt.Printf("%d\t", r)
	}
	fmt.Println()
	for _, r := range s {
		fmt.Print(string(r))
	}
	fmt.Printf("\n----\n")
}
