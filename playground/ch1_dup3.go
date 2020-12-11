package main

import (
    "io/ioutil"
    "os"
    "fmt"
    "strings"
)

func main() {
    m := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: error %v\n", err)
        }
        for _, line := range strings.Split(string(data), "\n") {
            m[line]++
        }
    }

    for line, count := range m {
        if count > 1 {
            fmt.Printf("%d\t%s\n", count, line)
        }
    }

}
