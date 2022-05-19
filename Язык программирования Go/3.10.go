// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
//  12,345
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var str bytes.Buffer
	n := len(s)
	ln := 4
	if n <= 3 {
		return s
	}
	switch {
	case n%3 == 1:
		str.WriteString(s[:1] + ",")
		ln = 1
	case n%3 == 2:
		str.WriteString(s[:2] + ",")
		ln = 2
	}

	for i := ln + 3; i <= n; i = i + 3 {
		if i < n {
			str.WriteString(s[ln:i] + ",")
		} else {
			str.WriteString(s[ln:i])
		}
		ln = ln + 3
	}
	return str.String()
}
