package radix

func ThreeWayQuickSort(arr []string) {
	rArr := []runeArray{}
	for _, s := range arr {
		rArr = append(rArr, newRuneArray(s))
	}
	threeWayQuickSort(rArr, 0, 0, len(arr)-1)
	for i, rArr := range rArr {
		arr[i] = rArr.String()
	}
}

func threeWayQuickSort(arr []runeArray, sortKeyIndex int, start int, end int) {
	if len(arr) <= 1 {
		return
	}
	if start >= end {
		return
	}
	maxLength := 0
	for _, rArr := range arr {
		if rArr.Len() > maxLength {
			maxLength = rArr.Len()
		}
	}
	if maxLength <= sortKeyIndex {
		return
	}

	lo := start
	var sortKeyValue rune
	for lo <= end {
		if sortKeyIndex < arr[lo].Len() {
			sortKeyValue = arr[lo].RuneAt(sortKeyIndex)
			break
		}
		lo++
	}
	hi := end
	i := lo + 1
	for i <= hi {
		if sortKeyIndex < arr[i].Len() {
			v := arr[i].RuneAt(sortKeyIndex)
			if v < sortKeyValue {
				temp := arr[lo]
				arr[lo] = arr[i]
				arr[i] = temp
				lo++
				i++
			} else if sortKeyValue < v {
				temp := arr[hi]
				arr[hi] = arr[i]
				arr[i] = temp
				hi--
			} else {
				i++
			}
		} else {
			temp := arr[lo]
			arr[lo] = arr[i]
			arr[i] = temp
			lo++
			i++
		}
	}
	threeWayQuickSort(arr, sortKeyIndex, start, lo-1)
	threeWayQuickSort(arr, sortKeyIndex+1, lo, hi)
	threeWayQuickSort(arr, sortKeyIndex, hi+1, end)
}
