ranges = []

with open('day20.in', 'r') as f:
    ranges = [tuple(map(int, line.strip().split('-'))) for line in f]

ranges.sort()
lowest = None
curr = 0
count = 0

for low, high in ranges:
    if curr < low:
        count += low - curr
        curr = high + 1
        if not lowest:
            lowest = curr
    elif curr < high:
        curr = high + 1

if curr < 4294967295:
    count += 4294967295 - curr

print("Day 20.1: %s" % lowest)
print("Day 20.2: %s" % count)
