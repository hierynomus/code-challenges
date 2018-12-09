import fileinput
from collections import defaultdict

forward = defaultdict(list)
reverse = defaultdict(list)

for l in fileinput.input():
    s = l.split(' ')
    forward[s[1]].append(s[7])
    reverse[s[7]].append(s[1])

nxt_start = forward.keys() - reverse.keys()
next_nodes = set(nxt_start)
result = []

while next_nodes:
    nxt_start = min(next_nodes)
    result.append(nxt_start)
    next_nodes.remove(nxt_start)
    if nxt_start in forward:
        nxt = forward[nxt_start]
        next_nodes.update([n for n in nxt if n not in result and all([x in result for x in reverse[n]])])
        del forward[nxt_start]

print(''.join(result))
