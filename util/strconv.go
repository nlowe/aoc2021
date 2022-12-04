package util

import (
	"fmt"
	"strconv"
)

func MustAtoI(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return result
}

func MustSingleDigitAToI(a rune) int {
	result, err := strconv.Atoi(string(a))
	if err != nil {
		panic(fmt.Errorf("code point %s (%v) is not a digit", string(a), a))
	}

	return result
}
