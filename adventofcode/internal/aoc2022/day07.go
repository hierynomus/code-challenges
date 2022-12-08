package aoc2022

import (
	"bufio"
	"strconv"
	"strings"
)

const (
	Part1MaxDirSize = 100000
	DiskSize        = 70000000
	MinFreeSpace    = 30000000
)

type Dir struct {
	Dirs  map[string]*Dir
	Files map[string]int
	Name  string
}

func NewDir(name string) *Dir {
	return &Dir{
		Name:  name,
		Files: map[string]int{},
		Dirs:  map[string]*Dir{},
	}
}

func (d *Dir) AddFile(name string, size int) {
	d.Files[name] = size
}

func (d *Dir) AddDir(name string) {
	dir := NewDir(name)
	d.Dirs[name] = dir
}

func (d *Dir) Size() int {
	size := 0
	for _, dir := range d.Dirs {
		size += dir.Size()
	}

	for _, s := range d.Files {
		size += s
	}

	return size
}

func (d *Dir) Flatten() []*Dir {
	dirs := []*Dir{d}
	for _, dir := range d.Dirs {
		dirs = append(dirs, dir.Flatten()...)
	}
	return dirs
}

func BuildTree(reader *bufio.Scanner, curDir *Dir) {
	for reader.Scan() {
		line := reader.Text()
		s := strings.Split(line, " ")
		switch s[0] {
		case "$":
			switch s[1] {
			case "ls":
				continue
			case "cd":
				if s[2] == ".." {
					return
				} else if s[2] == "/" {
					continue
				} else {
					BuildTree(reader, curDir.Dirs[s[2]])
				}
			default:
				panic("unknown command")
			}
		case "dir":
			curDir.AddDir(s[1])
		default:
			size, err := strconv.Atoi(s[0])
			if err != nil {
				panic(err)
			}
			curDir.AddFile(s[1], size)
		}
	}
}

func Day07(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	root := NewDir("/")
	BuildTree(reader, root)

	for _, dir := range root.Flatten() {
		s := dir.Size()
		if s < Part1MaxDirSize {
			part1 += s
		}
	}

	minDirSize := root.Size()
	currFreeSpace := DiskSize - minDirSize
	for _, dir := range root.Flatten() {
		if MinFreeSpace <= currFreeSpace+dir.Size() {
			// we could remove this dir
			if dir.Size() < minDirSize {
				// It's the smallest dir we could remove up to now
				minDirSize = dir.Size()
			}
		}
	}

	part2 = minDirSize

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
