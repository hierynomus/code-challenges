import fileinput
import operator
from collections import defaultdict, Counter

input = sorted([l.strip() for l in fileinput.input()])

guard_intervals = defaultdict(list)
for l in input:
    if 'Guard' in l:
        current_guard = l.split(' ')[3][1:]
    elif 'asleep' in l:
        sleepy_time = int(l.split(' ')[1][:-1].split(':')[1])
    elif 'wakes up' in l:
        wakey_time = int(l.split(' ')[1][:-1].split(':')[1])
        guard_intervals[current_guard].append(range(sleepy_time, wakey_time))

guard_mostcommon = {}
for guard, intervals in guard_intervals.items():
    c = Counter({})
    for i in intervals:
        c.update(i)
    guard_mostcommon[guard] = c.most_common(1)[0]

max_guard_counter = max(guard_mostcommon.items(), key=lambda i: i[1][1])
print(int(max_guard_counter[0]) * max_guard_counter[1][0])
# print(int(max_sleep_mins) * )
