from functools import reduce


def tie_knot(knot, pos, skip, lengths):
    for l in lengths:
        knot = knot[pos:] + knot[:pos]
        knot[:l] = knot[:l][::-1]
        knot = knot[256 - pos:] + knot[:256 - pos]
        pos = (pos + skip + l) % 256
        skip += 1
    return (knot, pos, skip)


def knot_hash(i):
    i += [17, 31, 73, 47, 23]
    knot = list(range(256))
    pos = 0
    skip = 0
    for _ in range(64):
        knot, pos, skip = tie_knot(knot, pos, skip, i)

    hexed = "".join(["%02x" % reduce(lambda x, y: x ^ y, z) for z in [knot[j:j + 16] for j in range(0, 256, 16)]])
    return hexed


def hex_it(s):
    return [ord(c) for c in s]
