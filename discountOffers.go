//Attempt to replicate https://github.com/ALockwood/DiscountOffers in go
//Original Challenge: https://www.codeeval.com/public_sc/48/
package main

import (
	"fmt"

	"github.com/alixaxel/go-gt/gt"
)

func main() {
	a := "foobArS" //3v,4c,7l
	testLine := "Jack Abraham,John Evans,Ted Dziuba;iPad 2 - 4-pack,Girl Scouts Thin Mints,Nerf Crossbow"

	//Test name parsing regexes
	fmt.Println(VowelCount(a))
	fmt.Println(ConsonantCount(a))
	fmt.Println(LetterCount(a))

	//Test customer and product parsing
	customers, products := CustomerProductParser(testLine)
	fmt.Println(customers)
	fmt.Println(products)

	//Test scoring
	fmt.Println(SuitabilityScorer("aeio", "Wirewood Symbiote")) //9

	//Test matrix + Hungarian.
	//Bad news: This is maxing the score (hungarian by default minimizes) and it panics when N == 3 (ex: 2,3,3:3,2,3:3,3,2)
	//ToDo: Troubleshoot the code here: https://github.com/alixaxel/go-gt/tree/master/gt or switch to https://github.com/clyphub/munkres
	g := new(gt.Matrix)
	g.N = 4
	g.A = []int64{
		0, 1, 9, 9,
		1, 0, 1, 9,
		9, 1, 0, 1,
		9, 9, 1, 0}
	p, _ := gt.Hungarian(g)
	fmt.Println(g.A)
	fmt.Println(p)
}
