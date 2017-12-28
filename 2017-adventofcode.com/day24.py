from collections import defaultdict, namedtuple
from heapq import heappop, heappush

mapping = defaultdict(list)
with open('day24.in', 'r') as f:
    for l in f:
        port = tuple(map(int, l.strip().split('/')))
        port = port[::-1] if port[0] > port[1] else port
        mapping[port[0]].append(port)
        mapping[port[1]].append(port)


def bridges(bridge):
    length, ports, strength, tail = bridge
    for port in [p for p in mapping[tail] if p not in ports]:
        new_tail = port[0] if tail == port[1] else port[1]
        new_bridge = (length + 1, ports | {port}, strength + port[0] + port[1], new_tail)
        yield from bridges(new_bridge)
    else:
        yield bridge


def solve():
    all_bridges = []
    for b in bridges((0, set(), 0, 0)):
        all_bridges.append(b)
    return all_bridges


all_bridges = solve()

print("Day 24.1:", max(all_bridges, key=lambda b: b[2])[2])
print("Day 24.2:", max(all_bridges, key=lambda b: (b[0], b[2]))[2])
