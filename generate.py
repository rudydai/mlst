from random import random
if __name__ == '__main__':
    for i in range(1,51):
        print '1 ' + str(i)
        print '2 ' + str(i)
        print '3 ' + str(i)
    for i in range(51,101):
        a = int(random()*2)
        if a:
            print str(i) + ' ' + str(i - 50)
        a = int(random()*2)
        if a:
            print str(i) + ' ' + str(i - 1)

