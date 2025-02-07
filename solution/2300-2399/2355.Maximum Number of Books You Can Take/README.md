# [2355. Maximum Number of Books You Can Take](https://leetcode.cn/problems/maximum-number-of-books-you-can-take)

[English Version](/solution/2300-2399/2355.Maximum%20Number%20of%20Books%20You%20Can%20Take/README_EN.md)

## 题目描述

<!-- 这里写题目描述 -->

<p>You are given a <strong>0-indexed</strong> integer array <code>books</code> of length <code>n</code> where <code>books[i]</code> denotes the number of books on the <code>i<sup>th</sup></code> shelf of a bookshelf.</p>

<p>You are going to take books from a <strong>contiguous</strong> section of the bookshelf spanning from <code>l</code> to <code>r</code> where <code>0 &lt;= l &lt;= r &lt; n</code>. For each index <code>i</code> in the range <code>l &lt;= i &lt; r</code>, you must take <strong>strictly fewer</strong> books from shelf <code>i</code> than shelf <code>i + 1</code>.</p>

<p>Return <em>the <strong>maximum</strong> number of books you can take from the bookshelf.</em></p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> books = [8,5,2,7,9]
<strong>Output:</strong> 19
<strong>Explanation:</strong>
- Take 1 book from shelf 1.
- Take 2 books from shelf 2.
- Take 7 books from shelf 3.
- Take 9 books from shelf 4.
You have taken 19 books, so return 19.
It can be proven that 19 is the maximum number of books you can take.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> books = [7,0,3,4,5]
<strong>Output:</strong> 12
<strong>Explanation:</strong>
- Take 3 books from shelf 2.
- Take 4 books from shelf 3.
- Take 5 books from shelf 4.
You have taken 12 books so return 12.
It can be proven that 12 is the maximum number of books you can take.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> books = [8,2,3,7,3,4,0,1,4,3]
<strong>Output:</strong> 13
<strong>Explanation:</strong>
- Take 1 book from shelf 0.
- Take 2 books from shelf 1.
- Take 3 books from shelf 2.
- Take 7 books from shelf 3.
You have taken 13 books so return 13.
It can be proven that 13 is the maximum number of books you can take.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= books.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= books[i] &lt;= 10<sup>5</sup></code></li>
</ul>

## 解法

<!-- 这里可写通用的实现逻辑 -->

**方法一：单调栈 + 动态规划**

设 $dp[i]$ 表示以 $books[i]$ 结尾时能取走的书的最大数量。

若从 $i$ 到 $0$ 可以取成一个公差为 $1$ 的等差数列，那么 $dp[i]$ 可以直接通过等差数列求和算出。

若从 $i$ 到 $0$ 不能取成一个公差为 $-1$ 的等差数列，即这个规律在某个 $j$ 处断掉了（$0 \le j \lt i$），那么一定有 $books[j] \lt books[i] - (i-j)$，也即 $books[j] - j \lt books[i] - i$，利用单调栈在新数组 $books[i] - i$ 的每个位置，找到左边第一个比它严格小的数的位置，可以求出符合题意的 $j$，此时 $dp[i]=dp[j] + \sum_{k=j+1}^{i} (books[k]-k)$。

答案为 $max(dp[i])$。

时间复杂度 $O(n)$。

<!-- tabs:start -->

### **Python3**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```python
class Solution:
    def maximumBooks(self, books: List[int]) -> int:
        nums = [v - i for i, v in enumerate(books)]
        n = len(nums)
        left = [-1] * n
        stk = []
        for i, v in enumerate(nums):
            while stk and nums[stk[-1]] >= v:
                stk.pop()
            if stk:
                left[i] = stk[-1]
            stk.append(i)
        ans = 0
        dp = [0] * n
        dp[0] = books[0]
        for i, v in enumerate(books):
            j = left[i]
            cnt = min(v, i - j)
            u = v - cnt + 1
            s = (u + v) * cnt // 2
            dp[i] = s + (0 if j == -1 else dp[j])
            ans = max(ans, dp[i])
        return ans
```

### **Java**

<!-- 这里可写当前语言的特殊实现逻辑 -->

```java
class Solution {
    public long maximumBooks(int[] books) {
        int n = books.length;
        int[] nums = new int[n];
        for (int i = 0; i < n; ++i) {
            nums[i] = books[i] - i;
        }
        int[] left = new int[n];
        Arrays.fill(left, -1);
        Deque<Integer> stk = new ArrayDeque<>();
        for (int i = 0; i < n; ++i) {
            while (!stk.isEmpty() && nums[stk.peek()] >= nums[i]) {
                stk.pop();
            }
            if (!stk.isEmpty()) {
                left[i] = stk.peek();
            }
            stk.push(i);
        }
        long ans = 0;
        long[] dp = new long[n];
        dp[0] = books[0];
        for (int i = 0; i < n; ++i) {
            int j = left[i];
            int v = books[i];
            int cnt = Math.min(v, i - j);
            int u = v - cnt + 1;
            long s = (long) (u + v) * cnt / 2;
            dp[i] = s + (j == -1 ? 0 : dp[j]);
            ans = Math.max(ans, dp[i]);
        }
        return ans;
    }
}
```

### **C++**

```cpp
using ll = long long;

class Solution {
public:
    long long maximumBooks(vector<int>& books) {
        int n = books.size();
        vector<int> nums(n);
        for (int i = 0; i < n; ++i) nums[i] = books[i] - i;
        vector<int> left(n, -1);
        stack<int> stk;
        for (int i = 0; i < n; ++i) {
            while (!stk.empty() && nums[stk.top()] >= nums[i]) stk.pop();
            if (!stk.empty()) left[i] = stk.top();
            stk.push(i);
        }
        vector<ll> dp(n);
        dp[0] = books[0];
        ll ans = 0;
        for (int i = 0; i < n; ++i) {
            int v = books[i];
            int j = left[i];
            int cnt = min(v, i - j);
            int u = v - cnt + 1;
            ll s = 1ll * (u + v) * cnt / 2;
            dp[i] = s + (j == -1 ? 0 : dp[j]);
            ans = max(ans, dp[i]);
        }
        return ans;
    }
};
```

### **Go**

```go
func maximumBooks(books []int) int64 {
	n := len(books)
	nums := make([]int, n)
	left := make([]int, n)
	for i, v := range books {
		nums[i] = v - i
		left[i] = -1
	}
	stk := []int{}
	for i, v := range nums {
		for len(stk) > 0 && nums[stk[len(stk)-1]] >= v {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}
	dp := make([]int, n)
	dp[0] = books[0]
	ans := 0
	for i, v := range books {
		j := left[i]
		cnt := min(v, i-j)
		u := v - cnt + 1
		s := (u + v) * cnt / 2
		dp[i] = s
		if j != -1 {
			dp[i] += dp[j]
		}
		ans = max(ans, dp[i])
	}
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### **TypeScript**

```ts

```

### **...**

```

```

<!-- tabs:end -->
