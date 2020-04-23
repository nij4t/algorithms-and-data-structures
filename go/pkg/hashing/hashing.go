package hashing

func CalculateHash(key string, size int) int {
	sum := 0
	for _, c := range key {
		sum += int(c)
	}
	return sum % size
}

func CalculateHashUsingPrimes(key string, size int) (index int) {
	sumOfPrime := 0
	for _, c := range key {
		sumOfPrime += getPrime(c)
	}
	return sumOfPrime % size
}

func getPrime(char rune) (prime int) {
	// TODO Implement
	return 0
}
