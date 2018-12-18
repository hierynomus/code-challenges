import fileinput
import numpy as np

def manhattan(x, y, p):
    return abs(x - p[0]) + abs(y - p[1])

input = [tuple(map(int, l.split(','))) for l in fileinput.input()]
grid = np.fromfunction(lambda x, y: sum([manhattan(x, y, p) for p in input]), (320, 800))
print((grid < 10000).sum())