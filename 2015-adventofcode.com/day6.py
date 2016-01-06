
grid = [[False for x in range(1000)] for y in range(1000)]
grid_2 = [[0 for x in range(1000)] for y in range(1000)]

def switch(x_left, x_right, y_top, y_bottom, onOff):
    for x in range(x_left, x_right + 1):
        for y in range(y_top, y_bottom + 1):
            grid[y][x] = onOff
            if onOff:
                grid_2[y][x] += 1
            elif grid_2[y][x] > 0:
                grid_2[y][x] -= 1

def toggle(x_left, x_right, y_top, y_bottom):
    for x in range(x_left, x_right + 1):
        for y in range(y_top, y_bottom + 1):
            grid[y][x] = not grid[y][x]
            grid_2[y][x] += 2


with open('day6.in', 'r') as f:
    for s in f:
        parts = s.split(' ')
        if s.startswith('turn'):
            direction = "on" in s
            x_1, y_1 = [int(p) for p in parts[2].split(',')]
            x_2, y_2 = [int(p) for p in parts[4].split(',')]
            switch(x_1, x_2, y_1, y_2, direction)
        elif s.startswith('toggle'):
            x_1, y_1 = [int(p) for p in parts[1].split(',')]
            x_2, y_2 = [int(p) for p in parts[3].split(',')]
            toggle(x_1, x_2, y_1, y_2)

    turned_on = 0
    brightness = 0
    for x in range(1000):
        for y in range(1000):
            if grid[y][x]:
                turned_on += 1
            brightness += grid_2[y][x]

    print("1: %s" % str(turned_on))
    print("2: %s" % str(brightness))
