def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a


a, b = map(int, input().split())
if a == 0 and b == 0:
    print(-1)
else:
    print(gcd(a, b))

