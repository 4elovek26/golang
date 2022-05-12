
package main

import (
	"fmt"
	"strings"
)

func main() {
	var scanstr, smallstr string
	fmt.Scan(&scanstr, &smallstr)
	fmt.Println(strings.Index(scanstr, smallstr))

}



