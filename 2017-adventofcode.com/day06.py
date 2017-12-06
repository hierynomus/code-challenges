with open('day06.in', 'r') as f:
    inp = list(map(int, f.readline().split()))

seen = {}
mem = inp[:]
count = 0
l = len(mem)
r = range(l)
s = ",".join(map(str, mem))
while s not in seen:
    seen[s] = count
    idx = max(r, key=mem.__getitem__)
    v = mem[idx]
    mem[idx] = 0
    d, m = v // l, v % l
    idx += 1
    for i in r:
        mem[(idx + i) % l] += d + (1 if m > 0 else 0)
        m -= 1
    count += 1
    s = ",".join(map(str, mem))

print("Day 6.1: ", count)
print("Day 6.2: ", count - seen[s])
