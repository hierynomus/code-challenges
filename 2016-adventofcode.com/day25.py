from itertools import islice

instructions = []


def reg_or_value(v, regs):
    return regs[v] if v in regs else int(v)


def toggle(instr):
    if len(instr) == 2:
        return ('dec' if 'inc' in instr else 'inc', instr[1])
    else:
        return ('cpy' if 'jnz' in instr else 'jnz', instr[1], instr[2])


def optimized_add(pointer, instructions, registers):
    # Pointer current location: inc X
    target_reg = instructions[pointer][1]
    if pointer + 4 < len(instructions):
        if 'dec' in instructions[pointer + 1] and 'dec' in instructions[pointer + 3]:
            if 'jnz' in instructions[pointer + 2] and '-2' in instructions[pointer + 2]:
                if 'jnz' in instructions[pointer + 4] and '-5' in instructions[pointer + 4]:
                    reg_1, reg_2 = instructions[pointer + 1][1], instructions[pointer + 3][1]
                    if reg_1 in instructions[pointer + 2] and reg_2 in instructions[pointer + 4]:
                        registers[target_reg] += (registers[reg_1] * registers[reg_2])
                        registers[reg_1] = 0
                        registers[reg_2] = 0
                        return 4
    registers[target_reg] += 1
    return 0


def compute(instructions, registers, pointer=0, do_print=False):
    while pointer < len(instructions):
        instr = instructions[pointer]
        if do_print:
            print(pointer, instr, registers)
        if instr[0] == 'cpy' and instr[2] in registers:
            registers[instr[2]] = reg_or_value(instr[1], registers)
        elif instr[0] == 'inc' and instr[1] in registers:
            pointer += optimized_add(pointer, instructions, registers)
        elif instr[0] == 'dec' and instr[1] in registers:
            registers[instr[1]] -= 1
        elif instr[0] == 'jnz' and reg_or_value(instr[1], registers) != 0:
            pointer += reg_or_value(instr[2], registers)
            continue
        elif instr[0] == 'tgl':
            toggle_loc = pointer + reg_or_value(instr[1], registers)
            if toggle_loc > 0 and toggle_loc < len(instructions):
                instructions[toggle_loc] = toggle(instructions[toggle_loc])
        elif instr[0] == 'out':
            yield reg_or_value(instr[1], registers)
        pointer += 1


def find(instructions):
    a = 0
    expected_out = [0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1]
    while a < 1568:
        out = list(islice(compute(instructions, {'a': a, 'b': 0, 'c': 0, 'd': 0}), len(expected_out)))
        if expected_out != out:
            a += 1
        else:
            return a


with open('day25.in', 'r') as f:
    instructions = [tuple(l.strip().split()) for l in f]

print("Day 25.1: %s" % find(instructions[:]))
