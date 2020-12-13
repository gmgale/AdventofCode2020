import requests
import http.cookies
import pyCookies as sec

taskURL = "https://adventofcode.com/2020/day/3/input"

raw_cookie = 'session='+sec.sessionID; 'domain='+sec.aocDomain
simple_cookie = http.cookies.SimpleCookie(raw_cookie)
cookie_jar = requests.cookies.RequestsCookieJar()
cookie_jar.update(simple_cookie)
r = requests.get(taskURL, cookies=cookie_jar)

respData = str(r.content)
respData = respData.strip("b'") # Strip start of content body
data = respData.split("\\n")
data = data[:-1]    # Strip end line

slope =[ # 0 element = right, 1 = down
    [1, 1],
    [3, 1],
    [5, 1],
    [7, 1],
    [1, 2],
]

totalTrees = []

for i in range(len(slope)):
                             
    total = 0                         
    map_width = len(data[0])
    map_height = len(data)

    y = 0
    x = 0  
    while y < map_height:
        if data[y][x] == "#":
            total +=1
        
        x = (x + slope[i][0]) % map_width
        y = y + slope[i][1]

    totalTrees.append(total)

answer = 1

for i in range(len(totalTrees)):
    answer *= totalTrees[i]

print("The answer is", answer)