def phi(n):
    result = n
    p = 2
    while p * p <= n:
        if n % p == 0:
            while n % p == 0:
                n //= p
            result -= result // p
        p += 1
    if n > 1:
        result -= result // n
    return result


def minPhi(y):
    x = 1
    while True:
        if phi(x) == y:
            return x
        x += 1

        
n = int(input())
print(minPhi(n))



