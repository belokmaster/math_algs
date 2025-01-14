def isSimpleDigit(x):
    if x == 1:
        return 0
    if x == 2:
        return 1
    if x % 2 == 0:
        return 0
    
    for i in range(3, int(x ** 0.5) + 1, 2):
        if x % i == 0:
            return 0
    return 1


n = int(input())
ans = 0 
for i in range(2, n + 1):
    ans += isSimpleDigit(i)
print(ans)

