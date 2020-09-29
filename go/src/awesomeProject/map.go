package main

import "fmt"

var (
	barVal = map[string]int{
		"alpha": 34, "bravo": 45, "charlie": 23, "delta": 87, "fave": 353, "echo": 56, "foxtrot": 12,
		"gold": 54, "hotel": 35, "ice_cream": 837, "july": 98, "keep": 8, "language": 98, "many": 23,
		"null": 24, "operator": 93, "queen": 94, "rain": 42, "street": 45, "turn": 76}
)

//map的类型切片
//func main() {
//	//map 的排序
//	fmt.Println("Unsorted: ")
//	for k, v := range barVal {
//		fmt.Printf("Key: %v, Value: %v\n", k, v)
//	}
//
//	//通过key值进行排序
//	keys := make([]string, len(barVal))
//	i := 0
//	for k, _ := range barVal{
//		keys[i] = k
//		i++
//	}
//	sort.Strings(keys)
//	fmt.Println()
//	fmt.Println("Keys Sorted: ")
//	for _, k := range keys {
//		fmt.Printf("Key: %v, Value: %v\n", k, barVal[k])
//	}
//
//	//通过value值进行排序 ---- 错误行不通
//	//values := make([]int, len(barVal))
//	//j := 0
//	//for _, v := range barVal {
//	//	values[i] = v
//	//	j++
//	//}
//	//fmt.Println("Value Sorted:")
//	//for _, k := range values {
//	//	fmt.Printf("Key: %v, Value: %v\n", barVal[k], k)
//	//}
//
//	// Version A:
//	items := make([]map[int]int, 5)
//	for i:= range items {
//		items[i] = make(map[int]int, 1)
//		items[i][1] = 2
//	}
//	fmt.Printf("Version A: Value of items: %v\n", items)
//
//	// Version B: NOT GOOD!
//	items2 := make([]map[int]int, 5)
//	for _, item := range items2 {
//		item = make(map[int]int, 1) // item is only a copy of the slice element.
//		item[1] = 2 // This 'item' will be lost on the next iteration.
//	}
//	fmt.Printf("Version B: Value of items: %v\n", items2)
//}

//将map的键值对调 invert_map.go
func main() {
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}
