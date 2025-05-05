package main

import (
	"fmt"

	crErrors "github.com/cockroachdb/errors"
)

func main() {
	err := func001()
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}
}

func func001() error {
	err := func002()
	if err != nil {
		return crErrors.Wrap(err, "wrap on func001")
	}
	return nil
}

func func002() error {
	err := func003()
	if err != nil {
		return crErrors.Wrap(err, "wrap on func002")
	}
	return nil
}

func func003() error {
	return crErrors.New("new on func003")
}
