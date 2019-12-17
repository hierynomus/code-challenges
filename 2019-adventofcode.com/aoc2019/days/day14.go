package days

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Day14 struct{}

type Chemical string

type Element struct {
	Amount int
	Name   Chemical
}

type Reaction struct {
	Input  []Element
	Output Element
}

func (r Reaction) String() string {
	return fmt.Sprintf("%v => %v", r.Input, r.Output)
}

type NanoFactory struct {
	Reactions map[Chemical]Reaction
	Surplus   map[Chemical]int
}

func (f *NanoFactory) React(needed Element) (Ore int) {
	NeededElements := map[Chemical]int{
		needed.Name: needed.Amount,
	}

	for len(NeededElements) > 0 {
		for elem, amount := range NeededElements {
			delete(NeededElements, elem)
			if elem == Chemical("ORE") {
				Ore += amount
				continue
			} else if amount == 0 {
				continue
			}
			reaction := f.Reactions[elem]
			times := 1
			surplusAmount := 0
			if a, ok := f.Surplus[elem]; ok {
				surplusAmount = a
			}
			if (amount-surplusAmount)%reaction.Output.Amount == 0 {
				times = (amount - surplusAmount) / reaction.Output.Amount
			} else if amount-surplusAmount > reaction.Output.Amount {
				times = 1 + (amount-surplusAmount)/reaction.Output.Amount
			}
			// fmt.Printf("%d :  %v\n", times, reaction)
			// Add to stock
			newSurplus := surplusAmount + times*reaction.Output.Amount - amount
			f.Surplus[elem] = newSurplus
			for _, need := range reaction.Input {
				amountNeeded := times * need.Amount
				if a, ok := f.Surplus[need.Name]; ok {
					if amountNeeded > a {
						amountNeeded -= a
						delete(f.Surplus, need.Name)
					} else {
						amountNeeded = 0
						f.Surplus[need.Name] -= amountNeeded
					}
				}
				if a, ok := NeededElements[need.Name]; ok {
					NeededElements[need.Name] = a + amountNeeded
				} else {
					NeededElements[need.Name] = amountNeeded
				}
			}
		}
	}
	return Ore
}

func parseElement(s string) Element {
	x := strings.Split(s, " ")
	i, err := strconv.Atoi(x[0])
	if err != nil {
		panic(err)
	}
	return Element{Amount: i, Name: Chemical(x[1])}
}

func ParseReactions(scanner *bufio.Scanner) []Reaction {
	reactions := []Reaction{}
	for scanner.Scan() {
		reactionString := strings.Split(scanner.Text(), " => ")
		input := []Element{}
		for _, e := range strings.Split(reactionString[0], ", ") {
			input = append(input, parseElement(e))
		}
		reaction := Reaction{
			Input:  input,
			Output: parseElement(reactionString[1]),
		}
		reactions = append(reactions, reaction)
	}
	return reactions
}

func (d *Day14) Solve(scanner *bufio.Scanner) (string, string) {
	OutputReactions := map[Chemical]Reaction{}
	for _, r := range ParseReactions(scanner) {
		if _, ok := OutputReactions[r.Output.Name]; ok {
			panic(fmt.Errorf("Duplicate output"))
		}
		OutputReactions[r.Output.Name] = r
	}

	factory := &NanoFactory{
		Reactions: OutputReactions,
		Surplus:   map[Chemical]int{},
	}

	OreNeeded := factory.React(Element{1, Chemical("FUEL")})

	factory = &NanoFactory{
		Reactions: OutputReactions,
		Surplus:   map[Chemical]int{},
	}

	OreInStock := 1000000000000
	FuelProduced := 0
	for OreInStock > OreNeeded {
		Fuel := OreInStock / OreNeeded
		OreInStock -= factory.React(Element{Fuel, Chemical("FUEL")})
		FuelProduced += Fuel
		// fmt.Printf("Produced %d Fuel for a total of %d (OreInStock: %d)\n", Fuel, FuelProduced, OreInStock)
	}

	return strconv.Itoa(OreNeeded), strconv.Itoa(FuelProduced)
}
