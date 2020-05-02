package main

import "fmt"

func findMedianSortedArrays(a []int, b []int) (float32, error) {
	aLength, bLength := len(a), len(b)
	if aLength > bLength {
		a, b, aLength, bLength = b, a, bLength, aLength
	}
	if bLength == 0 {
		return 0.0, fmt.Errorf("arrays mast be not null")
	}

	aMin, aMax, halfLen := 0, aLength, (aLength+bLength+1)/2
	for aMin <= aMax {
		i := (aMin + aMax) / 2
		j := halfLen - i
		if i < aLength && b[j-1] > a[i] {
			// i is too small, must increase it
			aMin = i + 1
		} else if i > 0 && a[i-1] > b[j] {
			// i is too big, must decrease it
			aMax = i - 1
		} else {
			// i is perfect
			maxOfLeft := 0
			if i == 0 {
				maxOfLeft = b[j-1]
			} else if j == 0 {
				maxOfLeft = a[i-1]
			} else {
				if a[i-1] > b[j-1] {
					maxOfLeft = a[i-1]
				} else {
					maxOfLeft = b[j-1]
				}
			}
			if (aLength+bLength)%2 == 1 {
				return float32(maxOfLeft), nil
			}

			minOfRight := 0
			if i == aLength {
				minOfRight = b[j]
			} else if j == bLength {
				minOfRight = a[j]
			} else {
				if a[i] < b[j] {
					minOfRight = a[i]
				} else {
					minOfRight = b[j]
				}
			}

			return float32(maxOfLeft+minOfRight) / 2.0, nil
		}
	}
	return 0.0, fmt.Errorf("no result")
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{1}
	arrays, err := findMedianSortedArrays(a, b)
	if err == nil {
		print(arrays)
	}
}
