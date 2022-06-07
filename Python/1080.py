maior = 0
pos = 0

for i in range(100):
    
    n = int(input())
    
    if n > maior:
        maior = i
        pos = i

print(f'{pos}')
print(pos)