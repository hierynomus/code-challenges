import fileinput
from collections import defaultdict

class Clock(object):
    def __init__(self):
        self.cur_time = 0
    
    def tick(self):
        self.cur_time += 1

    def __repr__(self):
        return "Clock[%d]" % self.cur_time


class Worker(object):
    def __init__(self, clock):
        self.work_time = 0
        self.item = None
        self.clock = clock
        pass
    
    def work_on(self, c):
        self.item = c
        self.work_time = self.clock.cur_time + ord(c) - 4

    def is_ready(self):
        return self.item is not None and self.work_time <= self.clock.cur_time
    
    def idle(self):
        return self.item is None

    def clear(self):
        self.item = None
        self.work_time = 0

    def __repr__(self):
        return "Worker[%d, %s]" % (max(0, self.work_time - self.clock.cur_time), self.item)

forward = defaultdict(list)
reverse = defaultdict(list)

for l in fileinput.input():
    s = l.split(' ')
    forward[s[1]].append(s[7])
    reverse[s[7]].append(s[1])

nxt_start = forward.keys() - reverse.keys()
next_nodes = set(nxt_start)
result = []

clock = Clock()
workers = [Worker(clock), Worker(clock), Worker(clock), Worker(clock), Worker(clock)]

for w in workers:
    if next_nodes:
        item = min(next_nodes)
        w.work_on(item)
        next_nodes.remove(item)

while not all([w.idle() for w in workers]):
    clock.tick()
    for aw in [w for w in workers if w.is_ready()]:
        ready_item = aw.item
        aw.clear()
        result.append(ready_item)
        nxt = forward[ready_item]
        next_nodes.update([n for n in nxt if n not in result and all([x in result for x in reverse[n]])])
        del forward[ready_item]
        for w in [x for x in workers if x.idle()]:
            if next_nodes:
                next_item = min(next_nodes)
                next_nodes.remove(next_item)
                w.work_on(next_item)

print(print(clock.cur_time))
