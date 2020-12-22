package main

import (
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, uri := range os.Args[1:] {
        go fetch(uri, ch)
    }

    for range os.Args[1:] {
        fmt.Println(<-ch)
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(uri string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(uri)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    // 创建文件句柄
    f, err := os.Create(url.QueryEscape(uri))
    if err != nil {
        ch <- err.Error()
    }
    nbytes, err := io.Copy(f, resp.Body)
    resp.Body.Close()
    // 关闭文件
    if closeErr := f.Close(); err == nil {
        err = closeErr
    }

    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", uri, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, uri)
}
