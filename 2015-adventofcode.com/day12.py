import json

total = 0
red_value = 0


def parse_and_add(s):
    global total
    i = int(s)
    total += i
    return i


def inspect_red_object(o):
    global red_value
    found_red = False
    obj_count = 0
    for k, v in o.items():
        if k == 'red' or v == 'red':
            found_red = True

        obj_count += count(v)

    if found_red:
        red_value += obj_count
        # Erase entire object
        return {'a': 0}
    else:
        return o


def count(o):
    if type(o) is int:
        return o
    elif type(o) is dict:
        total = 0
        for k, v in o.items():
            total += count(v)
        return total
    elif type(o) is list:
        return sum([count(x) for x in o])
    return 0


with open('day12.in', 'r') as f:
    json.load(f, object_hook=inspect_red_object, parse_int=parse_and_add)

print("Day 12.1: %s" % total)
print("Day 12.2: %s" % (total - red_value,))
