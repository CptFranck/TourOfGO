package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
)

func methodsBis()  {
	stringer()

	goError()

	goReaders()

	ImageExercice()
}

func stringer()  {
	stringerExample()
	
	stringerExercice()
}

func stringerExample() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

////////////////////////////////////////////////

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

////////////////////////////////////////////////

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string{
	// var string [4]string
	// for i, v := range ip{
	// 	string[i] = strconv.Itoa(int(v))
	// }
	// stringIPAddrs := strings.Join(string[:], ",")
	// return fmt.Sprintf("%v", stringIPAddrs)
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func stringerExercice() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func goError()  {
	ErrorExample()

	ErrorExercice()
}

///////////////////////////////////////////////////////

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}


func ErrorExample() {
	// test if an error has occurred
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

///////////////////////////////////////////////////////

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func ErrorExercice() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func goReaders() {
	basicReaders()

	ReaderExercice()

	ReaderExerciceBis()
}

func basicReaders()  {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
type Reader interface {
	Read(buf []byte) (n int, err error)
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (i int, e error) {
    for x := range b {
        b[x] = 'A'
    }
    return len(b), nil
}

func ReaderExercice() {
	reader.Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

func rot13(i byte) byte{
	switch{
	case 65 <= i && i < 78:
		fallthrough
	case 77 <= i && i < 110:
		i = i + 13
	case 78 <= i && i < 91:
		fallthrough
	case 110 <= i && i < 123:
		i = i - 13
	}
	return i
}

func (r13 rot13Reader) Read(b []byte) (int, error){
	n, err := r13.r.Read(b)
	for i := range b{
		b[i] = rot13(byte(b[i]))
	}
	return n, err
}

func ReaderExerciceBis() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

//////////////////////////////////////////////////

type Image struct{
	h int;
	w int;
	v int;
}


func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle  {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x,y int) color.Color  {
	return color.RGBA{uint8 (x), uint8 (i.v), uint8 (y), 255}
}

func ImageExercice() {
	m := Image{255,255,200}
	pic.ShowImage(m)
}