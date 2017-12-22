infected = set()

dirs = {
    0: lambda x, y: (x, y - 1),
    1: lambda x, y: (x + 1, y),
    2: lambda x, y: (x, y + 1),
    3: lambda x, y: (x - 1, y)
}

start_x, start_y = 0, 0
with open('day22.in', 'r') as f:
    y = 0
    for l in f:
        x = 0
        for c in l.strip():
            if c == '#':
                infected.add((x, y))
            x += 1
        y += 1
    start_x, start_y = x // 2, y // 2

mutate = infected.copy()
x, y = start_x, start_y
d = 0  # Up
infectious = 0
for _ in range(10000):
    if (x, y) in mutate:
        d = (d + 1) % 4
        mutate.remove((x, y))
    else:
        d = (d - 1) % 4
        infectious += 1
        mutate.add((x, y))
    x, y = dirs[d](x, y)

print("Day 22.1:", infectious)

evolved = {}
for i in infected:
    evolved[i] = '#'

x, y = start_x, start_y
infectious = 0
d = 0
for _ in range(10000000):
    if not (x, y) in evolved:
        evolved[(x, y)] = 'W'
        d = (d - 1) % 4
    else:
        state = evolved[(x, y)]
        if state == 'W':
            evolved[(x, y)] = '#'
            infectious += 1
        elif state == '#':
            evolved[(x, y)] = 'F'
            d = (d + 1) % 4
        else:
            del evolved[(x, y)]
            d = (d + 2) % 4
    x, y = dirs[d](x, y)

print("Day 22.2:", infectious)
