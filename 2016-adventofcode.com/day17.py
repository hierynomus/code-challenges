from hashlib import md5
from heapq import heappush, heappop

inp = "awrkjxxr"
next_states = []
hasher = md5(inp.encode('utf-8'))
heappush(next_states, (0, (0, 0), [], hasher))
open_chars = ['b', 'c', 'd', 'e', 'f']
solutions = []

while next_states:
    s = heappop(next_states)
    p = s[1]
    if p == (3, 3):
        solutions.append(s)
        continue

    possibles = s[3].hexdigest()[:4]
    if possibles[0] in open_chars and p[1] > 0:
        h = s[3].copy()
        h.update('U'.encode('utf-8'))
        heappush(next_states, (s[0] + 1, (p[0], p[1] - 1), s[2] + ['U'], h))
    if possibles[1] in open_chars and p[1] < 3:
        h = s[3].copy()
        h.update('D'.encode('utf-8'))
        heappush(next_states, (s[0] + 1, (p[0], p[1] + 1), s[2] + ['D'], h))
    if possibles[2] in open_chars and p[0] > 0:
        h = s[3].copy()
        h.update('L'.encode('utf-8'))
        heappush(next_states, (s[0] + 1, (p[0] - 1, p[1]), s[2] + ['L'], h))
    if possibles[3] in open_chars and p[0] < 3:
        h = s[3].copy()
        h.update('R'.encode('utf-8'))
        heappush(next_states, (s[0] + 1, (p[0] + 1, p[1]), s[2] + ['R'], h))

print("Day 17.1: %s" % ''.join(solutions[0][2]))
print("Day 17.2: %s" % len(solutions[-1][2]))
