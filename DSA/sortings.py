'''Sorting algorithms all together'''

__author__ = 'Yang Xuan (jumpthepig@gmail.com)'

import random
import time
from functools import wraps


def time_it(func):
    @wraps(func)
    def inner(*args):
        start = time.perf_counter()
        result = func(*args)

        delt = time.perf_counter() - start

        length_of_array = len(args[1])

        print(f"[{args[0].__name__}.{func.__name__}] sort length of [{length_of_array}] array costs: {delt*10e3:8.4f}ms ")
        return result
    return inner


class Sorting:

    @classmethod
    @time_it
    def bubble_sort(cls, nums):
        '''
        Time: O(N^2)
        In-place
        Stable
        '''
        for i in range(len(nums)-1):
            for j in range(len(nums)-1-i):
                if nums[j+1] < nums[j]:
                    nums[j+1], nums[j] = nums[j], nums[j+1]

def gen(num):
    '''generate list with random integers of num numbers'''
    return [random.randint(0, 100) for i in range(num)]


def is_sorted(nums):
    '''Check if a list is sorted'''
    for i in range(len(nums) - 1):
        if nums[i+1] < nums[i]:
            return False
    return True


def tests():
    test_nums = gen(100)
    Sorting.bubble_sort(test_nums)
    assert is_sorted(test_nums)


if __name__ == '__main__':
    tests()

