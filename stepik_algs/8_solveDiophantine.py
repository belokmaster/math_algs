'''
def solveDiophantine(a, b, c):
    for x in range(1, 100000):
        if (c - a * x) % b == 0:
            y = (c - a * x) // b
            return f"{x} {y}"
    return "Нет решений"
'''

def GcdExtended(a, b):
    if b == 0:
        return a, 1, 0
    
    g, x1, y1 = GcdExtended(b, a % b)
    x = y1
    y = x1 - (a // b) * y1
    
    return g, x, y


def solveDiophantine(a, b, c):
    g, x, y = GcdExtended(a, b)

    # Если c не делится на g, то уравнение ax + by = c не имеет целочисленных решений
    if c % g != 0:
        return "Нет решений"
    
    # Находим частное решение x0, y0 для уравнения ax + by = g
    x0 = x * (c // g)
    y0 = y * (c // g)
    
    # Находим минимальное значение k, чтобы x_current был положительным
    # условие, что x должно быть минимальным..
    k_min = -x0 // (b // g)
    
    # Перебираем значения k от k_min до k_min + |b|, чтобы найти первое положительное x_current
    for k in range(k_min, k_min + abs(b)):
        x_current = x0 + (b // g) * k
        if x_current > 0:
            y_current = y0 - (a // g) * k
            return f"{x_current} {y_current}"
            
    return "Нет решений"


a, b, c = map(int, input().split())
print(solveDiophantine(a, b, c))