import fileinput
from functools import lru_cache
from collections import defaultdict

def manhattan(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1]) + abs(p1[2] - p2[2])

@lru_cache(maxsize=None)
def collides(b1, b2):
    return manhattan(b1[0], b2[0]) <= b1[1] + b2[1]

bots = []

for l in fileinput.input():
    p, r = l.rstrip('\n').split(', ')
    r = int(r.split('=')[1])
    p = tuple(map(int, p.split('=')[1][1:-1].split(',')))
    bots.append((p, r))

collisions = defaultdict(set)
for b in bots:
    collisions[b] = set([c for c in bots if collides(b, c)])

max_colliding_bot = max(bots, key=lambda b: len(collisions[b]))

half_collision_size = len(collisions[max_colliding_bot])/2

collision = collisions[max_colliding_bot].copy()
remove = []
for c in collisions[max_colliding_bot]:
    if len(collisions[c]) < half_collision_size:
        remove.append(c)
    else:
        collision = collision.intersection(collisions[c])

max_bots = collision - set(remove)
print(len(max_bots))
max_dist = max([manhattan((0, 0, 0), b[0]) - b[1] for b in max_bots])
print(max_dist)