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
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{11, 156, 49, 1}, color.RGBA{111, 73, 232, 1}}

const (
	blackIndex = 0 // Первый цвет палитры
	greenIndex = 1 // Следующий цвет палитры
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	nframes, err := strconv.Atoi(r.Form.Get("nframes")) // Количество кадров анимации
	if err != nil {
		nframes = 20
	}
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Конва изображения охватывает [size..+size]
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)
	rand.New(rand.NewSource((time.Now().UTC().UnixNano())))
	freq := rand.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	colorIndex := uint8(rand.Intn(len(palette)-1)+1)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
