import re

disc_r = re.compile("Disc #([0-9]+) has ([0-9]+) positions; at time=0, it is at position ([0-9]+).")
discs = []


def disc_as_lambda(size, dist, offset):
    return lambda t: ((t + dist + offset) % size) == 0


def solve(discs):
    s_0, d_0, o_0 = discs[0]
    lambdas = [disc_as_lambda(*disc) for disc in discs]
    t = s_0 - d_0 - o_0
    while True:
        if all(l(t) for l in lambdas):
            return t
        t += s_0


with open('day15.in', 'r') as f:
    for l in f:
        m = disc_r.match(l)
        discs.append((int(m.group(2)), int(m.group(1)), int(m.group(3))))

print("Day 15.1: %s" % solve(discs))
discs.append((11, len(discs) + 1, 0))
print("Day 15.2: %s" % solve(discs))


# #1: (15 + 17x) % 17 =
# #2: (5 + 7x) %  = y / 7
# #3: 14 + 19x = y /19
# #4: 1 + 5x = y / 5
# #5: 1 + 3x = y / 3
# #6: 2 + 13x = y / 13
