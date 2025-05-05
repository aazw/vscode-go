package main

import (
	"fmt"

	stdErrors "errors"
)

func main() {
	err := func001()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Println()
		fmt.Printf("err: %+v\n", err)
		return
	}
}

func func001() error {
	err := func002()
	if err != nil {
		return fmt.Errorf("func001: %w", err)
	}
	return nil
}

func func002() error {
	err := func003()
	if err != nil {
		return fmt.Errorf("func002: %w", err)
	}
	return nil
}

func func003() error {
	return stdErrors.New("new on func003")
}
