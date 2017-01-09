package main

import "github.com/alockwood/goDiscountOffers/floatMatrix"

func BuildMatrix(cuPr CustomerProductList) *floatMatrix.FloatMatrix {
	var matrixDim int64

	if len(cuPr.Customers) > len(cuPr.Products) {
		matrixDim = int64(len(cuPr.Customers))
	} else {
		matrixDim = int64(len(cuPr.Products))
	}

	//Not great -- need to look into how to handle this
	if matrixDim == 0 {
		return floatMatrix.NewMatrix(0)
	}

	pcfm := floatMatrix.NewMatrix(matrixDim)

	for custIdx, cust := range cuPr.Customers {
		for prodIdx, prod := range cuPr.Products {
			pcfm.Set(int64(custIdx), int64(prodIdx), 0-SuitabilityScorer(cust, prod))
		}
	}
	return pcfm
}
