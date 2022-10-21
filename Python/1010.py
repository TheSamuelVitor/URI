l1 = input().split(" ")
l2 = input().split(" ")
cd1, qtd1, val1 = l1
cd2, qtd2, val2 = l2
total = (int(qtd1) * float(val1)) + (int(qtd2) * float(val2))
print("VALOR A PAGAR: R$ %0.2f" %total)