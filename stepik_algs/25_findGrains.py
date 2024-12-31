def find_grains(n): 
    if n <= 0: 
        return "0" 
    grains = [40, 60] 
    for i in range(2, n): 
        new_grain = sum(grains) / len(grains) 
        grains.append(new_grain) 
    return round(grains[n - 1])


try: 
    n = int(input()) 
    print(find_grains(n)) 
except ValueError: 
    print("0") 