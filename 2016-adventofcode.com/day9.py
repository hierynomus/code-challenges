def decompress(line, isV2, length=0):
    if not line:
        return length
    start, rest = line.split("(", 1) if "(" in line else (line, None)
    length += len(start)
    if not rest:
        return length
    compression, rest = rest.split(")", 1)
    chars, repeat = map(int, compression.split("x"))
    length += (decompress(rest[:chars], isV2) * repeat) if isV2 else (len(rest[:chars]) * repeat)
    return decompress(rest[chars:], isV2, length)


with open('day9.in', 'r') as f:
    line = f.readline().strip()
    print("Day 9.1: %s" % decompress(line, False))
    print("Day 9.2: %s" % decompress(line, True))
