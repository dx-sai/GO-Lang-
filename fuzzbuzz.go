package main

import "fmt"

/*
func maxMessages(tresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > tresh {
			return i
		}
	}
}

func main() {
	threshold := 550.0
	fmt.Printf("Max messages that can be sent with a threshold of %.2f pennies: %d\n", threshold, maxMessages(threshold))

}
*/

/*
func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0

	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
	return maxMessagesToSend
}

func main() {
	costMultiplier := 1.5    // Cost increases by 50% with each message
	maxCostInPennies := 1000 // Maximum cost threshold in pennies

	result := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
	fmt.Printf("Max messages that can be sent with a multiplier of %.2f and max cost of %d pennies: %d\n",
		costMultiplier, maxCostInPennies, result)
}

*/

func fizzbuzz() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
