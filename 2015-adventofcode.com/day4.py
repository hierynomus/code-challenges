import hashlib

input = "iwrupvqb"
nr = 1
hasher = hashlib.md5()
hasher.update(input)

def digest(nr):
    copy = hasher.copy()
    copy.update(str(nr))
    return copy.hexdigest()


dig = digest(nr)
while dig[:5] != '00000':
    nr += 1
    dig = digest(nr)

print("1: %s" % str(nr))


nr = 1
dig = digest(nr)
while dig[:6] != '000000':
    nr += 1
    dig = digest(nr)

print("2: %s" % str(nr))
