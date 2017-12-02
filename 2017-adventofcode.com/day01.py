with open('day01.in', 'r') as f:
    s = f.readline().strip()
    out1 = sum([int(s[i]) if s[i - 1] == s[i] else 0 for i in range(0, len(s))])
    print("Day 1.1: ", out1)
    out2 = sum([int(s[i]) if s[i] == (s + s)[i + (len(s) // 2)] else 0 for i in range(0, len(s))])
    print("Day 1.2: ", out2)
