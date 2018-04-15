package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hello")
	os.Exit(1)
	fmt.Println("selamat datang")
}
