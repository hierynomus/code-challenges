coord_x = 0
coord_y = 0
day2 = None


def walk(start_x, start_y, walk, direction, locations):
    global day2
    walks = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    (new_x, new_y) = (start_x, start_y)
    for step in range(walk):
        (new_x, new_y) = (new_x + walks[direction][0], new_y + walks[direction][1])

        if not day2 and (new_x, new_y) in locations:
            print("Found day2!")
            day2 = (new_x, new_y)

        locations.add((new_x, new_y))

    return (new_x, new_y)


with open('day1.in', 'r') as f:
    direction = 0  # 0=North, 1=East, 2=South, 3=West
    locations = set()
    for i in f.readline().split(', '):
        turn = 1 if i[0] == 'R' else - 1
        steps = int(i[1:])
        direction = (direction + turn) % 4
        (coord_x, coord_y) = walk(coord_x, coord_y, steps, direction, locations)

print("Day 1.1: %d" % (abs(coord_x) + abs(coord_y)))
print("Day 1.2: %d" % (abs(day2[0]) + abs(day2[1])))
