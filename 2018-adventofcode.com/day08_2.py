import fileinput

class Node(object):
    def __init__(self, nr_nodes, nr_meta):
        self.nr_nodes = nr_nodes
        self.nr_meta = nr_meta
        self.children = []
        self.meta = []

    def __repr__(self):
        return "Node[children=" + str(self.children) + ",meta=" + str(self.meta) + "]"

    def node_sum(self):
        return sum(self.meta) if not self.children else self.child_sum()

    def child_sum(self):
        s = 0
        for m in self.meta:
            if m > 0 and m <= len(self.children):
                s += self.children[m - 1].node_sum()
        return s

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
print(root.node_sum())