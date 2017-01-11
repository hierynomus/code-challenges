from __future__ import print_function


class LiteralValue(object):
    def __init__(self, val):
        self.val = val

    def __int__(self):
        return self.val

    def __str__(self):
        return "LiteralValue(%s)" % self.val


class Register(object):
    def __init__(self, reg, vm):
        self.vm = vm
        self.reg = reg

    def __int__(self):
        return self.vm.read_register(self.reg)

    def store(self, val):
        return self.vm.store_register(self.reg, int(val))

    def __str__(self):
        return "Register(%s, val=%s)" % (self.reg, int(self))


class TerminationException(Exception):
    pass


class InstructionSet(object):
    @staticmethod
    def i0_halt(vm):
        vm.debug('halt')
        raise TerminationException()

    @staticmethod
    def i1_set(vm):
        reg, val = vm.read(), vm.read()
        vm.debug("set %s %s" % (reg, val))
        reg.store(val)

    @staticmethod
    def i2_push(vm):
        val = vm.read()
        vm.debug("push %s" % val)
        vm.stack.append(int(val))

    @staticmethod
    def i3_pop(vm):
        reg = vm.read()
        vm.debug("pop %s" % reg)
        reg.store(vm.stack.pop())

    @staticmethod
    def i4_eq(vm):
        reg, a, b = vm.read(), vm.read(), vm.read()
        vm.debug("eq %s %s %s" % (reg, a, b))
        reg.store(1 if int(a) == int(b) else 0)

    @staticmethod
    def i5_gt(vm):
        reg, a, b = vm.read(), vm.read(), vm.read()
        vm.debug("gt %s %s %s" % (reg, a, b))
        reg.store(1 if int(a) > int(b) else 0)

    @staticmethod
    def i6_jmp(vm):
        new_pointer = vm.read()
        vm.debug("jmp %s" % new_pointer)
        vm.pointer = int(new_pointer)

    @staticmethod
    def i7_jt(vm):
        cond, new_pointer = vm.read(), vm.read()
        vm.debug("jt %s %s" % (cond, new_pointer))
        if int(cond):
            vm.pointer = int(new_pointer)

    @staticmethod
    def i8_jf(vm):
        cond, new_pointer = vm.read(), vm.read()
        vm.debug("jf %s %s" % (cond, new_pointer))
        if not int(cond):
            vm.pointer = int(new_pointer)

    @staticmethod
    def i9_add(vm):
        reg, a, b = vm.read(), vm.read(), vm.read()
        vm.debug("add %s %s %s" % (reg, a, b))
        reg.store((int(a) + int(b)) % 32768)

    @staticmethod
    def i10_mult(vm):
        reg, a, b = vm.read(), vm.read(), vm.read()
        vm.debug("mult %s %s %s" % (reg, a, b))
        reg.store((int(a) * int(b)) % 32768)

    @staticmethod
    def i11_mod(vm):
        reg, a, b = vm.read(), vm.read(), vm.read()
        vm.debug("mod %s %s %s" % (reg, a, b))
        reg.store(int(a) % int(b))

    @staticmethod
    def i12_and(vm):
        reg, a, b = vm.read(), vm.read(), vm.read()
        vm.debug("and %s %s %s" % (reg, a, b))
        reg.store(int(a) & int(b))

    @staticmethod
    def i13_or(vm):
        reg, a, b = vm.read(), vm.read(), vm.read()
        vm.debug("or %s %s %s" % (reg, a, b))
        reg.store(int(a) | int(b))

    @staticmethod
    def i14_not(vm):
        reg, a = vm.read(), vm.read()
        vm.debug("not %s %s" % (reg, a))
        mask = 0
        for i in range(15):
            mask |= 1 << i
        i = ~int(a) & mask
        reg.store(i)

    @staticmethod
    def i15_rmem(vm):
        reg, mem = vm.read(), vm.read()
        vm.debug("rmem %s %s" % (reg, mem))
        reg.store(vm.memory[int(mem)])

    @staticmethod
    def i16_wmem(vm):
        mem, a = vm.read(), vm.read()
        vm.debug("wmem %s %s" % (mem, a))
        vm.memory[int(mem)] = int(a)

    @staticmethod
    def i17_call(vm):
        a = vm.read()
        vm.debug("call %s" % a)
        vm.stack.append(vm.pointer)
        vm.pointer = int(a)

    @staticmethod
    def i18_ret(vm):
        if not vm.stack:
            raise TerminationException()
        v = vm.stack.pop()
        vm.debug("ret")
        vm.pointer = v

    @staticmethod
    def i19_out(vm):
        print(chr(int(vm.read())), end='')

    @staticmethod
    def i20_in(vm):
        reg = vm.read()
        vm.debug("in %s" % reg)
        if not vm.input:
            while True:
                line = raw_input("> ") + '\n'
                if line.startswith('x'):
                    print("Registers: %s" % vm.registers)
                elif line.startswith('debug'):
                    vm.debug_enabled = True
                elif line.startswith('register'):
                    reg, val = [vm.as_val(int(x)) for x in line[8:].split()]
                    vm.debug("set %s %s" % (reg, val))
                    reg.store(val)
                else:
                    print(line)
                    vm.input = line
                    break

        c = ord(vm.input[0])
        vm.input = vm.input[1:]
        reg.store(c)

    @staticmethod
    def i21_noop(vm):
        vm.debug("noop")
        pass


class VirtualMachine(object):
    instructions = {
        0: InstructionSet.i0_halt,
        1: InstructionSet.i1_set,
        2: InstructionSet.i2_push,
        3: InstructionSet.i3_pop,
        4: InstructionSet.i4_eq,
        5: InstructionSet.i5_gt,
        6: InstructionSet.i6_jmp,
        7: InstructionSet.i7_jt,
        8: InstructionSet.i8_jf,
        9: InstructionSet.i9_add,
        10: InstructionSet.i10_mult,
        11: InstructionSet.i11_mod,
        12: InstructionSet.i12_and,
        13: InstructionSet.i13_or,
        14: InstructionSet.i14_not,
        15: InstructionSet.i15_rmem,
        16: InstructionSet.i16_wmem,
        17: InstructionSet.i17_call,
        18: InstructionSet.i18_ret,
        19: InstructionSet.i19_out,
        20: InstructionSet.i20_in,
        21: InstructionSet.i21_noop
    }

    def __init__(self, program, debug=False):
        self.memory = []
        self.stack = []
        self.registers = [0 for i in range(8)]
        self.read_to_memory(program)
        self.pointer = 0
        self.debug_enabled = debug
        self.input = None

    def debug(self, msg):
        if self.debug_enabled:
            print("%s: %s (Registers: %s)" % (self.instruction_address, msg, self.registers))

    def read_to_memory(self, program):
        """Fully read the program file to newly initialized memory"""
        with open(program, 'r') as f:
            while True:
                lobyte = f.read(1)
                if not lobyte:
                    break
                hibyte = f.read(1)
                try:
                    self.memory.append(256*ord(hibyte)+ord(lobyte))
                except:
                    raise Exception("memory: %s, bytes (%s, %s)", self.memory, lobyte, hibyte)

    def store_register(self, reg, val):
        self.registers[reg % 32768] = val

    def read_register(self, reg):
        return self.registers[reg % 32768]

    def read_instruction(self):
        # Used for debug logging.
        self.instruction_address = self.pointer
        return self.read()

    def read(self):
        """Read the next instruction or value from memory at pointer location and increment pointer"""
        v = self.memory[self.pointer]
        self.pointer += 1
        return self.as_val(v)

    def as_val(self, v):
        if v < 32768:
            return LiteralValue(v)
        elif v < 32768 + 8:
            return Register(v, self)
        else:
            raise Exception("%s is invalid" % v)

    def execute(self, instruction):
        VirtualMachine.instructions[int(instruction)](self)

    def preload_input(self, input_file):
        with open(input_file, 'r') as f:
            self.input = f.read()
        return self

    def run(self):
        while True:
            instruction = self.read_instruction()
            try:
                self.execute(instruction)
            except TerminationException:
                print("Terminating!")
                break
        pass


VirtualMachine('spec/challenge.bin', debug=False).preload_input('directions').run()
