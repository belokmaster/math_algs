def fib(x):
    a, b = 1, 1
    while a < x:
        a, b = b, a + b
        if a == x:
            return 1
    return 0


n = int(input())
print(fib(n))