# -*-coding:utf-8 -*-
with open('data.txt', 'w') as f:
    li=[]
    for i in range(1,1000001):
        li.append(str(i)+','+'name'+str(i)+'\n')
    f.writelines(li)