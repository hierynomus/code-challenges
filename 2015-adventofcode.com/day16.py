gift = {
    "children": 3,
    "cats": 7,
    "samoyeds": 2,
    "pomeranians": 3,
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5,
    "trees": 3,
    "cars": 2,
    "perfumes": 1
}


def to_aunt(l):
    aunt, d = l.split(': ', 1)
    knowledge = {}
    for k in d.split(','):
        a, b = k.strip().split(':')
        knowledge[a] = int(b)

    return {
        "aunt": int(aunt.split()[1]),
        "knowledge": knowledge
    }


aunts = []

with open('day16.in', 'r') as f:
    for l in f:
        aunts.append(to_aunt(l))

for aunt in aunts:
    for k, v in aunt['knowledge'].items():
        if gift[k] != v:
            break
    else:
        print("Day 16.1: %s" % aunt['aunt'])

for aunt in aunts:
    for k, v in aunt['knowledge'].items():
        if k in ['cats', 'trees'] and gift[k] >= v:
            break
        elif k in ['pomeranians', 'goldfish'] and gift[k] <= v:
            break
        elif k not in ['cats', 'trees', 'pomeranians', 'goldfish'] and gift[k] != v:
            break
    else:
        print("Day 16.2: %s" % aunt['aunt'])
