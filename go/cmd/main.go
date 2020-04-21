package main

import (
	"github.com/nij4t/cs-lab/pkg/bitfields"
	"github.com/nij4t/cs-lab/pkg/hashing"
)

func main() {
	data := []string{"applejuice", "banana", "avocado", "mango", "staple", "grape", "potato", "rudders", "cheddar", "tomato"}

	for _, item := range data {
		println(hashing.CalculateHash(item, 10))
	}

	bitfields.Log("info", bitfields.INFO)
	bitfields.Log("notice", bitfields.NOTICE)
	bitfields.Log("warn", bitfields.WARN)
	bitfields.Log("error", bitfields.ERROR|bitfields.TIMESTAMP)
	bitfields.Log("debug", bitfields.DEBUG)
}
