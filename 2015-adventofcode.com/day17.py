from itertools import chain, combinations

def powerset(iterable):
    "powerset([1,2,3]) --> () (1,) (2,) (3,) (1,2) (1,3) (2,3) (1,2,3)"
    s = list(iterable)
    return chain.from_iterable(combinations(s, r) for r in range(len(s)+1))

target = 150

containers = []

with open('day17.in', 'r') as f:
    for l in f:
        containers.append(int(l.strip()))

total = 0

for c in powerset(containers):
    if sum(c) == target:
        total += 1

print("1: %s" % total)

min_length = -1
min_total = 0
for c in powerset(containers):
    if sum(c) == target and min_length == -1:
        min_length = len(c)
        min_total += 1
    elif sum(c) == target and min_length == len(c):
        min_total += 1
    elif min_length != -1 and len(c) > min_length:
        break

print("2: %s" % min_total)
