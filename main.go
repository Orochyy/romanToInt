package main

import "fmt"

var romanInts = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var exceptions = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	n := len(s)
	result := 0
	i := 0

	for ; i < n-1; i++ {
		if val, found := exceptions[s[i:i+2]]; found {
			result += val
			i++
		} else {
			result += romanInts[s[i]]
		}
	}

	if i == n-1 {
		result += romanInts[s[n-1]]
	}

	return result
}

func main() {

	s := "CM"
	q := romanToInt(s)
	fmt.Println(q)

}
