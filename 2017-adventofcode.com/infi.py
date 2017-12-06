import sys

robots = []
hits = set()

with open(sys.argv[1], 'r') as f:
    line = f.readline().strip()
    start, moves = line.split('](', 1)
    for sp in start[1:].split(']['):
        robots.append(tuple(map(int, sp.split(','))))
    nr = len(robots)
    moves = [tuple(map(int, m.split(','))) for m in moves[:-1].split(')(')]

    for m in range(len(moves)):
        pos = robots[m % nr]
        move = moves[m]
        newpos = (pos[0] + move[0], pos[1] + move[1])
        robots[m % nr] = newpos
        if (m % nr) == nr - 1:
            if len(set(robots)) < nr:
                positions = robots[:]
                for r in set(robots):
                    positions.remove(r)
                for p in positions:
                    hits.add(p)
        # print(m, pos, "->", move, "=", robots)

print("Infi 1:", len(hits))
for y in range(max(hits, key=lambda x: x[1])[1] + 1):
    for x in range(max(hits, key=lambda x: x[0])[0] + 1):
        if (x, y) in hits:
            print('*', end='')
        else:
            print(' ', end='')
    print('')
