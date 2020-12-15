import re

testData = open(r'C:\Users\gmgal\go\src\github.com\gmgale\AdventofCode2020\Day_14\testData.txt',"r")

memory = []
for i in range(35):
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
            m = d[7:-2]
            for i in range(len(m)):
                mask.append(m[i])
            break


        address = 0
        out = re.compile('\[.\]')
        mr = out.search(d)
        if mr:
            address = str(mr.group(0))
            address = address[1: len(address) -1]
        address = int(address)

        value = 0
        out = re.compile('=.+')
        mr = out.search(d)
        if mr:
            value = str(mr.group(0))
            valueArr = str(bin(int(value[1:]))).split()

            m = []
            for i in range((37)-(len(value)-2)):
                m.append(0)

            for i in range(2, len(value)):
                m.append(int(value[i]))
            value = m[2:]

        # Now apply Mask to Value and save to memory[address]
        for i in range(len(value)):
            if mask[i] == "X":
                memory[i] = int(value[i])
                continue
            if mask[i] == value[i]:
                print("Copy")
                memory[i] = int(value[i])
                continue
            if mask[i] != value[i]:
                print("mask")
                memory[i] = int(mask[i])
                continue
            
        print(memory)


