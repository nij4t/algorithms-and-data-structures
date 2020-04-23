# lru_cache is a decorator that optimizes the performance of
# function calls by saving its result inside a dictionary
# 
# Example:
# 
# @lru_cache
# def fibonacci(n):
# ...
#
def lru_cache(config):
    # TODO: Implement configuration
    def decorator(fn):
        cache = {}
        def wrapper(*args, **kwargs):
            func_sig = fn.__name__,args
            if func_sig in cache:
                print('found in cache', func_sig, cache[func_sig])
                return cache[func_sig]
            cache[func_sig] = fn(*args, **kwargs)
            return cache[func_sig] 
        return wrapper
    return decorator

