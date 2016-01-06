size = 100
generations = 100

def empty_grid():
    return [['.' for x in range(size + 2)] for y in range(size + 2)]

def count_neighbours(grid, y, x):
    total_lit = 0
    for y2 in range(y - 1, y + 2):
        for x2 in range (x - 1, x + 2):
            lamp = grid[y2][x2]
            # print("(%s, %s) is %s" % (y2, x2, lamp))
            if x2 == x and y2 == y:
                continue
            if lamp == '#':
                total_lit += 1
    return total_lit

    # return sum([1 if grid[y2][x2] == '#' and y != y2 and x != x2 else 0 for x2 in range(x - 1, x + 2) for y2 in range(y - 1, y + 2)])

def print_grid(grid):
    print('\n'.join([''.join(grid[y]) for y in range(size + 2)]))

def sum_grid(grid):
    return sum([1 if grid[y2][x2] == '#' else 0 for x2 in range(size + 2) for y2 in range(size + 2)])

def iterate(grid):
    new_grid = empty_grid()
    for y in range(1, size + 1):
        for x in range(1, size + 1):
            neighbours = count_neighbours(grid, y, x)
            # print("(%s, %s) has %s lit neighbours" % (y, x, neighbours))
            if grid[y][x] == '#' and (neighbours == 2 or neighbours == 3):
                new_grid[y][x] = '#'
            elif grid[y][x] == '.' and neighbours == 3:
                new_grid[y][x] = '#'

    return new_grid

def iterate_stuck(grid):
    new_grid = empty_grid()
    for y in range(1, size + 1):
        for x in range(1, size + 1):
            neighbours = count_neighbours(grid, y, x)
            # print("(%s, %s) has %s lit neighbours" % (y, x, neighbours))
            if grid[y][x] == '#' and (neighbours == 2 or neighbours == 3):
                new_grid[y][x] = '#'
            elif grid[y][x] == '.' and neighbours == 3:
                new_grid[y][x] = '#'
            elif x in [1, size] and y in [1, size]:
                new_grid[y][x] = '#'

    return new_grid

grid = empty_grid()

with open('day18.in', 'r') as f:
    y = 0
    for l in f:
        y += 1
        x = 0
        for c in l.strip():
            x += 1
            if c == '#':
                grid[y][x] = '#'

g2 = grid
for i in range(generations):
    g2 = iterate(g2)

print("1: %s" % sum_grid(g2))

# Make the lights stuck
g3 = grid
g3[1][1] = '#'
g3[1][size] = '#'
g3[size][1] = '#'
g3[size][size] = '#'
for i in range(generations):
    g3 = iterate_stuck(g3)

print("2: %s" % sum_grid(g3))

