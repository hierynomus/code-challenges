import numpy as np
from functools import reduce

with open('day10.in', 'r') as f:
    line = f.readline().strip()
    inp_1 = list(map(int, line.split(',')))
    inp_2 = [ord(c) for c in line] + [17, 31, 73, 47, 23]
knot = np.arange(256)


def tie_knot(knot, pos, skip, lengths):
    for l in lengths:
        knot = np.roll(knot, -pos)
        knot[:l] = knot[:l][::-1]
        knot = np.roll(knot, pos)
        pos = (pos + skip + l) % 256
        skip += 1
    return (knot, pos, skip)


knot_1, _, _ = tie_knot(knot, 0, 0, inp_1)
print("Day 10.1:", knot_1[0] * knot_1[1])

knot_2 = np.arange(256)
cur_pos = 0
skip = 0
for _ in range(64):
    knot_2, cur_pos, skip = tie_knot(knot_2, cur_pos, skip, inp_2)

dense_hash = ["%02x" % reduce(lambda x, y: x ^ y, knot_2[i:i + 16]) for i in range(0, 256, 16)]
print("Day 10.2:", "".join(dense_hash))
