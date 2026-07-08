package main

import "strconv"

func hexToDecimal(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}
