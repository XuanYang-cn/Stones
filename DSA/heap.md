# HEAP

## Maximum heap

parent node is greater than the children nodes, but the order of right child and left child doesn't matter.

Using a one-dimensional array to construct a Binary heap, because heap has to be complete.

parent node with index $i$, and its left child's index would be $2i+1$, and its right child's index would be $2i+2$

### Building a heap

Fist insert the data to the heap and check whether the heap properties are met, if the heap properties are violated, we reconstruct the heap in order to make it a valid heap.

- it is an $O(N)$ process to construct a heap
- reconstructi it if the heap properties are violated, and it takes $O(logN)$ time. $O(N) + O(logN) = O(N)$

### Deleting an item

Get rid of the item we want to delete, and then put the last item there, and reconstruct the heap to make sure the heap properties are valid.

Operation: deleting the root node O(1) + reconstruciton O(logN) = O(log(N))

Operation: deleting the artitrary node O(N) + reconstruciton O(logN) = O(N)

### Heapsort

- Comparison-based, in-place, un-stable sorting algorithm

Running time: O(NlogN)

Keep swapping the root with the last element and maintain heap properties.

After swapping with the root, we consider the last item to be sorted: nolonger part of the tree.

Find the minimum/maximun: O(1) very fast.

Insert nwe item: we can insert at the next acailable palce, so increment the array index and insert it O(1), then make sure the heap properties are met, it may take O(logN) time.

why logN? Because a node has at most logN parents so at most logN swaps are needed.

|Operation|Time Complexity|
|:---:|:----:|
|Find minimum/maximum|O(1)|
|Remove min/max|O(logN)|
|Insert|O(logN)|

