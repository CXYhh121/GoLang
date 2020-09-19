package main

import "fmt" //package implementing formatted I/O

//println a int value
func helloGo() {
	var ret = 100
	fmt.Println(ret)
}

//return sum
func sum(a int, b int) int {
	return a + b
}

//return a float sum value
func floatSum(c float32, d float32) (t1 float32, t2 float32) {
	return c + d, c - d
}

//main function
func main() {
	const Pi = 3.1415926
	fmt.Println(Pi)
	var x, y = floatSum(3.123, 434.32)
	fmt.Println(x, y)
	helloGo()
	fmt.Println("hello world!")
}
