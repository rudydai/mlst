'''
Algorithm to randomly generate combination of nodes and test in
polynomial time if they are an MLST. Code written in Python,
Peter Gao has volunteered to translate this code into Go.


'''
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

