import math 

def bernoulli(n, k, p): 
    c = math.factorial(n) // (math.factorial(k) * math.factorial(n - k)) 
    probability = c * (p ** k) * ((1 - p) ** (n - k)) 
    return round(probability, 2) 


n, k, p = map(float, input().split()) 
n, k = int(n), int(k) 
print(bernoulli(n, k, p)) 

