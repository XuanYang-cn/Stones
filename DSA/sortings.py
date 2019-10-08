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

    @classmethod
    @time_it
    def selection_sort(cls, nums):
        '''
        Time: O(N^2)
        In-place
        Not stable
        Method: find the smallest and swap
        '''
        for i in range(len(nums)-1):
            index = i  # i is the first number to be sort at the moment

            for j in range(i+1, len(nums)):
                if nums[j] < nums[index]:
                    index = j

            if index != i:
                nums[i], nums[index] = nums[index], nums[i]

    @classmethod
    @time_it
    def insertion_sort(cls, nums):
        '''
        Time: O(N^2)
        In-place
        Stable
        Method: If smaller shift
        '''
        for i in range(1, len(nums)):
            j = i
            while j > 0 and nums[j-1] > nums[j]:
                nums[j-1], nums[j] = nums[j], nums[j-1]
                j -= 1

    @classmethod
    def _quick_sort_partition(cls, nums, low, high):
        pivot_index = (low + high) // 2
        nums[pivot_index], nums[high] = nums[high], nums[pivot_index]

        index = low
        for i in range(low, high):
            if nums[i] < nums[high]:
                nums[i], nums[index] = nums[index], nums[i]
                index += 1

        nums[index], nums[high] = nums[high], nums[index]
        return index

    @classmethod
    def _quick_sort(cls, nums, low, high):
        if low >= high:
            return

        pivot  = cls._quick_sort_partition(nums, low, high)
        cls._quick_sort(nums, low, pivot-1)
        cls._quick_sort(nums, pivot+1, high)

    @classmethod
    @time_it
    def quick_sort(cls, nums):
        '''
        Time: O(NlogN)
        In-place
        Not stable
        Method: divide and conquer algorithm
            pick an element from the array: pivot
            partition phase: rearrange the array so that all elements with value less
                than pivot item come before the pivot, while all elements with values
                greater than the pivot come after it, equal elements can go either way.
        '''
        cls._quick_sort(nums, 0, len(nums)-1)

    @classmethod
    def merge_sort(cls, nums):
        '''
        Time: O(NlogN)
        Not in-place, requier O(N) auxiliary space
        Stable
        Method: divide the array into two subarrays recursively.
            Sort these subarrays recursively with mergesort again
            If there is only a single item left in the subarray ->
                we consider it to be sorted by definition
            Merge the subarray to get the final sorted array
        '''

        def merge(left, right):
            i, j, result = 0, 0, []

            while i < len(left) and j < len(right):
                if left[i] <= right[j]:
                    result.append(left[i])
                    i += 1
                else:
                    result.append(right[j])
                    j += 1

            result.extend(left[i:])
            result.extend(right[j:])
            return result

        if len(nums) <= 1:
            return nums

        middle = len(nums) // 2
        left = cls.merge_sort(nums[:middle])
        right = cls.merge_sort(nums[middle:])

        return merge(left, right)

    @classmethod
    @time_it
    def heapsort(cls, array):
        def sink(array, index, upto):

            while index <= upto:
                left = index * 2 + 1
                right = index * 2 + 2
                k = index
                if left <= upto:
                    if right <= upto and array[right] < array[left]:
                        k = right
                    else:
                        k = left
                if array[index] > array[k]:
                    array[index], array[k] = array[k], array[index]
                    index = k
                else:
                    break

        for i in range((len(array) - 1)//2, -1, -1):
            sink(array, i, len(array) - 1)

        result, upto = [], len(array) - 1
        for i in range(len(array)):
            result.append(array[0])
            array[0], array[upto] = array[upto], array[0]

            upto -= 1
            sink(array, 0, upto)
        return result


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
    # Test bubble sort
    test_nums = gen(100)
    Sorting.bubble_sort(test_nums)
    assert is_sorted(test_nums)

    # Test selection sort
    test_nums01 = gen(100)
    Sorting.selection_sort(test_nums01)
    assert is_sorted(test_nums01)

    # Test insertion sort
    test_nums02 = gen(100)
    Sorting.insertion_sort(test_nums02)
    assert is_sorted(test_nums02)

    # Test quick sort
    test_nums03 = gen(100)
    Sorting.quick_sort(test_nums03)
    assert is_sorted(test_nums03)

    # Test merge sort
    test_nums04 = gen(100)
    assert is_sorted(Sorting.merge_sort(test_nums04))

    # Test heap sort
    test_nums05 = gen(100)
    assert is_sorted(Sorting.heapsort(test_nums05))
    
if __name__ == '__main__':
    tests()
