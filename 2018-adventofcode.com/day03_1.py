import fileinput
import numpy as np

input = [l.rstrip('\n') for l in fileinput.input()]
fabric = np.full([1000, 1000], 0)

for l in input:
    m = l.split(' ')
    sx, sy = map(int, m[2][:-1].split(','))
    dx, dy = map(int, m[3].split('x'))
    fabric[sx:dx+sx,sy:dy+sy] += 1

print((fabric > 1).sum())