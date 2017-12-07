class Program(object):
    def __init__(self, l):
        s = l.strip().split(' -> ')
        self.name, size = s[0].split()
        self.weight = int(size[1:-1])
        self.children = [] if len(s) == 1 else list(s[1].split(', '))
        self.correction = 0
        self.parent = None

    def __repr__(self):
        return "%s (%s + %s) -> %s" % (self.name, self.weight, self.correction, self.children)

    def fix_children(self, programs):
        self.children = list(map(lambda c: programs[c], self.children))
        for c in self.children:
            c.parent = self

    def full_weight(self):
        return self.weight + sum([c.full_weight() for c in self.children]) + self.correction

    def correct_unbalance(self):
        unbalanced = None
        for c in self.children:
            ub = c.correct_unbalance()
            if ub:
                unbalanced = ub

        if not len(self.children) or unbalanced:
            return unbalanced

        cw = sorted([(c, c.full_weight()) for c in self.children], key=lambda x: x[1])
        wrong = cw[0] if cw[0][1] != cw[1][1] else cw[-1]
        diff = cw[1][1] - wrong[1]
        if diff:
            wrong[0].correction = diff
            return wrong[0]
        return None


programs = {}
with open('day07.in', 'r') as f:
    for l in f.readlines():
        p = Program(l)
        programs[p.name] = p

for _, p in programs.items():
    p.fix_children(programs)

root = [x for _, x in programs.items() if not x.parent][0]
print("Day 7.1:", root.name)

unbalanced = root.correct_unbalance()
print("Day 7.2:", unbalanced.weight + unbalanced.correction)
