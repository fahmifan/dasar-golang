package main

import (
	"fmt"
)

// func main() {
// 	orderSomeFood("pizza")
// 	orderSomeFood("burger")
// }

func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih silahkan menunggu")
	if menu == "pizza" {
		fmt.Println("Pilihan tepat!, ", " ")
		fmt.Println("pizza di tempat kami paling enak")
		return
	}
	fmt.Println("Pilihan anda: ", menu)
}
