inp = 289326
# Each 'ring' around the center has 8 more positions than the previous one:
# n * 8
# ring 1 = 8 pos
# ring 2 = 16 pos (24)
# ring 3 = 24 pos (48)
# ring 4 = 32 pos (80)
# etc
ring = 1
nr = inp - 1  # make '0' based
while ring * 8 < nr:
    nr -= ring * 8
    ring += 1

prev_ring = ring - 1
pos_on_ring = inp - 1 - (8 * (prev_ring + prev_ring * (prev_ring - 1) // 2))
pos_on_side = pos_on_ring % (2 * ring)  # a side of a ring has 2 * ring positions
print(ring, pos_on_ring, pos_on_side, abs(ring - pos_on_side))
print("Day 3.1: ", abs(ring - pos_on_side) + ring)

x, y = ring, abs(ring - pos_on_side)
