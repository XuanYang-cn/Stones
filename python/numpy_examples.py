'''Some examples of numpy'''

__author__ = 'Yang Xuan (jumpthepig@gmail.com)'


import numpy as np
import time
from functools import wraps


n = 1_000_000


def time_it(func, repeat=None):
    @wraps(func)
    def inner(*args, **kw):
        if repeat:
            for _ in range(repeat):
                start = time.perf_counter()
                func(*args, **kw)
                delt = time.perf_counter() - start
                print(f"{func.__name__:18s}: {delt*1000:8.6}ms")
        return
    return inner



@time_it(repeat=3)
def pure_python_gen(n):
    x = list(range(n))
    return


@time_it(repeat=3)
def numpy_gen(n):
    y = np.arange(n)


if __name__ == '__main__':
    pure_python_gen(n)
    numpy_gen(n)

