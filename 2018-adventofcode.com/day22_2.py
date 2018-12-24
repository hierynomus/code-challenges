import fileinput
from functools import lru_cache
from heapq import heappop, heappush, heapify

depth = 0
target_x = target_y = 0

ROCKY = 0
WET = 1
NARROW = 2

NONE = 0
CLIMBING_GEAR = 1
TORCH = 2

kind_tool = {
    ROCKY: set([CLIMBING_GEAR, TORCH]),
    WET: set([NONE, CLIMBING_GEAR]),
    NARROW: set([NONE, TORCH])
}

for l in fileinput.input():
    if 'depth' in l:
        depth = int(l.split(':')[1].strip())
    elif 'target' in l:
        target_x, target_y = map(int, l.split(':')[1].strip().split(','))

def neighbours_4(x, y):
    if x > 0:
        yield (x - 1, y)
    if y > 0:
        yield (x, y - 1)
    if x < target_x + 100:
        yield (x + 1, y)
    if y < target_y + 100:
        yield (x, y + 1)

def manhattan(x, y):
    return abs(x - target_x) + abs(y - target_y)

@lru_cache(maxsize=None)
def determine_el(depth, x, y):
    if x == 0 and y == 0:
        return depth % 20183
    elif x == target_x and y == target_y:
        return depth % 20183
    elif x == 0:
        return (y * 48271 + depth) % 20183
    elif y == 0:
        return (x * 16807 + depth) % 20183
    else:
        return (determine_el(depth, x - 1, y) * determine_el(depth, x, y - 1) + depth) % 20183

queue = [(0, manhattan(0, 0), (0, 0), determine_el(depth, 0, 0) % 3, TORCH)]
min_time = 1000000
seen = {}
while queue:
    time, dist, point, kind, tool = heappop(queue)
    if (point, tool) in seen and seen[(point, tool)] <= time:
        continue
    seen[(point, tool)] = time
    if dist == 0:
        print(time)
        break
    for x, y in neighbours_4(*point):
        el = determine_el(depth, x, y)
        nk = el % 3
        mh = manhattan(x, y)
        t = tool
        dt = 1
        if nk != kind and not tool in kind_tool[nk]:
            t = list(kind_tool[kind] - set([tool]))[0]
            dt += 7

        heappush(queue, (time + dt, mh, (x, y), nk, t))
