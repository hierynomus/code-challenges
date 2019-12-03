from collections import defaultdict, namedtuple
from itertools import combinations, count, permutations
import fileinput

# import numpy as np
import sys

def fuel_for(mass):
    fuel = mass // 3 - 2
    total_fuel = 0
    while fuel > 0:
        total_fuel += fuel
        fuel = fuel // 3 - 2
    return total_fuel

def input():
    with fileinput.input() as f:
        for line in f:
            yield fuel_for(int(line.rstrip("\n")))


print(sum(input()))
