package main

import "fmt"

func main() {
	fmt.Printf("status: %s\n", Unknown.String())
	fmt.Printf("status: %s\n", Active.String())
	fmt.Printf("status: %s\n", Inactive.String())
}
