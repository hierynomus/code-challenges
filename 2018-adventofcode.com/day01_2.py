from collections import defaultdict, namedtuple
from itertools import combinations, count, permutations, cycle
import fileinput
# import numpy as np
import sys

def input():
    with fileinput.input() as f:
        for line in f:
            yield line.rstrip('\n')

l = list(map(int, input()))
s = 0
seen = set([s])
for i in cycle(l):
    s += i
    if s not in seen:
        seen.add(s)
    else:
        print(s)
        break
            