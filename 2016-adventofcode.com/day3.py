possible1 = 0
possible2 = 0
part2 = []


def possible(x, y, z):
    tr = sorted([x, y, z])
    return tr[0] + tr[1] > tr[2]


with open('day3.in', 'r') as f:
    for l in f:
        arr = [int(c) for c in l.strip().split(' ') if c.strip()]
        part2.append(arr)
        if possible(*arr):
            possible1 += 1

for i in range(0, len(part2), 3):
    for j in range(0, 3):
        if possible(part2[i][j], part2[i + 1][j], part2[i + 2][j]):
            possible2 += 1

print("Day 3.1: %s" % possible1)
print("Day 3.1: %s" % possible2)
