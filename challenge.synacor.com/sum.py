from itertools import permutations

chars = [2, 3, 5, 7, 9]


def calc(perm):
    return perm[0] + perm[1] * perm[2]**2 + perm[3]**3 - perm[4]

for p in permutations(chars):
    print "%s + %s * %s^2 + %s^3 - %s = %s" % (p[0], p[1], p[2], p[3], p[4], calc(p))
