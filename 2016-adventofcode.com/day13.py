import heapq
from collections import namedtuple

Coord = namedtuple('Coord', ['x', 'y'])
Step = namedtuple('Step', ['length', 'coord', 'prev'])
inp = 1352


def solve_1(target):
    explore = []
    path = Step(0, Coord(1, 1), None)
    seen_coords = [Coord(1, 1)]
    while path.coord != target:
        neighbours = find_neighbours(path.coord)
        for neighbour in neighbours:
            if neighbour not in seen_coords:
                seen_coords.append(neighbour)
                heapq.heappush(explore, Step(path.length + 1, neighbour, path))
        path = heapq.heappop(explore)
    return path


def find_neighbours(c):
    return [d for d in [Coord(c.x - 1, c.y), Coord(c.x + 1, c.y), Coord(c.x, c.y - 1), Coord(c.x, c.y + 1)] if not is_wall(d.x, d.y) and d.x > 0 and d.y > 0]


def is_wall(x, y):
    nr = x * x + 3 * x + 2 * x * y + y + y * y
    nr += inp
    return bin(nr).count('1') % 2


print("Day 13.1: %s" % solve_1(Coord(31, 39)).length)
