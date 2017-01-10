//Attempt to replicate https://github.com/ALockwood/DiscountOffers in go
//Original Challenge: https://www.codeeval.com/public_sc/48/
package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		//PrintScores("/InputSample/InputSample.txt")
		PrintScores(argsWithoutProg[0])
	} else {
		fmt.Println("You must pass in a path to a file containing customer and product combinations for evaluation!")
	}
}
