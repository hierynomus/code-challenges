from functools import partial


def spin(x, p):
    return p[-x:] + p[:-x]


def exchange(x, y, p):
    p[x], p[y] = p[y], p[x]
    return p


def partner(x, y, p):
        i_x = p.index(x)
        i_y = p.index(y)
        p[i_x], p[i_y] = p[i_y], p[i_x]
        return p


def dance(programs, moves):
    p = programs[:]
    for move in moves:
        p = move(p)
    return p


programs = [x for x in "abcdefghijklmnop"]

with open("day16.in", 'r') as f:
    s = f.readline().strip().split(',')
    moves = []
    for m in s:
        if m.startswith('s'):
            moves.append(partial(spin, int(m[1:])))
        elif m.startswith('x'):
            x, y = map(int, m[1:].split('/'))
            moves.append(partial(exchange, x, y))
        elif m.startswith('p'):
            moves.append(partial(partner, m[1], m[3]))

d = dance(programs, moves)
print("Day 16.1:", "".join(d))

d = programs
seen = {''.join(d): 0}
i_seen = {0: ''.join(d)}

for i in range(1, 10000):
    d = dance(d, moves)
    s = ''.join(d)
    if s not in seen:
        seen[s] = i
        i_seen[i] = s
    else:
        period = i - seen[s]
        print("Day 16.2:", i_seen[1000000000 % period])
        break
