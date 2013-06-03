import math

def times(f,n):
	tmp = 1
	for i in range(0,n):
		tmp *= math.factorial(f)
	return tmp
 
print times(5,6)
print math.factorial(30)
a =  math.factorial(100) / times(10,10)
b = str(a)
print len(b)
tmp = ""
for i in range(0,92):
	print i
	tmp += "0"

print b
print tmp

print 8*6


print 9.461 * 93
print 0.12 * 0.30 * 0.000001 
print 2.53 * 3.6
print 9.1 * math.pow(10,88) *4 / 3  /math.pi

print pow(9.1 * math.pow(10,88) *4 / 3  /math.pi, 1.0/3)