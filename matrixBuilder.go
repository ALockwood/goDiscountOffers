package main

func BuildMatrix(cuPr CustomerProductList) *Matrix {
	var matrixDim int64

	if len(cuPr.Customers) > len(cuPr.Products) {
		matrixDim = int64(len(cuPr.Customers))
	} else {
		matrixDim = int64(len(cuPr.Products))
	}

	//Not great -- need to look into how to handle this
	if matrixDim == 0 {
		return NewMatrix(0)
	}

	pcm := NewMatrix(matrixDim)

	for custIdx, cust := range cuPr.Customers {
		/*fmt.Print("cIdx: ")
		fmt.Print(custIdx)
		fmt.Println("")
		fmt.Println(cust)*/
		for prodIdx, prod := range cuPr.Products {
			/*fmt.Print("pIdx: ")
			fmt.Print(prodIdx)
			fmt.Println("")
			fmt.Println(prod)*/
			pcm.Set(int64(custIdx), int64(prodIdx), int64(SuitabilityScorer(cust, prod)))
		}
	}

	return pcm
}
