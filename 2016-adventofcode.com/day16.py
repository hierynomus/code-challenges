from bitarray import bitarray

inp = "01000100010010111"
data = inp


def checksum(cs):
    return [1 if a == b else 0 for a, b in zip(cs[::2], cs[1::2])]


def solve(initial_state, disk_size):
    arr = bitarray(initial_state)
    while len(arr) < disk_size:
        inv = arr.copy()
        inv.reverse()
        inv.invert()
        arr.append(0)
        arr.extend(inv)
    cs = checksum(arr[:disk_size])
    while len(cs) % 2 == 0:
        cs = checksum(cs)

    return ''.join(['0' if not c else '1' for c in cs])


print("Day 16.1: %s" % solve(inp, 272))
print("Day 16.1: %s" % solve(inp, 35651584))
