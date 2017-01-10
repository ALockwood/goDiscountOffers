package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	customerProductSplitterChar string = ";"
	itemSplitterChar            string = ","
)

type CustomerProductList struct {
	Customers []string
	Products  []string
}

func getCustomersAndProductsStreamer(relativePathFileName string) <-chan CustomerProductList {
	pcChan := make(chan CustomerProductList)

	go func() {
		result, file := fileExists(relativePathFileName)
		defer file.Close()

		if result {
			//if the file is open, attempt to read lines...
			s := bufio.NewScanner(file)
			for s.Scan() {
				line := s.Text()
				c, p := splitCustomersAndProducts(line)

				pcChan <- CustomerProductList{c, p}
			}
		}
		close(pcChan)
	}()
	return pcChan
}

func splitCustomersAndProducts(lineItem string) (customers []string, products []string) {
	splitLine := strings.Split(lineItem, customerProductSplitterChar)

	if len(splitLine) != 2 {
		fmt.Println("Line parsing failed to find at least one product and/or customer.")
		return nil, nil
	}

	return strings.Split(splitLine[0], itemSplitterChar), strings.Split(splitLine[1], itemSplitterChar)
}

func fileExists(relativeFileName string) (bool, *os.File) {
	//Get the current path and append to the relative path passed in
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	fullFileName := path.Join(currentDir, relativeFileName)

	file, err := os.Open(fullFileName)

	if err != nil {
		fmt.Println("Failed to open file! ", err)
		return false, nil
	}

	return true, file
}
