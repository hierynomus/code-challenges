# From: https://pastebin.com/BMd61PUv
walks = {
    'Up': (0, -1),
    'Down': (0, 1),
    'Left': (-1, 0),
    'Right': (1, 0)
}


def manhattan(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])


A = []
B = []
x, y = 0, 0

with open('warmup.in', 'r') as f:
    for i in f.readline().split(', '):
        if i in walks:
            w = walks[i]
            x += w[0]
            y += w[1]
        elif i == 'A':
            A.append((x, y))
        elif i == 'B':
            B.append((x, y))

        else:
            # Done
            pass

max_dist = max(A + B, key=lambda x: manhattan((0, 0), x))
print("1: ", manhattan((0, 0), max_dist))

max_ab_dist = -1
for a in A:
    for b in B:
        d = manhattan(a, b)
        if d > max_ab_dist:
            max_ab_dist = d

print("2: ", max_ab_dist)
