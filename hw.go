package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// // 1.1
	// fmt.Println(os.Args)

	// // 1.2
	// for i, arg := range os.Args {
	// 	fmt.Println(i, arg)
	// }

	// 1.3
	// ...

	// 1.4
	// dup3()

	// 1.5
	// lissajous(os.Stdout) TODO

	// 1.7 | 1.8 | 1.9
	// fetch()

	// 1.10
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchall(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch

	}
	f, _ := os.Create("out.txt")
	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds()) // 1.10
}

// 1.4
func dup3() {
	counts := make(map[string]int)
	var filenames string
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] > 1 && !strings.Contains(filenames, filename) {
				filenames += filename
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(line, n)
		}
	}
	fmt.Println(filenames)
}

// 1.5 NOT DONE
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-sizwe..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Println(err)
	}
}

// 1.7
func fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") { // 1.8
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body) // 1.7
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println(resp.Status) // 1.9
	}
}

// 1.10
func fetchall(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

// 1.11
