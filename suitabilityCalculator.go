package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/alockwood/goDiscountOffers/munkres"
)

const (
	evenProductNameMultiplier       float64 = 1.5
	greatestCommonDivisorMultiplier float64 = 1.5
)

//Suitability Calculator Rules (from challenge site):
//Product Name - Even Letters
//  #CustomerNameVowelsCount x 1.5
//Product Name - Odd Letters
//  #CustomerNameConsonantsCount
//gcd(ProductName.LetterCount & CustomerName.LetterCount) >1 then ss x 1.5

func SuitabilityScorer(customerName string, productName string) float64 {

	var tmpSuitabilityScore float64

	if LetterCount(productName)%2 == 0 {
		tmpSuitabilityScore = float64(VowelCount(customerName)) * evenProductNameMultiplier
	} else {
		tmpSuitabilityScore = float64(ConsonantCount(customerName))
	}

	gcdResult := new(big.Int)
	p := big.NewInt(int64(LetterCount(productName)))
	c := big.NewInt(int64(LetterCount(customerName)))
	gcdResult.GCD(nil, nil, p, c)

	if gcdResult.Cmp(big.NewInt(1)) == 1 {
		tmpSuitabilityScore *= greatestCommonDivisorMultiplier
	}

	return tmpSuitabilityScore
}

func PrintScores(relativePathSourceFile string) {
	cpStream := getCustomersAndProductsStreamer("/InputSample/InputSample.txt")

	for cpPair := range cpStream {
		tmpMatrix := BuildMatrix(cpPair)
		fmt.Printf("%.2f\n", math.Abs(munkres.GetMunkresMinScore(tmpMatrix)))
	}
}
