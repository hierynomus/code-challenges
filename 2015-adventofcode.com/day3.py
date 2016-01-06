from collections import namedtuple

Location = namedtuple('Location', ['x', 'y'])

def move(direction, location, houses):
    if '^' == direction:
        location = location._replace(y=location.y - 1)
    elif 'v' == direction:
        location = location._replace(y=location.y + 1)
    elif '<' == direction:
        location = location._replace(x=location.x - 1)
    elif '>' == direction:
        location = location._replace(x=location.x + 1)

    houses.add(location)
    return location

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
        if santa_turn:
            santa = move(direction, santa, houses)
        else:
            robo = move(direction, robo, houses)

        santa_turn = not santa_turn

    return houses

with open('day3.in', 'r') as f:
    directions = f.read()
    print("1: %s" % str(len(santa(directions))))
    print("2: %s" % str(len(santa_and_robo(directions))))



