import heapq
from collections import namedtuple
from time import sleep

Coord = namedtuple('Coord', ['x', 'y'])
Step = namedtuple('Step', ['length', 'coord', 'prev'])

inp = 1352
seen_in_50 = set()


def solve(target):
    explore = []
    path = Step(0, Coord(1, 1), None)
    seen_coords = [path.coord]
    seen_in_50.add(path.coord)

    while path.coord != target:
        neighbours = find_neighbours(path.coord)
        if path.length < 50:
            seen_in_50.update(neighbours)
            # print("Path %s; Seen %s" % (str(path), [str(s) for s in seen_in_50]))
        for neighbour in neighbours:
            if neighbour not in seen_coords:
                seen_coords.append(neighbour)
                draw(seen_coords)
                sleep(0.1)
                heapq.heappush(explore, Step(path.length + 1, neighbour, path))
        path = heapq.heappop(explore)
    return path


def find_neighbours(c):
    return [d for d in [Coord(c.x - 1, c.y), Coord(c.x + 1, c.y), Coord(c.x, c.y - 1), Coord(c.x, c.y + 1)] if not is_wall(d.x, d.y) and d.x >= 0 and d.y >= 0]


def is_wall(x, y):
    nr = x * x + 3 * x + 2 * x * y + y + y * y
    nr += inp
    return bin(nr).count('1') % 2


def draw(seen, reset=True):
    if reset:
        print("\033[51A")
    maze = [[('#' if is_wall(x, y) else ('.' if Coord(x, y) not in seen else 'O')) for x in range(50)] for y in range(50)]
    print("\n".join([''.join(y) for y in maze]))


draw([], False)
path_to_target = solve(Coord(31, 39))
print("Day 13.1: %s" % path_to_target.length)
print("Day 13.2: %s" % len(seen_in_50))
