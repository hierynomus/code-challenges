import fileinput
import networkx as nx
import sys

def neighbours(x, y):
    yield (x, y - 1)
    yield (x - 1, y)
    yield (x + 1, y)
    yield (x, y + 1)

class Combatant(object):
    def __init__(self, x, y, power, cave):
        self.x = x
        self.y = y
        self.hp = 200
        self.power = power
        self.cave = cave

    def open_spots(self):
        return [n for n in neighbours(self.x, self.y) if n in self.cave]

    def pos(self):
        return (self.x, self.y)

    def is_enemy(self, o):
        return False
    
    def enemies(self, players):
        for p in players:
            if p.alive() and self.is_enemy(p):
                yield p

    def in_range(self, players):
        n = [x for x in neighbours(self.x, self.y)]
        return [p for p in self.enemies(players) if p.pos() in n]

    def attack(self, p_range):
        # Order by HP, y, x to get lowest HP in reading order
        o = sorted(p_range, key=lambda p: (p.hp, p.y, p.x))[0]
        # print(self, ": Attacks", o)
        o.hp -= self.power
        if not o.alive():
            # print(o, "DIED")
            o.cave.add_node(o.pos())
            for n in o.open_spots():
                o.cave.add_edge(o.pos(), n)

    def move(self, players):
        targets = [n for p in self.enemies(players) for n in p.open_spots()]
        shortest_path = None
        for n in self.open_spots():
            sps = [(nx.shortest_path_length(self.cave, n, t), n, t) for t in targets if nx.has_path(self.cave, n, t)]
            if sps:
                sp = min(sps, key=lambda t: (t[0], t[1][1], t[1][0]))
                if not shortest_path or sp[0] < shortest_path[0]:
                    shortest_path = sp

        # Move and update cave system
        if shortest_path:
            # print(self, ": Move to", shortest_path[1])
            old = self.pos()
            self.cave.add_node(old)
            for n in neighbours(self.x, self.y):
                if n in self.cave:
                    self.cave.add_edge(old, n)
            self.x, self.y = shortest_path[1]
            self.cave.remove_node(self.pos())
        # else:
        #     print(self, "cannot move")

    def take_turn(self, players):
        if not [p for p in self.enemies(players)]:
            return False
        r = self.in_range(players)
        if not r:
            self.move(players)
            r = self.in_range(players)
        if r:
            self.attack(r)
        return True
    
    def __repr__(self):
        return type(self).__name__ + '[' + ','.join(map(str, [self.hp, self.x, self.y])) + ']'
    
    def alive(self):
        return self.hp > 0

class Elf(Combatant):
    def __init__(self, x, y, power, cave):
        super().__init__(x, y, power, cave)

    def is_enemy(self, o):
        return isinstance(o, Goblin)

class Goblin(Combatant):
    def __init__(self, x, y, cave):
        super().__init__(x, y, 3, cave)
    
    def is_enemy(self, o):
        return isinstance(o, Elf)


cave = []
for l in sys.stdin:
    cave.append([c for c in l.rstrip('\n')])

elf_attack = 4
while True:
    players = []
    g = nx.Graph()
    for y in range(1, len(cave) - 1):
        for x in range(1, len(cave[y]) - 1):
            c = cave[y][x]
            if c == '.':
                g.add_node((x, y))
                for n in neighbours(x, y):
                    if cave[n[1]][n[0]] == '.':
                        g.add_edge((x, y), n)
            if c == 'G':
                players.append(Goblin(x, y, g))
            elif c == 'E':
                players.append(Elf(x, y, elf_attack, g))

    elfs = len([p for p in players if isinstance(p, Elf)])

    round = 0
    x = True
    while x:
        # print(round)
        for p in sorted(players, key=lambda p: (p.y, p.x)):
            if p.alive():
                x = p.take_turn(players)
                players = [p for p in players if p.alive()]
                if not x:
                    break
        # print(players)
        if x:
            round += 1

    if elfs == len([p for p in players if isinstance(p, Elf)]):
        print(round * sum(map(lambda p: p.hp, players)))
        break
    else:
        elf_attack += 1
