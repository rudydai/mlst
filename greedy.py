import sys

def solve():
	d = {} #graph representation: if i is index for node, d[i] is list of nodes attached to it: ie adjacency list
	s1 = set([]) #nodes attached by the solution-in-progress
	s2 = set([]) #list of unique nodes in graph
	edges = set([])
	i = 0;

	def solveGraph(d, s1, s2):
		edges = set([])

		while len(s1) < len(d):
			max = 0
			temp = None
		
			if len(s1) == 0:
				for x in d:
					if len(d[x]) > max:
						max = len(d[x])
						temp = x
			else:
				for x in s1:
					if len(d[x]) > max:
						max = len(d[x])
						temp = x

			s1.add(temp)
			for y in d[temp]:
				s1.add(y)
				edges.add(str(temp) + " " +  str(y))
				if temp in d[y]:
					d[y].remove(temp)
				for v in d:
					if y in d[v]:
						d[v].remove(y)
					if temp in d[v]:
						d[v].remove(temp)

			d[temp] = []
		return edges

	f = open(sys.argv[1], 'r')
	g  = open(sys.argv[2] , 'w')

	numGraphs = int(f.readline())
	g.write(str(numGraphs) + '\n')

	for _ in range(numGraphs):
		d = {}
		s1 = set([])
		s2 = set([])

		numEdges = int(f.readline())
		for _ in range(numEdges):
			edgeLine = f.readline()
			b = edgeLine.split()
			if b[0] in s2:
				d[b[0]].append(b[1])
			else:
				s2.add(b[0])
				d[b[0]] = [b[1]]

		edges = solveGraph(d, s1, s2)
		g.write(str(len(edges)) + "\n")
		for e in edges:
			g.write(str(e) + '\n')

	g.close()

if len(sys.argv) < 3:
	print "error.  need 2 arguments: first arg is input file, second arg is output file"
else:
	solve()