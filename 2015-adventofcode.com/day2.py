area = 0
ribbon = 0

with open('day2.in', 'r') as f:
    for gift in f:
        l, w, h = [int(a) for a in gift.split('x')]
        lw, wh, hl = l * w, w * h, h * l
        area += 2 * lw + 2 * wh + 2 * hl + min([lw, wh, hl])

        lengths = sorted([l, w, h])
        ribbon += 2 * lengths[0] + 2 * lengths[1] + l * w * h

print("Day 2.1: %s" % area)
print("Day 2.2: %s" % ribbon)
