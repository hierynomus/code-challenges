import fileinput
import numpy as np

input = [l.rstrip('\n') for l in fileinput.input()]
fabric = np.full([1000, 1000], 0)
claims = {}
for l in input:
    m = l.split(' ')
    sx, sy = map(int, m[2][:-1].split(','))
    dx, dy = map(int, m[3].split('x'))
    claims[m[0][1:]] = (sx, sy, dx, dy)
    fabric[sx:dx+sx,sy:dy+sy] += 1

for claim, c in claims.items():
    if (fabric[c[0]:c[2]+c[0],c[1]:c[3]+c[1]] > 1).sum() == 0:
        print (claim)
        break

