package main

import "fmt"
import "runtime"

func getAvg(numbers []int, ch chan float64) {
	sum := 0

	for _, each := range numbers {
		sum += each
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	max := numbers[0]

	for _, e := range numbers {
		if max < e {
			max = e
		}
	}

	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)

	numbers := []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers : ", numbers)

	ch1 := make(chan float64)
	// ch1 passed here is a refference of ch1, it's pointer
	go getAvg(numbers, ch1)

	ch2 := make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)

		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}
