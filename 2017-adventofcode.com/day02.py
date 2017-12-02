from itertools import combinations

rows = []
with open('day02.in', 'r') as f:
    for l in f.readlines():
        rows.append(sorted([int(i) for i in l.strip().split()], reverse=True))

checksum = 0
evenly_sum = 0
for row in rows:
    checksum += row[0] - row[-1]
    evenly_sum += [c[0] // c[1] for c in combinations(row, 2) if c[0] % c[1] == 0][0]

print("Day 2.1: ", checksum)
print("Day 2.2: ", evenly_sum)
