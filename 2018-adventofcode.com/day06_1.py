import fileinput 
from itertools import combinations
from collections import defaultdict

def neighbours(p):
    for x in [-1, 1]:
        yield (p[0] + x, p[1])
        yield (p[0], p[1] + x)

def grow(old_frontier, all_points):
    frontier = {}
    for p, v in old_frontier.items():
        if v != '.':
            for n in neighbours(p):
                if n not in all_points:
                    if n in frontier and frontier[n] != v:
                        frontier[n] = '.'
                    else:
                        frontier[n] = v
    return frontier

def value_hist(d):
    cts = defaultdict(int)
    for p, v in d.items():
        cts[v] += 1
    return cts

input = [tuple(map(int, l.split(','))) for l in fileinput.input()]

pts = {}
for i, p in enumerate(input):
    pts[p] = i

f = dict(pts)
for _ in range(400):
    f = grow(f, pts)
    hist = value_hist(f)
    f_keys = hist.keys()
    pts.update(f)


print(max([v for k, v in value_hist(pts).items() if v != '.' and k not in f_keys]))
