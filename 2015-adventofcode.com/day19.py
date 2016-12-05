import re

rewrites = []
rev_rewrites = []
input = None


def recombinate(elements, index, target):
    new_elements = elements[:]
    new_elements[index] = target
    return new_elements


with open('day19.in', 'r') as f:
    for l in f:
        if '=>' in l:
            parts = l.strip().split()
            rewrites.append((parts[0], parts[2]))
            rev_rewrites.append((parts[2][::-1], parts[0][::-1]))
        elif len(l.strip()) > 0:
            input = l.strip()

molecules = set()

elements = [x for x in re.split('([A-Z][a-z]?)', input) if x]

for source, target in rewrites:
    for i, e in enumerate(elements):
        if e == source:
            molecules.add(''.join(recombinate(elements, i, target)))

print("Day 19.1: %s" % len(molecules))

rev_input = input[::-1]
index = 0
replacements = 0
while rev_input != 'e':
    to_check = rev_input[index:]
    for target, source in rev_rewrites:
        if to_check.startswith(target):
            rev_input = rev_input[0:index] + source + rev_input[index + len(target):]
            replacements += 1
            index = -1
            break
    index += 1

print("Day 19.2: %s" % replacements)
