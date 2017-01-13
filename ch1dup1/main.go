// Dup1 prints the text of each line that appears more that
// once in the standard input, preceded by its count.
package main

import (
    "buffio"
    "fmt"
    "os"
)

funcd main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++

    }
    // NOTE: ignoring potential error from input.Err()
    for line, n := range counts {
        if n > 1{
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}