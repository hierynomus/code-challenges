inp = (".#.", "..#", "###")


def transpose(matrix):
    return tuple(map(lambda t: ''.join(t), zip(*matrix[::-1])))


def all_rotations(inp):
    for i in [inp, inp[::-1], tuple([x[::-1] for x in inp])]:
        yield i
        for _ in range(3):
            i = transpose(i)
            yield i


def subgrid(grid, x, y, l):
    result = []
    for yi in range(y, y + l):
        result.append(grid[yi][x:x + l])
    return tuple(result)


def expand(grid):
    subgrid_size = 2 if len(grid) % 2 == 0 else 3

    new_grid = []
    for y in range(0, len(grid), subgrid_size):
        new_subgrids = [rules[subgrid(grid, x, y, subgrid_size)] for x in range(0, len(grid), subgrid_size)]
        for row in zip(*new_subgrids):
            new_grid.append(''.join(row))

    return tuple(new_grid)


rules = {}
with open('day21.in', 'r') as f:
    for l in f:
        f, t = l.strip().split(' => ')
        t = tuple(t.split('/'))
        for i in all_rotations(tuple(f.split('/'))):
            rules[i] = t


grid = inp
for _ in range(5):
    grid = expand(grid)

print("Day 21.1:", sum(map(lambda s: s.count('#'), grid)))

for _ in range(13):
    grid = expand(grid)

print("Day 21.2:", sum(map(lambda s: s.count('#'), grid)))
