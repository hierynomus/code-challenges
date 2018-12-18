import fileinput

class Watch(object):
    def __init__(self, registers):
        self.registers = registers
        self.opcodes = {
            2: self.addr,
            14: self.addi,
            4: self.muli,
            6: self.mulr,
            7: self.andr,
            11: self.andi,
            1: self.borr,
            8: self.bori,
            12: self.setr,
            15: self.seti,
            3: self.gtri,
            5: self.gtir,
            13: self.gtrr,
            9: self.eqri,
            0: self.eqir,
            10: self.eqrr
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

    def opcode(self, instruction):
        self.opcodes[instruction[0]](instruction[1], instruction[2], instruction[3])

watch = Watch([0, 0, 0, 0])
program = []
with fileinput.input() as f:
    while True:
        line = next(f, None)
        if not line:
            break
        if 'Before:' in line:
            start = list(map(int, line[9:-2].split(',')))
            code = list(map(int, next(f).rstrip('\n').split()))
            end = list(map(int, next(f)[9:-2].split(',')))
        elif line.rstrip('\n'):
            program.append(list(map(int, line.rstrip('\n').split())))

for i in program:
    watch.opcode(i)

print(watch.registers[0])
