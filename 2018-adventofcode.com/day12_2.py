import fileinput

start = []
rules = {}

def evolve(generation, left_idx):
    next_gen = ''
    if generation[0:2] == '#.':
        next_gen += '#'
        left_idx -= 1
    next_gen += rules['..' + generation[0:3]]
    next_gen += rules['.' + generation[0:4]]
    for i in range(2, len(generation) - 2):
        next_gen += rules[generation[i - 2 : i + 3]]
    next_gen += rules[generation[-4:] + '.']
    next_gen += rules[generation[-3:] + '..']
    if generation[-2:] == '##':
        next_gen += '#'
    return next_gen, left_idx


for line in fileinput.input():
    if 'initial state:' in line:
        start = line.rstrip('\n').split(':')[1].strip()
    elif '=>' in line:
        f, t = line.rstrip('\n').split(' => ')
        rules[f] = t

generation = start
left_idx = 0
for i in range(20):
    generation, left_idx = evolve(generation, left_idx)
        
r_100 = 0
for i in range(len(generation)):
    r_100 += (i + left_idx) if generation[i] == '#' else 0

print(r_100 + (50000000000 - 100) * generation.count('#'))