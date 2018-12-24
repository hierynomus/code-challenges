import numpy as np
import fileinput

depth = 0
target_x = target_y = 0

for l in fileinput.input():
    if 'depth' in l:
        depth = int(l.split(':')[1].strip())
    elif 'target' in l:
        target_x, target_y = map(int, l.split(':')[1].strip().split(','))

def determine_gi(cave_el, x, y):
    if x == 0 and y == 0:
        return 0
    elif x == target_x and y == target_y:
        return 0
    elif x == 0:
        return y * 48271
    elif y == 0:
        return x * 16807
    else:
        return cave_el[y, x - 1] * cave_el[y - 1, x]


cave_gi = np.zeros((target_y + 1, target_y + 1))
cave_el = np.zeros((target_y + 1, target_y + 1))
cave = np.zeros((target_y + 1, target_y + 1))

for x in range(target_x + 1):
    for y in range(target_y + 1):
        gi = determine_gi(cave_el, x, y)
        el = (gi + depth) % 20183
        cave_gi[y, x] = gi
        cave_el[y, x] = el
        cave[y, x] = el % 3

print(cave.sum())