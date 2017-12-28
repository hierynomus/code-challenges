from collections import namedtuple, defaultdict

TuringState = namedtuple('TuringState', ['source', 'target', 'when', 'write', 'move'])


def parse_state(state_name, f):
    cur_val = int(f.readline().strip()[-2])
    write_val = int(f.readline().strip()[-2])
    move = 1 if 'right' in f.readline() else -1
    next_state = f.readline().strip()[-2]
    return TuringState(state_name, next_state, cur_val, write_val, move)


states = {}
with open('day25.in', 'r') as f:
    start = f.readline().strip()[-2]
    runs = int(f.readline().strip().split()[-2])
    for line in f:
        line = line.strip()
        if line.startswith('In state'):
            state_name = line.split()[-1][:-1]
            d = {}
            states[state_name] = d
            state = parse_state(state_name, f)
            d[state.when] = state
            state = parse_state(state_name, f)
            d[state.when] = state

tape = defaultdict(int)
state = start
cursor = 0
for _ in range(runs):
    transition = states[state][tape[cursor]]
    tape[cursor] = transition.write
    cursor += transition.move
    state = transition.target

print("Day 25.1:", sum(tape.values()))
