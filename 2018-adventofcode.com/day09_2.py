players = 431
last_marble = 7095000

class Marble(object):
    def __init__(self, v):
        self.value = v
        self.right = self
        self.left = self

    def insert(self, m):
        right = self.right
        m.left = self
        m.right = right
        right.left = m
        self.right = m

    def next(self):
        return self.right

    def prev(self):
        return self.left

    def remove(self):
        left = self.left
        right = self.right
        left.right = right
        right.left = left
        return self.value

    def __repr__(self):
        return "Marble[" + str(self.left.value) + ", " + str(self.value) + ", " + str(self.right.value) + "]"

current_marble = Marble(0)
scores = [0] * players
player = 0

for marble in range(1, last_marble + 1):
    if marble % 23 == 0:
        scores[player] += marble
        current_marble = current_marble.left.left.left.left.left.left
        scores[player] += current_marble.remove()
    else:
        current_marble = current_marble.right.right
        m = Marble(marble)
        current_marble.insert(m)
    player = (player + 1) % players

print(max(scores))