import networkx as nx
from itertools import combinations, permutations

with open('day24.in', 'r') as f:
    maze = [l.strip() for l in f]

height, width = len(maze), len(maze[0])
grid = nx.grid_2d_graph(height, width)
locs = {}
for y in range(height):
    for x in range(width):
        c = maze[y][x]
        if c == '#':
            grid.remove_node((y, x))
        elif c.isdigit():
            locs[int(c)] = (int(c), (y, x))

shortest_paths = {}
for a, b in combinations(locs.values(), 2):
    shortest_paths[(a[0], b[0])] = shortest_paths[(b[0], a[0])] = nx.shortest_path_length(grid, a[1], b[1])

path_lengths = {}
for p in permutations(locs):
    if p[0] == 0:
        length = sum([shortest_paths[(p[i - 1], p[i])] for i in range(1, len(p))])
        path_lengths[p] = length

print("Day 24.1: %s" % min(path_lengths.values()))
print("Day 24.2: %s" % min([v + shortest_paths[(k[-1], 0)] for k, v in path_lengths.items()]))
