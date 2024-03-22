package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func moreType() {
	// go pointers works as C pointers
	pointers()

	// Same for structs
	structType()

	// Array
	arrayType()
	
	// SliceArray
	sliceArray()
	
	// rangeLoop
	rangeLoop()

	// Dictionnaire
	maps()

	// Fonction
	functionValues()
}

// & is the address operator
// * is value operator of a pointer  
func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	X,Y int
}

func structType(){
	structInitiation()
	// Inline structs
	structLiteral()
}

func structInitiation(){
	// Create struct
	v := Vertex{1, 2}
	fmt.Println(v)

	// changing attributs
	v.X = 4
	fmt.Println(v.X)
	// changing attributs via pointer
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func structLiteral(){
	
	v1 := Vertex{1, 2}  // has type Vertex
	v2 := Vertex{X: 1}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	p  := &Vertex{1, 2} // has type *Vertex

	fmt.Println(v1, v2, v3, p)
}

func arrayType(){
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func sliceArray()  {
	
	basicSlice()
	// SliceArrayLiterals
	sliceLiteral()

	// Slice capacity and length
	sliceCapacityAndLength()

	// The zero value of slice is nil
	nilSlice()

	// Make SliceArray
	makeSlice()

	// Append SliceArray
	appendSlice()

	sliceExample()

	sliceExercice()

}

func basicSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int
	var sBis []int = primes[1:4]
	sTer := primes[1:4]

	fmt.Println(s)
	fmt.Println(sBis)
	fmt.Println(sTer)

	// Slice works as array reference
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiteral()  {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	
	// These slice expressions are equivalent: 
	// a[0:10]
	// a[:10] with a := [11]int
	// a[0:]
	// a[:]
	sBis := []int{2, 3, 5, 7, 11, 13}

	sBis = sBis[1:4]
	fmt.Println(sBis)

	sBis = sBis[:2]
	fmt.Println(sBis)

	sBis = sBis[1:]
	fmt.Println(sBis)
}

func sliceCapacityAndLength()  {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeSlice() {

	a := make([]int, 5)
	
	printSliceBis("a", a)

	b := make([]int, 0, 5)
	printSliceBis("b", b)

	c := b[:2]
	printSliceBis("c", c)

	d := c[2:5]
	printSliceBis("d", d)
}

func printSliceBis(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sliceExample() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeLoop() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func sliceExercice()  {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := range image {
		image[i] = make([]uint8, dx)
	}
	
	for x :=range image{
		for y :=range image[x]{
			image[x][y] = uint8(x)*uint8(y)/10 + uint8(y)^2*uint8(x)
		}
	}
	
	return image
}

type floatVertex struct {
	Lat, Long float64
}

// The zero for this type is nil
var m map[string]floatVertex

func maps() {
	mapInitiation()

	mutatingMap()

	exerciceMap()
}

func mapInitiation() {
	m = make(map[string]floatVertex)
	m["Bell Labs"] = floatVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Literal version
	var mBis = map[string]floatVertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(mBis)

	// Literal Continued
	mBis = map[string]floatVertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(mBis)
}

func mutatingMap()  {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	myMap := map[string]int{}
	words := strings.Fields(s)
	for _, word := range words{
		myMap[word] +=1
	}
	return myMap
}

func exerciceMap() {
	wc.Test(WordCount)
}


func functionValues() {
	basicFunctionInitiation()
	
	functionCloser()
	
	fibonacciExercice()
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func basicFunctionInitiation(){
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionCloser() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n, nPlusOne, nPlusTwo := 0, 0, 0
	
	return func() int {
		nPlusTwo = nPlusOne + n
		n = nPlusOne
		nPlusOne = nPlusTwo
		if n == 0 && nPlusOne == 0 {
			nPlusOne = 1
		}
		return n
	}
}

func fibonacciExercice() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}