import hashlib

inp = "iwrupvqb".encode('utf-8')
nr = 1
hasher = hashlib.md5()
hasher.update(inp)


def digest(nr):
    copy = hasher.copy()
    copy.update(str(nr).encode('utf-8'))
    return copy.hexdigest()


dig = digest(nr)
while dig[:5] != '00000':
    nr += 1
    dig = digest(nr)

print("Day 4.1: %s" % str(nr))


nr = 1
dig = digest(nr)
while dig[:6] != '000000':
    nr += 1
    dig = digest(nr)

print("Day 4.2: %s" % str(nr))
