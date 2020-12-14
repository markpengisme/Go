package main

import (
	"errors"
	"fmt"
)

func b() {
	// boolean
	var b1 bool                            // 定義變數
	var b2, b3, b4 bool = true, true, true // 定義變數並賦值
	var b5, b6, b7 = true, true, true      // 省略型別
	b8, b9, b10 := true, true, true        // 簡短宣告，只能用在func裡
	// 宣告未使用會出錯
	fmt.Printf("b1 %t\n", b1)
	fmt.Println("b2~b4", b2, b3, b4)
	fmt.Println("b5~b7", b5, b6, b7)
	fmt.Println("b8~b10", b8, b9, b10)
}
func n() {
	// 整數、浮點數、複數、常數

	// 分組宣告
	var (
		n0      = 0      // int(default)
		n1 int  = 1      // int: int8,16,32(rune),int64
		n2 uint = 2      // uint: unit8,16,32(byte),uint64
		f0      = 0.0    // float32, float64(default)
		C0      = 1 + 1i // complex64, complex128(default)
	)
	const (
		PI  = 3.14
		MAX = 100
	)
	const (
		// iota 列舉從0開始
		c1, c2, c3 = iota, iota, iota // 同行值相同
		c4         = iota
		c5         = iota
		c6         = iota
		c7         //常數宣告省列預設和上一個值相同
	)
	fmt.Println("n0~n2", n0, n1, n2)
	fmt.Printf("f0: %T\n", f0)
	fmt.Printf("c0: %T\n", C0)
	fmt.Println("PI ,MAX:", PI, MAX)
	fmt.Println("c1~c7:", c1, c2, c3, c4, c5, c6, c7)
}

func s() {
	// 字串型別
	var s0 string = "!"
	var s1 = "?"
	s2, s3 := "yes", "no"
	fmt.Println("s0~s3", s0, s1, s2, s3)

	// 修改字串 string -> []byte -> byte
	s4 := "hat"
	c := []byte(s4)
	c[0] = 'f'
	s4 = string(c)
	fmt.Println("s4", s4)

	s5 := "bat"
	fmt.Println("s4+s5", s4+s5) // 字串相加
	m := ` good
	bye~` // 多行字串
	fmt.Println("m", m)
}

func e() {
	// Error型別
	err := errors.New("Find Error: you suck")
	if err != nil {
		fmt.Println(err)
	}
}

func as() {
	// array(參考型別)
	var arr0 [5]int                               // 宣告
	arr1 := [5]int{}                              // 簡短宣告
	arr2 := [5]int{1, 2, 3}                       // 賦值，其餘為0
	arr3 := [...]int{1, 2, 3}                     // 自動計算長度
	arr4 := [2][2]int{[2]int{1, 2}, [2]int{3, 4}} // 2d array
	arr5 := [2][2]int{{1, 2}, {3, 4}}             // 簡化宣告
	fmt.Println("arr0", arr0)
	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)
	fmt.Println("arr3", arr3)
	fmt.Println("arr4", arr4)
	fmt.Println("arr5", arr5)

	// slice: 動態陣列(參考型別)
	var slice0 []int
	slice1 := []byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Println("slice0", slice0)
	fmt.Printf("slice1 %s\n", slice1)
	fmt.Printf("slice1[:2] %s\n", slice1[:2])
	fmt.Printf("slice1[2:4] %s\n", slice1[2:4])
	fmt.Printf("slice1[4:] %s\n", slice1[4:])
	fmt.Printf("slice1[:] %s\n", slice1[:])

	slice2 := arr2[0:3:5] // len = 4-2, cap = 5-2
	slice3 := make([]int, 3, 5)
	fmt.Println("slice2,len,cap", slice2, len(slice2), cap(slice2))
	num := copy(slice3, slice2) // copy_num = copy(des, src)
	fmt.Println("slice3,len,cap,copyitem", slice3, len(slice3), cap(slice3), num)
	slice4 := append(slice3, 1, 2, 3) // append
	fmt.Println("slice4,len,cap", slice4, len(slice4), cap(slice4))
	slice5 := slice4
	slice5[0] = 10
	fmt.Println("slice4,len,cap", slice4, len(slice4), cap(slice4))
}

func m() {
	// map(參考型別)
	var map1 = make(map[string]int) // [key]value
	map2 := make(map[string]int)
	map3 := map[string]int{"pineapple": 3}
	map1["apple"] = 1
	map2["pen"] = 2
	fmt.Println("apple", map1["apple"])
	fmt.Println("pen", map2["pen"])
	fmt.Println("pineapple", map3["pineapple"])
	fmt.Printf("%T", map3)
}

func main() {

	// Hello world! %v可以印出任何型別
	fmt.Printf("Hello, world%v\n", "!")
	b()
	n()
	s()
	e()
	as()
	m()

}
