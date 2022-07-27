class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        foundDuplicate = False
        numSet = {}
        for n in nums:
            if n in numSet: 
                foundDuplicate = True
                break
            else:
                numSet[n] = 1
        
        return foundDuplicate