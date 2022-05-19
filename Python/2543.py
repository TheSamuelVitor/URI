qtd = 0

while(True):
    try:
        x, y = input().split()
        x = int(x)
        for i in range(x):
            a,b = input().split()
            if a == y and b == '0':
                qtd += 1
        print(qtd)

    except EOFError:
        break