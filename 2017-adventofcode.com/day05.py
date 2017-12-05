mem = []
with open('day05.in', 'r') as f:
    for l in f.readlines():
        mem.append(int(l.strip()))

orig_mem = mem[:]
idx, count = 0, 0
mem_len = len(mem)
while idx < mem_len and idx >= 0:
    jmp = mem[idx]
    count += 1
    mem[idx] += 1
    idx += jmp

print("Day 5.1: ", count)

mem = orig_mem[:]
count, idx = 0, 0
while idx < mem_len and idx >= 0:
    jmp = mem[idx]
    count += 1
    mem[idx] += 1 if jmp < 3 else -1
    idx += jmp

print("Day 5.2: ", count)
