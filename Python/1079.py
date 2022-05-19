x = int(input())
for i in range(x):
    a,b,c = input().split()
    a = float(a)
    b = float(b)
    c = float(c)
    print(f'{((a*2)+(b*3)+(c*5))/10:.1f}')