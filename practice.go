package main

import "fmt"

func concat(s1, s2 string) string {
	return s1 + s2
}

func main() {

	accountage := 2.8
	accountageInt := int(accountage)
	fmt.Println(accountageInt)
	avgoprate, displaymessage := .23, "This is the one"
	fmt.Println(
		avgoprate,
		displaymessage,
	)
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = minutesInHour * secondsInMinute

	fmt.Println("no of seconds in hour :", secondsInHour)

	const name = "saul Goodman"
	const openRate = 20.585
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent",
		name,
		openRate)
	fmt.Println(msg)

	messageLen := 10
	maxMessageLen := 20
	if messageLen <= maxMessageLen {
		fmt.Println("Mesage  sent")
	} else {
		fmt.Println("Message not Sent")
	}

	fmt.Println(concat("This is \n", "Vamsi"))

	sendsSoFar := 430
	const sendsToAdd = 25

	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")

	test(messageToSend{
		phoneNumber: 231456,
		message:     "Thankyou for signing up",
	})
	test(messageToSend{
		phoneNumber: 321456,
		message:     "Love to have you aboard",
	})
	test(messageToSend{
		phoneNumber: 4123245,
		message:     "Wer are so excited to have you",
	})
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message:'%s' to: %v", m.message, m.phoneNumber)
	fmt.Println("=================")
}

type car struct {
	make  string
	model string
}

type truck struct {
	car
	bedsize int
}
