package aoc2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type empty struct{}
type Paper map[string]struct{}
type Fold struct {
	axis  string
	where int
}

func Day13(reader *bufio.Scanner) (string, string) {
	var part1 int
	var part2 string

	paper := make(Paper)

	// Points
	for reader.Scan() {
		l := reader.Text()
		if strings.TrimSpace(l) == "" {
			break
		}
		p := aoc.ReadPoint(l)
		paper[p.Coords()] = empty{}
	}

	folds := []Fold{}
	// Folds
	for reader.Scan() {
		l := reader.Text()
		parts := strings.Split(l, " ")
		fld := strings.Split(parts[2], "=")
		folds = append(folds, Fold{fld[0], aoc.ToInt(fld[1])})
	}

	folded := FoldPaper(paper, folds[0])
	part1 = len(folded)

	for _, f := range folds[1:] {
		folded = FoldPaper(folded, f)
	}

	pts := []aoc.Point{}
	for p := range folded {
		pts = append(pts, aoc.ReadPoint(p))
	}

	grid := aoc.AsGrid(pts, rune(' '), rune('#'))
	part2 = aoc.RenderRuneGrid(grid)

	return strconv.Itoa(part1), part2
}

func FoldPaper(paper Paper, fold Fold) Paper {
	folded := make(Paper)
	if fold.axis == "x" {
		for xy := range paper {
			p := aoc.ReadPoint(xy)
			if p.X < fold.where {
				folded[p.Coords()] = empty{}
			} else {
				np := aoc.Point{X: fold.where - (p.X - fold.where), Y: p.Y}
				folded[np.Coords()] = empty{}
			}
		}
	} else if fold.axis == "y" {
		for xy := range paper {
			p := aoc.ReadPoint(xy)
			if p.Y < fold.where {
				folded[p.Coords()] = empty{}
			} else {
				np := aoc.Point{X: p.X, Y: fold.where - (p.Y - fold.where)}
				folded[np.Coords()] = empty{}
			}
		}
	}

	return folded
}
