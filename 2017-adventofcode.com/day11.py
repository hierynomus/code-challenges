moves = {
    'ne': lambda x, y, z: (x + 1, y, z - 1),
    'se': lambda x, y, z: (x + 1, y - 1, z),
    's': lambda x, y, z: (x, y - 1, z + 1),
    'sw': lambda x, y, z: (x - 1, y, z + 1),
    'nw': lambda x, y, z: (x - 1, y + 1, z),
    'n': lambda x, y, z: (x, y + 1, z - 1)
}


def grid_manhattan(x, y, z):
    return (abs(x) + abs(y) + abs(z)) / 2


with open('day11.in', 'r') as f:
    inp = f.readline().strip().split(',')

x = y = z = 0
max_d = 0
for m in inp:
    x, y, z = moves[m](x, y, z)
    max_d = max(max_d, grid_manhattan(x, y, z))

print("Day 11.1:", grid_manhattan(x, y, z))
print("Day 11.2:", max_d)
