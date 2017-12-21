from math import sqrt


def add_t(t_1, t_2):
    return tuple([t_1[x] + t_2[x] for x in range(len(t_1))])


pos = []
vel = []
acc = []

with open('day20.in', 'r') as f:
    for l in f:
        p, v, a = map(lambda x: tuple(map(int, x.split('=')[1][1:-1].split(','))), l.strip().split(', '))
        pos.append(p)
        vel.append(v)
        acc.append(a)

nr_particles = len(pos)

print("Day 20.1:", min(range(nr_particles), key=lambda i: sqrt(acc[i][0] ** 2 + acc[i][1] ** 2 + acc[i][2] ** 2)))

collided = [False] * nr_particles
seen_pos = {}
for i in range(100):
    for j in range(nr_particles):
        if collided[j]:
            continue

        v = vel[j]
        a = acc[j]
        p = pos[j]
        v = (v[0] + a[0], v[1] + a[1], v[2] + a[2])
        p = (p[0] + v[0], p[1] + v[1], p[2] + v[2])
        pos[j] = p
        vel[j] = v
        if p not in seen_pos:
            seen_pos[p] = j
        else:
            collided[seen_pos[p]] = True
            collided[j] = True

print("Day 20.2", sum([1 if not collided[x] else 0 for x in range(nr_particles)]))
