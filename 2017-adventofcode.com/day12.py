import networkx as nx

G = nx.Graph()

with open('day12.in', 'r') as f:
    for l in f:
        s, ds = l.strip().split(' <-> ')
        for d in ds.split(', '):
            G.add_edge(int(s), int(d))
            G.add_edge(int(d), int(s))

print("Day 12.1:", len(nx.node_connected_component(G, 0)))
print("Day 12.2:", len(list(nx.connected_components(G))))
