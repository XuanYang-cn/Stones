'''Some examples of numpy'''

__author__ = 'Yang Xuan (jumpthepig@gmail.com)'


import numpy as np
import matplotlib.pyplot as plt
import matplotlib
import time
from functools import wraps


matplotlib.use('Qt5Agg')
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


def plot():
    cvalues = [20.1, 20.8, 21.9, 22.5, 22.7, 22.3, 21.8, 21.2, 20.9, 20.1]

    C =  np.array(cvalues)
    print(C)
    print(C * 9 / 5 + 32)

    plt.plot(C)
    plt.show()


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

