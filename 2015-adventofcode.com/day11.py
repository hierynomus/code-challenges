inp = "hepxcrrq"


def increment_ords(ords):
    i = len(ords) - 1
    while i >= 0:
        if ords[i] == 25:
            ords[i] = 0
            i -= 1
        elif ords[i] in [7, 10, 13]:
            # We need to skip 'o', 'i', 'l'
            ords[i] += 2
            break
        else:
            ords[i] += 1
            break
    return ords


def has_straight(password):
    for x, y, z in [(password[i], password[i + 1], password[i + 2]) for i in range(0, len(password) - 2)]:
        if y == x + 1 and z == y + 1:
            return True
    return False


def has_two_in_row(password):
    ch = password[0]
    found = set()
    for c in password[1:]:
        if c == ch:
            found.add(c)
        else:
            ch = c

    return len(found) > 1


new_input = [ord(c) - 97 for c in inp]
while not has_two_in_row(new_input) or not has_straight(new_input):
    new_input = increment_ords(new_input)

print("Day 11.1: %s" % "".join([chr(c + 97) for c in new_input]))

# Increment once, because it expired
new_input = increment_ords(new_input)
while not has_two_in_row(new_input) or not has_straight(new_input):
    new_input = increment_ords(new_input)

print("Day 11.2: %s" % "".join([chr(c + 97) for c in new_input]))
