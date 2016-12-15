import hashlib
import re

triplet_r = re.compile("([a-z0-9])\\1\\1")
inp = "jlmsuwbz".encode('UTF-8')
# inp = "abc".encode('UTF-8')
hasher = hashlib.md5(inp)


def calc_key(idx, hashes):
    if idx not in hashes:
        cp = hasher.copy()
        cp.update(str(idx).encode('UTF-8'))
        hashes[idx] = cp.hexdigest()
    return hashes[idx]


def calc_stretched_key(idx, hashes):
    if idx in hashes:
        return hashes[idx]
    stretch = calc_key(idx, hashes)
    for _ in range(2016):
        stretch = hashlib.md5(stretch.encode('UTF-8')).hexdigest()
    hashes[idx] = stretch
    return stretch


def find_triplet(checksum):
    count = 1
    ch = checksum[0]
    for c in checksum[1:]:
        if c == ch:
            count += 1
        else:
            ch = c
            count = 1
        if count == 3:
            return ch


def has_quintet(checksum, ch):
    return ch * 5 in checksum


def solve(key_func):
    keys = {}
    idx = 0
    hash_cache = {}

    while len(keys) < 64:
        key = key_func(idx, hash_cache)
        trip = find_triplet(key)
        if trip:
            # print("Found potential key at %s: %s" % (idx, (trip, key)))
            for i in range(idx + 1, idx + 1001):
                if has_quintet(key_func(i, hash_cache), trip):
                    print("Found key at %s: %s" % (idx, key))
                    keys[idx] = key
                    break
        idx += 1
    return keys


print("Day 14.1: %s" % sorted(list(solve(calc_key)))[-1])
print("Day 14.2: %s" % sorted(list(solve(calc_stretched_key)))[-1])
