def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a


def modInverse(a, m):
    m0, x0, x1 = m, 0, 1
    if gcd(a, m) != 1:
        return None  
    while a > 1:
        q = a // m
        a, m = m, a % m
        x0, x1 = x1 - q * x0, x0
    return x1 + m0 if x1 < 0 else x1


def chineseTheorem(n, c, m):
    M = 1
    for mi in m:
        M *= mi
    
    x = 0
    for i in range(n):
        Mi = M // m[i]
        yi = modInverse(Mi, m[i])
        if yi is None:
            raise ValueError("Модули не взаимно простые")  
        x += c[i] * Mi * yi
    
    return x % M


n = int(input())
c = list(map(int, input().split()))
m = list(map(int, input().split()))
result = chineseTheorem(n, c, m)
print(result)
