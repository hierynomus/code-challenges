package aoc2024

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type DiskBlock struct {
	P, N   *DiskBlock
	Length int
	Idx    int
}

func (d *DiskBlock) IsSpace() bool {
	return d.Idx == -1
}

func (d *DiskBlock) IsFile() bool {
	return d.Idx >= 0
}

func Day09(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	line := aoc.ReadIntArrayLine(reader, "")
	disk, end := BuildDisk(line)
	// Compact the disk
	DefragPart1(disk, end)
	part1 = Checksum(disk)

	disk, end = BuildDisk(line)
	DefragPart2(disk, end)
	part2 = Checksum(disk)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func Checksum(disk *DiskBlock) int {
	checksum := 0
	pos := 0
	curr := disk
	for curr != nil {
		if curr.IsFile() {
			for i := 0; i < curr.Length; i++ {
				checksum += pos * curr.Idx
				pos++
			}
		} else {
			pos += curr.Length
		}
		curr = curr.N
	}

	return checksum
}

func DefragPart1(current *DiskBlock, end *DiskBlock) {
	for current != nil && end != nil && current != end {
		if end.IsSpace() || end.Length == 0 {
			end = end.P
			continue
		}

		if current.IsSpace() {
			len := current.Length
			endLen := end.Length
			if len >= endLen { // File fits in space
				// current - 1 -> current -> current + 1 -> ... -> end -1 -> end
				// current - 1 -> end -> current -> current + 1 -> ... -> end -1
				newEnd := end.P
				newEnd.N = end.N
				current.P.N, end.P = end, current.P
				current.Length -= endLen
				if current.Length == 0 {
					// Remove current as its depleted
					end.N, current.N.P = current.N, end
				} else {
					end.N, current.P = current, end
				}

				end = newEnd
			} else {
				// File partially fits
				current.Idx = end.Idx
				end.Length -= len
				// To Next
				current = current.N
			}
		} else {
			// To Next
			current = current.N
		}
	}
}

func DefragPart2(start *DiskBlock, end *DiskBlock) {
	for end != start {
		if start.IsFile() {
			start = start.N
			continue
		} else if start.IsSpace() && start.Length == 0 {
			start = start.N
			continue
		}

		if end.IsSpace() {
			end = end.P
			continue
		}

		// Start is space, end is file
		n := start
		moved := false
		// Find the first available space
		for n != end {
			if n.IsSpace() && n.Length >= end.Length {
				file := &DiskBlock{Length: end.Length, Idx: end.Idx}
				// n - 1 -> n -> n + 1 -> ... -> end -1 -> end
				// n - 1 -> end -> n -> n + 1 -> ... -> end -1
				end.Idx = -1 // Clear out
				n.P.N, file.P = file, n.P
				n.Length -= end.Length
				if n.Length == 0 {
					// Remove n as its depleted
					file.N, n.N.P = n.N, file
				} else {
					file.N, n.P = n, file
				}

				end = end.P
				moved = true
				break
			}

			n = n.N
		}

		if !moved {
			end = end.P
		}
	}
}

func BuildDisk(line []int) (*DiskBlock, *DiskBlock) {
	var disk = &DiskBlock{Length: line[0], Idx: 0}
	curBlock := disk
	for i := 1; i < len(line); i++ {
		if i%2 == 0 {
			newBlock := &DiskBlock{Length: line[i], Idx: i / 2}
			curBlock.N, newBlock.P = newBlock, curBlock
			curBlock = newBlock
		} else {
			newBlock := &DiskBlock{Length: line[i], Idx: -1}
			curBlock.N, newBlock.P = newBlock, curBlock
			curBlock = newBlock
		}
	}

	return disk, curBlock
}
