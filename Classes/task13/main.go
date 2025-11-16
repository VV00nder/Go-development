package main

import "fmt"

func Stats(arr [5]int) (sum int, maxIdx int, maxVal int, positives []int, midl float64) {
	maxVal = -100000000000
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if arr[i] > 0 {
			positives = append(positives, arr[i])
		}
		if arr[i] > maxVal {
			maxVal = arr[i]
			maxIdx = i
		}
	}
	midl = float64(float64(sum) / float64(5))
	fmt.Printf("%v\n%v\n%v\n%v\n%.2f", sum, maxIdx, maxVal, positives, midl)
	return sum, maxIdx, maxVal, positives, midl
}

func main() {
	var a [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}
	Stats(a)

}
