package main

import "fmt"

//func main() {
//	var arr1 [5]int
//	for i := 0; i < len(arr1); i++ {
//		arr1[i] = i * 2
//	}
//
//	for i := 0; i < len(arr1); i++ {
//		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
//	}
//
//	//数组常量，通过数组常量来初始化数组，可初始化数组任意位置上的值
//	// var arrAge = [5]int{18, 20, 15, 22, 16}
//	// var arrLazy = [...]int{5, 6, 7, 8, 22}
//	// var arrLazy = []int{5, 6, 7, 8, 22}	//注：初始化得到的实际上是切片slice
//	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
//	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice
//
//	for i := 0; i < len(arrKeyValue); i++ {
//		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
//	}
//}

//多维数组 multidim_array.go

//const (
//	WIDTH = 5
//	HEIGHT = 5
//)
//
//type pixel int
//var screen [WIDTH][HEIGHT]pixel
//
//func main() {
//	for y := 0; y < HEIGHT; y++ {
//		for x := 0; x < WIDTH; x++ {
//			screen[x][y] = pixel(x * y)
//			fmt.Printf("this postions (%d,%d) value is: %d\n", x, y, screen[x][y])
//		}
//	}
//
//	array := [3]float64{7.0, 8.5, 9.3}
//	x := Sum(&array)
//	// to pass a pointer to the array
//	fmt.Printf("The sum of the array is: %f\n", x)
//}
//
//func Sum(a *[3]float64) (sum float64) {
//	for _, v := range a {// derefencing *a to get back to the array is not necessary
//		sum += v
//	}
//	return
//}

//切片，array_slices.go

//func main() {
//	var arr1 [6]int
//	var slice1 []int = arr1[2:5] //item at index 5 not included
//
//	//load the array with integers: 0,1,2,3,4,5
//	for i := 0; i < len(arr1); i++ {
//		arr1[i] = i
//	}
//
//	// print the slice
//	for i := 0; i < len(slice1); i++ {
//		fmt.Printf("Slice at %d is: %d\n", i, slice1[i])
//	}
//
//	fmt.Printf("The length of arr1 is: %d\n", len(arr1))
//	fmt.Printf("The length of slices1 is: %d\n", len(slice1))
//	fmt.Printf("The capacity of slice1 is: %d\n", cap(slice1))
//
//	//grow the slice
//	slice1 = slice1[0:4]
//	for i := 0; i < len(slice1); i++ {
//		fmt.Printf("Slice at %d is: %d\n", i, slice1[i])
//	}
//	fmt.Printf("The length of slices1 is: %d\n", len(slice1))
//	fmt.Printf("The capacity of slice1 is: %d\n", cap(slice1))
//
//	// grow the slice beyond capacity
//	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
//}

//func main() {
//	var slice1 []int = make([]int, 10)
//	//load the array/slice:
//	for i := 0; i < len(slice1); i++ {
//		slice1[i] = 5 * i
//	}
//
//
//	//print slice1
//	for i := 0; i < len(slice1); i++ {
//		fmt.Printf("Slice at %d is: %d\n", i, slice1[i])
//	}
//	fmt.Printf("The length of slice1 is: %d\n", len(slice1))
//	fmt.Printf("The capacity of slice1 is: %d\n", cap(slice1))
//}

//for range 结构
//func main() {
//	var slice []int = make([]int, 5)
//
//	for i := 0; i < len(slice); i++ {
//		slice[i] = 2 * i
//	}
//
//	for ix, value := range slice {
//		fmt.Printf("Slice at %d is:%d\n", ix, value)
//	}
//
//	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
//	for ix, season := range seasons {
//		fmt.Printf("Season %d is: %s\n", ix, season)
//	}
//
//	var season string
//	for _, season = range seasons {
//		fmt.Printf("%s\n", season)
//	}
//}

//reslice 切片重组
//func main() {
//	slice1 := make([]int, 0, 10)
//	// load the slice, cap(slice) is 10
//	for i := 0; i < cap(slice1); i++ {
//		slice1 = slice1[0:i+1]
//		slice1[i] = i
//		fmt.Printf("The length of slice is %d\n", len(slice1))
//	}
//
//	// print the slice:
//	for i := 0; i < len(slice1); i++ {
//		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
//	}
//}

// 切片的复制与重构 copy_append_slice.go
func main() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copy %d elements\n", n)

	sl3 := append(slFrom, 3, 5, 6)
	fmt.Println(sl3)

	slice1 := []byte{2, 3, 5}
	sl4 := AppendByte(slice1, 3, 5, 6, 8)
	fmt.Println(sl4)
}

//模拟实现append函数
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
