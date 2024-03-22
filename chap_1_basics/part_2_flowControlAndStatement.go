package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func flowControlAndStatement(){
	forLoop()

	ifStatement()

	exerciceForIf()

	switchStatement()

	differStatement()
}

func forLoop() {
	sum := 0
	// No parenthesis needed, only ;
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// For loop have default index and incrementation
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// Without ; for becomes while loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Infinite loop
	// for {
	// }
}

func ifStatement(){
	// If statement use no parenthesis
	fmt.Println(sqrt(2), sqrt(-4))

	// With an else
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
	// can't use v here
}

func exerciceForIf(){
	fmt.Println(Sqrt(2))
	
	fmt.Println(math.Sqrt(2))
}

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		v := z
		z -= (z*z - x) / (2*z)
		fmt.Printf("for i = %v, z =%v \n", i, z)
		if z == v || almostEqual(v, z){
			break
		}
	}
	return z
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
    return math.Abs(a - b) <= float64EqualityThreshold
}

func switchStatement() {
	// switch structure
	switchFindingCurrentOs()
	switchEvaluationOrder()
	//switch without condition, used for cleaned if statement
	switchAsIfStatement()
}

func switchFindingCurrentOs()  {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switchEvaluationOrder(){
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchAsIfStatement()  {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Differ statement differ action after the return of function
func differStatement()  {
	defer fmt.Println("world")
	fmt.Println("hello")

	StackingDiffer()
}

func StackingDiffer(){
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}