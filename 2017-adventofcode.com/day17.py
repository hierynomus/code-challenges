inp = 366

circ_buf = [0, 1]
pos = 0
for i in range(2, 2018):
    new_pos = (pos + inp) % i
    circ_buf = circ_buf[:new_pos] + [i] + circ_buf[new_pos:]
    pos = new_pos + 1

print("Day 17.1:", circ_buf[pos])

p = 0
follow_0 = 1
for i in range(1, 50000001):
    p = ((p + inp) % i) + 1
    if p == 1:
        follow_0 = i

print("Day 17.2:", follow_0)
