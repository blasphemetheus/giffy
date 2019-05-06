package main
import (
  "image"
  "image/color"
  "image/gif"
  "io"
  "math"
  "math/rand"
  "os"
)

// var someColor = color.RGBA{0xRR, 0xGG, 0xBB, 0xff} which is #RRGGBB
var green = color.RGBA{0x32, 0xCD, 0x32, 0xff}
var orange = color.RGBA{0xff, 0xA5, 0x00, 0xff}
var red = color.RGBA{0xff, 0x00, 0x00, 0xff}
var yellow = color.RGBA{0xff, 0xff, 0x33, 0xff}
var indigo = color.RGBA{0x4b, 0x00, 0x82, 0xff}
var blue = color.RGBA{0x1e, 0x90, 0xff, 0xff}
var grey = color.RGBA{0xa9, 0xa9, 0xa9, 0xff}

var palette = []color.Color{color.Black, color.White,
  green, orange, red, yellow, indigo, blue, grey}

const (
  blackIndex = 0
  whiteIndex = 1
  greenIndex = 2
  orangeIndex = 3
  redIndex = 4
  yellowIndex = 5
  indigoIndex = 6
  blueIndex = 7
  greyIndex = 8
)

func main() {
  lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
  const (
    cycles = 5 // num complete x oscillator revolutions
    res = 0.001 // angular resolution
    size = 100 // image canvas covers [-size to +size]
    nframes = 64 // number of animation frames
    delay = 8 // delay between frames in 10ms units
  )

  freq := rand.Float64() * 3.0 // relative frequency of y oscillator
  anim := gif.GIF{LoopCount: nframes}
  phase := 0.0 // phase difference
  for i := 0; i < nframes; i++ {
    rect := image.Rect(0, 0, 2*size+1, 2*size+1)
    img := image.NewPaletted(rect, palette)
    for t := 0.0; t < cycles*2*math.Pi; t += res {
      x := math.Sin(t)
      y := math.Sin(t*freq + phase)
      img.SetColorIndex(size + int(x*size+0.5), size + int(y*size+0.5), greenIndex)
    }
    phase += 0.1
    anim.Delay = append(anim.Delay, delay)
    anim.Image = append(anim.Image, img)
  }
  gif.EncodeAll(out, &anim) // ignore encoding errors here
}
