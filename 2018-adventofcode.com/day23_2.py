import fileinput
from collections import Counter

def manhattan(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1]) + abs(p1[2] - p2[2])

bots = []

for l in fileinput.input():
    p, r = l.rstrip('\n').split(', ')
    r = int(r.split('=')[1])
    p = tuple(map(int, p.split('=')[1][1:-1].split(',')))
    bots.append((p, r))

def full_axis(bots, axis):
    a = []  # (x, end?, bot)
    for b in bots:
        a.append((b[0][axis]-b[1], 0, b))
        a.append((b[0][axis]+b[1], 1, b))

    return sorted(a)

all_axis = [full_axis(bots, a) for a in range(0, 3)]

def max_overlap(axis):
    cur_bots = []
    max_bots = []
    for s in axis:
        if s[1] == 0:
            cur_bots.append(s[2])
            if len(cur_bots) > len(max_bots):
                max_bots = cur_bots[:]
        else:
            cur_bots.remove(s[2])
    return max_bots

max_bots = bots[:]
for axis in [2, 0, 1]:
    mb = max_overlap(all_axis[axis])
    max_bots = [b for b in max_bots if b in mb]

print(len(max_bots))
max_dist = max([manhattan((0, 0, 0), b[0]) - b[1] for b in max_bots])
print(max_dist)

# for b in sorted(max_bots, key=lambda b: b[0]):
#     print("[%d, %d, %d]"%(b[0][0], b[0][1], b[0][2]))