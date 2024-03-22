package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

func main() {
	concurrency()

	// mutex()
}

func concurrency() ()  {
	goRoutines()

	channels()

	goSelect()

	binaryTree()
}

func goRoutines()  {
	basicsRoutines()

	syncRoutines()
}

// use the main goroutines without
// "go" key world, to prevent the main
// goroutines to close the other one
func basicsRoutines()  {
	debut := time.Now()
    go run("Hatim", false)
    go run("Robert",false)
    run("Alex", false) 
	fin := time.Now()
    fmt.Println(fin.Sub(debut))
}

var wg sync.WaitGroup
// instanciation de notre structure WaitGroup
func syncRoutines()  {
	debut := time.Now()
    wg.Add(1)
    go run("Hatim", true)
    wg.Add(1)
    go run("Robert", true)
    wg.Add(1)
    go run("Alex", true)
    wg.Wait()
    fin := time.Now()
    fmt.Println(fin.Sub(debut))
}

func run(name string, sync bool) {
	if sync {
		defer wg.Done()
	}
    for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
        fmt.Println(name, " : ", i)
    }
}

//////////////////////////////////////////////

func channels()  {
	channelsBasic()
	
	bufferedChannels()

	rangeAndClose()
}

// by default channel are locked 
// witch means when channel receive value in go routine
// it will first use it, before take another value from another go routine
// 1) ch <- v   Send v to channel ch.
// 2) v := <-ch Receive from ch, and assign value to v.
func channelsBasic()  {
	s := []int{7, 2, 8, -9, 4, 0}
	
	// use make to create channel

	c := make(chan int)

	// fmt.Printf("%v\n",s[:len(s)/2])
	go func ()  {
		sum(s[:len(s)/2], c)
		time.Sleep(1*time.Millisecond)
	} ()
	// fmt.Printf("%v\n",s[len(s)/2:])
	go func ()  {
		sum(s[len(s)/2:], c)
		time.Sleep(10*time.Millisecond)
	} ()
	time.Sleep(1*time.Second)
	x, y := <-c, <-c // receive from c
	//17 -5 12
	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	println("valeur ajoutée à c")
}

// no error with channel sync
// cause channel used in one go routine
func bufferedChannels()  {
	ch := make(chan int, 2)
	// error channel buffer empty
	// fmt.Println(<-ch)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	ch <- 4
	// error channel buffer full 
	// ch <- 5
}


// when using range, use close()
// to prevent it from crashing
func rangeAndClose(){
	c := make(chan int, 10)
	// cBis := [3]int{1,2,3}
	go fibonacci(cap(c), c)
	// i value contained in channel
	for i := range c {
		// until c closed, do :
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func goSelect() {
	selectBasics()

	selectComplexe()

	selectExample()
}

// select take first response from one of go routine
func selectBasics()  {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for{
			time.Sleep(50*time.Millisecond)
			c1 <- "c1"
		}
		}()
		
	go func() {
		for{
			time.Sleep(100*time.Millisecond)
			c2 <- "c2"
		}
	}()

	for i := 0; i < 10; i++ {
		select{
			case res1 := <-c1:
				println(res1)
			case res2 := <-c2:
				println(res2)
		}
	}
}

func selectComplexe() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	
	fibonacciBis(c, quit)
}

func fibonacciBis(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
		}
	}
}

func selectExample() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func binaryTree()  {
	// ch := make(chan int)
	// t := tree.New(1)
	// go walk(t, ch)

	// for v := range ch {
	// 	println(v)
	// }

	// tBis := tree.New(1)
	// println(same(t, tBis))
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

func walk(t *tree.Tree, ch chan int)  {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int)  {
	if t != nil {
		ch <- t.Value
		_walk(t.Left, ch)
		_walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go walk(t1, ch1)
	go walk(t2, ch2)
	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func same(t1, t2 *tree.Tree) bool{
	ch1, ch2 := make(chan int), make(chan int)
	var s1, s2 [10]int

	go walk(t1, ch1)
	go walk(t2, ch2)

	channelToSlice(ch1, s1[:])
	channelToSlice(ch2, s2[:])

	fmt.Printf("%v \n %v \n",s1,s2)

	var control int
	for _, v:= range s1{
		control = 0
		for _, u := range s2{
			if v == u{
				control = 1
			}
		}
		if control == 0{
			return false
		}
	}
	return true
}

func channelToSlice(ch chan int, s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = <-ch
	}
}
