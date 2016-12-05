from itertools import permutations

persons = set()
happy_dict = dict()


def total_happiness(seating):
    total = 0
    for x, y in [(seating[i], seating[(i + 1) % len(seating)]) for i in range(len(seating))]:
        total += happy_dict[(x, y)] + happy_dict[(y, x)]
    return total


def best_seating(l):
    return max([perm for perm in permutations(l)], key=total_happiness)


with open('day13.in', 'r') as f:
    for l in f:
        parts = l.strip().split(' ')
        p1 = parts[0]
        p2 = parts[10][:-1]
        happiness = int(parts[3])
        persons.add(p1)
        persons.add(p2)
        if 'lose' in l:
            happy_dict[(p1, p2)] = -happiness
        elif 'gain' in l:
            happy_dict[(p1, p2)] = happiness
        else:
            raise Exception()

arrangement = best_seating(persons)
happiness_diff = total_happiness(arrangement)
print("Day 13.1: %s" % (happiness_diff,))

for p in persons:
    happy_dict[(p, 'Me')] = 0
    happy_dict[('Me', p)] = 0

persons.add('Me')

arrangement = best_seating(persons)
happiness_diff = total_happiness(arrangement)
print("Day 13.2: %s" % (happiness_diff,))
