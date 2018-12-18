import fileinput
import numpy as np

LIMIT = 50

def show(woods):
    for i in woods:
        print(''.join(i))

def neighbours_8(x, y):
    for dy in [-1, 0, 1]:
        for dx in [-1, 0, 1]:
            if dx == 0 and dy == 0:
                continue
            nx, ny = x + dx, y + dy
            if nx < 0 or ny < 0 or nx >= LIMIT or ny >= LIMIT:
                continue
            yield (nx, ny)

def evolve(woods):
    new_woods = np.full((LIMIT, LIMIT), '.')
    for y in range(LIMIT):
        for x in range(LIMIT):
            c = woods[y, x]
            if c == '.':
                trees = 0
                for n in neighbours_8(x, y):
                    trees += 1 if woods[n[1], n[0]] == '|' else 0
                    if trees >= 3:
                        new_woods[y, x] = '|'
                        break
                else:
                    new_woods[y, x] = c
            elif c == '|':
                lumber = 0
                for n in neighbours_8(x, y):
                    lumber += 1 if woods[n[1], n[0]] == '#' else 0
                    if lumber >= 3:
                        new_woods[y, x] = '#'
                        break
                else:
                    new_woods[y, x] = c
            elif c == '#':
                trees, lumber = 0, 0
                for n in neighbours_8(x, y):
                    v = woods[n[1], n[0]]
                    trees += 1 if v == '|' else 0
                    lumber += 1 if v == '#' else 0
                    if trees >= 1 and lumber >= 1:
                        new_woods[y, x] = '#'
                        break
                else:
                    new_woods[y, x] = '.'
    return new_woods


woods = np.full((LIMIT, LIMIT), '.')
y = 0
for l in fileinput.input():
    for i, c in enumerate(l.rstrip('\n')):
        woods[y, i] = c
    y += 1

show(woods)

for t in range(10):
    woods = evolve(woods)

print(('|' == woods).sum() * ('#' == woods).sum())
