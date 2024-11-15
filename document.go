package main

func fact[P ~int | ~float64](n P) P {
	if n <= 1 {
		return 1
	}
	return fact[P](n-1) * n
}

func main() {
	println(fact(5))
	println(fact(5.1))
}
