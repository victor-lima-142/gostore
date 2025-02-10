package utils

import (
	"strconv"
)

// StringToUint takes a string and converts it to a uint. If the string is invalid, or is a negative number, 0 is returned.
func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

// StringArrToUintArr takes a slice of strings and converts each element to a uint.
// If an element is invalid, or is a negative number, 0 is used in its place.
// The function returns a slice of uints.
func StringArrToUintArr(s []string) []uint {
	var arr []uint
	for _, v := range s {
		arr = append(arr, StringToUint(v))
	}
	return arr
}
