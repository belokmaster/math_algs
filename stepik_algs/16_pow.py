def pow(base, exponent, mod):
	result = 1
	base = base % mod 
	while exponent > 0:
		if exponent % 2 == 1:
			result = (result * base) % mod
		base = (base * base) % mod   
		exponent //= 2
	return result


base, exponent, mod = map(int, input().split()) 
print(pow(base, exponent, mod)) 