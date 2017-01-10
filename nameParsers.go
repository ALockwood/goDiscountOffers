package main

import "regexp"

const (
	//Consonants defines characters are considered consonants for the name parser.
	consonants string = "BCDFGHJKLMNPQRSTVWXZbcdfghjklmnpqrstvwxz"
	//Vowels defines characters considered vowels for the name parser.
	vowels string = "AEIOUYaeiouy"
)

//VowelCount accepts a string and returns the number of vowels found.
func VowelCount(name string) int {
	return getCount(name, vowels)
}

//ConsonantCount accepts a string and returns the number of consonants found.
func ConsonantCount(name string) int {
	return getCount(name, consonants)
}

//LetterCount accepts a string and returns the sum of vowels and consonants found.
func LetterCount(name string) int {
	return getCount(name, vowels+consonants)
}

//getCount is a generic method used by the various vowel/consonant/letter count methods.
//Accepts the string to be counted and the characters used for comparison.
func getCount(name string, regex string) int {
	re := regexp.MustCompile("[" + regex + "]")
	return len(re.FindAllString(name, -1))
}
