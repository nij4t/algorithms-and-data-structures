# lru_cache is a decorator function with a memoizing
# callable that saves most recent calls. 
# It can save time when an expensive or I/O bound 
# function is periodically called with the same arguments.
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

