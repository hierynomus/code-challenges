from aoc2017 import knot_hash, hex_it

inp = "nbysizxe"


def flood(grid, y, x, seen):
    if (x, y) in seen:
        return
    elif grid[y][x] == '1':
        seen.add((x, y))
        if y > 0:
            flood(grid, y - 1, x, seen)
        if y < 127:
            flood(grid, y + 1, x, seen)
        if x > 0:
            flood(grid, y, x - 1, seen)
        if x < 127:
            flood(grid, y, x + 1, seen)


rows = []
s = 0
for i in range(128):
    k = knot_hash(hex_it("%s-%d" % (inp, i)))
    b = '{:0128b}'.format(int(k, 16))
    rows.append(b)
    s += sum(map(int, b))

print("Day 14.1:", s)

seen = set()
p = 0
group = 0
while p != 128 * 128:
    y, x = p // 128, p % 128
    if (x, y) not in seen and rows[y][x] == '1':
        flood(rows, y, x, seen)
        group += 1
    p += 1

print("Day 14.2:", group)
