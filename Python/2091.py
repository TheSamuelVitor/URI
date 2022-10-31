#Numero Solitario

while True:
    tamanhoLista = int(input())
    if tamanhoLista == 0:
        break
    else:
        elementosLista = input().split()
        lista = []
        for x in range(tamanhoLista):
            lista.append(int(elementosLista[x]))
        lista.sort()
        x = 0
        while x <= tamanhoLista-1:
            if x == tamanhoLista-1:
                print(lista[x])
                break
            elif lista[x] == lista[x+1]:
                x += 2
            else:
                print(lista[x])
                break