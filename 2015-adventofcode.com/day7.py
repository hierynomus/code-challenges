from bitarray import bitarray

DEBUG = False

def debug(msg):
    if DEBUG:
        print(msg)

class mybitarray(bitarray):
    def __lshift__(self, count):
        return self[count:] + type(self)('0') * count
    def __rshift__(self, count):
        return type(self)('0') * count + self[:-count]
    def __repr__(self):
        return "{}('{}')".format(type(self).__name__, self.to01())


class Signal:
    def __init__(self, n):
        self.signal = self.__signal(n)

    def get_signal(self):
        return self.signal

    def __signal(self, n):
        binary = bin(n)[2:]
        if len(binary) > 16:
            raise Exception()
        ba = (16 - len(binary)) * mybitarray('0')
        ba.extend(binary)
        return ba


class Wire:
    def __init__(self, name):
        self.name = name
        self.signal = None
        self.input = None

    def set_input(self, i):
        self.input = i

    def get_signal(self):
        if self.signal == None:
            self.signal = self.input.get_signal()
        debug("%s: %s" % (self.name, self.signal.to01()))
        return self.signal

    def reset(self):
        self.signal = None

class AndGate:
    def __init__(self, in_wire_1, in_wire_2, out_wire):
        self.in_wire_1 = in_wire_1
        self.in_wire_2 = in_wire_2
        self.out_wire = out_wire
        out_wire.set_input(self)
        self.name = "%s AND %s -> %s" % (in_wire_1.name, in_wire_2.name, out_wire.name)

    def get_signal(self):
        out_signal = self.in_wire_1.get_signal() & self.in_wire_2.get_signal()
        debug("%s: %s & %s -> %s" % (self.name, self.in_wire_1.get_signal().to01(), self.in_wire_2.get_signal().to01(), out_signal.to01()))
        return out_signal

class OrGate:
    def __init__(self, in_wire_1, in_wire_2, out_wire):
        self.in_wire_1 = in_wire_1
        self.in_wire_2 = in_wire_2
        self.out_wire = out_wire
        out_wire.set_input(self)
        self.name = "%s OR %s -> %s" % (in_wire_1.name, in_wire_2.name, out_wire.name)

    def get_signal(self):
        out_signal = self.in_wire_1.get_signal() | self.in_wire_2.get_signal()
        debug("%s: %s | %s -> %s" % (self.name, self.in_wire_1.get_signal().to01(), self.in_wire_2.get_signal().to01(), out_signal.to01()))
        return out_signal


class LShiftGate:
    def __init__(self, in_wire, out_wire, shift):
        self.in_wire = in_wire
        self.out_wire = out_wire
        out_wire.set_input(self)
        self.shift = shift
        self.name = "%s LSHIFT %s -> %s" % (in_wire.name, shift, out_wire.name)

    def get_signal(self):
        out_signal = self.in_wire.get_signal() << self.shift
        debug("%s: %s << %s -> %s" % (self.name, self.in_wire.get_signal().to01(), self.shift, out_signal.to01()))
        return out_signal


class RShiftGate:
    def __init__(self, in_wire, out_wire, shift):
        self.in_wire = in_wire
        self.out_wire = out_wire
        out_wire.set_input(self)
        self.shift = shift
        self.name = "%s RSHIFT %s -> %s" % (in_wire.name, shift, out_wire.name)

    def get_signal(self):
        out_signal = self.in_wire.get_signal() >> self.shift
        debug("%s: %s >> %s -> %s" % (self.name, self.in_wire.get_signal().to01(), self.shift, out_signal.to01()))
        return out_signal


class NotGate:
    def __init__(self, in_wire, out_wire):
        self.in_wire = in_wire
        self.out_wire = out_wire
        out_wire.set_input(self)
        self.name = "NOT %s -> %s" % (self.in_wire.name, self.out_wire.name)

    def get_signal(self):
        in_signal = self.in_wire.get_signal()
        in_signal_copy = in_signal.copy()
        in_signal_copy.invert()
        debug("%s: NOT %s -> %s" % (self.name, in_signal.to01(), in_signal_copy.to01()))
        return in_signal_copy


class JointGate:
    def __init__(self, in_wire, out_wire):
        self.in_wire = in_wire
        self.out_wire = out_wire
        out_wire.set_input(self)
        self.name = "%s -> %s" % (self.in_wire.name, self.out_wire.name)

    def get_signal(self):
        in_signal = self.in_wire.get_signal()
        in_signal_copy = in_signal.copy()
        debug("%s: %s -> %s" % (self.name, in_signal.to01(), in_signal_copy.to01()))
        return in_signal_copy



def signal_to_int(s):
    out = 0
    for bit in s.to01():
        out = (out << 1) | int(bit)
    return out

wires = dict()

def get_or_create_wire(n):
    n = n.strip()
    if n.isdigit():
        w = Wire('<anon>')
        w.set_input(Signal(int(n)))
        return w
    elif n in wires:
        return wires[n]
    else:
        w = Wire(n)
        wires[n] = w
        return w

with open('day7.in', 'r') as f:
    for l in f:
        parts = l.split(' ')
        if 'AND' in l:
            w1 = get_or_create_wire(parts[0])
            w2 = get_or_create_wire(parts[2])
            out = get_or_create_wire(parts[4])
            AndGate(w1, w2, out)
        elif 'OR' in l:
            w1 = get_or_create_wire(parts[0])
            w2 = get_or_create_wire(parts[2])
            out = get_or_create_wire(parts[4])
            OrGate(w1, w2, out)
        elif 'LSHIFT' in l:
            w = get_or_create_wire(parts[0])
            shift = int(parts[2])
            out = get_or_create_wire(parts[4])
            LShiftGate(w, out, shift)
        elif 'RSHIFT' in l:
            w = get_or_create_wire(parts[0])
            shift = int(parts[2])
            out = get_or_create_wire(parts[4])
            RShiftGate(w, out, shift)
        elif 'NOT' in l:
            w = get_or_create_wire(parts[1])
            out = get_or_create_wire(parts[3])
            NotGate(w, out)
        elif parts[0].isdigit():
            w = get_or_create_wire(parts[2])
            w.set_input(Signal(int(parts[0])))
        else:
            w = get_or_create_wire(parts[0])
            out = get_or_create_wire(parts[2])
            JointGate(w, out)

a_signal = signal_to_int(wires['a'].get_signal())
print("1: " + str(a_signal))

wires['b'].set_input(Signal(a_signal))
for n, w in wires.iteritems():
    w.reset()

a_signal = signal_to_int(wires['a'].get_signal())
print("2: " + str(a_signal))

