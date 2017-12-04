count_1 = 0
count_2 = 0
with open('day04.in', 'r') as f:
    for l in f.readlines():
        arr = l.strip().split()
        sorted_arr = ["".join(sorted(x)) for x in arr]
        if len(arr) == len(set(arr)):
            count_1 += 1
        if len(sorted_arr) == len(set(sorted_arr)):
            count_2 += 1

print("Day 4.1: ", count_1)
print("Day 4.2: ", count_2)
