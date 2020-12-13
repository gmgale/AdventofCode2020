import requests
import http.cookies
import pyCookies as sec

taskURL = "https://adventofcode.com/2020/day/3/input"
# smallList = [
# "..##.......",
# "#...#...#..",
# ".#....#..#.",
# "..#.#...#.#",
# ".#...##..#.",
# "..#.##.....",
# ".#.#.#....#",
# ".#........#",
# "#.##...#...",
# "#...##....#",
# ".#..#...#.#"]

raw_cookie = 'session='+sec.sessionID; 'domain='+sec.aocDomain
simple_cookie = http.cookies.SimpleCookie(raw_cookie)
cookie_jar = requests.cookies.RequestsCookieJar()
cookie_jar.update(simple_cookie)
r = requests.get(taskURL, cookies=cookie_jar)

respData = str(r.content)
respData = respData.strip("b'") # Strip start of content body
data = respData.split("\\n")
data = data[:-1]    # Strip end line

x = 0                             
total = 0                         
map_width = len(data[0])
map_height = len(data)

print(data)

for y in range(map_height): 
    if data[y][x] == '#':         
        total += 1              
    x = (x + 3) % map_width

print("The answer is : " + str(total))

