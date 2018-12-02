from collections import defaultdict, namedtuple
from itertools import combinations, count, permutations
import fileinput
# import numpy as np
import sys

def input():
    with fileinput.input() as f:
        for line in f:
            yield line.rstrip('\n')

print(sum(map(int, input())))
