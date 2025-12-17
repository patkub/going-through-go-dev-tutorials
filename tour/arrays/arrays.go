package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slice an array, a[low : high] including low and excluding high
	// slice of the indices 1 to  3
	var s []int = primes[1:4]
	fmt.Println(s)

	// slices pass by reference
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	b := names[0:2]
	c := names[1:3]
	fmt.Println(b, c)

	c[0] = "XXX"
	fmt.Println(b, c)
	fmt.Println(names)

	// slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct {
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
	fmt.Println(s2)

	// slice defaults
	s3 := []int{2, 3, 5, 7, 11, 13}

	s3 = s3[1:4]
	fmt.Println(s3)

	s3 = s3[:2]
	fmt.Println(s3)

	s3 = s3[1:]
	fmt.Println(s3)

	// slice length and capacity
	s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4)

	// Slice the slice to give it zero length.
	s4 = s4[:0]
	printSlice(s4)

	// Extend its length.
	s4 = s4[:4]
	printSlice(s4)

	// Drop its first two values.
	s4 = s4[2:]
	printSlice(s4)

	// Nil slices
	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var s5 []int
	fmt.Println(s5, len(s5), cap(s5))
	if s5 == nil {
		fmt.Println("nil!")
	}

	// slices with make
	a2 := make([]int, 5)
	printSlice2("a", a2)

	b2 := make([]int, 0, 5)
	printSlice2("b", b2)

	c2 := b2[:2]
	printSlice2("c", c2)

	d2 := c2[2:5]
	printSlice2("d", d2)
}
