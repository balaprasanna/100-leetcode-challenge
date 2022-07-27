class Solution:
    def twoSum(self, nums, target):
        
        halfDict = {}
        
        for idx, num in enumerate(nums):
            if target-num in halfDict:
                return halfDict[target-num], idx
            else:
                halfDict[num] = idx