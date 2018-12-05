import fileinput

input = fileinput.input().readline().strip()
result = []
i = 0
r = -1
l = len(input)
while (i < l):
    if not result:
        result.append(input[i])
        r += 1
        i += 1
    if abs(ord(result[r]) - ord(input[i])) != 32:
        result.append(input[i])
        # print(result)
        i += 1
        r += 1
    else:
        # print("destroy", input[i], result[r])
        del result[r]
        i += 1
        r -= 1

print(len(result))
        