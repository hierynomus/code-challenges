import fileinput

def manhattan(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1]) + abs(p1[2] - p2[2])

bots = []

for l in fileinput.input():
    p, r = l.rstrip('\n').split(', ')
    r = int(r.split('=')[1])
    p = tuple(map(int, p.split('=')[1][1:-1].split(',')))
    bots.append((p, r))

max_bot = max(bots, key=lambda b: b[1])
print(sum([1 if manhattan(max_bot[0], b[0]) <= max_bot[1] else 0 for b in bots]))