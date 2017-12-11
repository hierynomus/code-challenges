with open('day09.in', 'r') as f:
    inp = f.readline().strip()


def generator(inp):
    for c in inp:
        yield c


group = 0
group_sum = 0
garbage_count = 0
garbage = False
g = generator(inp)
for c in g:
    if '!' == c:
        next(g)  # Skip over next char
    elif '>' == c and garbage:
        garbage = False
    elif garbage:
        garbage_count += 1
    elif '<' == c:
        garbage = True
    elif '{' == c:
        group += 1
    elif '}' == c:
        group_sum += group
        group -= 1
    else:
        pass

print("Day 9.1:", group_sum)
print("Day 9.2:", garbage_count)
