package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)


func main() {
	// Go basics
	packageVariableFunctions()
	
	// Flow and control statement
	flowControlAndStatement()

	// More types
	moreType()
}


func packageVariableFunctions()  {
	fmt.Println("My favorite number is", rand.Intn(10))
	
	// Use Uppercase for exported name
	fmt.Println(math.Pi)
	
	// Functions
	defineFunctions()
	
	// Key words
	keyWords()

	// Go available's type :
	goTypes()

	// Type conversion
	convertType()

	// Defining constant 
	defineConstant()
}

func defineFunctions()  {
	// Define function
	fmt.Println(add(42, 13))
	fmt.Println(addBis(42, 13))
	fmt.Println(split(17))
}

// Define function with func keyword, followed by args and args type in parenthesis,
// then the return value(s) and his(them) type
func add(x int, y int) int {
	return x + y
}

// If args are of same type, no need to write it again 
// there is no limit of return values
func addBis(x, y int) int {
	return x + y
}

// If return values have names, we can use naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func keyWords(){
	// Var key word define list of variables of ending type keyword
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)

	// If type is not define, it will take the type of the initializer
	var j int = 1
	var cBis, pythonBis, javaBis = true, false, "no!"
	fmt.Println(j, cBis, pythonBis, javaBis)
	
	// := short assignment statement can be used in place of a var declaration with implicit type
	k := 3
	cTer, pythonTer, javaTer := true, false, "no!"
	fmt.Println(k, cTer, pythonTer, javaTer)
}

func goTypes(){
	var ToBe bool = false
	var MaxInt uint64 = 1<<64 - 1 // int, int8/byte, int16, int32/rune, int64
	var MaxFloat float64 = 1<<64 - 1 + 0.05 // float32, float64
	var z complex128 = cmplx.Sqrt(-5 + 12i) // complex64, complex128

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", MaxFloat, MaxFloat)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Default values
	var i int // 0
	var f float64 // 0
	var b bool // false
	var s string // ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func convertType(){
	// The expression T(v) converts the value v to the type T
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func defineConstant()  {
	// Define constant with const key word and = operator
	const Big = 1 << 100
	const Small = Big >> 99

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	 return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}