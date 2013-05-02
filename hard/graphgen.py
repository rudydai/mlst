import random

maxNumEdges = 2000
maxNumNodes = 100

f = open('temp2', 'w')

edges = []
# for node in range(maxNumNodes):
# 	for other in range(maxNumNodes):
# 		if (node, other) not in edges and (other, node) not in edges and other != node:
# 			edges.append((node, other))
# 		if len(edges) == maxNumEdges:
# 			break
# 	if len(edges) == maxNumEdges:
# 		break

for node in range(maxNumNodes):
	for _ in range(int(maxNumEdges/maxNumNodes)):
		other = random.randint(0, maxNumNodes - 1)
		while (node, other) in edges or (other, node) in edges or other == node:
			other = random.randint(0, maxNumNodes - 1)
		edges.append((node, other))
		if len(edges) == maxNumEdges:
			break
	if len(edges) == maxNumEdges:
		break

f.write('1\n')
f.write(str(len(edges)) + '\n')
for edge in edges:
	f.write(str(edge[0]) + ' ' + str(edge[1]) + '\n')
f.close()