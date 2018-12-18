import fileinput
import numpy as np

def new_flood(ground, x, y, ymax):
    frontiers = [(x, y)]
    while frontiers:
        x, y = frontiers.pop()
        # Falling
        while ground[y + 1, x] in ['.', '|'] and y <= ymax:
            ground[y, x] = '|'
            y += 1
        # Spreading
        if ground[y + 1, x] in ['#', '~']:
            left = x
            right = x
            standing = True
            while ground[y, left - 1] != '#':
                if ground[y + 1, left] in ['#', '~']:
                    left -= 1
                else:
                    standing = False
                    break
            while ground[y, right + 1] != '#':
                if ground[y + 1, right] in ['#', '~']:
                    right += 1
                else:
                    standing = False
                    break
            ground[y, left: right + 1] = '~' if standing else '|'
            if not standing and ground[y + 1, left] == '.':
                frontiers.append((left, y + 1))
            if not standing and ground[y + 1, right] == '.':
                frontiers.append((right, y + 1))
            if standing:
                frontiers.append((x, y - 1))


def show(ground):
    clay = np.where(ground == '#')
    ymin, ymax, xmin, xmax = min(clay[0]), max(clay[0]), min(clay[1]), max(clay[1])
    for l in ground[0 : ymax + 1, xmin-1: xmax + 1]:
        print(''.join(l))

ground = np.full((2000,2000), '.')

for line in fileinput.input():
    a, b = line.rstrip('\n').split(', ')
    bs, be = map(int, b[2:].split('..'))
    if a[0] == 'y':
        ground[int(a[2:]), bs: be + 1] = '#'
    else:
        ground[bs:be+1, int(a[2:])] = '#'

clay = np.where(ground == '#')
ymin, ymax = min(clay[0]), max(clay[0])
ground[0, 500] = '+'
new_flood(ground, 500, 1, ymax)
bounded = ground[ymin:ymax+1,]
print((bounded == '~').sum() + (bounded == '|').sum())
