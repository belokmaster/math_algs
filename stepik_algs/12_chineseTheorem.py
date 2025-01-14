def modInverse(x, m):
    for i in range(m):
        if x*i % m == 1:
            return i
    return None


def chineseTheorem(n, c, m):
    M = 1
    for mi in m:
        M *= mi
    
    x = 0
    for i in range(n):
        Mi = M // m[i] # Вычисляем M_i как M, делённое на текущий модуль
        yi = modInverse(Mi, m[i]) # Находим обратный элемент M_i по модулю m[i]
        if yi is None:
            raise ValueError("Модули не взаимно простые")  
        x += c[i] * Mi * yi
    
    return x % M


n = int(input())
c = list(map(int, input().split()))
m = list(map(int, input().split()))
result = chineseTheorem(n, c, m)
print(result)
