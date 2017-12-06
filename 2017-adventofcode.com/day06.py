with open('day06.in', 'r') as f:
    inp = list(map(int, f.readline().split()))

seen = {}
mem = inp[:]
count = 0
mem_len = len(mem)
r = range(mem_len)
t = tuple(mem)
while t not in seen:
    seen[t] = count
    idx = max(r, key=mem.__getitem__)
    v = mem[idx]
    mem[idx] = 0
    d, m = divmod(v, mem_len)
    idx += 1
    for i in r:
        mem[(idx + i) % mem_len] += d + (1 if m > 0 else 0)
        m -= 1
    count += 1
    t = tuple(mem)

print("Day 6.1: ", count)
print("Day 6.2: ", count - seen[t])
