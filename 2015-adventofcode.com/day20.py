import math

def smallest_prime_factor(n):
    i = 2
    while i * i <= n:
        if n % i:
            i += 1
        else:
            return i
    return n

gifts_needed = 200
houses = {1: 10}
i = 1
while houses[i] < gifts_needed:
    i += 1
    smallest = smallest_prime_factor(i)
    if smallest == i:
        houses[i] = 10 + (i * 10)
    else:
        houses[i] = houses[i / smallest] + (i * 10)

print houses

# def visiting_elves(house):
#     large_divisors = []
#     for elf in xrange(1, int(math.sqrt(house) + 1)):
#         if house % elf == 0:
#             yield elf
#             if elf*elf != house:
#                 large_divisors.append(house / elf)
#     for divisor in reversed(large_divisors):
#         yield divisor

# input = 34000000

# def presents(house, nr_delivered, elf_generator):
#     return sum([elf * nr_delivered for elf in elf_generator])

# # Part 1
# h = 1
# p = presents(h, 10, visiting_elves(h))

# while p < input:
#     # print("%s: %s presents" % (h, p))
#     h += 1
#     p = presents(h, 10, visiting_elves(h))

# print("1: %s" % (h,))

# # Part 2
# def lazy_visiting_elves(house):
#     high_elves = []
#     for elf in xrange(1, int(math.sqrt(house) + 1)):
#         if house % elf == 0:
#             if house / elf < 50:
#                 yield elf
#             if elf*elf != house:
#                 high_elves.append(house / elf)
#     for high_elf in reversed(high_elves):
#         if high_elf * 50 > house:
#             yield high_elf

# p = presents(h, 11, lazy_visiting_elves(h))
# while p < input:
#     # print("%s: %s presents" % (h, p))
#     h += 1
#     p = presents(h, 11, lazy_visiting_elves(h))

# print("2: %s" % (h,))
