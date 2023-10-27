package radix

import (
	"unicode/utf8"
)

func LSDSort(arr []string) {
	if len(arr) == 0 {
		return
	}
	maxLength := 0
	for _, s := range arr {
		len := utf8.RuneCountInString(s)
		if len > maxLength {
			maxLength = len
		}
	}
	if maxLength == 0 {
		return
	}
	for i := maxLength - 1; i >= 0; i-- {
		Sort(arr, i, 0, len(arr)-1)
	}
}

func Sort(arr []string, sortKeyIndex int, start int, end int) {
	// Radix sort can only be sorted on a closed set of characters.

	if len(arr) == 0 {
		return
	}
	if start >= end {
		return
	}

	arrayOfAlphabetArray := convertToArrayOfAlphabetArray(arr[start : end+1])

	alphabetsMap := LexicographicallySortedAlphabetsMap()
	count := make([]int, len(alphabetsMap)+2)
	for _, alphabetArray := range arrayOfAlphabetArray {
		index := 1
		if sortKeyIndex < len(alphabetArray) {
			key := alphabetArray[sortKeyIndex].Alphabet()
			if k, ok := alphabetsMap[key]; ok {
				index = k + 1
			}
		}
		count[index]++
	}

	for i := range count {
		if i > 0 {
			count[i] += count[i-1]
		}
	}

	auxArray := make([][]Alphabet, len(arrayOfAlphabetArray))
	for _, alphabetArray := range arrayOfAlphabetArray {
		index := 0
		if sortKeyIndex < len(alphabetArray) {
			key := alphabetArray[sortKeyIndex].Alphabet()
			if k, ok := alphabetsMap[key]; ok {
				index = k
			}
		}
		i := count[index]
		auxArray[i] = alphabetArray
		count[index]++
	}

	for i, a := range convertToArrayOfStrings(auxArray) {
		arr[start+i] = a
	}
}

func convertToArrayOfAlphabetArray(arr []string) [][]Alphabet {
	arrayOfAlphabetArray := [][]Alphabet{}
	for _, s := range arr {
		alphabetArray := []Alphabet{}
		for _, r := range s {
			a := alphabet(r)
			alphabetArray = append(alphabetArray, a)
		}
		arrayOfAlphabetArray = append(arrayOfAlphabetArray, alphabetArray)
	}
	return arrayOfAlphabetArray
}

func convertToArrayOfStrings(arr [][]Alphabet) []string {
	arrayOfStrings := []string{}
	for _, s := range arr {
		bArray := []byte{}
		for _, r := range s {
			b := make([]byte, utf8.RuneLen(r.Rune()))
			utf8.EncodeRune(b, r.Rune())
			bArray = append(bArray, b...)
		}
		arrayOfStrings = append(arrayOfStrings, string(bArray))
	}
	return arrayOfStrings
}
