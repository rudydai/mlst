'''
Algorithm to randomly generate combination of nodes and test in
polynomial time if they are an MLST. Code written in Python,
Peter Gao has volunteered to translate this code into Go.


'''
import random
from random import random

# Read input
d = {}
edges = []
i = 0
f = open('temp2', 'r')
for line in f:
    if i > 1:
        b = line.split()
        if b[0] in d.keys(): d[b[0]].append(b[1])
        else: d[b[0]] = [b[1]]

        if b[1] in d.keys(): d[b[1]].append(b[0])
        else: d[b[1]] = [b[0]]

        edges.append((b[0], b[1]))

    i += 1

def check_if_mlst():
    pass

# Randomly generate graphs and check if they are an MLST
while True:
    # Method 1: Randomly select an edge from each vertex
    mlst = []
    for node in d.keys():
        mlst.append((node, d[node][int(random()*len(d[node]))]))
    check_if_mlst()

    # Method 2: Randomly select n - 1 edges
    i = 0
    mlst = []
    while i < len(d) - 1:
        random_edge = int(random()*len(edges))
        if edges[random_edge] not in mlst:
            mlst.append(edges[random_edge])
            i += 1
    check_if_mlst()

    def dfs(node, mlst, visited):
        if node in visited:
            return

        visited.add(node)

        import random
        random.shuffle(d[node])
        connecting_nodes = d[node]

        for n in connecting_nodes:
            if n not in visited:
                mlst.append((node, n))
                dfs(n, mlst, visited)


    visited = set([])
    mlst = []
    import random as rand
    dfs(rand.sample(set(d.keys()), 1)[0], mlst, visited)
    print mlst

    def num_leaves(mlst):
        num_leaves = {}
        for edge in mlst:
            if edge[0] in num_leaves.keys():
                num_leaves[edge[0]] += 1
            else:
                num_leaves[edge[0]] = 1
            if edge[1] in num_leaves.keys():
                num_leaves[edge[1]] += 1
            else:
                num_leaves[edge[1]] = 1

        return len([x for x in num_leaves.values() if x == 1])

    def find_path(mlst, node, ending_node, visited=set([])):
        if node in visited:
            return None

        visited.add(node)

        for edge in mlst:
            # Base Case
            if edge[0] == node and edge[1] == ending_node:
                return ((edge[0], edge[1]),)

            path = find_path(mlst, edge[1], ending_node, visited)
            if edge[0] == node and path:
                return ((edge[0], edge[1]),) + path

        return None

    print num_leaves(mlst)
    print find_path(mlst, '1', '90')
    print '\n'
    print find_path(mlst, '90', '1')
    break

    '''
    while True:
        random_edge = int(random()*len(edges))
        if random_edge not in mlst or (random_edge[1], random_edge[0]) not in mlst:
            # Add the edge, forming a cycle, then remove one from the cycle
            mlst.append(random_edge) # Need to maybe append the reverse?
            num_leaves = num_leaves(mlst)
            remove_edge = random_edge

            for edge in find_cycle(mlst, edge):
                mlst.remove(edge)
                if num_leaves(mlst) > num_leaves
                    num_leaves = num_leaves(mlst)
                    remove_edge = edge

            mlst.remove(random_edge)
    '''


