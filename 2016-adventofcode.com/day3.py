possible1 = 0
possible2 = 0
triangles = []


def possible(x, y, z):
    tr = sorted([x, y, z])
    return tr[0] + tr[1] > tr[2]


with open('day3.in', 'r') as f:
    triangles = [[int(c) for c in l.strip().split(' ') if c.strip()] for l in f]

for i in range(0, len(triangles), 3):
    for j in range(0, 3):
        if possible(triangles[i][j], triangles[i + 1][j], triangles[i + 2][j]):
            possible2 += 1
        if possible(*triangles[i + j]):
            possible1 += 1

print("Day 3.1: %s" % possible1)
print("Day 3.1: %s" % possible2)
