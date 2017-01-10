package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/alockwood/munkres"
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

func buildMatrix(cuPr CustomerProductList) *munkres.FloatMatrix {
	var matrixDim int64

	if len(cuPr.Customers) > len(cuPr.Products) {
		matrixDim = int64(len(cuPr.Customers))
	} else {
		matrixDim = int64(len(cuPr.Products))
	}

	//Not great -- need to look into how to handle this
	if matrixDim == 0 {
		return munkres.NewMatrix(0)
	}

	pcfm := munkres.NewMatrix(matrixDim)

	for custIdx, cust := range cuPr.Customers {
		for prodIdx, prod := range cuPr.Products {
			pcfm.SetElement(int64(custIdx), int64(prodIdx), 0-SuitabilityScorer(cust, prod))
		}
	}
	return pcfm
}

func PrintScores(relativePathSourceFile string) {
	cpStream := getCustomersAndProductsStreamer("/InputSample/InputSample.txt")

	for cpPair := range cpStream {
		tmpMatrix := buildMatrix(cpPair)
		fmt.Printf("%.2f\n", math.Abs(munkres.GetMunkresMinScore(tmpMatrix)))
	}
}
