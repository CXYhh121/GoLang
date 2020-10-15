package main

import (
	"fmt"
	"io"
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

//func main() {
//	inputReader := bufio.NewReader(os.Stdin)
//	fmt.Println("Please input your name: ")
//	input, err := inputReader.ReadString('\n')
//
//	if err != nil {
//		fmt.Println("There were errors reading, exit program...")
//		return
//	}
//	fmt.Printf("Your name is %s", input)
//	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
//	switch input {
//	case "chenxiyue\r\n":  fmt.Println("Welcome chenxiyue!")
//	case "zhangweixin\r\n":   fmt.Println("Welcome zhangweixin!")
//	case "zhailijiao\r\n":     fmt.Println("Welcome zhailijiao!")
//	default: fmt.Printf("You are not welcome here! Goodbye!\n")
//	}
//
//	fmt.Printf("Your Name is %s", input)
//	switch input {
//	case "chenxiyue":
//		fallthrough
//	case "zhangweixin":
//		fallthrough
//	case "zhailijiao":
//		fmt.Printf("Welcome %s\n", input)
//	default:
//		fmt.Printf("You are not welcome here! Goodbye!\n")
//	}
//
//}

//读文件
//func main() {
//	inputFile, inputError := os.Open("input.dat")
//	if inputError != nil {
//		fmt.Printf("An error occurred on opening the inputfile\n" +
//			"Does the file exist?\n" +
//			"Have you got acces to it?\n")
//		return // exit the function on error
//	}
//	defer inputFile.Close()
//
//	inputReader := bufio.NewReader(inputFile)
//	for {
//		inputString, readerError := inputReader.ReadString('\n')
//		fmt.Printf("The input was: %s", inputString)
//		if readerError == io.EOF {
//			return
//		}
//	}
//}

//文件拷贝
func main() {
	CopyFile("target.txt", "source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
