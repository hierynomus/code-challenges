from collections import defaultdict, namedtuple, Counter
from itertools import combinations, count, permutations

import fileinput
# import numpy as np
import sys

def hamming_distance(s1, s2):
    return len([c for (c, d) in zip(s1, s2) if c != d])

input = [l.rstrip('\n') for l in fileinput.input()]
for i, j in combinations(input, 2):
    if hamming_distance(i, j) == 1:
        print(''.join([x for (x, y) in zip(i, j) if x == y]))
        break