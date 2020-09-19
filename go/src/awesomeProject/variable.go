package main

//例4.3 function_calls_function.go
//输出：GOG 可以看出main函数中的变量作用域是全局，普通函数的作用域是在函数体内
//var as string
//
//func main() {
//	as = "G"
//	print(as)
//	f1()
//}
//
//func f1() {
//	as := "O"
//	print(as)
//	f2()
//}
//func f2() {
//	print(as)
//}

//例：练习4.2 global_scope.go
//输出：GOO 说明直接在函数中给相同名称的全局变量用=赋值，修改了全局变量的值，而不是新创建了一个变量
//var as = "G"
//
//func main() {
//	n()
//	m()
//	n()
//}
//
//func n() { print(as) }
//func m() {
//	as = "O"
//	print(as)
//}

//例：练习4.1 local_scope.go
//输出：GOG 说明在函数体中给相同名称的变量用:=赋值是创建了一个作用域只在该函数体中的变量，修改该变量不影响全局变量的值
//var as = "G"
//
//func main() {
//	n()
//	m()
//	n()
//}
//
//func n() { print(as) }
//func m() {
//	a := "O"
//	print(a)
//}

//import (
//	"fmt"
//	"math"
//)
//
//var Pi float64
//
//func init() {
//	Pi = 4 * math.Atan(1)  // init() function computes pi
//}
//
//func main()  {
//	fmt.Println(Pi)
//}

//import "fmt"
//import "runtime"
//import "os"
//
//func main() {
//	var goos string = runtime.GOOS
//	fmt.Printf("This operating system is: %s\n", goos)
//	path := os.Getenv("PATH")
//	fmt.Printf("Path is: %s\n", path)
//}

//var as, bd int
//
//var (
//	a1 int = 15
//	a2 string = "this is a string value"
//	a3 float32 = 3.232
//	a4 float64 = 4.8000000
//	a5 bool = true
//)
//
//func main() {
//	fmt.Println(as, bd)
//	fmt.Println(a1, a2, a3, a4, a5)
//}
