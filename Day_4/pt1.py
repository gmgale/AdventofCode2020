testData = open(r'C:\Users\gmgal\go\src\github.com\gmgale\AdventofCode2020\Day_4\testData.txt',"r")

testStr = testData.read()
reqFeilds = ["byr", "iyr", "eyr", "hgt","hcl","ecl","pid"]

passports = testStr.split("\n\n")
testData.close()

total = 0

for i in range(len(passports)):
    data = passports[i].split(" ")
    feilds = []
    count = 0

    for j in range(len(data)): 
        if data[j][0:3] in reqFeilds:
            count +=1
            print(count)
            if count == len(reqFeilds):
                total +=1
            continue
        else:
            break


print(total)


