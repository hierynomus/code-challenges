inp_a, inp_b = 679, 771
modulus = 2147483647


def nr_generator(inp, mult, modulus, divisor):
    x = inp
    while True:
        x = (x * mult) % modulus
        if x % divisor == 0:
            yield x


def count_same(gen_a, gen_b, rounds):
    counter = 0
    for _ in range(rounds):
        a = next(gen_a)
        b = next(gen_b)
        if (a ^ b) & 0xFFFF == 0:
            counter += 1
    return counter


print("Day 15.1:", count_same(nr_generator(inp_a, 16807, modulus, 1), nr_generator(inp_b, 48271, modulus, 1), 40000000))
print("Day 15.2:", count_same(nr_generator(inp_a, 16807, modulus, 4), nr_generator(inp_b, 48271, modulus, 8), 5000000))
