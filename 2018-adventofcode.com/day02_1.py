from collections import defaultdict, namedtuple, Counter
from itertools import combinations, count, permutations
import fileinput
# import numpy as np
import sys


twos = 0
threes = 0
input = [l.rstrip('\n') for l in fileinput.input()]
for i in input:
    two = False
    three = False
    for k, v in Counter(i).items():
        two = two or v == 2
        three = three or v == 3
    twos += 1 if two else 0
    threes += 1 if three else 0

print(twos * threes)