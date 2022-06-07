teste = 1
while True:
    qtd_meteoros = 0
    x1, y1, x2, y2 = map(int, input().split())
    if x1==x2==y1==y2:
        break
    else:
        qtd = int(input())
        for i in range(qtd):
            x,y = map(int, input().split())
            if x>=x1 and x<=x2 and y<=y1 and y>=y2:
                qtd_meteoros += 1
        print('Teste {}'.format(teste))
        print('{}'.format(qtd_meteoros))
        teste += 1