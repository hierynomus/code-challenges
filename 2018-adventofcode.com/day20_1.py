import fileinput
import networkx as nx

directions = {
    'N': lambda x, y: (x, y - 1),
    'E': lambda x, y: (x + 1, y),
    'W': lambda x, y: (x - 1, y),
    'S': lambda x, y: (x, y + 1)
}


# def expand(g, line, sx, sy):
#     frontiers = [(sx, sy)]
#     x, y = sx, sy
#     i = 0
#     while i < len(line):
#         c = line[i]
#         i += 1
#         if c in 'NEWS':
#             print(c)
#             n_f = []
#             for x, y in frontiers:
#                 nx, ny = directions[c](x, y)
#                 g.add_edge((x, y), (nx, ny))
#                 print(g.edges)
#                 n_f.append((nx, ny))
#             frontiers = n_f
#         elif c == '(':
#             e = line.rfind(')')
#             frontiers = expand(g, line[i:e], x, y)
#             i += e
#         elif c == '|':
#             frontiers.append((x, y))
#             x, y = sx, sy



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
            print((x, y), '->', (nx, ny))
            g.add_edge((x, y), (nx, ny))
            x, y = nx, ny
        elif c == '|':
            x, y = sx, sy

for line in fileinput.input():
    input = line.rstrip('\n')
    g = nx.Graph()
    expand(g, input, 0, 0)
    print(max([d for _, d in nx.single_source_shortest_path_length(g, (0, 0)).items()]))
