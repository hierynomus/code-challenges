from itertools import permutations

places = set()
dist = dict()

def total_distance(points):
    return sum([dist[(point, points[index + 1])] for index, point in enumerate(points[:-1])])

def travelling_salesman(points):
    return min([perm for perm in permutations(points)], key=total_distance)

def max_travelling_salesman(points):
    return max([perm for perm in permutations(points)], key=total_distance)


with open('day9.in', 'r') as f:
    for l in f:
        parts = l.strip().split(' ')
        f, t, d = parts[0], parts[2], int(parts[4])
        places.add(f)
        places.add(t)
        dist[(f, t)] = d
        dist[(t, f)] = d

visit = travelling_salesman(list(places))
d = total_distance(visit)
print("1: %s" % str(d))

visit = max_travelling_salesman(list(places))
d = total_distance(visit)
print("2: %s" % str(d))
