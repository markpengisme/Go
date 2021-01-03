/*
	go basic practice
	ref.[Go 語言基礎](https://willh.gitbook.io/build-web-application-with-golang-zhtw/02.0)
*/

package main

import (
	"errors"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"
	// . "fmt" // 引用套件函式，呼叫時省略套件名
	// f "fmt" // 引用套件函式，呼叫時使用別名
	// _ "fmt" // 不引用套件函式，只呼叫套件的init()
)

/* Go basic */
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
		c7         //常數宣告省略預設和上一個值相同，所以=iota
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

	// 修改字串 string -> []byte -> string
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

/* Go flow&func */
// flow
func flowControl() {
	x := 3
	// if-elseif-else
	if x == 3 {
		fmt.Println("x is equal to 3")
	} else if x < 3 {
		fmt.Println("x is less than 3")
	} else {
		fmt.Println("x is greater than 3")
	}

	// +宣告變數
	if y := 1 + 1; y > 3 {
		fmt.Println("y is greater than 3")
	} else {
		fmt.Println("y is less than 3")
	}

	// for & continue & break
	var sum int
	sum = 0
	for index := 0; index < 10; index++ {
		if sum > 30 {
			break
		} else {
			sum += index
			continue
		}
	}
	fmt.Println("sum is equal to ", sum)

	// while: for省略前後表達式
	sum = 0
	for sum < 1000 {
		sum++
	}
	fmt.Println("sum is equal to ", sum)

	// iterate slice
	s := []byte{'a', 'b', 'c'}
	for i, v := range s {
		fmt.Printf("index=%d value=%c\n", i, v)
	}
	// iterate map
	m := map[int]byte{0: 'a', 1: 'b', 2: 'c'}
	for k, v := range m {
		fmt.Printf("key=%d value=%c\n", k, v)
	}
	// 用 _ 丟棄不要的變數
	for _, v := range s {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	// switch 如果z沒指明預設為匹配true
	// case 預設為break, 但如果有 fallthrough 會強制執行下個class
	z := 2
	switch z {
	case 1, 2, 3:
		fmt.Println("The z = 1~3")
		fallthrough
	case 4:
		fmt.Println("The z was <= 4")
		fallthrough
	case 5:
		fmt.Println("The z was <= 5")
		fallthrough
	case 6:
		fmt.Println("The z was <= 6")
	case 7:
		fmt.Println("The z was = 7")
	case 8:
		fmt.Println("The z was = 8")
	default:
		fmt.Println("default case")
	}
}

// function
//// 3參數 2回傳
func foo1(a, b, c int) (int, int) {
	return a + b + c, a * b * c
}

//// 可以直接命名回傳
func foo2(a, b, c int) (x int, y int) {
	x, y = a+b+c, a*b*c
	return
}

//// 可變參數函式(不定數量參數)
func foo3(arg ...int) int {
	sum := 0
	for _, n := range arg {
		sum += n
	}
	return sum
}

//// 傳指標(&=位置, *=實際數值)
func foo4(a *int) {
	*a++
}

// defer(延遲)，當函式執行到最後時，這些 defer 語句會按照逆序執行，適合用於關閉資源
func d() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// 函數變數
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

type testIntFunc func(int) bool

func testInt(a int, f testIntFunc) bool {
	return f(a)
}
func testIntSV(a int, f func(int) bool) bool {
	return f(a)
}
func multiFunc() {
	w, x := foo1(3, 2, 1)
	fmt.Println("a+b+c =", w, " a*b*c =", x)
	y, z := foo2(3, 2, 1)
	fmt.Println("a+b+c =", y, " a*b*c =", z)
	sum := foo3(1, 23456, 543, 21, 456, 43, 2453, 675, 43, 4256)
	fmt.Println("sum =", sum)
	foo4(&sum)
	fmt.Println("sum =", sum)
	d()
	fmt.Println("1 is odd?", testInt(1, isOdd))
	fmt.Println("1 is even?", testIntSV(1, isEven))
}

/* Go struct */
// struct
type person struct {
	name string
	age  int
}

// 匿名欄位(嵌入欄位)，重複欄位最外層的優先存取
type student struct {
	person
	speciality string
}

type skills []string // 自訂型別

type teacher struct {
	person     // 匿名欄位，struct
	skills     // 匿名欄位，自訂型別
	int        // 匿名欄位，int
	speciality string
}

func structPerson() {
	var alice person
	alice.name, alice.age = "Alice", 18
	bob := person{name: "Bob", age: 22}
	carol := person{"Carol", 27}
	dave := student{person{"Dave", 21}, "CS"}
	eve := teacher{person{"Eve", 40}, skills{"CS", "Mangement"}, 999, "CS"}
	fmt.Printf("Name:%s, Age:%d\n", alice.name, alice.age)
	fmt.Printf("Name:%s, Age:%d\n", bob.name, bob.age)
	fmt.Printf("Name:%s, Age:%d\n", carol.name, carol.age)
	fmt.Printf("Name:%s, Age:%d, Speciality:%s\n", dave.name, dave.age, dave.speciality)
	fmt.Printf("Name:%s, Age:%d, Skills:%s, int?:%d, Speciality:%s\n",
		eve.name, eve.age, eve.skills, eve.int, eve.speciality)
}

/* Go OO */

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Method(帶有接收者的函數):func (r ReceiverType) funcName(parameters) (results)

//// area example
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func shapeArea() {
	r1 := rectangle{12, 2}
	c1 := circle{5}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", c1.area())

}

//// color box example
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type color byte
type box struct {
	width, height, depth float64
	color                color
}
type boxList []box

func (b box) Volume() float64 {
	return b.width * b.height * b.depth
}

////// 指標作為 receiver
func (b *box) SetColor(c color) {

	(*b).color = c // 要改實際的值可以省略成 b.color
}
func (c color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}
func (bl boxList) BiggestColor() color {
	maxVolume := 0.00
	paintColor := color(WHITE)
	for _, box := range bl {
		if boxVolume := box.Volume(); boxVolume > maxVolume {
			maxVolume = boxVolume
			paintColor = box.color
		}
	}
	return paintColor
}
func someBoxes() {
	boxes := boxList{
		box{4, 4, 4, RED},
		box{10, 10, 1, WHITE},
		box{1, 1, 20, BLACK},
		box{10, 10, 1, BLUE},
		box{10, 30, 1, YELLOW},
	}
	for _, box := range boxes {
		fmt.Printf("box Volume=%.0f Color=%s\n", box.Volume(), box.color.String())
	}
	fmt.Printf("Paint color:%s\n", boxes.BiggestColor().String())
	for _, box := range boxes {
		(&box).SetColor(boxes.BiggestColor()) // 要改實際的值可以省略成 box.SetColor
	}

}

// Method 繼承(inheritance) & 覆寫(override)

type human struct {
	name string
	age  int
}

type employee struct {
	human
	company string
}

type boss struct {
	human
	company string
}

func (h human) SayHi() {
	fmt.Printf("Hi, I am %s, %d years old.\n", h.name, h.age)
}
func (h human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}
func (b boss) SayHi() {
	fmt.Printf("Shut the fuck off. I'm the boss\n")
}

//// Method繼承與覆寫使用
func humanInher() {
	alice := employee{human{"Alice", 45}, "Golang Inc"}
	bob := boss{human{"Bob", 30}, "Golang Inc"}
	alice.SayHi()
	bob.SayHi()
}

// interface: 一組 method 簽名的組合，我們透過 interface 來定義物件的一組行為。

type men interface {
	SayHi()
	Sing(lyrics string)
}

func inter() {
	mike := human{"Mike", 25}
	alice := employee{human{"Alice", 45}, "Golang Inc"}
	bob := boss{human{"Bob", 30}, "Golang Inc"}

	// 如果我們定義了一個 interface 的變數，那麼這個變數裡面可以存實現這個 interface 的任意型別的物件。
	var i men
	i = mike
	i.SayHi()
	i.Sing("~Mike")

	i = alice
	alice.SayHi()
	alice.Sing("~Alice")

	i = bob
	i.SayHi()
	bob.Sing("~Bob")

	// 空interface可以儲存任何值
	var a interface{}
	a = 5
	a = "Hello"
	a = true
	fmt.Println(a)
}

type element interface{}
type list []element

func (h human) String() string {
	// 如果需要某個型別能被 fmt 套件以特殊的格式輸出，你就必須實現 Stringer 這個介面
	// 	type Stringer interface {
	// 		String() string
	//    }
	return "(name: " + h.name + " - age: " + strconv.Itoa(h.age) + " years)"
}

//// 判斷 interface 儲存的東西
func interfaceStore() {
	l := make(list, 3)
	l[0] = 1
	l[1] = "Hello"
	l[2] = human{"Mark", 25}

	// Comma-ok 斷言: value, ok = element.(T)
	for index, element := range l {
		if value, ok := element.(int); ok {
			fmt.Printf("l[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("l[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(human); ok {
			fmt.Printf("l[%d] is a Human and its value is %s\n", index, value)
		} else {
			fmt.Printf("l[%d] is of a different type\n", index)
		}
	}

	// Switch
	for index, element := range l {
		switch value := element.(type) {
		case int:
			fmt.Printf("l[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("l[%d] is a string and its value is %s\n", index, value)
		case human:
			fmt.Printf("l[%d] is a Human and its value is %s\n", index, value)
		default:
			fmt.Printf("l[%d] is of a different type", index)
		}
	}
}

//// 嵌入欄位
/**
// # container/heap
// type Interface interface {
//     sort.Interface //嵌入欄位 sort.Interface
//     Push(x interface{}) //a Push method to push elements into the heap
//     Pop() interface{} //a Pop elements that pops elements from the heap
// }

// # io.ReadWriter
// type ReadWriter interface {
//     Reader
//     Writer
// }
**/

/* Go concurrent(並行) */

// go routine

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func gor() {
	go say("world") //開一個新的 Goroutines 執行
	say("hello")    //當前 Goroutines 執行
}

// channels: goroutine 之間資料通訊，預設情況下，channel 接收和傳送資料都是阻塞的
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func ch() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

//// buffer channel
func bch() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

//// 消費者 range 取得 channel 值，生產者 close channel
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rbch() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

//// Select: 可以監聽多個 channel 上的資料流動。
func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		// 超時
		case <-time.After(5 * time.Second):
			println("timeout")
			return
		}
	}
}

func sch() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func main() {

	// Hello world! %v可以印出任何型別
	fmt.Printf("Hello, world%v\n", "!")

	/* Go basic */
	b()
	n()
	s()
	e()
	as()
	m()

	/* Go flow&func */
	flowControl()
	multiFunc()

	/* Go struct */
	structPerson()

	/* Go OO */
	shapeArea()
	someBoxes()
	humanInher()
	inter()
	interfaceStore()

	/* Go concurrent(並行) */
	gor()
	ch()
	bch()
	rbch()
	sch()
}
