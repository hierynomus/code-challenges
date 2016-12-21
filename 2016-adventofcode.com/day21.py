from functools import partial
from itertools import permutations


def swap_letter(s, a, b):
    return [b if c == a else (a if c == b else c) for c in s]


def swap_pos(s, p1, p2):
    s[p1], s[p2] = s[p2], s[p1]
    return s


def rotate_left(s, pos):
    pos = pos % len(s)
    return s[pos:] + s[:pos] if pos > 0 else s


def rotate_right(s, pos):
    pos = pos % len(s)
    return s[-pos:] + s[:-pos] if pos > 0 else s


def reverse_pos(s, p1, p2):
    return ((s[:p1] if p1 > 0 else []) + list(reversed(s[p1:p2 + 1])) + (s[p2 + 1:] if p2 + 1 < len(s) else []))


def move_pos(s, p1, p2):
    s.insert(p2, s.pop(p1))
    return s


def rotate_letter(s, a):
    idx = s.index(a)
    return rotate_right(s, idx + 1) if idx < 4 else rotate_right(s, idx + 2)


def rev_rotate_letter(s, a):
    t = s[:]
    for i in range(0, len(s)):
        rot_left = rotate_left(t, i)
        rot_letter = rotate_letter(rot_left, a)
        if s == rot_letter:
            return rot_left
    raise Exception("could not find correct rotation for %s %s" % (s, a))


def solve(s, instructions):
    for instruction in instructions:
        s = instruction(s)
    return s


def solve_2(s, instructions):
    for t in permutations(s):
        if solve(list(t), instructions) == s:
            return t


instructions = []
reverse_instructions = []

with open('day21.in', 'r') as f:
    for line in [l.strip().split() for l in f]:
        if line[0] == 'rotate' and line[1] == 'right':
            instructions.append(partial(rotate_right, pos=int(line[2])))
            reverse_instructions.append(partial(rotate_left, pos=int(line[2])))
        elif line[0] == 'rotate' and line[1] == 'left':
            instructions.append(partial(rotate_left, pos=int(line[2])))
            reverse_instructions.append(partial(rotate_right, pos=int(line[2])))
        elif line[0] == 'rotate':
            instructions.append(partial(rotate_letter, a=line[6]))
            reverse_instructions.append(partial(rev_rotate_letter, a=line[6]))
        elif line[0] == 'swap' and line[1] == 'letter':
            instructions.append(partial(swap_letter, a=line[2], b=line[5]))
            reverse_instructions.append(partial(swap_letter, a=line[5], b=line[2]))
        elif line[0] == 'swap' and line[1] == 'position':
            instructions.append(partial(swap_pos, p1=int(line[2]), p2=int(line[5])))
            reverse_instructions.append(partial(swap_pos, p1=int(line[5]), p2=int(line[2])))
        elif line[0] == 'reverse':
            instructions.append(partial(reverse_pos, p1=int(line[2]), p2=int(line[4])))
            reverse_instructions.append(partial(reverse_pos, p1=int(line[2]), p2=int(line[4])))
        elif line[0] == 'move':
            instructions.append(partial(move_pos, p1=int(line[2]), p2=int(line[5])))
            reverse_instructions.append(partial(move_pos, p1=int(line[5]), p2=int(line[2])))

print("Day 21.1: %s " % ''.join(solve(list('abcdefgh'), instructions)))
print("Day 21.2: %s " % ''.join(solve(list('fbgdceah'), reverse_instructions[::-1])))
