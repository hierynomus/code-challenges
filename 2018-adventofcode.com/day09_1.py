# players = 431
# last_marble = 70950
players = 9
last_marble = 25

marbles = [0]
scores = [0] * players
marble_idx = 0
player = 0
for marble in range(1, last_marble + 1):
    if marble % 23 == 0:
        scores[player] += marble
        marble_idx = (marble_idx - 8) % len(marbles) + 1
        scores[player] += marbles[marble_idx]
        marbles = marbles[:marble_idx] + marbles[marble_idx + 1:]
    else:
        marble_idx = (marble_idx + 1) % len(marbles) + 1
        marbles = marbles[:marble_idx] + [marble] + marbles[marble_idx:]
    player = (player + 1) % players
    print(marbles)

print(max(scores))