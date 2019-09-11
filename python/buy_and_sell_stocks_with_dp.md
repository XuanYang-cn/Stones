# 股票问题-动态规划详解

使用状态机的方法解决所有买卖股票问题，状态机设计、状态转换其实是动态规划（Dynamic Programming， DP）的核心问题。

本文通过分析出这六个问题最通用的解法（买卖股票IV），再结合六个不同的买卖股票问题分别进行简化，逐一攻破动态规划的思路，让你最终成为股市大牛（FAUX）。

其中，买卖股票IV是最一般情况，限制最大买卖次数为k次，虽然题目中没有提到k有效的范围，但我们可以很容易的推测出k隐含的条件是 **0 < k <= N/2**, N为天数（数组的长度）。因为买和卖分别需要花1天，有效的交易次数最多为 N/2，所以当 k > N/2 时，第IV题会退化为第II题，也就是 k = +Infinity。通过对k隐含条件的分析，我们可以大幅减少生成 DP table 的额外空间大小.

买卖股票的六个问题中，I、II、III、IV之间最主要的区别是买卖次数 k 的限定。

|买卖股票问题|k|
|:---:|:---:|
|I|1|
|II|+Infinity|
|III|2|
|IV|K|
|冷冻期|+Infinity|
|手续费|+Infinity|

冷冻期和手续费两题相当于II的延伸，因此在最后进行分析

## 状态和状态转移方程

#### 状态

动态规划的思想就是穷举所有状态下的最优解。这就需要我们分析出具体问题下哪些是“状态”，哪些是能够让“状态”之间互相转移的“行为”，并且构建出状态转移方程。一旦转移方程构建出来了，这道问题就解决了八成。

对这题来说，第一个状态是第几天（i），第二个状态是允许最大的交易次数（k），第三个状态是当前持有股票状态（不妨设 0 为不持有股票，1 为持有股票）。使用 3 维数组就可以描述所有的状态:

    dp[i][k]{0, 1}, 0<=i<=N-1, 1<k<=K


每天可以选择的“行为”有： 买入（buy）、卖出（sell）、无操作（rest）

因此对于一个一般状态来说：

    dp[i][k][0]表示现在是第i天，最多还剩k次交易次数，手中不持有股票时的最大收益

    dp[1][3][1]表示现在是第一天，最多还可以交易3次，手中持有股票时的最大收益

    # 状态转移方程的思路
    for i in range(N):
        for k in range(1, K):
          dp[i][k][0] = max(buy, sell, rest)
          dp[i][k][1] = max(buy, sell, rest)

#### 状态转移方程
下图则表示了状态转移的过程：

**图片**

根据这个图写出状态转移方程：

    # 第i天，不持有股票的最大收益 = 前一天无股票并【无操作】或
                                前一天持有股票并【买入】- 【买入金额】
                                中的收益高者
    dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] - prices[i])   （1）

    # 第i天，持有股票的最大收益 = 前一天持有股票并【无操作】或
                              前一天无股票并【卖出】 + 【卖出金额】
                              中的收益高者
    dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] + prices[i]) （2）

在状态转移方程（1）中，**dp[i-1][k][0]** 表示第 i - 1 天时手上无股票，并且在第 i 天选择【无操作】的收益， **dp[i-1][k][1] - prices[i]** 表示第 i - 1 天时手上有股票并且在第 i 天选择【买入股票】的收益。

在状态转移方程（2）中， **max(dp[i-1][k][1]** 表示第 i - 1 天手上有股票并且在第 i 天选择【无操作】的收益， **dp[i-1][k-1][0] + prices[i]** 表示第 i - 1 天手上没有股票，在第 i 天选择【卖出股票】的收益，这里的 **k - 1** 表示进行了一次买卖操作， 也可以选择在购买股票的时候将 k 减 1，只要保证一次买入加卖出算作一次操作即可。

#### 初始状态

确定了状态转移方程之后，现在来分析初始状态

    dp[-1][k][0] = 0  # 交易未开始时没有持有股票，收益是0
    dp[-1][k][1] = -Infinity  # 交易未开始就持有股票，用负无穷表示行为不可能

    # k 从 1 开始计算
    dp[i][0][0] = 0  # 没有交易次数时无股票， 收益是0
    dp[i][0][1] = -Infinity  # 没有交易次数时持有股票，用负无穷表示行为不可能


#### 总结方程

```python
# 初始状态， 这里 -1， -Infinity在程序里的表现方法见下面代码部分
dp[-1][k][0] = dp[i][0][0] = 0
dp[-1][k][1] = dp[i][0][1] = -infinity

# 状态转移方程
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
```
我们最后需要求出来的就是 **dp[N-1][k][0]** 也就是在最后一天手中没有股票时的最大收益

## 具体问题具体分析

#### 买卖股票的最佳时机I

k = 1

状态转移方程就变为了：

```python
# 初始状态
dp[-1][1][0] = dp[i][0][0] = 0
dp[-1][1][1] = dp[i][0][1] = -infinity

# 状态转移方程
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i])

# 因为 dp[i-1][0][0] = 0 且其他的状态转移与 k 无关，所以继续简化
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], - prices[i])

# 又因为当前状态只与之前两个状态相关，所以可以只用两个变量来存储，减少空间复杂度
dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
dp_i_1 = max(dp_i_1, -prices[i])

# 这时候初始状态就变成
dp_i_0 = 0
dp_i_1 = - prices[0]  # 第一天购买股票的收益
```

代码为：
```python
def max_profit_I(prices):
    n = len(prices)
    if n == 0:
        return 0
    for i in range(n):
        if i - 1 == -1:
            dp_i_0 = 0
            dp_i_1 = -prices[i]
        else:
            dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
            dp_i_1 = max(dp_i_1, -prices[i])
    return dp_i_0
```
#### 买卖股票的最佳时机II

k = + Infinity， 相当于 k - 1 = + Infinity， k 对状态转移方程没有影响

所以状态方程可以简化为：

```python
# 初始状态
dp[-1][0]  = 0
dp[-1][1]  = -infinity

# 状态转移方程
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])

# 又因为当前状态只与之前两个状态相关，所以可以只用两个变量来存储，减少空间复杂度
# 但是这里需要临时变量存储 dp_i_0 的值，因为它在第一个方程中产生了变化
temp = dp_i_0
dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
dp_i_1 = max(dp_i_1, temp - prices[i])

# 这时候初始状态就变成
dp_i_0 = 0  # 此时 i = -1
dp_i_1 = -infinity  # 此时 i = -1
```

代码为：

```python
def max_profit_II(prices):
    n = len(prices)
    if n == 0:
        return 0
    dp_i_0 = 0
    dp_i_1 = float("-inf")
    for i in range(0, n):
        temp = dp_i_0
        dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = max(dp_i_1, temp - prices[i])
    return dp_i_0
```

#### 买卖股票的最佳时机III

k = 2, 状态转移方程如下

```python
# 初始状态
dp[-1][1][0] = dp[i][0][0] = 0
dp[-1][1][1] = dp[i][0][1] = -infinity

# 状态转移方程(A)
dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
dp[i][1][1] = max(dp[i-1][1][1], - prices[i])     

# 也可以直接用 k 表示 (B)
dp[-1][k][0] = dp[i][0][0] = 0
dp[-1][k][1] = dp[i][0][1] = -infinity

dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

```

代码如下：
```python
# 如果是方程（A），那么答案的写法就和之前差不多
def max_profit_III_A(prices):
    dp_i10 = 0, dp_i11 = - float("Inf")
    dp_i20 = 0, dp_i21 = - float("Inf")
    for price in prices:
        dp_i20 = max(dp_i20, dp_i21 + price);
        dp_i21 = max(dp_i21, dp_i10 - price);
        dp_i10 = max(dp_i10, dp_i11 + price);
        dp_i11 = max(dp_i11, - price);
    return dp_i20;

# 如果是方程（B），那么需要把对 k 的循环添加进去，而且生成的 DP table 是三维的，需要注意生成方法。
def max_profit_III_B(prices):
    n = len(prices)
    max_k = 2
    if n == 0:
        return 0

    # 生成三维数组，最外层 size 为 n，中间层 size 为 K + 1， 最内层 size 为 2
    dp = [[[0 for _ in range(2)] for _ in range(max_k+1)] for _ in range(n)]

    for i in range(n):
        for k in range(max_k, 0, -1):
            if i == 0:
                dp[i][k][0] = 0
                dp[i][k][1] = - prices[i]
                continue

            dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
            dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
    return dp[n-1][max_k][0]
```

#### 买卖股票的最佳时机IV

k = K

状态转移方程：
```python
dp[-1][k][0] = dp[i][0][0] = 0
dp[-1][k][1] = dp[i][0][1] = -infinity

# 状态转移方程
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
```

代码写法和第三题类似，注意 k 是有隐藏约束的

```python
def max_profit_IV(k, prices):
    max_k = k
    n = len(prices)
    if n == 0:
        return 0
    if max_k > n//2:
        return max_profit_II(prices)  # 当 k 大于 N/2 时，变成了第二题。

    dp = [[[0 for _ in range(2)] for _ in range(max_k + 1)] for _ in range(n)]
    for i in range(n):
        for k in range(max_k, 0, -1):
            if i == 0:
                dp[i][k][0] = 0
                dp[i][k][1] = -prices[i]
                continue
            dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
            dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
    return dp[n-1][max_k][0]

```
#### 最佳买卖股票时机含冷冻期

k = + Infinity

冷冻期就意味着购买股票时，当天的状态转移与昨天无关，但是与前天相关，于是状态转移方程可以变为

```python
# 状态转移方程
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][1] - prices[i])
```

代码

```python
def max_profit_cooldown(prices):
    n = len(prices)
    if n == 0:
        return 0
    dp_i_0, dp_i_1 = 0, float("-Inf")
    dp_pre_0 = 0  # 存储前天不持有股票的利润
    for i in range(n):
        temp = dp_i_0
        dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = max(dp_i_1, dp_pre_0 - prices[i])
        dp_pre_0 = temp
    return dp_i_0
```
#### 买卖股票的最佳时期含手续费

k = + Infinity

有手续费意味着每次买入的时候需要再减去手续费

状态转移方程为：

```python
# 状态转移方程
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][1] - prices[i] - fee)
```

代码为
```python
def max_profit_fee(prices, fee):
    n = len(prices)
    if n == 0:
        return 0
    dp_i_0, dp_i_1 = 0, float("-Inf")

    for i in range(n):
        temp = dp_i_0
        dp_i_0 = max(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = max(dp_i_1, temp - prices[i] - fee)
    return dp_i_0
```

## 总结

哒哒，至此六道股票问题就已经完全解决了
