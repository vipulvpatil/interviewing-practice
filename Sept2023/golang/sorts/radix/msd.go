package radix

type MSDSort struct {
	alphabetsMap map[alphabet]int
	auxArray     [][]Alphabet
}

func NewMSDSort(arr []string) {
	m := &MSDSort{}
	m.alphabetsMap = LexicographicallySortedAlphabetsMap()
	m.auxArray = make([][]Alphabet, len(arr))
	arrayOfAlphabetArray := convertToArrayOfAlphabetArray(arr)
	m.msdSort(arrayOfAlphabetArray, 0, 0, len(arr)-1)
	for i, a := range convertToArrayOfStrings(arrayOfAlphabetArray) {
		arr[i] = a
	}
}

func (m *MSDSort) msdSort(arr [][]Alphabet, sortKeyIndex int, start int, end int) {
	// Radix sort can only be sorted on a closed set of characters.
	if len(arr) == 0 {
		return
	}
	if start >= end {
		return
	}
	maxLength := 0
	for i := start; i <= end; i++ {
		alpArr := arr[i]
		if len(alpArr) > maxLength {
			maxLength = len(alpArr)
		}
	}
	if maxLength <= sortKeyIndex {
		return
	}
	alphabetsMap := LexicographicallySortedAlphabetsMap()
	count := make([]int, len(m.alphabetsMap)+2)
	for i := start; i <= end; i++ {
		alpArr := arr[i]
		index := 1
		if sortKeyIndex < len(alpArr) {
			key := alpArr[sortKeyIndex].Alphabet()
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
	for i := start; i <= end; i++ {
		alpArr := arr[i]
		index := 0
		if sortKeyIndex < len(alpArr) {
			key := alpArr[sortKeyIndex].Alphabet()
			if k, ok := alphabetsMap[key]; ok {
				index = k
			}
		}

		j := count[index]
		m.auxArray[start+j] = alpArr
		count[index]++
	}
	for i := start; i <= end; i++ {
		arr[i] = m.auxArray[i]
	}
	m.msdSort(arr, sortKeyIndex+1, start, start+count[0]-1)
	for i := 0; i < len(count)-1; i++ {
		m.msdSort(arr, sortKeyIndex+1, start+count[i], start+count[i+1]-1)
	}
}
