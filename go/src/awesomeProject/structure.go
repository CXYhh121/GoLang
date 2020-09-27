package main

import "fmt"

// for 结构

//func main() {
//	for i := 0; i < 5; i++ {
//		fmt.Printf("This is the %d iteration\n", i)
//	}
//
//	//for i := 0; i < 2; i++ {
//	//	for j := 0; j < 5; j++ {
//	//		println(j)
//	//	}
//	//}
//
//	str := "Go is a beautiful language!"
//	fmt.Printf("The length of str is: %d\n", len(str))
//	for ix := 0; ix < len(str); ix++ {
//		fmt.Printf("Character on position %d is: %c\n", ix, str[ix])
//	}
//
//	str2 := "日本語"
//	fmt.Printf("The length of str2 is: %d\n", len(str2))
//	for ix :=0; ix < len(str2); ix++ {
//		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
//	}
//
//	for i := 0 ; i < 7; i++ {
//		for j := 0; j < i; j++ {
//			print("G")
//		}
//		println()
//	}
//}

//Break 与 continue
//func main() {
//	i := 10
//	for {
//		i = i - 1
//		fmt.Printf("The variable i is now: %d\n", i)
//		if i < 0 {
//			break
//		}
//	}
//    //break
//	for i := 0; i < 3; i++ {
//		for j := 0;j < 10; j++ {
//			if j > 5 {
//				break
//			}
//			print(j)
//		}
//		print(" ")
//	}
//	println()
//    //continue
//	for i := 0; i < 10; i++ {
//		if i == 5 {
//			continue
//		}
//		print(i)
//		print(" ")
//	}
//}

//标签与goto
func main() {

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE

	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
}
