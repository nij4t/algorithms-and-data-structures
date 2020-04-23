from lib import lru_cache

@lru_cache
def fib(num):
    if num < 3:
        return 1
    return fib(num-1)+fib(num-2)

print(fib(100))
