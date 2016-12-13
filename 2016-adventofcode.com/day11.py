import re
import heapq
from collections import namedtuple

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
    if state.elevator < 0 or state.elevator > 3:
        return False
    for e in state.locations:
        if e.chip == e.generator:
            continue
        for f in state.locations:
            if e.chip == f.generator:
                return False
    return True


def update_chip(el, delta):
    return Element(el.chip + delta, el.generator)


def update_generator(el, delta):
    return Element(el.chip, el.generator + delta)


def move(state, coll, i, f_i, j=None, f_j=None):
    for d in [-1, 1]:
        locs = state.locations[:]
        locs[i] = f_i(locs[i], d)
        if j:
            locs[j] = f_j(locs[j], d)
        s = State(state.elevator + d, sorted(locs))
        if is_valid(s):
            coll.append(s)


def next_states(state):
    ns = []
    for i in range(len(state.locations)):
        if state.locations[i].chip == state.elevator:
            move(state, ns, i, update_chip)
            if state.locations[i].generator == state.elevator:
                move(state, ns, i, update_chip, i, update_generator)
            for j in range(i + 1, len(state.locations)):
                if state.locations[j].chip == state.elevator:
                    move(state, ns, i, update_chip, j, update_chip)
        if state.locations[i].generator == state.elevator:
            move(state, ns, i, update_generator)
            for j in range(i + 1, len(state.locations)):
                if state.locations[j].generator == state.elevator:
                    move(state, ns, i, update_generator, j, update_generator)

    return ns


def solve(elements):
    state = (0, State(0, sorted(list(elements.values()))))
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
