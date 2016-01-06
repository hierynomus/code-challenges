
input = "1113122113"
expanded = input

def expand(s):
    last_char = s[0]
    count = 1
    new_s = ""
    for c in s[1:]:
        if c == last_char:
            count += 1
        else:
            new_s += str(count) + str(last_char)
            last_char = c
            count = 1

    new_s += str(count) + str(last_char)

    return new_s


for i in range(1, 41):
    expanded = expand(expanded)

print("1: %s" % len(expanded))

for i in range(1, 11):
    expanded = expand(expanded)
print("2: %s" % len(expanded))
