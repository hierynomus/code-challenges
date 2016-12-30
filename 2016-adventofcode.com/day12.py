class Computer(object):
    def __init__(self, registers={'a': 0, 'b': 0, 'c': 0, 'd': 0}):
        self.registers = registers
        self.pointer = 0

    def value(self, v):
        return self.registers[v] if v in self.registers else int(v)

    def smart_add_from_reg(self, x):
        if len(self.program) > self.pointer + 2:
            next_instruction = self.program[self.pointer + 1]
            if next_instruction[0] == 'dec':
                last_instruction = self.program[self.pointer + 2]
                if last_instruction[0] == 'jnz' and last_instruction[2] == '-2' and last_instruction[1] == next_instruction[1]:
                    self.registers[x] += self.registers[next_instruction[1]]
                    self.registers[next_instruction[1]] = 0
                    self.pointer += 2
                    return True
        return False

    def inc(self, x):
        if not self.smart_add_from_reg(x):
            self.registers[x] += 1

    def dec(self, x):
        self.registers[x] -= 1

    def cpy(self, x, y):
        self.registers[y] = self.value(x)

    def jnz(self, x, y):
        if self.value(x):
            self.pointer += self.value(y) - 1  # The last increment is done by the program loop

    def load(self, program):
        self.program = [tuple(instruction.strip().split()) for instruction in program]
        return self

    def execute(self, debug=False):
        """Executes a program, which is a list of instructions"""
        while self.pointer < len(self.program):
            instruction = self.program[self.pointer]
            if debug:
                print("{:3}: {:25} {}".format(self.pointer, instruction, self.registers))
            getattr(self, instruction[0])(*instruction[1:])
            self.pointer += 1
        return self


with open('day12.in', 'r') as f:
    program = f.readlines()

print("Day 12.1: {}".format(Computer().load(program).execute().registers['a']))
print("Day 12.2: {}".format(Computer({'a': 0, 'b': 0, 'c': 1, 'd': 0}).load(program).execute().registers['a']))
