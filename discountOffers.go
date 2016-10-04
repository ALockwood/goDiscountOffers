//Attempt to replicate https://github.com/ALockwood/DiscountOffers in go
//Original Challenge: https://www.codeeval.com/public_sc/48/
package main

import (
	"fmt"
	"os"

	"github.com/alixaxel/go-gt/gt"
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
	g := gt.NewMatrix(3)
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
	p, _ := gt.Hungarian(g)
	//fmt.Println(g.A)
	g.Print()
	fmt.Println(p)
	//fmt.Println(x)

	//Test reading from a file
	//So relative paths seem to be a pain in Go...
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	GetCustomersAndProducts(pwd + "/InputSample/InputSample.txt")

}
