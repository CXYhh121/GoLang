package main

import "fmt"

//func main() {
//	var i1 = 5
//	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
//
//	var intP *int
//	intP = &i1
//	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
//
//	//string
//	str := "hello world"
//	fmt.Printf("here is the string str: %s\n", str)
//	var p *string = &str
//	*p = "ciao"
//
//	fmt.Printf("here is the pointer: %p\n", p)
//	fmt.Printf("here is the string *p: %s\n", *p)
//	fmt.Printf("here is the string str: %s\n", str)
//}

//if_else.go
//
//func main() {
//	var first int = 10
//
//	if first <= 0 {
//		fmt.Printf("first is less than or equal to 0\n")
//	} else if first > 0 && first < 5 {
//		fmt.Printf("first is between 0 and 5\n")
//	} else {
//		fmt.Printf("first is 5 or greater\n")
//	}
//
//	if cond := 5; cond > 10 {
//		fmt.Printf("cond is greater than 10\n")
//	} else {
//		fmt.Printf("cond is not greater 10\n")
//	}
//}

// 测试多返回值函数的错误
// string_conversion.go
//func main() {
//	var orig string = "ABC"
//	//var an int
//	var newS string
//	// vae err error
//
//	fmt.Printf("the size of ints is: %d\n", strconv.IntSize)
//	//anInt, err = strconv.Atoi(origStr)
//	an, err := strconv.Atoi(orig)
//	if err != nil {
//		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
//		return
//	}
//	fmt.Printf("The integer is %d\n", an)
//	an = an + 5
//	newS = strconv.Itoa(an)
//	fmt.Printf("The new string is: %s\n", newS)
//}

// switch
//func main() {
//	var num1 int = 99
//
//	switch num1 {
//	case 98, 99:
//		fmt.Println("It's equal to 98 or 99")
//		fallthrough
//	case 100:
//		fmt.Println("It's equal to 100")
//	default:
//		fmt.Println("It's not equal to 98 or 100")
//	}
//}

//func main() {
//	var num1 int = 7
//
//	switch {
//	case num1 < 0:
//		fmt.Println("Number is negative")
//	case num1 > 0 && num1 < 10:
//		fmt.Println("Number is between 0 and 10")
//	default:
//		fmt.Println("Number is 10 or greater")
//	}
//
//	k := 6
//	switch k {
//	case 4:
//		fmt.Println("was <= 4")
//		fallthrough
//	case 5:
//		fmt.Println("was <= 5")
//		fallthrough
//	case 6:
//		fmt.Println("was <= 6")
//		fallthrough
//	case 7:
//		fmt.Println("was <= 7")
//		fallthrough
//	case 8:
//		fmt.Println("was <= 8")
//		fallthrough
//	default:
//		fmt.Println("default case")
//	}
//}

// season.go

func main() {
	month := 9

	switch month {
	case 3, 4, 5:
		fmt.Printf("%d\n month in this year is Spring!", month)
	case 6, 7, 8:
		fmt.Printf("%d\n month in this year is Summer", month)
	case 9, 10, 11:
		fmt.Printf("%d\n month in this year is autumn", month)
	case 12, 1, 2:
		fmt.Printf("%d\n month in this year is winter", month)
	}
}
