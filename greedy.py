d = {}
s1 = set([])
s2 = set([])
edges = set([])
i =0;
f = open('temp2', 'r')
for line in f:
	if i >1:
		b = line.split()
		if b[0] in s2:
			d[b[0]].append(b[1])
		else:
			s2.add(b[0])
			d[b[0]] = [b[1]]
	i+=1
	

while len(s1) <100:
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
	
g  = open('hard2.out' , 'w')
g.write('1\n')

g.write(str(len(edges)) + "\n")
for e in edges:
	g.write(str(e) + '\n')
g.close()


		
	