package main

import "regexp"

const Consonants string = "ABCDFGHJKLMNPQRSTVWXZ"
const Vowels string = "[AEIOUYaeiouy]"

func VowelCount(name string) int {
	re := regexp.MustCompile(Vowels)
	return len(re.FindAllString(name, -1))
}

func ConsonantCount(name string) int {
	return 0
}

func LetterCount(name string) int {
	return 0
}
