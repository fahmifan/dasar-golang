package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty")
	}

	return true, nil
}

func main() {
	var name string
	fmt.Print("Input your name : ")
	fmt.Scanln(&name)

	if valid, err := validate(name); !valid {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("halo", name)
}
