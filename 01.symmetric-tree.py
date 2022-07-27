# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        if root is None:
            return True
        else:
            return self.helper(root.left, root.right)
    
    def helper(self, left, right):
        # base case if both left/right are null
        if left == None and right == None:
            return True
        
        if left == None or right == None or left.val != right.val:
            return False
        
        return self.helper(left.left, right.right) & self.helper(left.right, right.left)
        
