import fileinput
import networkx as nx

directions = {
    'N': lambda x, y: (x, y - 1),
    'E': lambda x, y: (x + 1, y),
    'W': lambda x, y: (x - 1, y),
    'S': lambda x, y: (x, y + 1)
}

def expand(g, line, sx, sy):
    i = 0
    x, y = sx, sy
    while i < len(line):
        c = line[i]
        i += 1
        if c == '(':
            i += expand(g, line[i:], x, y)
        elif c == ')':
            return i
        elif c in 'NEWS':
            nx, ny = directions[c](x, y)
            g.add_edge((x, y), (nx, ny))
            x, y = nx, ny
        elif c == '|':
            x, y = sx, sy

for line in fileinput.input():
    input = line.rstrip('\n')
    g = nx.Graph()
    expand(g, input, 0, 0)
    print(sum([1 for _, d in nx.single_source_shortest_path_length(g, (0, 0)).items() if d >= 1000]))
