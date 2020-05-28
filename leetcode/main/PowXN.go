package main

func main() {
	myPow(2.0, 10)
}
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	result := 1.0
	contribute := x
	for n > 0 {
		if n%2 == 1 {
			result *= contribute
		}
		contribute *= contribute
		n /= 2
	}
	return result
}
