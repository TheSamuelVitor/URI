dinheiro = float(input())

notas100 = dinheiro//100
notas50 = (dinheiro - (notas100*100)) // 50
notas20 = (dinheiro - (notas100*100) - (notas50*50)) // 20
notas10 = (dinheiro - (notas100*100) - (notas50*50) - (notas20*20)) // 10
notas5 = (dinheiro - (notas100*100) - (notas50*50) - (notas20*20) - (notas10*10)) // 5
notas2 = (dinheiro - (notas100*100) - (notas50*50) - (notas20*20) - (notas10*10) - (notas5*5)) // 2

dinheiro = (dinheiro - (notas100*100) - (notas50*50) - (notas20*20) - (notas10*10) - (notas5*5) - (notas2*2)) * 100

moedas100 = dinheiro//100
moedas50 = (dinheiro - (moedas100*100)) // 50
moedas25 = (dinheiro - (moedas100*100) - (moedas50*50)) // 25
moedas10 = (dinheiro - (moedas100*100) - (moedas50*50) - (moedas25*25)) // 10
moedas5 = (dinheiro - (moedas100*100) - (moedas50*50) - (moedas25*25) - (moedas10*10)) // 5
moedas1 = (dinheiro - (moedas100*100) - (moedas50*50) - (moedas25*25) - (moedas10*10) - (moedas5*5) )

print('NOTAS:')
print(f'{notas100:.0f} nota(s) de R$ 100.00')
print(f'{notas50:.0f} nota(s) de R$ 50.00')
print(f'{notas20:.0f} nota(s) de R$ 20.00')
print(f'{notas10:.0f} nota(s) de R$ 10.00')
print(f'{notas5:.0f} nota(s) de R$ 5.00')
print(f'{notas2:.0f} nota(s) de R$ 2.00')

print('MOEDAS:')
print(f'{moedas100:.0f} moeda(s) de R$ 1.00')
print(f'{moedas50:.0f} moeda(s) de R$ 0.50')
print(f'{moedas25:.0f} moeda(s) de R$ 0.25')
print(f'{moedas10:.0f} moeda(s) de R$ 0.10')
print(f'{moedas5:.0f} moeda(s) de R$ 0.05')
print(f'{moedas1:.0f} moeda(s) de R$ 0.01')

