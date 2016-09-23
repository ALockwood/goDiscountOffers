package main

import "regexp"

//Consonants defines characters are considered consonants for the name parser.
const Consonants string = "BCDFGHJKLMNPQRSTVWXZbcdfghjklmnpqrstvwxz"

//Vowels defines characters considered vowels for the name parser.
const Vowels string = "AEIOUYaeiouy"

//VowelCount accepts a string and returns the number of vowels found.
func VowelCount(name string) int {
	return getCount(name, Vowels)
}

//ConsonantCount accepts a string and returns the number of consonants found.
func ConsonantCount(name string) int {
	return getCount(name, Consonants)
}

//LetterCount accepts a string and returns the sum of vowels and consonants found.
func LetterCount(name string) int {
	return getCount(name, Vowels+Consonants)
}

//getCount is a generic method used by the various vowel/consonant/letter count methods.
//Accepts the string to be counted and the characters used for comparison.
func getCount(name string, regex string) int {
	re := regexp.MustCompile("[" + regex + "]")
	return len(re.FindAllString(name, -1))
}
