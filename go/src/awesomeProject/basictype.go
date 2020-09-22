package main

//Go中不允许不同类型之间的混合使用，但是对于常量的类型就限制的很少，因此是允许
//常量混合使用，举个例子看看：
//func main() {
//	var a int
//	var b int32
//	a = 15
//	// b = a + a //编译失败， 因为go语言不允许不同类型之间混合使用
//	b = b + 5   //编译成功，因为5是常量所以可以通过编译
//
//	fmt.Printf("this is a int value:%d\n", a)
//}

//Go语言只允许进行显示类型转换，不允许隐式类型转换
//func main() {
//	var n int16 = 34
//	var m int32
//
//	//compiler error:cannot use n(type int16)as type int32 in assignment
//	//m = n   隐式类型转换
//	m = int32(n)   //显示类型转换
//	fmt.Printf("32 bit int is: %d\n", m)
//	fmt.Printf("16 bit int is: %d\n", n)
//}

//数字值转换，当进行类似 a32bitInt = int32(a32Float) 的转换时，小数点后的数字将被丢弃。
//这种情况一般发生当从取值范围较大的类型转换为取值范围较小的类型时，或者你可以写一个专门用于处理类型转换的函数来确保没有发生精度的丢失。
//下面这个例子展示如何安全地从 int 型转换为 int8：

//func Uint8FromInt(n int) (uint8, error) {
//	if 0 <= n && n <= math.MaxUint8 {
//		return uint8(n), nil
//	}
//	return 0, fmt.Errorf("%d is out of the uint8 range", n)
//}
//
//func IntFromUint64(x float64) int {
//	if math.MinInt32 <= x && x <= math.MaxInt32 {
//		whole, fraction := math.Modf(x)
//		if fraction >= 0.5 {
//			whole++
//		}
//		return int(whole)
//	}
//	panic(fmt.Sprintf("%g is out of the int32 range", x))
//}
//
//func main() {
//	var ret, err = Uint8FromInt(32)
//	var i = IntFromUint64(3435.2325)
//
//	fmt.Printf("Uint8 value is %d, error is %s\n", ret, err)
//	fmt.Printf("int value is %d\n", i)
//}

//随机数，rand包实现了伪随机数的生成

//func main()  {
//	for i := 0; i < 10; i++ {
//		a := rand.Int()
//		fmt.Printf("arand:%d /", a)
//	}
//	for i := 0; i < 5; i++ {
//		r := rand.Intn(8)
//		fmt.Printf("rrand:%d /", r)
//	}
//	fmt.Println()
//	timeNs := int64(time.Now().Nanosecond())
//	rand.Seed(timeNs)
//	for i := 0; i < 10; i++ {
//		fmt.Printf("frand:%2.2f /", 100*rand.Float32())
//	}
//}

//string

//func main() {
//	str := "beginning of the string" +
//		" second part of the string"
//	fmt.Printf("this string is: %s\n", str)
//
//	var str2 string = "hello"
//	// strings.Join(str2, "world")
//	str2 += ",world!"
//
//	fmt.Printf("this str2: %s\n", str2)
//}
