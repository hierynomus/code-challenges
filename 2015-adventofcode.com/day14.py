
class Reindeer:
    def __init__(self, name, speed, fly_time, rest_time):
        self.name = name
        self.speed = speed
        self.fly_time = fly_time
        self.rest_time = rest_time

    def distance_travelled(self, time):
        distance = 0
        flying = True
        while time > 0:
            if flying and time >= self.fly_time:
                time -= self.fly_time
                distance += self.speed * self.fly_time
                flying = False
            elif flying and time < self.fly_time:
                distance += self.speed * time
                break
            elif not flying and time >= self.rest_time:
                time -= self.rest_time
                flying = True
            else:
                break
        return distance

reindeer = []

with open('day14.in', 'r') as f:
    for l in f:
        parts = l.strip().split()
        reindeer.append(Reindeer(parts[0], int(parts[3]), int(parts[6]), int(parts[13])))

print("1: %s" % max([r.distance_travelled(2503) for r in reindeer]))

points = {}
for r in reindeer:
    points[r.name] = 0

for i in range(1, 2504):
    r = max([r for r in reindeer], key=lambda x: x.distance_travelled(i))
    points[r.name] += 1

print("2: %s" % max([points[r.name] for r in reindeer]))
