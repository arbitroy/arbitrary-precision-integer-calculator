package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BigInt struct {
	digits []int
	sign   int
}

// create number from input
func NewBigInt(s string) *BigInt {
	// Remove any leading/trailing spaces
	s = strings.TrimSpace(s)

	// Handle negative numbers
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}

	//convert string to digits in reverse order
	digits := make([]int, 0)
	for i := 0; i < len(s); i++ {
		//convert char to int
		digits[len(s)-1-i] = int(s[i] - '0')
	}

	return &BigInt{
		digits: digits,
		sign:   sign,
	}
}

// Display number as string
func (b *BigInt) String() string {
	// Convert back to string representation
	result := ""
	if b.sign < 0 {
		result = "-"
	}

	// reverse the digits back
	for i := len(b.digits) - 1; i >= 0; i-- {
		result += strconv.Itoa(b.digits[i])
	}
	return result
}



func main() {
	fmt.Println("Hello calculator")
}
