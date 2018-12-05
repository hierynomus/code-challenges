import fileinput

input = fileinput.input().readline().strip()
l = len(input)
result_lens = []
for c in range(0, 26):
    result = []
    i = 0
    r = -1
    to_remove = [97 + c, 65 + c]
    while (i < l):
        if ord(input[i]) in to_remove:
            i += 1
        elif not result:
            result.append(input[i])
            r += 1
            i += 1
        elif abs(ord(result[r]) - ord(input[i])) != 32:
            result.append(input[i])
            # print(result)
            i += 1
            r += 1
        else:
            # print("destroy", input[i], result[r])
            del result[r]
            i += 1
            r -= 1
    result_lens.append(len(result))
print(min(result_lens))
        