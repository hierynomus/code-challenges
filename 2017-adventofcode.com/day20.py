from math import sqrt

pos = []
vel = []
acc = []

with open('day20.in', 'r') as f:
    for l in f:
        p, v, a = map(lambda x: tuple(map(int, x.split('=')[1][1:-1].split(','))), l.strip().split(', '))
        pos.append(p)
        vel.append(v)
        acc.append(a)

print("Day 20.1:", min(range(len(acc)), key=lambda i: sqrt(acc[i][0] ** 2 + acc[i][1] ** 2 + acc[i][2] ** 2)))

