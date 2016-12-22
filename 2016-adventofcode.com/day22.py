import re
import numpy as np

np.set_printoptions(threshold=np.inf, linewidth=np.inf)
fs_re = re.compile("/dev/grid/node-x([0-9]+)-y([0-9]+)[ 0-9]+T[ ]*([0-9]+)T[ ]+([0-9]+)T[ ]+.*")


class Node(object):
    def __init__(self, x, y, used, free):
        self.x = x
        self.y = y
        self.used = used
        self.free = free
        self.cap = used + free
        self.grid = None

    def neighbours(self):
        for dx, dy in [(-1, 0), (1, 0), (0, 1), (0, -1)]:
            px, py = self.x + dx, self.y + dy
            if px > 0 and py > 0 and px < self.grid.shape[1] and py < self.grid.shape[0]:
                yield grid[py][px]

    def fits_data(self, node):
        return self.cap >= node.used

    def __repr__(self):
        if self.used == 0:
            return '{:^7}'.format('___')
        elif any([not n.fits_data(self) for n in self.neighbours()]):
            return '{:^7}'.format('|||')
        elif all([n.fits_data(self) for n in self.neighbours()]):
            return '{:^7}'.format('.')
        return '{:^7}'.format('{}/{}'.format(self.free, self.used))


nodes = []
grid = None

with open('day22.in', 'r') as f:
    max_x, max_y = 0, 0
    for line in f:
        m = fs_re.match(line.strip())
        if m:
            node = Node(*map(int, (m.group(1), m.group(2), m.group(3), m.group(4))))
            nodes.append(node)
            max_x = max_x if max_x > node.x else node.x
            max_y = max_y if max_y > node.y else node.y
    grid = np.transpose(np.array(nodes).reshape(max_x + 1, max_y + 1))
    for n in nodes:
        n.grid = grid

viable_pairs = []
for a in nodes:
    if a.used:
        viable_pairs.extend([(a, b) for b in nodes if a != b and a.used < b.free])

print("Day 22.1: %s" % len(viable_pairs))
print("Day 22.2: \n%s" % grid)
