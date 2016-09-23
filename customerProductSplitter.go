package main

import "strings"

const CustomerProductSplitterChar string = ";"
const ItemSplitterChar string = ","

func customerProductSplitter(customerProductLine string) []string {
	return strings.Split(customerProductLine, CustomerProductSplitterChar)
}

func itemSplitter(itemList string) []string {
	return strings.Split(itemList, ItemSplitterChar)
}

func CustomerProductParser(customerProductList string) ([]string, []string) {
	tmpCustomerProductSlice := customerProductSplitter(customerProductList)

	if len(tmpCustomerProductSlice) == 2 {
		return itemSplitter(tmpCustomerProductSlice[0]), itemSplitter(tmpCustomerProductSlice[1])
	}

	return nil, nil
}
