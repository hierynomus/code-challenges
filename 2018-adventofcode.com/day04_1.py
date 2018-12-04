import fileinput
import operator
from collections import defaultdict, Counter

input = sorted([l.strip() for l in fileinput.input()])

guard_intervals = defaultdict(list)
guard_sleep = defaultdict(int)
for l in input:
    if 'Guard' in l:
        current_guard = l.split(' ')[3][1:]
    elif 'asleep' in l:
        sleepy_time = int(l.split(' ')[1][:-1].split(':')[1])
    elif 'wakes up' in l:
        wakey_time = int(l.split(' ')[1][:-1].split(':')[1])
        guard_intervals[current_guard].append(range(sleepy_time, wakey_time))
        guard_sleep[current_guard] += wakey_time - sleepy_time

max_sleep_mins = max(guard_sleep.items(), key=operator.itemgetter(1))[0]
c = Counter({})
for r in guard_intervals[max_sleep_mins]:
    c.update(r)

print(int(max_sleep_mins) * c.most_common(1)[0][0])
