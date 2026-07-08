package main

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	result := 1
	for i := 1; i <= 20; i++ {
		result *= i
	}
	return result
}
