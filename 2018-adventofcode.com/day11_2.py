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
calc_grid = np.copy(grid)
max_val = (0, (1, 0, 0))
for i in range(2, 300):
    for y in range(0, 300 - i):
        for x in range(0, 300 - i):
            calc_grid[y, x] += (np.sum(grid[y + i - 1, x: x + i]) + np.sum(grid[y: y + i - 1, x + i - 1]))
            if calc_grid[y, x] > max_val[0]:
                max_val = (calc_grid[y, x], (i, x, y))
        

print("%d,%d,%d"%(max_val[1][1] + 1, max_val[1][2] + 1, max_val[1][0]))