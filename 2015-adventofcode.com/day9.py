from itertools import permutations

places = set()
dist = dict()


def total_distance(points):
    return sum([dist[(point, points[index + 1])] for index, point in enumerate(points[:-1])])


def travelling_salesman(points, func):
    return func([perm for perm in permutations(points)], key=total_distance)


with open('day9.in', 'r') as f:
    for l in f:
        parts = l.strip().split(' ')
        f, t, d = parts[0], parts[2], int(parts[4])
        places.add(f)
        places.add(t)
        dist[(f, t)] = d
        dist[(t, f)] = d

visit = travelling_salesman(list(places), min)
print("Day 9.1: %s" % total_distance(visit))

visit = travelling_salesman(list(places), max)
print("Day 9.2: %s" % total_distance(visit))
