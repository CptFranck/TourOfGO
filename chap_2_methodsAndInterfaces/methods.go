package main

import (
	"fmt"
	"math"
)

func main() {
	methods()
	
	methodsBis()
}


func methods()  {
	receiver()

	interfaces()
}

func receiver(){
	basicsReceiver()

	pointersReceiver()
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func basicsReceiver()  {
	v := Vertex{3, 4}
	// usual function (pointer used is not necessary cause we do not modified vertex data)
	fmt.Println(Abs(v))
	// receiver function
	fmt.Println(v.Abs())
	// receiver works with primitives
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

////////////////////////////////////////////////

// usual use
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// receiver use
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// receiver use with primitive
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

////////////////////////////////////////////////

func pointersReceiver()  {
	// pointer receiver function does
	// not need value address to be called 
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) 
	// interpreted as (&v).scale(10)
	v.scale(10)
	fmt.Println(v.Abs())

	// in usual functions address need to
	// be specified
	v = Vertex{3, 4}
	scale(&v, 10)
	fmt.Println(Abs(v))

	p := &v
	p.scale(10)
	fmt.Println(p.Abs())
}

////////////////////////////////////////////////

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

////////////////////////////////////////////////

func interfaces() {
	basicInterfaces()

	strongerInterfaces()

	emptyInterfaces()

	typeAssertions()

	switchType()
}

type Abser interface {
	Abs() float64
}

func basicInterfaces()  {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex implements Abser (work also with Vertex)
	fmt.Println(a.Abs())

	// Abs function works cause Abs as been implemented with *Vertex
	// but if it has not been the case, it would print an error
	a = v
	fmt.Println(a.Abs())
}

func strongerInterfaces()  {

	// default value and type of interface is nil
	var i InterfaceForPrint
	describe(i)

	// here interface is type of T with nil value
	var t *StructWithString
	i = t
	describe(i)
	i.MethodPrint()
	
	// here interface is type of T with "Hello" value
	i = &StructWithString{"Hello"}
	describe(i)
	i.MethodPrint()
	
	// here interface is type of Float with Pi value
	i = Float(math.Pi)
	describe(i)
	i.MethodPrint()
}

////////////////////////////////////////////////

type InterfaceForPrint interface {
	MethodPrint()
}

type StructWithString struct {
	S string
}

func (t *StructWithString) MethodPrint() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type Float float64

func (f Float) MethodPrint() {
	fmt.Println(f)
}

////////////////////////////////////////////////

// print value and package type variable of interface I
func describe(i InterfaceForPrint) {
	fmt.Printf("(%v, %T)\n", i, i)
}


func emptyInterfaces(){
	var i interface{}
	emptyDescribe(i)

	i = 42
	emptyDescribe(i)

	i = "hello"
	emptyDescribe(i)
}

func emptyDescribe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func typeAssertions(){
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// allow to test interface value type
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}

func switchType() {
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

