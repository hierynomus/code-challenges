from aoc2017 import tie_knot, knot_hash, hex_it

with open('day10.in', 'r') as f:
    line = f.readline().strip()
    inp_1 = list(map(int, line.split(',')))
    inp_2 = hex_it(line)
knot = list(range(256))

knot_1, _, _ = tie_knot(knot, 0, 0, inp_1)
print("Day 10.1:", knot_1[0] * knot_1[1])
print("Day 10.2:", "".join(knot_hash(inp_2)))
