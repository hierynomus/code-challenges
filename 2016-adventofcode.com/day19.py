import math


class Elf(object):
    def __init__(self, value):
        self.value = value
        self.prev = None
        self.next = None
        self.deleted = False

    def delete(self):
        self.prev.next = self.next
        self.next.prev = self.prev
        self.deleted = True

    def __repr__(self):
        return "Elf[{}]".format(self.value)


def solve_1(elves):
    elves_left = list(map(Elf, range(1, elves + 1)))
    while len(elves_left) > 1:
        new_elves_left = elves_left[::2]
        if len(elves_left) % 2 == 1:
            elves_left = new_elves_left[1:]
        else:
            elves_left = new_elves_left
    return elves_left[0]


def solve_2(nr_elves):
    elves = list(map(Elf, range(1, nr_elves + 1)))
    for i in range(nr_elves):
        elves[i].next = elves[(i + 1) % nr_elves]
        elves[i].prev = elves[(i - 1) % nr_elves]

    curr_elf, steal = elves[0], elves[int(math.floor(nr_elves / 2))]
    for i in range(nr_elves - 1):
        steal.delete()
        steal = steal.next
        if nr_elves % 2:
            steal = steal.next
        curr_elf = curr_elf.next
        nr_elves -= 1

    return curr_elf


print("Day 19.1: %s" % solve_1(3004953))
print("Day 19.2: %s" % solve_2(3004953))
