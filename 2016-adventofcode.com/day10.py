bot_part_1 = None


class Output(object):
    """An output bin"""

    def __init__(self, idx):
        self.chips = []
        self.idx = idx

    def receive_chip(self, value):
        self.chips.append(value)


class Bot(object):

    def __init__(self, idx):
        self.chips = []
        self.rule = None
        self.idx = idx
        pass

    def receive_chip(self, value):
        global bot_part_1
        self.chips.append(value)
        self.chips.sort()
        if 17 in self.chips and 61 in self.chips and not bot_part_1:
            bot_part_1 = self.idx
        if len(self.chips) == 2:
            self.process()

    def assign_rule(self, lower, higher):
        self.rule = (lower, higher)
        if len(self.chips) == 2:
            self.process()

    def process(self):
        if self.rule:
            self.rule[0].receive_chip(self.chips[0])
            self.rule[1].receive_chip(self.chips[1])


bots = {}
outputs = {}


def get_actor(actor_idx, actor_type):
    if actor_type == "bot":
        if actor_idx not in bots:
            bots[actor_idx] = Bot(actor_idx)
        return bots[actor_idx]
    else:
        if actor_idx not in outputs:
            outputs[actor_idx] = Output(actor_idx)
        return outputs[actor_idx]


with open('day10.in', 'r') as f:
    for l in f:
        line = l.strip().split()
        if line[0] == "value":
            bot_nr, chip = map(int, (line[5], line[1]))
            get_actor(bot_nr, line[4]).receive_chip(chip)
        if line[0] == "bot":
            get_actor(int(line[1]), line[0]).assign_rule(get_actor(int(line[6]), line[5]), get_actor(int(line[11]), line[10]))

print("Day 10.1: %s" % bot_part_1)
print("Day 10.2: %s" % (outputs[0].chips[0] * outputs[1].chips[0] * outputs[2].chips[0]))
