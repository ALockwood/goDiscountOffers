//Attempt to replicate https://github.com/ALockwood/DiscountOffers in go
//Original Challenge: https://www.codeeval.com/public_sc/48/
package main

import "fmt"

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

}
