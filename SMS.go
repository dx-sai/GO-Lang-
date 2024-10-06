package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costForCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costForSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return costForCustomer + costForSpouse, nil
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil

}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("===========")
	fmt.Println("msg To Customer", msgToCustomer)
	fmt.Println("msg To Spouse", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Total cost : $%.4f\n", totalCost)
}

func main() {
	test("Happy Birthday!", "Love you!")
	test("Anniversary Surprise", "Can’t wait to see you!")
	test("A long message that exceeds the maximum length allowed.", "Another short message.")
}
