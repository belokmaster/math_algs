import math

def countWords(k, n):
    if k > n:
        return 0
    
    combinations = math.factorial(n) // (math.factorial(k) * math.factorial(n - k))
    permutations = math.factorial(k)
    total_words = combinations * permutations
    return total_words


k, n = list(map(int, input().split()))
print(countWords(k, n))