package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	for k, v := range r.Form {
		switch k {
		case "cycles":
			num, err := strconv.Atoi(v[0])
			if err != nil {
				continue
			}
			cycles = num
		case "res":
			num, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				continue
			}
			res = num
		case "size":
			num, err := strconv.Atoi(v[0])
			if err != nil {
				continue
			}
			size = num
		case "nframes":
			num, err := strconv.Atoi(v[0])
			if err != nil {
				continue
			}
			nframes = num
		case "delay":
			num, err := strconv.Atoi(v[0])
			if err != nil {
				continue
			}
			delay = num
		}
	}
	lissajous(w)
}

var palette = []color.Color{color.Black, color.RGBA{
	R: 0,
	G: 0xFF,
	B: 0,
	A: 0xFF,
}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

var (
	cycles  = 5     // number of complete x oscillator revolutions
	res     = 0.001 // angular resolution
	size    = 100   // image canvas covers [-size..+size]
	nframes = 64    // number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

func lissajous(out io.Writer) {

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
