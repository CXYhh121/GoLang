package main

import "fmt"
import "./src/uc"

func main() {
	str1 := "USING package uc!"
	fmt.Println(uc.UpperCase(str1))
}

