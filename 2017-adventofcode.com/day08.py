from collections import defaultdict
with open('day08.in', 'r') as f:
    instructions = [l.strip().split() for l in f]

registers = defaultdict(int)
max_ever = 0
for i in instructions:
    r = i[0]
    diff = int(i[2]) if i[1] == 'inc' else -int(i[2])
    reg_to_check = i[4]
    op = i[5]
    val_to_check = int(i[6])

    upd = False
    if op == '==':
        upd = (registers[reg_to_check] == val_to_check)
    elif op == '<=':
        upd = (registers[reg_to_check] <= val_to_check)
    elif op == '>=':
        upd = (registers[reg_to_check] >= val_to_check)
    elif op == '<':
        upd = (registers[reg_to_check] < val_to_check)
    elif op == '>':
        upd = (registers[reg_to_check] > val_to_check)
    elif op == '!=':
        upd = (registers[reg_to_check] != val_to_check)

    if upd:
        registers[r] += diff
        if registers[r] > max_ever:
            max_ever = registers[r]

print("Day 8.1:", max([v for _, v in registers.items()]))
print("Day 8.2:", max_ever)
