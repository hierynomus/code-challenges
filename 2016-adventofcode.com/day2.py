from collections import namedtuple

Key = namedtuple('Key', ['x', 'y'])

keypad_1 = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
sequence_1 = []
moves_1 = {
    'U': lambda k: Key(k.x, max(0, k.y - 1)),
    'D': lambda k: Key(k.x, min(2, k.y + 1)),
    'L': lambda k: Key(max(0, k.x - 1), k.y),
    'R': lambda k: Key(min(2, k.x + 1), k.y)
}

keypad_2 = {
    '1': {'D': '3'},
    '2': {'R': '3', 'D': '6'},
    '3': {'U': '1', 'L': '2', 'D': '7', 'R': '4'},
    '4': {'L': '3', 'D': '8'},
    '5': {'R': '6'},
    '6': {'L': '5', 'U': '2', 'R': '7', 'D': 'A'},
    '7': {'L': '6', 'U': '3', 'R': '8', 'D': 'B'},
    '8': {'L': '7', 'U': '4', 'R': '9', 'D': 'C'},
    '9': {'L': '8'},
    'A': {'U': '6', 'R': 'B'},
    'B': {'L': 'A', 'U': '7', 'R': 'C', 'D': 'D'},
    'C': {'U': '8', 'L': 'B'},
    'D': {'U': 'A'}
}
sequence_2 = []


with open('day2.in', 'r') as f:
    key_1 = Key(1, 1)
    key_2 = '5'
    for l in f:
        for c in l.strip():
            key_1 = moves_1[c](key_1)
            key_2 = keypad_2[key_2][c] if c in keypad_2[key_2] else key_2
        sequence_1.append(keypad_1[key_1.y][key_1.x])
        sequence_2.append(key_2)

print("Day 2.1: %s" % ''.join([str(i) for i in sequence_1]))
print("Day 2.2: %s" % ''.join(sequence_2))
