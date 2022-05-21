total = 0
qtd_coelhos = 0
qtd_ratos = 0
qtd_sapos = 0

n = int(input())
for i in range(n):
    qtd,tipo = input().split()
    total += int(qtd)
    if tipo == 'C':
        qtd_coelhos += int(qtd)
    elif tipo == 'R':
        qtd_ratos += int(qtd)
    elif tipo == 'S':
        qtd_sapos += int(qtd)

print(f'Total: {total} cobaias')
print(f'Total de coelhos: {qtd_coelhos}')
print(f'Total de ratos: {qtd_ratos}')
print(f'Total de sapos: {qtd_sapos}')
print(f'Percentual de coelhos: {(qtd_coelhos/total)*100:.2f} %')
print(f'Percentual de ratos: {(qtd_ratos/total)*100:.2f} %')
print(f'Percentual de sapos: {(qtd_sapos/total)*100:.2f} %')