package main

import (
	"github.com/nij4t/alorithms-and-data-structures/pkg/hashing"
)

func main() {
	data := []string{"applejuice", "banana", "avocado", "mango", "staple", "grape", "potato", "rudders", "cheddar", "tomato"}

	for _, item := range data {
		println(hashing.CalculateHash(item, 20))
	}
}
