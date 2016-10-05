//Attempt to replicate https://github.com/ALockwood/DiscountOffers in go
//Original Challenge: https://www.codeeval.com/public_sc/48/
package main

import (
	"fmt"
	//"github.com/alixaxel/go-gt/gt"
)

func main() {
	a := "foobArS" //3v,4c,7l

	//Test name parsing regexes
	fmt.Println(VowelCount(a))
	fmt.Println(ConsonantCount(a))
	fmt.Println(LetterCount(a))

	//Test scoring
	fmt.Println(SuitabilityScorer("aeio", "Wirewood Symbiote")) //9

	//Test matrix + Hungarian. Appears to maximize cost rather than minimize; good for this use case.
	//Original Source here: https://github.com/ThePaw/go-gt FIX REQUIRED: Swap lines 112 and 113 of Hungarian.go (TODO: Fork & push or pull req'd files into subdir)
	//g := new(gt.Matrix)
	g := NewMatrix(3)
	//g.N = 3
	// g.A = []int64{
	// 	2, 1, 1,
	// 	1, 2, 1,
	// 	1, 1, 2}

	// g.A = []int64{
	// 	2, 3, 3,
	// 	3, 2, 3,
	// 	3, 3, 2}

	g.Set(0, 0, 2)
	g.Set(0, 1, 3)
	g.Set(0, 2, 3)
	g.Set(1, 0, 3)
	g.Set(1, 1, 2)
	g.Set(1, 2, 3)
	g.Set(2, 0, 3)
	g.Set(2, 1, 3)
	g.Set(2, 2, 2)

	// g.A = []int64{
	// 	9, 1, 9, 9,
	// 	1, 9, 1, 9,
	// 	9, 1, 9, 1,
	// 	9, 9, 1, 9}
	//p, x := gt.Hungarian(g)
	p, _ := Hungarian(g)
	//fmt.Println(g.A)
	g.Print()
	fmt.Println(p)
	//fmt.Println(x)

	//Test reading from a file
	//GetCustomersAndProducts("/InputSample/InputSample.txt")

	//Test streaming from a file
	done := make(chan struct{})
	defer close(done)

	t := GetCustomersAndProductsStreamer(done, "/InputSample/InputSample.txt")
	for r := range t {
		fmt.Println("CUSTOMER LIST")
		fmt.Println(r.Customers)
		fmt.Println("PRODUCT LIST")
		fmt.Println(r.Products)
	}
}
