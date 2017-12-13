scanner = {}
with open('day13.in', 'r') as f:
    for l in f:
        layer, depth = map(int, l.strip().split(': '))
        scanner[layer] = (depth, 2 * depth - 2)

nr_layers = max(scanner.keys())


def pass_firewall(T):
    severity = 0
    caught = False
    for layer in range(nr_layers + 1):
        if layer in scanner:
            d, r = scanner[layer]
            pos = (layer + T) % r
            # print(layer, ":", pos, d)
            if pos == 0:
                caught = True
                severity += (layer * d)
    return (caught, severity)


def clock():
    n = 0
    while True:
        yield n
        n += 1


print("Day 13.1:", pass_firewall(0)[1])
# print(pass_firewall(168108))

for T in clock():
    if not pass_firewall(T)[0]:
        print("Day 13.2:", T)
        break
