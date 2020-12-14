import re

testData = open(r'C:\Users\gmgal\go\src\github.com\gmgale\AdventofCode2020\Day_14\testData.txt',"r")

memory = []
for i in range(2^32):
    memory.append(0)

mask = []

EOF = False

while EOF != True:
    while True:
        d = testData.readline()
        if not d:
            EOF = True
            break

        
        if d[0:4] == "mask":
            m = d[7:]
            for i in range(len(m)):
                mask.append(m[i])
            break

        address = 0
        out = re.compile('\[.\]')
        mr = out.search(d)
        if mr:
            address = str(mr.group(0))
            address = address[1: len(address) -1]

        value = 0
        out = re.compile('=.+')
        mr = out.search(d)
        if mr:
            value = str(mr.group(0))
            value = str(bin(int(value[2:])))

            m = []
            for i in range((2^32)-(len(value)-2)):
                m.append(0)

            for i in range(2, len(value)):
                m.append(int(value[i]))

        
    
        

        


