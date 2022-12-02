package aoc2022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeGuide(t *testing.T) {
	tests := []struct {
		p1      RockPaperScissors
		outcome rune
		want    RockPaperScissors
	}{
		{ROCK, 'X', SCISSORS},
		{ROCK, 'Y', ROCK},
		{ROCK, 'Z', PAPER},
		{PAPER, 'X', ROCK},
		{PAPER, 'Y', PAPER},
		{PAPER, 'Z', SCISSORS},
		{SCISSORS, 'X', PAPER},
		{SCISSORS, 'Y', SCISSORS},
		{SCISSORS, 'Z', ROCK},
	}

	for _, tt := range tests {
		got := DecodeGuide(tt.outcome, tt.p1)
		assert.Equal(t, tt.want, got)
	}
}
