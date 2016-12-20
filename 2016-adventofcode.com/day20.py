
def solve(ranges):
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
    return lowest, count


ranges = []

with open('day20.in', 'r') as f:
    for line in f:
        low, high = map(int, line.strip().split('-'))
        ranges.append((low, high))

ranges.sort()
lowest, count = solve(ranges)
print("Day 20.1: %s" % lowest)
print("Day 20.2: %s" % count)
