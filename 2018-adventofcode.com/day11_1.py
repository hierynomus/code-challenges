import numpy as np
from operator import itemgetter

input = 9221

def pwr_lvl(y, x):
    x += 1
    y += 1
    rack_id = x + 10
    pwr = rack_id * y
    pwr += input
    pwr *= rack_id
    pwr = (pwr % 1000) // 100
    pwr -= 5
    return pwr

grid = np.fromfunction(pwr_lvl, (300, 300), dtype=int)

square_values = {}

for y in range(0, 297):
    for x in range(0, 297):
        square_values[(x + 1, y + 1)] = np.sum(grid[y: y + 3, x: x + 3])

print(','.join(map(str, max(square_values.items(), key=itemgetter(1))[0])))