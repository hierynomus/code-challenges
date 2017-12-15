inp_a, inp_b = 679, 771
modulus = 2147483647


def nr_generator(seed, mult, modulus, divisor):
    x = seed
    divisor -= 1
    while True:
        x = (x * mult) % modulus
        if x & divisor == 0:
            yield x


def nr_generator_opt(x, a, m, divisor):
    """https://programmingpraxis.com/2014/01/14/minimum-standard-random-number-generator/"""
    divisor -= 1
    while True:
        t = a * x
        p = t >> 31
        q = t & 0x7FFFFFFF
        x = p + q
        if x > m:
            x -= m
        if x & divisor == 0:
            yield x


def count_same(gen_a, gen_b, rounds):
    counter = 0
    for _ in range(rounds):
        a = next(gen_a)
        b = next(gen_b)
        if (a ^ b) & 0xFFFF == 0:
            counter += 1
    return counter


# assert count_same(nr_generator(inp_a, 16807, modulus, 1), nr_generator_opt(inp_a, 16807, modulus, 1), 10000) == 10000
print("Day 15.1:", count_same(nr_generator_opt(inp_a, 16807, modulus, 1), nr_generator_opt(inp_b, 48271, modulus, 1), 40000000))
print("Day 15.2:", count_same(nr_generator_opt(inp_a, 16807, modulus, 4), nr_generator_opt(inp_b, 48271, modulus, 8), 5000000))
