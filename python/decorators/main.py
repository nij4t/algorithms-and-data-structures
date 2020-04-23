from time import time 

def timer(func):
    def wrapper(*args, **kwargs):
        before = time()
        res = func(*args, **kwargs)
        after = time()
        print('elapsed:', after - before)
        return res
    return wrapper

def log(prefix):
    def decorator(fn):
        def wrapper(*args, **kwargs):
            res = fn(*args, **kwargs)
            print(prefix, str(res))
            return res
        return wrapper
    return decorator


def add(x, y):
    return x + y

add = timer(add)

@timer
def sub(x, y):
    return x - y

@log('log:')
def mul(x, y):
    return x*y

print(add(1,2))
print(sub(1,2))

mul(4,4)

