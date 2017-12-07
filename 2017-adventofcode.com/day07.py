parent_child = {}
program_size = {}
with open('day07.in', 'r') as f:
    for l in f.readlines():
        tower_line = l.strip().split(' -> ')
        if len(tower_line) > 1:
            (program, size), children = tower_line[0].split(), list(tower_line[1].split(', '))
        else:
            program, size = tower_line[0].split()
            children = []
        program_size[program] = int(size[1:-1])
        parent_child[program] = children

# Part 1
root = None
potential_roots = set(parent_child.keys())
for _, children in parent_child.items():
    for c in children:
        potential_roots.remove(c)

assert len(potential_roots) == 1
root = potential_roots.pop()
print("Day 7.1:", root)


# Part 2
tower_weight = {}


def calc_weight(program):
    children = parent_child[program]
    self_weight = program_size[program]
    child_weights = []
    if len(children) == 0:
        tower_weight[program] = self_weight
        return self_weight
    for c in children:
        if c not in tower_weight:
            if not calc_weight(c):
                return False
        child_weights.append((c, tower_weight[c]))

    child_weights.sort(key=lambda t: t[1])
    weight_diff = child_weights[0][1] - child_weights[-1][1]
    if weight_diff:
        wrong_child = child_weights[0] if child_weights[0][1] != child_weights[1][1] else child_weights[-1]
        wrong_weight = program_size[wrong_child[0]]
        correct_weight = wrong_weight + weight_diff if wrong_weight < child_weights[1][1] else wrong_weight - weight_diff
        print("Day 7.2:", correct_weight)
        return False
    tower_weight[program] = self_weight + sum([x[1] for x in child_weights])
    return True


calc_weight(root)
