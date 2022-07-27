# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def mergeTrees(self, t1: Optional[TreeNode], t2: Optional[TreeNode]) -> Optional[TreeNode]:
        if t1 is None and t2 is None:
            return None
        
        if t1 is None:
            return t2
        
        if t2 is None:
            return t1
        
        newRoot = TreeNode(val=t1.val + t2.val)
         # merge both left sub trees
        newRoot.left = self.mergeTrees(t1.left, t2.left)
         # merge both right sub trees
        newRoot.right = self.mergeTrees(t1.right, t2.right)
        return newRoot