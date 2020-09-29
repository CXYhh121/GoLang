package main

//regexp 包

//func main() {
//	//目标字符串
//	searchIn := "John: 34689.347 William: 234325.343 Steve: 3454.345"
//	pat := "[0-9]+.[0-9]+"
//
//	f := func(s string) string {
//		v, _ := strconv.ParseFloat(s, 32)
//		return strconv.FormatFloat(v*2, 'f', 2, 43)
//	}
//
//	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
//		fmt.Println("Match found!!!")
//	}
//
//	re, _ := regexp.Compile(pat)
//	//将匹配到的部分替换为###.##
//	str := re.ReplaceAllString(searchIn, "###.##")
//	fmt.Printf("this string is: %s\n", str)
//	//参数为函数时
//	str2 := re.ReplaceAllStringFunc(searchIn, f)
//	fmt.Printf("this func string is: %s\n", str2)
//}

//锁和sync包  sync.Mutex互斥锁   sync.RWMutex读写锁 RLock可以支持一个写操作和多个读操作
