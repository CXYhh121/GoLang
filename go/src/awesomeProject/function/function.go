package main

import (
	"fmt"
	"time"
)

//func (st *Stack) Pop() int {
//	v := 0
//	for ix := len(st) - 1; ix >= 0; ix-- {
//		if v = st[ix]; v != 0 {
//			st[ix] = 0
//			return v
//		}
//	}
//}
//var num int = 10
//var numX2, numX3 int
//
//func main() {
//	numX2, numX3 = getX2AndX3(num)
//	PrintValues()
//	numX2, numX3 = getX2AndX3_2(num)
//	PrintValues()
//
//	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
//	// var i1 int = MultiPly3Nums(2, 5, 6)
//	// fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)
//}
//
//func PrintValues() {
//	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numX2, numX3)
//}
//
//func getX2AndX3(input int) (int, int) {
//	return 2 * input, 3 * input
//}
//
//func getX2AndX3_2(input int) (x2 int, x3 int) {
//	x2 = 2 * input
//	x3 = 3 * input
//	// return x2, x3
//	return
//}
//
//func MultiPly3Nums(a int, b int, c int) int {
//	// var product int = a * b * c
//	// return product
//	return a * b * c
//}

//返回值与函数的参数

//func main() {
//	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
//	// var i1 int = MultiPly3Nums(2, 5, 6)
//	// fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)
//}
//
//func MultiPly3Nums(a, b, c int) int {
//	return a * b * c
//	// var product int = a * b * c
//	// return product
//}

//空白符 blank identifier
//空白符可以用来匹配一些不需要的值，然后直接丢弃不使用
//func main() {
//	var i1 int
//	var f1 float32
//	i1, _, f1 = ThreeValues()
//	fmt.Printf("The int: %d, the float: %f \n", i1, f1)
//}
//
//func ThreeValues() (int, int, float32) {
//	return 5, 6, 7.5
//}
//func MinMax(a, b int) (min int, max int){
//	if a < b {
//		min = a
//		max = b
//	} else { // a > b or a = b
//		min = b
//		max = a
//	}
//	return
//}
//
//func Multiply(a, b int, reply *int) {
//	*reply = a * b
//}
//
//func main() {
//	var min, max int
//	min, max = MinMax(78, 65)
//	fmt.Printf("Minmium is: %d, Maxnmium is:%d\n", min, max)
//
//	//传指针，允许在外部函数中对主函数的值进行修改
//	n := 0
//	reply := &n
//	Multiply(5, 6, reply)
//	fmt.Printf("Multiply 5 * 6 = %d\n", n)
//}

//传递变长参数
//func main() {
//	x := min(1, 3, 2, 0)
//	fmt.Printf("The minimum is: %d\n", x)
//	slice := []int{7,9,3,5,1}
//	x = min(slice...)
//	fmt.Printf("The minimum in the slice is: %d", x)
//}
//
//func min(s ...int) int {
//	if len(s)==0 {
//		return 0
//	}
//	min := s[0]
//	for _, v := range s {
//		if v < min {
//			min = v
//		}
//	}
//	return min
//}

// defer 将需要执行的操作推迟到函数返回之前才执行，一般用于释放某些已分配资源的释放，降低内存泄漏的风险
//func main()  {
//	function1()
//	functionA()
//	functionB()
//	doDBOperations()
//}
//
//func function1() {
//	fmt.Printf("In function1 at the top\n")
//	defer function2()
//	fmt.Printf("In funtion1 at the bottom!\n")
//}
//
//func function2() {
//	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
//}
//
//func functionA() {
//	i := 0
//	defer fmt.Printf("This is a first: %d\n",i)
//	i++
//	defer fmt.Printf("This is a second: %d\n", i)
//	return
//}
//
//func functionB() {
//	for i := 0; i < 5; i++ {
//		defer fmt.Printf("%d ", i)
//	}
//}
//
//func connectToDB() {
//	fmt.Println("ok, connected to db")
//}
//
//func disconnectFromDB() {
//	fmt.Println("ok, disconnected from db")
//}
//
//func doDBOperations() {
//	connectToDB()
//	fmt.Println("Defering the database disconnect.")
//	defer disconnectFromDB() //function called here with defer
//	fmt.Println("Doing some DB operations ...")
//	fmt.Println("Oops! some crash or network error ...")
//	fmt.Println("Returning from function here!")
//	return //terminate the program
//	// deferred function executed here just before actually returning, even if
//	// there is a return or abnormal termination before
//}

//递归函数
//fibonacci.go

//func main() {
//	result := 0
//	for i := 0; i <= 20; i++ {
//		result = fibonacci(i)
//		fmt.Printf("This fibonacci number %d is :%d\n", i, result)
//	}
//}
//
//func fibonacci(n int) (res int) {
//	if n <= 1 {
//		return 1
//	} else {
//		res = fibonacci(n - 1) + fibonacci(n - 2)
//	}
//	return
//}

//mut_recurs.go
//func main() {
//	fmt.Printf("%d is even: is %t\n", 16, even(16))
//	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
//	//17 is odd: is true
//	fmt.Printf("%d is odd : is %t\n", 18, odd(18))
//	//18 is odd: is false
//}
//
//func even(nr int) bool {
//	if nr == 0 {
//		return true
//	}
//	return odd(RevSign(nr) - 1)
//}
//
//func odd(nr int) bool {
//	if nr == 0 {
//		return false
//	}
//	return even(RevSign(nr) - 1)
//}
//
//func RevSign(nr int) int {
//	if nr < 0 {
//		return -nr
//	}
//	return nr
//}

// 将函数作为参数，称之为回调函数

//func main() {
//	callback(5, Add)
//}
//func Add(a, b int) {
//	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
//}
//
//func callback(y int, f func(int, int)) {
//	f(y, 2)
//}

//应用闭包：将函数作为返回值
//func Add2() func(b int) int {
//	return func(b int) int {
//		return b + 2
//	}
//}
//
//func Adder(a int) func(b int) int {
//	return func(b int) int {
//		return a + b
//	}
//}
//
//func main() {
//	//make an Add2 function,give it a name p2, and call it:
//	p2 := Add2()
//	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
//	//make a specialAdder function,a gets value 2:
//	TwoAdder := Adder(2)
//	fmt.Printf("The result is: %v\n", TwoAdder(3))
//
//	start := time.Now()
//	fibonacci(20)
//	end := time.Now()
//	delta := end.Sub(start)
//	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
//}

//计算fibonacci数列的执行时间
//通过内存缓存提升fibonacci数列的执行效率
const LIM = 50

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
