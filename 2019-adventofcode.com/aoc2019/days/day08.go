package days

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
)

type Day08 struct{}
type Layer []rune

const (
	Width       = 25
	Height      = 6
	Black       = '0'
	White       = '1'
	Transparent = '2'
)

func (d *Day08) Solve(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf(""))
	}

	all := scanner.Text()
	layers := []Layer{}

	for i := 0; i < len(all); i += Width * Height {
		layer := all[i : i+Width*Height]
		layers = append(layers, []rune(layer))
	}

	var hist aoc.RuneHistogram

	for _, l := range layers {
		h := aoc.MakeRuneHistogram(l)

		if hist == nil {
			hist = h
			continue
		}

		if h['0'] < (hist)['0'] {
			hist = h
		}
	}

	output := "\n"

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			for _, l := range layers {
				layerPixel := l[Width*y+x]
				if layerPixel == Black {
					output += " "
					break
				} else if layerPixel == White {
					output += "X"
					break
				}
			}
		}

		output += "\n"
	}

	return strconv.Itoa(hist['1'] * hist['2']), output
}
