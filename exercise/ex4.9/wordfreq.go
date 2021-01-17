package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    seen := map[string]int{}
    input := bufio.NewScanner(os.Stdin)
    input.Split(bufio.ScanWords)
    for input.Scan() {
        word := input.Text()
        seen[word]++
    }
    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
        os.Exit(1)
    }
    for w, n := range seen {
        fmt.Printf("%-10s %d\n", w, n)
    }
}
