f  = open('hard1.out' , 'w')

f.write('1\n')
f.write('99\n')
for x in range(1,100):
	f.write('0 %i\n' %x)
f.close()