import math

def permutationNumber(n, perm):
    available = list(range(1, n + 1))
    number = 0
    
    for i in range(n):
        current = perm[i]
        position = available.index(current)
        number += position * math.factorial(n - i - 1)
        available.pop(position)
    
    return number


n = int(input()) 
perm = list(map(int, input().split())) 
print(permutationNumber(n, perm))
