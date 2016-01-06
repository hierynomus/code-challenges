
DEBUG = False

def debug(msg):
    if DEBUG:
        print(msg)


instructions = [""]

with open('day23.in', 'r') as f:
    for l in f:
        instructions.append(l.strip())

def jmp(instruction, index):
    if '-' in instruction:
        return -int(instruction[index:])
    else:
        return int(instruction[index:])


def evaluate(instructions, registers):
    pointer = 1
    while pointer < len(instructions):
        i = instructions[pointer]
        debug("%s: %s (%s)" % (pointer, i, registers))
        if "jio" in i and registers[i[4]] == 1:
            pointer += jmp(i, 8)
        elif "jie" in i and registers[i[4]] % 2 == 0:
            pointer += jmp(i, 8)
        elif 'jmp' in i:
            pointer += jmp(i, 5)
        elif 'inc' in i:
            registers[i[4]] += 1
            pointer += 1
        elif 'tpl' in i:
            registers[i[4]] *= 3
            pointer += 1
        elif 'hlf' in i:
            registers[i[4]] /= 2
            pointer += 1
        else:
            pointer += 1
    return registers

print("1: %s" % evaluate(instructions, {'a': 0, 'b': 0})['b'])
print("2: %s" % evaluate(instructions, {'a': 1, 'b': 0})['b'])
