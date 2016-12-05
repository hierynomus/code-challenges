from itertools import combinations
from functools import reduce


def quantum_entanglement(gifts):
    return reduce(lambda x, y: x * y, gifts)


def minimum_balanced_combos(gifts, groups):
    balance_weight = sum(gifts) / groups
    combos = []
    for i in range(2, len(gifts)):
        for c in combinations(gifts, i):
            if sum(c) == balance_weight:
                combos.append(c)
        if len(combos) > 0:
            break
    return combos


gifts = []
with open('day24.in', 'r') as f:
    for l in f:
        gifts.append(int(l.strip()))

gifts.sort()

combo = min(minimum_balanced_combos(gifts, 3), key=quantum_entanglement)
print("Day 24.1: %s" % quantum_entanglement(combo))
combo = min(minimum_balanced_combos(gifts, 4), key=quantum_entanglement)
print("Day 24.2: %s" % quantum_entanglement(combo))


