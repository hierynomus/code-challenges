from collections import deque

instructions = []
with open('day18.in', 'r') as f:
    for l in f:
        instructions.append(tuple(l.strip().split()))


class Program(object):
    def __init__(self, program_id, instructions, recv_queue, send_queue):
        self.registers = {}
        for c in map(chr, range(97, 123)):
            self.registers[c] = 0
        self.program_id = program_id
        self.registers['p'] = program_id
        self.instructions = instructions
        self.send_queue = send_queue
        self.recv_queue = recv_queue
        self.pointer = 0
        self.nr_sent = 0

    def reg_or_value(self, v):
        return self.registers[v] if v in self.registers else int(v)

    def run(self):
        while 0 <= self.pointer < len(self.instructions):
            i = self.instructions[self.pointer]
            if i[0] == 'set':
                self.registers[i[1]] = self.reg_or_value(i[2])
            elif i[0] == 'snd':
                self.send_queue.append(self.reg_or_value(i[1]))
                self.nr_sent += 1
            elif i[0] == 'add':
                self.registers[i[1]] += self.reg_or_value(i[2])
            elif i[0] == 'mul':
                self.registers[i[1]] *= self.reg_or_value(i[2])
            elif i[0] == 'mod':
                self.registers[i[1]] = self.registers[i[1]] % self.reg_or_value(i[2])
            elif i[0] == 'rcv':
                if len(self.recv_queue) > 0:
                    self.registers[i[1]] = self.recv_queue.popleft()
                else:
                    break
            elif i[0] == 'jgz':
                if self.reg_or_value(i[1]) > 0:
                    self.pointer += self.reg_or_value(i[2])
                    continue
            self.pointer += 1


send = deque()

p_0 = Program(0, instructions, deque(), send)
p_0.run()
print("Day 18.1:", send.pop())

send_to_0 = deque()
send_to_1 = deque()
p_0 = Program(0, instructions, send_to_0, send_to_1)
p_1 = Program(1, instructions, send_to_1, send_to_0)

while True:
    p_0.run()
    p_1.run()
    if len(send_to_0) + len(send_to_1) == 0:
        break

print("Day 18.2:", p_1.nr_sent)
