from itertools import count

firewall = []


def as_rule(layer, depth):
    return lambda T: layer * depth if ((T + layer) % (2 * (depth - 1))) == 0 else -1


def pass_firewall(T):
    for rule in firewall:
        yield rule(T)


def pass_undetected():
    for c in count():
        if not any(map(lambda x: x >= 0, pass_firewall(c))):
            yield c


with open('day13.in', 'r') as f:
    for l in f:
        layer, depth = map(int, l.strip().split(': '))
        firewall.append(as_rule(layer, depth))


print("Day 13.1:", sum([max(0, severity) for severity in pass_firewall(0)]))
print("Day 13.2:", next(pass_undetected()))
