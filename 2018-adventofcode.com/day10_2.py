import fileinput

class Light(object):
    def __init__(self, line):
        self.sx, self.sy = map(int, line.rstrip('\n')[10:24].split(','))
        self.dx, self.dy = map(int, line.rstrip('\n')[-7:-1].split(','))

    def at_time(self, t):
        return (self.sx + t * self.dx, self.sy + t * self.dy)
        
    def __repr__(self):
        return str((self.x, self.y))

lights = [Light(l) for l in fileinput.input()]

min_box = None
time = 0
while True:
    time += 1
    t_lights = [l.at_time(time) for l in lights]
    min_x = min([x[0] for x in t_lights])
    max_x = max([x[0] for x in t_lights])
    min_y = min([x[1] for x in t_lights])
    max_y = max([x[1] for x in t_lights])
    if not min_box or min_box[1] > ((max_x - min_x) * (max_y - min_y)):
        min_box = (time, (max_x - min_x) * (max_y - min_y), max_x, max_y)
    else:
        break

print(min_box[0])