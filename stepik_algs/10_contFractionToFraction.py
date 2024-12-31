def contFractionToFraction(n, coefficients):
    numerator = coefficients[-1] 
    denominator = 1
    for i in range(n - 1, -1, -1): 
        numerator, denominator = coefficients[i] * numerator + denominator, numerator 
    return numerator, denominator 


n = int(input()) 
coeff = list(map(int, input().split())) 
numerator, denominator = contFractionToFraction(n, coeff) 
print(numerator, denominator) 




