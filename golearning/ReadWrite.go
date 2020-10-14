package main

import (
	"bufio"
	"fmt"
	"os"
)

//var (
//	firstName, lastName, s string
//	i int
//	f float32
//	//input = "56.12 / 5212 / GO"
//	format = "%f / %d / %s"
//)
//
//var inputReader *bufio.Reader
//var input string
//var err error
//
//
//func main() {
//	//fmt.Println("Please enter you full name: ")
//	//fmt.Scanln(&firstName, &lastName)
//	////fmt.Scanf("%s %s", &firstName, &lastName)
//	//fmt.Printf("Hi %s %s!\n", firstName, lastName)
//	//fmt.Sscanf(input, format, &f, &i, &s)
//	//fmt.Println("From the string we read: ", f, i, s)
//
//	inputReader = bufio.NewReader(os.Stdin)
//	fmt.Println("Please enter some input: ")
//	input, err = inputReader.ReadString('\n')
//	if err == nil {
//		fmt.Printf("The input was: %s\n", input)
//	}
//}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name: ")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exit program...")
		return
	}

	fmt.Printf("Your Name is %s", input)
	switch input {
	case "chenxiyue":
		fallthrough
	case "zhangweixin":
		fallthrough
	case "zhailijiao":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

}
