numbers = {1:[1], 17:[2], 0:[3], 10:[4], 18:[5], 11:[6], 6:[7]}
#numbers = {0:[1], 3:[2], 6:[3]}
#numbers = {1:[1], 3:[2], 2:[3], 0:[4]}
#numbers = {2:[1], 1:[2], 3:[3], 0:[4]}
#numbers = {1:[1], 2:[2], 3:[3], 0:[4]}
#numbers = {2:[1], 3:[2], 1:[3], 0:[4]}
#numbers = {3:[1], 2:[2], 1:[3], 0:[4]}
#numbers = {3: [1], 1:[2], 2:[3], 0:[4]}
N = 6
for i in range(len(numbers.keys())+1, 2021): # range(numbers.keys())+1,2021) for first part
    if len(numbers[N]) == 1:
        N = 0
        numbers[N] = [numbers[N][-1], i] # only works if 0 is already in numbers, we add
    else:
        N = numbers[N][1] - numbers[N][0]
        if N in numbers.keys():
            numbers[N] = [numbers[N][-1], i]
        else:
            numbers[N] = [i]
print(N)