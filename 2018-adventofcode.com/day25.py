import fileinput
from itertools import chain

def manhattan(p1, p2):
    return sum([abs(a - b) for a, b in zip(p1, p2)])

points = [tuple(map(int, l.rstrip('\n').split(','))) for l in fileinput.input()]

constellations = []

for p in points:
    to_merge = []
    keep = []
    for c in constellations:
        for p2 in c:
            if manhattan(p, p2) <= 3:
                to_merge.append(c)
                break
        else:
            keep.append(c)
    if to_merge:
        constellations = keep
        n = list(chain(*to_merge))
        n.append(p)
        constellations.append(n)
    else:
        constellations.append([p])

print(len(constellations))
