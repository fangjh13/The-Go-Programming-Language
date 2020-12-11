package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    m := make(map[string]int)
    if len(os.Args[1:]) > 0 {
        for _, f := range os.Args[1:] {
            f, err := os.Open(f)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2 error: %v\n", err)
                continue
            }
            readLines(f, m)
            f.Close()
        }
    } else {
        readLines(os.Stdin, m)
    }

    for line, n := range m {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func readLines(f *os.File, m map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        m[input.Text()]++
    }
}
