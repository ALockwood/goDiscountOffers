package main

import "regexp"

const Consonants string = "BCDFGHJKLMNPQRSTVWXZbcdfghjklmnpqrstvwxz"
const Vowels string = "AEIOUYaeiouy"

func VowelCount(name string) int {
	return getCount(name, Vowels)
}

func ConsonantCount(name string) int {
	return getCount(name, Consonants)
}

func LetterCount(name string) int {
	return getCount(name, Vowels+Consonants)
}

func getCount(name string, regex string) int {
	re := regexp.MustCompile("[" + regex + "]")
	return len(re.FindAllString(name, -1))
}
