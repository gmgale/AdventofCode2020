#numbers = {1:[1], 20:[2], 8:[3], 12:[4], 0:[5], 14:[6]}
#numbers = {0:[1], 3:[2], 6:[3]}
#numbers = {1:[1], 3:[2], 2:[3], 0:[4]}
#numbers = {2:[1], 1:[2], 3:[3], 0:[4]}
#numbers = {1:[1], 2:[2], 3:[3], 0:[4]}
#numbers = {2:[1], 3:[2], 1:[3], 0:[4]}
#numbers = {3:[1], 2:[2], 1:[3], 0:[4]}
#numbers = {3: [1], 1:[2], 2:[3], 0:[4]}
numbers = {1:[1], 17:[2], 0:[3], 10:[4], 18:[5], 11:[6], 6:[7]}
N = 6
for i in range(len(numbers.keys())+1,30000001): 
    if len(numbers[N]) == 1:
        N = 0
        numbers[N] = [numbers[N][-1], i] 
    else:
        N = numbers[N][1] - numbers[N][0]
        if N in numbers.keys():
            numbers[N] = [numbers[N][-1], i]
        else:
            numbers[N] = [i]
print(N)