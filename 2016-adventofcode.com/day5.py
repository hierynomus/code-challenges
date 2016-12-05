import hashlib

inp = "reyedfim"
hasher = hashlib.md5()
hasher.update(inp)


def digest(nr):
    copy = hasher.copy()
    copy.update(str(nr))
    return copy.hexdigest()


def next_hash(idx):
    dig = digest(idx)
    while dig[:5] != '00000':
        idx += 1
        dig = digest(idx)
    return (idx, dig)


password1 = ""
password2 = [None] * 8
idx = 0

# Password 1
for _ in range(0, 8):
    idx, dig = next_hash(idx)
    password1 += str(dig[5])
    pos = int(dig[5], 16)
    if pos < 8 and not password2[pos]:
        password2[pos] = dig[6]
    idx += 1

while None in password2:
    idx, dig = next_hash(idx)
    pos = int(dig[5], 16)
    if pos < 8 and not password2[pos]:
        password2[pos] = dig[6]
    idx += 1


print("Day 5.1: %s" % password1)
print("Day 5.2: %s" % password2)
