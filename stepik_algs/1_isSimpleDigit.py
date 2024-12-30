def isSimpleDigit(x):
    if x == 1:
        return 1
    if x == 2:
        return "простое"
    for i in range(3, int(x ** 0.5) + 1, 2):
        if x % i == 0:
            return "составное"
    return "простое"


x = int(input())
print(isSimpleDigit(x))