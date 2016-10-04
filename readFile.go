package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const CustomerProductSplitterChar string = ";"
const ItemSplitterChar string = ","

type CustomerProductList struct {
	Customers []string
	Products  []string
}

func GetCustomersAndProducts(fileName string) *CustomerProductList {
	e := new(CustomerProductList)
	result, file := fileExists(fileName)
	defer file.Close()

	if result {
		r := bufio.NewReader(file)
		line, readErr := r.ReadString(10) //0x0A == newline \n
		if readErr == nil {
			c, p := splitCustomersAndProducts(line)
			fmt.Println(c)
			fmt.Println(p)
		} else {
			fmt.Println("Failed to read from file. ", readErr)
		}
	}

	return e
}

func splitCustomersAndProducts(lineItem string) (customers []string, products []string) {
	splitLine := strings.Split(lineItem, CustomerProductSplitterChar)

	if len(splitLine) != 2 {
		fmt.Println("Line parsing failed to find at least one product and/or customer.")
		return nil, nil
	}

	return strings.Split(splitLine[0], ItemSplitterChar), strings.Split(splitLine[1], ItemSplitterChar)
}

func fileExists(fileName string) (bool, *os.File) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Failed to open file! ", err)
		return false, nil
	}

	return true, file
}
