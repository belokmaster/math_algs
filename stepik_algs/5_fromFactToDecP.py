def fact(n):
    ans = 1
    for i in range(1, n + 1):
        ans *= i
    return ans


def fromFactToDecP(num):
    num = str(num)
    ans = 0
    for i in range(len(num)):
        dif = len(num) - i
        ans += int(num[i]) * fact(dif)
    return ans


x = int(input())
print(fromFactToDecP(x))