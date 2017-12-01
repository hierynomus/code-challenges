with open('day01.in', 'r') as f:
    s = f.readline().strip()
    out1 = sum([int(s[i]) if s[i - 1] == s[i] else 0 for i in range(0, len(s))])
    print("Day 1.1: ", out1)
    s2 = s + s
    steps = int(len(s) / 2)
    out2 = sum([int(s2[i]) if s2[i] == s2[i + steps] else 0 for i in range(0, len(s))])
    print("Day 2.2: ", out2)
