import fileinput

# UP = 0, RIGHT = 1, DOWN = 2, LEFT = 3
move = {
    0: lambda t: (t[0], t[1] - 1),
    1: lambda t: (t[0] + 1, t[1]),
    2: lambda t: (t[0], t[1] + 1),
    3: lambda t: (t[0] - 1, t[1])
}

intersection = {
    0: lambda d: (d - 1) % 4,
    1: lambda d: d,
    2: lambda d: (d + 1) % 4
}

tracks = []

class Cart(object):
    def __init__(self, x, y, direction):
        self.pos = (x, y)
        self.direction = direction
        self.intersections = 0
        self.destroyed = False

    def __repr__(self):
        return "Cart[%s, %d]" % (self.pos, self.direction)

    def move(self, tracks):
        
        self.pos = move[self.direction](self.pos)
        t = tracks[self.pos[1]][self.pos[0]]
        if t == '|' and self.direction % 2 == 1:
            print("ERROR %s wrong direction" % self)
        
        if t == '-' and self.direction % 2 == 0:
            print("ERROR %s wrong direction" % self)
        
        # print(self, t)
        if t == '+':
            self.direction = intersection[self.intersections % 3](self.direction)
            self.intersections += 1
        elif t == '\\':
            self.direction = ((self.direction - 1) if self.direction % 2 == 0 else (self.direction + 1)) % 4
        elif t == '/':
            self.direction = ((self.direction + 1) if self.direction % 2 == 0 else (self.direction - 1)) % 4
    
    def destroy(self):
        self.destroyed = True

carts = []

y = 0
for line in fileinput.input():
    l = [c for c in line.rstrip('\n')]
    for x,c in enumerate(l):
        if c == '^':
            carts.append(Cart(x, y, 0))
            l[x] = '|'
        elif c == '<':
            carts.append(Cart(x, y, 3))
            l[x] = '-'
        elif c == 'v':
            carts.append(Cart(x, y, 2))
            l[x] = '|'
        elif c == '>':
            carts.append(Cart(x, y, 1))
            l[x] = '-'
    tracks.append(l)
    y += 1

print('\n'.join([''.join(t) for t in tracks]))

print(sorted(carts, key=lambda c: (c.pos[1], c.pos[0])))
time = 0
while len(carts) > 1:
    pos = dict([(c.pos, c) for c in carts])
    for c in sorted(carts, key=lambda c: (c.pos[1], c.pos[0])):
        if c.destroyed:
            continue
        prev_pos = c.pos
        c.move(tracks)
        if c.pos not in pos:
            pos[c.pos] = c
            del pos[prev_pos]
        else:
            c.destroy()
            carts.remove(c)
            o = pos[c.pos]
            o.destroy()
            del pos[c.pos]
            carts.remove(o)
            print("Boom!", time, c , o)
    time += 1
    if time == 65058:
        print(carts)

# carts[0].move(tracks)
print(carts[0].pos)
    