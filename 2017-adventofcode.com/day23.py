instructions = []
with open('day23.in', 'r') as f:
    for l in f:
        instructions.append(tuple(l.strip().split()))


class Program(object):
    def __init__(self, instructions, debug=True):
        self.registers = {}
        for c in map(chr, range(97, 105)):
            self.registers[c] = 0
        if not debug:
            self.registers['a'] = 1
        self.debug = debug
        self.instructions = instructions
        self.pointer = 0
        self.multiplied = 0

    def reg_or_value(self, v):
        return self.registers[v] if v in self.registers else int(v)

    def run(self):
        r = 0
        while 0 <= self.pointer < len(self.instructions):
            if r > 60:
                break
            r += 1
            i = self.instructions[self.pointer]
            if not self.debug:
                print(i, self.registers)
            if i[0] == 'set':
                self.registers[i[1]] = self.reg_or_value(i[2])
            elif i[0] == 'sub':
                self.registers[i[1]] -= self.reg_or_value(i[2])
            elif i[0] == 'mul':
                self.registers[i[1]] *= self.reg_or_value(i[2])
                self.multiplied += 1
            elif i[0] == 'jnz':
                if self.reg_or_value(i[1]) != 0:
                    self.pointer += self.reg_or_value(i[2])
                    continue
            self.pointer += 1


# p_0 = Program(instructions)
# p_0.run()
# print("Day 23.1:", p_0.multiplied)

p_1 = Program(instructions, False)
p_1.run()
