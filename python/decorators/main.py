from time import time 

def timer(func):
    def f(*args, **kwargs):
        before = time()
        res = func(*args, **kwargs)
        after = time()
        print('elapsed: ', after - before)
        return res
    return f

def add(x, y):
    return x + y

add = timer(add)

@timer
def sub(x, y):
    return x - y

print(add(1,2))
print(sub(1,2))