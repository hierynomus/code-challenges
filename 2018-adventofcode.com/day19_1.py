import fileinput

class Watch(object):
    def __init__(self):
        self.ip = 0
        self.ip_register = None
        self.registers = [0] * 6
        self.instructions = []
        self.opcodes = {
            'addr': self.addr,
            'addi': self.addi,
            'muli': self.muli,
            'mulr': self.mulr,
            'andr': self.andr,
            'andi': self.andi,
            'borr': self.borr,
            'bori': self.bori,
            'setr': self.setr,
            'seti': self.seti,
            'gtri': self.gtri,
            'gtir': self.gtir,
            'gtrr': self.gtrr,
            'eqri': self.eqri,
            'eqir': self.eqir,
            'eqrr': self.eqrr
        }
    
    def addr(self, a, b, c):
        self.registers[c] = self.registers[a] + self.registers[b]

    def addi(self, a, b, c):
        self.registers[c] = self.registers[a] + b

    def mulr(self, a, b, c):
        self.registers[c] = self.registers[a] * self.registers[b]

    def muli(self, a, b, c):
        self.registers[c] = self.registers[a] * b

    def andr(self, a, b, c):
        self.registers[c] = self.registers[a] & self.registers[b]

    def andi(self, a, b, c):
        self.registers[c] = self.registers[a] & b

    def borr(self, a, b, c):
        self.registers[c] = self.registers[a] | self.registers[b]

    def bori(self, a, b, c):
        self.registers[c] = self.registers[a] | b

    def setr(self, a, b, c):
        self.registers[c] = self.registers[a]

    def seti(self, a, igncred, c):
        self.registers[c] = a

    def gtrr(self, a, b, c):
        self.registers[c] = 1 if self.registers[a] > self.registers[b] else 0

    def gtir(self, a, b, c):
        self.registers[c] = 1 if a > self.registers[b] else 0

    def gtri(self, a, b, c):
        self.registers[c] = 1 if self.registers[a] > b else 0

    def eqrr(self, a, b, c):
        self.registers[c] = 1 if self.registers[a] == self.registers[b] else 0

    def eqir(self, a, b, c):
        self.registers[c] = 1 if a == self.registers[b] else 0

    def eqri(self, a, b, c):
        self.registers[c] = 1 if self.registers[a] == b else 0

    def load(self, instructions):
        for instruction in instructions:
            if '#ip' in instruction:
                self.ip_register = int(instruction.split(' ')[1])
            else:
                s = instruction.split(' ')
                opcode = s[0]
                a, b, c = map(int, s[1:])
                self.instructions.append((opcode, a, b, c))

    def execute_opcode(self, opcode, a, b, c):
        self.registers[self.ip_register] = self.ip
        self.opcodes[opcode](a, b, c)
        self.ip = self.registers[self.ip_register]

    def execute(self):
        while self.ip >= 0 and self.ip < len(self.instructions):
            opcode, a, b, c = self.instructions[self.ip]
            print(self.ip, opcode, a, b, c, self.registers)
            self.execute_opcode(opcode, a, b, c)
            self.ip += 1


watch = Watch()
watch.load([l.rstrip('\n') for l in fileinput.input()])
watch.execute()
print(watch.registers[0])
