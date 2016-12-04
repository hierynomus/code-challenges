import re
from collections import defaultdict

regexp = re.compile("(?P<room_id>[a-z-]+)-(?P<sector>[0-9]+)\[(?P<checksum>[a-z]{5})\]")
sector1 = 0
sector2 = 0
rooms = []


def real_room(room_id, checksum):
    d = defaultdict(int)
    for c in [c for c in room_id if c != '-']:
        d[c] += 1
    sorted_d = sorted(d.items(), key=lambda x: (-x[1], x[0]))
    check = ''.join([t[0] for t in sorted_d[:5]])
    return check == checksum


def room_name(room_id, sector):
    room_name = ''
    for c in room_id:
        if c == '-':
            room_name += ' '
        else:
            room_name += chr((((ord(c) - 97) + sector) % 26) + 97)
    return room_name


with open('day4.in', 'r') as f:
    for r in f:
        match = regexp.search(r)
        room = match.groupdict()
        if real_room(room["room_id"], room["checksum"]):
            sector1 += int(room["sector"])
            if 'northpole' in room_name(room['room_id'], int(room['sector'])):
                sector2 = int(room['sector'])


print("Day 4.1: %s" % sector1)
print("Day 4.2: %s" % sector2)
