instructions = []


def reg_or_value(v, regs):
    return int(v) if v.isdigit() else regs[v]


def compute(registers):
    pointer = 0
    while pointer < len(instructions):
        instr = instructions[pointer]
        parts = instr.split()
        if parts[0] == 'cpy':
            registers[parts[2]] = reg_or_value(parts[1], registers)
        elif parts[0] == 'inc':
            registers[parts[1]] += 1
        elif parts[0] == 'dec':
            registers[parts[1]] -= 1
        elif parts[0] == 'jnz' and reg_or_value(parts[1], registers) != 0:
            pointer += int(parts[2])
            continue

        pointer += 1
    return registers


with open('day12.in', 'r') as f:
    instructions = [l.strip() for l in f]

print("Day 12.1: %s" % compute({'a': 0, 'b': 0, 'c': 0, 'd': 0})['a'])
print("Day 12.2: %s" % compute({'a': 0, 'b': 0, 'c': 1, 'd': 0})['a'])
