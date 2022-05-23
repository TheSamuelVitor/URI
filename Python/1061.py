a, di = input().split()
di = int(di)
hi,mi,si = input().split(':')

b, df = input().split()
df = int(df)
hf,mf,sf = input().split(':')

hi = int(hi)
mi = int(mi)
si = int(si)

hf = int(hf)
mf = int(mf)
sf = int(sf)

if df > di and hf > hi and mf > mi and sf > si:

    dd = df-di
    dh = hf-hi
    dm = mf-mi
    ds = sf-si



print(f'{dd} dia(s)')
print(f'{dh} hora(s)')
print(f'{dm} minuto(s)')
print(f'{ds} segundo(s)')