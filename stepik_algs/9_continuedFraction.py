def continuedFraction(a, b):
    coeff = []
    while b != 0:
        quotient = a // b
        coeff.append(quotient)
        a, b = b, a % b
    return coeff


a, b = map(int, input().split())
result = continuedFraction(a, b)
result_str = "[" + str(result[0]) + ";" + ",".join(map(str, result[1:])) + "]"
print(result_str)