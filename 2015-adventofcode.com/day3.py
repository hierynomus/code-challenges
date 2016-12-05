from collections import namedtuple

Location = namedtuple('Location', ['x', 'y'])
lookup = {'^': Location(0, -1), 'v': Location(0, 1), '>': Location(1, 0), '<': Location(-1, 0)}


def move(direction, location, houses):
    new_loc = Location(x=location.x + lookup[direction].x, y=location.y + lookup[direction].y)
    houses.add(new_loc)
    return new_loc


def santa(directions):
    houses = set()
    santa = Location(0, 0)
    houses.add(santa)
    for direction in directions:
        santa = move(direction, santa, houses)

    return houses


def santa_and_robo(directions):
    houses = set()
    santa_turn = True
    santa = Location(0, 0)
    houses.add(santa)
    robo = Location(0, 0)
    for direction in directions:
        santa, robo = (move(direction, santa, houses), robo) if santa_turn else (santa, move(direction, robo, houses))
        santa_turn = not santa_turn

    return houses


with open('day3.in', 'r') as f:
    directions = f.read().strip()
    print("Day 3.1: %s" % str(len(santa(directions))))
    print("Day 3.2: %s" % str(len(santa_and_robo(directions))))
