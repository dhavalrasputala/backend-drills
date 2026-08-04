## Day 8: Subarray Sum Equals K (LeetCode 560)
## Status: LOOKED AT SOLUTION

## The Pattern: Prefix Sum + Hash Map
## Explanation: 
 - Instead of checking every possible subarray (O(n^2)), we keep a running total (prefix sum).
 - If the current running total is `total`, and we have seen `total - k` before, 
 - it means the subarray between those two points adds up to exactly `k`.
 - We use a Hash Map to store how many times we've seen each prefix sum.

## Step-by-Step Logic:
 1. Initialize a dictionary `prefix_sums` with {0: 1} (We've seen a sum of 0 exactly once).
 2. Initialize `total` = 0 and `count` = 0.
 3. Loop through each number `n` in the array:
    a. Add `n` to `total`.
    b. Check if `total - k` is in the dictionary. If it is, add its value to `count`.
    c. Add `total` to the dictionary (increment its count by 1).
 4. Return `count`.

# Time Complexity: O(N) - We loop through the array once.
# Space Complexity: O(N) - We store prefix sums in a hash map.
```
class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        sub_num = {0:1}
        total = count = 0

        for n in nums:
            total += n
            
            if total - k in sub_num:
                count += sub_num[total-k]
            
            sub_num[total] = 1 + sub_num.get(total, 0)
        
        return count
```        

# Reflection: 
- I didn't realize that `current_sum - k` means there is a subarray ending here that equals k.
- Next time I see "contiguous subarray" or "sum equals k", I will immediately think: Prefix Sums + Hash Map.
