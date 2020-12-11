// 要输出出现重复行的文件的文件的名称，首先要知道某行可能在两个文件中，引入另一个map记录
// 每行对应的文件名，是一个slice
package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    counts := make(map[string]int)
    record := make(map[string][]string)
    if len(os.Args[1:]) > 0 {
        for _, filename := range(os.Args[1:]) {
            f, err := os.Open(filename)
            if err != nil {
                fmt.Fprintf(os.Stderr, "error open %s\t error: %v", err)
                continue
            }
            countLines(f, counts, record)
            f.Close()
        }
    } else {
        countLines(os.Stdin, counts, record)
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%s\t%d\t%v\n", line, n, record[line])
        }
    }
}

func countLines(f *os.File, counts map[string]int, record map[string][]string) {
    // 使用另一个`record`存储字符串和文件名
    input := bufio.NewScanner(f)
    for input.Scan() {
        text := input.Text()
        counts[text]++
        if !in(f.Name(), record[text]) {
            record[text] = append(record[text], f.Name())
        }
    }
}

func in(filename string, filenames []string) bool {
    // 判断在对应字符的文件名列表中这个文件名是否已经存在
    for _, f := range filenames {
        if f == filename {
            return true
        }
    }
    return false
}

