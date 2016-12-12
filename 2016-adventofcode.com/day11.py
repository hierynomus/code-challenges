import re
import heapq
from collections import namedtuple
import itertools

State = namedtuple('State', ['elevator', 'locations'])
Element = namedtuple('Element', ['chip', 'generator'])

chip_re = re.compile("([a-z]+)-compatible microchip")
generator_re = re.compile("([a-z]+) generator")


def is_solved(state):
    if state.elevator < 3:
        return False
    for element in state.locations:
        if element.chip != 3 or element.generator != 3:
            return False
    return True


def is_valid(state):
    for e in state.locations:
        if e.chip == e.generator:
            continue
        for f in state.locations:
            if e.chip == f.generator:
                return False
    return True


def next_states(state):
    ns = []
    for i in range(len(state.locations)):
        if state.locations[i].chip == state.elevator:
            if state.elevator > 0:
                locs = state.locations[:]
                locs[i] = Element(locs[i].chip - 1, locs[i].generator)
                ns.append(State(state.elevator - 1, sorted(locs)))
            if state.elevator < 3:
                locs = state.locations[:]
                locs[i] = Element(locs[i].chip + 1, locs[i].generator)
                ns.append(State(state.elevator + 1, sorted(locs)))
            if state.locations[i].generator == state.elevator:
                if state.elevator > 0:
                    locs = state.locations[:]
                    locs[i] = Element(locs[i].chip - 1, locs[i].generator - 1)
                    ns.append(State(state.elevator - 1, sorted(locs)))
                if state.elevator < 3:
                    locs = state.locations[:]
                    locs[i] = Element(locs[i].chip + 1, locs[i].generator + 1)
                    ns.append(State(state.elevator + 1, sorted(locs)))
            for j in range(i + 1, len(state.locations)):
                if state.locations[j].chip == state.elevator:
                    if state.elevator > 0:
                        locs = state.locations[:]
                        locs[i] = Element(locs[i].chip - 1, locs[i].generator)
                        locs[j] = Element(locs[j].chip - 1, locs[j].generator)
                        ns.append(State(state.elevator - 1, sorted(locs)))
                    if state.elevator < 3:
                        locs = state.locations[:]
                        locs[i] = Element(locs[i].chip + 1, locs[i].generator)
                        locs[j] = Element(locs[j].chip + 1, locs[j].generator)
                        ns.append(State(state.elevator + 1, sorted(locs)))
        if state.locations[i].generator == state.elevator:
            if state.elevator > 0:
                locs = state.locations[:]
                locs[i] = Element(locs[i].chip, locs[i].generator - 1)
                ns.append(State(state.elevator - 1, sorted(locs)))
            if state.elevator < 3:
                locs = state.locations[:]
                locs[i] = Element(locs[i].chip, locs[i].generator + 1)
                ns.append(State(state.elevator + 1, sorted(locs)))
            for j in range(i + 1, len(state.locations)):
                if state.locations[j].generator == state.elevator:
                    if state.elevator > 0:
                        locs = state.locations[:]
                        locs[i] = Element(locs[i].chip, locs[i].generator - 1)
                        locs[j] = Element(locs[j].chip, locs[j].generator - 1)
                        ns.append(State(state.elevator - 1, sorted(locs)))
                    if state.elevator < 3:
                        locs = state.locations[:]
                        locs[i] = Element(locs[i].chip, locs[i].generator + 1)
                        locs[j] = Element(locs[j].chip, locs[j].generator + 1)
                        ns.append(State(state.elevator + 1, sorted(locs)))

    return [s for s in ns if is_valid(s)]


def solve(elements):
    state = (0, State(0, list(elements.values())))
    seen = []
    to_view = []
    while not is_solved(state[1]):
        seen.append(state[1])
        for ns in next_states(state[1]):
            if ns not in seen:
                seen.append(ns)
                heapq.heappush(to_view, (state[0] + 1, ns))
        state = heapq.heappop(to_view)
    return state


start_elements = {}

with open('day11.in', 'r') as f:
    lines = f.readlines()
    for i in range(len(lines)):
        chips = chip_re.findall(lines[i])
        for chip in chips:
            if chip not in start_elements:
                start_elements[chip] = Element(i, None)
            else:
                start_elements[chip] = Element(i, start_elements[chip].generator)
        generators = generator_re.findall(lines[i])
        for generator in generators:
            if generator not in start_elements:
                start_elements[generator] = Element(None, i)
            else:
                start_elements[generator] = Element(start_elements[generator].chip, i)

print("Day 11.1: %s" % solve(start_elements)[0])
start_elements['elerium'] = Element(0, 0)
start_elements['dilithium'] = Element(0, 0)
print("Day 11.2: %s" % solve(start_elements)[0])
