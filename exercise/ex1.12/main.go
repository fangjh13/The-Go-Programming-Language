
package main

import (
    "log"
    "io"
    "net/http"
    "image"
    "image/color"
    "image/gif"
    "math"
    "math/rand"
    "strconv"
    "fmt"
    "time"
)


func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    http.HandleFunc("/", handler)

    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    cycles := 5
    cyclesStr := r.FormValue("cycles")
    if cyclesStr != "" {
        var err error
        cycles, err = strconv.Atoi(cyclesStr)
        if err != nil {
            fmt.Fprintf(w, "bad cycles param: %s, %v", cyclesStr, err)
            return
        }
    }
    lissajous(w, cycles)
}


func lissajous(out io.Writer, cycles int) {
    const (
        res = 0.001
        size = 100
        nframes = 64
        delay = 8
    )
    var palette = []color.Color{color.White, color.Black}
    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
