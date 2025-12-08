import itertools
import math
import networkx as nx

junctions = [tuple(map(int, ln.split(","))) for ln in iter(open(0).readline, "")]

edges = iter(sorted(
    itertools.combinations(range(len(junctions)), 2),
    key=lambda ij: sum((ci - cj)**2 for ci, cj in zip(junctions[ij[0]], junctions[ij[1]])))
)

graph = nx.Graph(itertools.islice(edges, 1000))
graph.add_nodes_from(range(len(junctions)))
part1 = math.prod(sorted(map(len, nx.connected_components(graph)), reverse=True)[:3])
print(part1)

for i, j in edges:
    graph.add_edge(i,j)
    if nx.is_connected(graph):
        print(junctions[i][0]*junctions[j][0])
        break
