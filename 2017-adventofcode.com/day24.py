from collections import defaultdict, namedtuple
from heapq import heappop, heappush

Bridge = namedtuple('Bridge', ['weight', 'ports', 'tail'])

mapping = defaultdict(list)

with open('day24.in', 'r') as f:
    for l in f:
        port = tuple(map(int, l.strip().split('/')))
        mapping[port[0]].append(port)
        mapping[port[1]].append(port)


def solve():
    bridges = [Bridge(0, [(None, 0)], 0)]
    built_bridges = []
    while bridges:
        state = heappop(bridges)
        connectors = [c for c in mapping[state.tail] if c not in state.ports]
        for c in connectors:
            new_tail = c[0] if state.tail == c[1] else c[1]
            heappush(bridges, Bridge(state.weight + c[0] + c[1], state.ports + [c], new_tail))
        else:
            built_bridges.append(state)

    return built_bridges


all_bridges = solve()

print("Day 24.1:", max(all_bridges, key=lambda b: b.weight).weight)
print("Day 24.2:", max(all_bridges, key=lambda b: (len(b.ports), b.weight)).weight)
