import fileinput

class Node(object):
    def __init__(self, nr_nodes, nr_meta):
        self.nr_nodes = nr_nodes
        self.nr_meta = nr_meta
        self.children = []
        self.meta = []

    def __repr__(self):
        return "Node[children=" + str(self.children) + ",meta=" + str(self.meta) + "]"

    def meta_sum(self):
        return sum(self.meta) + sum([c.meta_sum() for c in self.children])s

def parse(line, idx):
    nr_children = line[idx]
    nr_meta = line[idx + 1]
    n = Node(nr_children, nr_meta)
    idx += 2
    if nr_children:
        for c in range(nr_children):
            child, idx = parse(line, idx)
            n.children.append(child)
    n.meta = line[idx: idx + nr_meta]
    return n, idx + nr_meta

input = list(map(int, fileinput.input().readline().strip().split(' ')))
root, idx = parse(input, 0)
print(root.meta_sum())