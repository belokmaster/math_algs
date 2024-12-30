def solveDiophantine(a, b, c):
    for x in range(1, 100000):
        if (c - a * x) % b == 0:
            y = (c - a * x) // b
            return f"{x} {y}"
    return "Нет решений"


a, b, c = map(int, input().split())
print(solveDiophantine(a, b, c))