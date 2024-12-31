def fastPowMod(base, exponent, modulus): 
    result = 1 
    base = base % modulus 
    while exponent > 0: 
        if exponent % 2 == 1: 
            result = (result * base) % modulus 
        exponent = exponent // 2 
        base = (base * base) % modulus 
    return result 


n = int(input()) 
numbers = list(map(int, input().split())) 

base = numbers[0] 
exponents = numbers[1:-1] 
modulus = numbers[-1] 
result = fastPowMod(base, max(exponents), modulus) 
print(result) 
