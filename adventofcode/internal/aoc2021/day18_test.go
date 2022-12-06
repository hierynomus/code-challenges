package aoc2021

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestDay18_sample(t *testing.T) {
	inp := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`

	d := day.TestDay(t, Day18)
	d.WithInput(inp, "4140", "")
}

// [[6,[5,[4,[3,2]]]],1]
//         /\
//        /\ 1
//       6 /\
//        5 /\
//         4 /\
//          3 2
func TestExplode(t *testing.T) {
	cases := map[string]string{
		"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}

	for sfn, res := range cases {
		t.Run(sfn, func(t *testing.T) {
			snailFish, i := ParseSnailFish(sfn, 0, nil)
			if i != len(sfn) {
				t.Fatalf("ParseSnailFishNumber(%q) = %v, %v", sfn, snailFish, i)
			}
			assert.True(t, snailFish.Explode())
			assert.Equal(t, res, snailFish.String())
		})
	}
}

func TestSplit(t *testing.T) {
	cases := map[string]string{
		"[10,9]": "[[5,5],9]",
		"[9,11]": "[9,[5,6]]",
	}

	for sfn, res := range cases {
		t.Run(sfn, func(t *testing.T) {
			snailFish, i := ParseSnailFish(sfn, 0, nil)
			if i != len(sfn) {
				t.Fatalf("ParseSnailFishNumber(%q) = %v, %v", sfn, snailFish, i)
			}
			assert.True(t, snailFish.Split())
			assert.Equal(t, res, snailFish.String())
		})
	}
}

func TestReduce(t *testing.T) {
	cases := map[string]string{
		"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]":                                                                                 "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]":                                         "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		"[[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]],[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]]":                 "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		"[[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]],[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]]": "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
		"[[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]],[7,[5,[[3,8],[1,4]]]]]":                                 "[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]",
		"[[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]],[[2,[2,2]],[8,[8,1]]]]":                                 "[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
		"[[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]],[2,9]]":                                                     "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		"[[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]],[1,[[[9,3],9],[[9,0],[0,7]]]]]":                                     "[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]]",
		"[[[[[7,8],[6,7]],[[6,8],[0,8]]],[[[7,7],[5,0]],[[5,5],[5,6]]]],[[[5,[7,4]],7],1]]":                                     "[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]",
		"[[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]],[[[[4,2],2],6],[8,7]]]":                                             "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
	}

	for sfn, res := range cases {
		t.Run(sfn, func(t *testing.T) {
			snailFish, i := ParseSnailFish(sfn, 0, nil)
			if i != len(sfn) {
				t.Fatalf("ParseSnailFishNumber(%q) = %v, %v", sfn, snailFish, i)
			}
			snailFish.Reduce()
			assert.Equal(t, res, snailFish.String())
		})
	}
}
func TestMagnitude(t *testing.T) {
	cases := map[string]int64{
		"[9,1]":                             29,
		"[1,9]":                             21,
		"[[9,1],[1,9]]":                     129,
		"[[1,2],[[3,4],5]]":                 143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]": 1384,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]": 3488,
	}

	for sfn, res := range cases {
		t.Run(sfn, func(t *testing.T) {
			snailFish, i := ParseSnailFish(sfn, 0, nil)
			if i != len(sfn) {
				t.Fatalf("ParseSnailFishNumber(%q) = %v, %v", sfn, snailFish, i)
			}
			snailFish.Reduce()
			assert.Equal(t, res, snailFish.Magnitude())
		})
	}
}