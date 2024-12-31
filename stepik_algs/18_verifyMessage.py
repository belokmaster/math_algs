def verifyMessage(e, n, x, y):
    calculated_x = pow(y, e, n)
    
    if calculated_x == x:
        return "Сообщение правильное"
    else:
        return "Сообщение фальсифицировано"


e = int(input())
n = int(input())
x = int(input())
y = int(input())

print(verifyMessage(e, n, x, y))