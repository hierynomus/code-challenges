row = 2947
col = 3029

factor = 252533
divide = 33554393
start_value = 20151125


def nr_operations(row, col):
    return (((row + col - 2) * (row + col - 1)) / 2) + (col - 1)


def modular_exponent(ith_number):
    exp_mod = pow(factor, int(ith_number), divide)
    return (start_value * exp_mod) % divide


print("Day 25.1: %s" % modular_exponent(nr_operations(row, col)))
