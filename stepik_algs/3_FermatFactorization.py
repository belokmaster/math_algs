import math

def FermatFactorization(n):
    # Если число четное, сразу возвращаем 2 и n/2
    if n % 2 == 0:
        return [2, n // 2]
    
    # Начальное значение для x выбирается как ближайшее целое число,
	# которое больше или равно квадратному корню из n.
    x = math.ceil(math.sqrt(n))
    y2 = x * x - n         # Вычисляем y^2 = x^2 - n
    y = int(math.sqrt(y2)) # Находим целую часть квадратного корня из y^2
    
    # Пока y^2 не является точным квадратом, увеличиваем x и пересчитываем y
    while y * y != y2:
        x += 1 # увеличиваем иксик и пересчитываем игрик
        y2 = x * x - n
        y = int(math.sqrt(y2))
        
    # Возвращаем найденные множители, которые являются (x - y) и (x + y).
    return [x - y, x + y]

n = int(input())
arr = FermatFactorization(n)
print(min(arr), max(arr))




