def grayCode(n):
    if n == 0:
        return ['']
    else:
        smaller_gray_codes = grayCode(n - 1)
        first_half = ['0' + code for code in smaller_gray_codes]
        second_half = ['1' + code for code in reversed(smaller_gray_codes)]
        return first_half + second_half

    
n = int(input())
codes = grayCode(n)
for code in codes:
    print(code)


