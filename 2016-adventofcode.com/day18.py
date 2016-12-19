inp = ".^^^.^.^^^^^..^^^..^..^..^^..^.^.^.^^.^^....^.^...^.^^.^^.^^..^^..^.^..^^^.^^...^...^^....^^.^^^^^^^"


def trap_or_safe(x, y):
    return '^' if x != y else '.'


def next_line(line):
    nl = ""
    for i in range(len(line)):
        if i == 0:
            nl += trap_or_safe('.', line[i + 1])
        elif i == len(line) - 1:
            nl += trap_or_safe(line[i - 1], '.')
        else:
            nl += trap_or_safe(line[i - 1], line[i + 1])
    return nl


def solve(nr):
    line = inp
    count = line.count('.')
    for _ in range(nr - 1):
        line = next_line(line)
        count += line.count('.')
    return count


print("Day 18.1: %s" % solve(40))
print("Day 18.2: %s" % solve(400000))
