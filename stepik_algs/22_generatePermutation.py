import itertools

n = int(input()) 
sequence = range(n) 
permutations = itertools.permutations(sequence) 
for perm in permutations: 
    print(''.join(map(str, perm))) 