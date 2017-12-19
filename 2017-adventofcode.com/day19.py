directions = {
    's': (0, 1),
    'n': (0, -1),
    'w': (-1, 0),
    'e': (1, 0)
}

grid = []
with open('day19.in', 'r') as f:
    for l in f:
        grid.append([c for c in l[1:-1]])

y_0 = 0
x_0 = grid[y_0].index('|')
d = 's'
x, y = x_0, y_0
seen = []
count = 0
while True:
    n_x, n_y = x + directions[d][0], y + directions[d][1]
    count += 1
    c = grid[n_y][n_x]
    if c == '+' and d in ['n', 's']:
        print(d, n_x, n_y)
        d = 'e' if n_x + 1 < len(grid[n_y]) and grid[n_y][n_x + 1] == '-' else 'w'
    elif c == '+' and d in ['e', 'w']:
        print(d, n_x, n_y)
        d = 's' if n_y + 1 < len(grid) and grid[n_y + 1][n_x] == '|' else 'n'
    elif c == '|' or c == '-':
        pass
    elif c != ' ':
        seen.append(c)
        print(seen)
    else:
        break
    x, y = n_x, n_y

print("Day 19.1:", ''.join(seen))
print("Day 19.2:", count)

