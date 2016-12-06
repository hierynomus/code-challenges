from collections import Counter

with open('day6.in', 'r') as f:
    part1 = None
    for l in f:
        word = l.strip()
        if not part1:
            part1 = [''] * len(word)
        for i in range(len(word)):
            part1[i] += word[i]
    print("Day 6.1: %s" % ''.join([Counter(c).most_common()[0][0] for c in part1]))
    print("Day 6.2: %s" % ''.join([Counter(c).most_common()[-1][0] for c in part1]))
