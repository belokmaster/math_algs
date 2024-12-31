def fib(n):
    if n == 0: 
        return 0 
    elif n == 1: 
        return 1 
    
    a, b = 0, 1 
    if n > 0: 
        for i in range(2, n + 1): 
            a, b = b, a + b 
        return b 
    else: 
        for i in range(-1, n - 1, -1): 
            a, b = b, a - b 
        return a 


n = int(input())
print(fib(n))
