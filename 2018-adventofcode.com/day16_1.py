import fileinput

class Watch(object):
    def __init__(self, registers):
        self.registers = registers
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

    def seti(self, b, igncred, c):
        self.registers[c] = b

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

    def matching_opcodes(self, begin, instruction, end):
        c = 0
        for _, opcode in self.opcodes.items():
            self.registers = begin[:]
            opcode(instruction[1], instruction[2], instruction[3])
            c += 1 if self.registers == end else 0
        return c

watch = Watch([0, 0, 0, 0])
count = 0
with fileinput.input() as f:
    while True:
        line = next(f, None)
        if not line:
            break
        if 'Before:' in line:
            start = list(map(int, line[9:-2].split(',')))
            code = list(map(int, next(f).rstrip('\n').split()))
            end = list(map(int, next(f)[9:-2].split(',')))
            codes = watch.matching_opcodes(start, code, end)
            count += 1 if codes >= 3 else 0

print(count)
