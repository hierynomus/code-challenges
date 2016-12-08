import re
import numpy as np


class Display(object):
    def __init__(self, display=np.array([['.'] * 50] * 6)):
        self.display = display

    def rect(self, x, y):
        d = self.display.copy()
        for i in range(x):
            for j in range(y):
                d[j][i] = '#'
        return Display(d)

    def show(self):
        for y in self.display:
            print(''.join(y))
        print("\n")

    def rotate_row(self, y, pos):
        d = self.display.copy()
        d[y] = np.roll(d[y], pos)
        return Display(d)

    def rotate_col(self, x, pos):
        d = self.display.copy()
        d = np.transpose(d)
        d[x] = np.roll(d[x], pos)
        return Display(np.transpose(d))

    def count(self):
        return (self.display == '#').sum()


rect = re.compile("(?P<x>[0-9]+)x(?P<y>[0-9]+)")
movement = re.compile("[xy]=(?P<idx>[0-9]+) by (?P<pos>[0-9]+)")
d = Display()

with open('day8.in', 'r') as f:
    for l in f:
        if l.startswith("rect"):
            m = rect.search(l).groupdict()
            d = d.rect(int(m['x']), int(m['y']))
        elif l.startswith("rotate col"):
            m = movement.search(l).groupdict()
            d = d.rotate_col(int(m['idx']), int(m['pos']))
        elif l.startswith("rotate row"):
            m = movement.search(l).groupdict()
            d = d.rotate_row(int(m['idx']), int(m['pos']))
        d.show()

print("Day 8.1: %s" % d.count())
print("Day 8.2:")
d.show()
