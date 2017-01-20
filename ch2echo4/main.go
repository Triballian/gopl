// Echo4 prints its command-line arguemnts.
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/Triballian/gopl/ch2/tempconv"
)

var n = flag.Bool("n", false, "omit trailing newlines")
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
}
