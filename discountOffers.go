//Attempt to replicate https://github.com/ALockwood/DiscountOffers in go
//Original Challenge: https://www.codeeval.com/public_sc/48/
package main

import (
	"fmt"

	//"github.com/alixaxel/go-gt/gt"
	"github.com/alockwood/goDiscountOffers/floatMatrix"
	"github.com/alockwood/goDiscountOffers/munkres"
)

func main() {
	a := "foobArS" //3v,4c,7l

	//Test name parsing regexes
	fmt.Println(VowelCount(a))
	fmt.Println(ConsonantCount(a))
	fmt.Println(LetterCount(a))

	//Test scoring
	fmt.Println(SuitabilityScorer("aeio", "Wirewood Symbiote")) //9

	g := floatMatrix.NewMatrix(3)
	// g.A = []int64{
	// 	2, 1, 1,
	// 	1, 2, 1,
	// 	1, 1, 2}

	// g.A = []int64{
	// 	2, 3, 3,
	// 	3, 2, 3,
	// 	3, 3, 2}

	g.Set(0, 0, -6)
	g.Set(0, 1, -7)
	g.Set(0, 2, -6)
	g.Set(1, 0, -5)
	g.Set(1, 1, -6)
	g.Set(1, 2, -7)
	g.Set(2, 0, -6)
	g.Set(2, 1, -5)
	g.Set(2, 2, -9)

	fmt.Println(munkres.ComputeMunkres(g))

	//Test streaming from a file
	t := GetCustomersAndProductsStreamer("/InputSample/InputSample.txt")

	for r := range t {
		fmt.Println(r.Customers, r.Products)
		tmpMatrix := BuildMatrix(r)
		tmpMatrix.Print()
		fmt.Println(munkres.ComputeMunkres(tmpMatrix))
	}
}
