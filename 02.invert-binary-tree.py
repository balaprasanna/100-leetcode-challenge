# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        self.invert2(root)
        return root
    
    def invert(self, node):        
        # base case
        if node is None:
            return
        
        if node.left is None and node.right is None:
            return node
        
        left = self.invert(node.left)
        right = self.invert(node.right)
        node.left, node.right = right, left
        return node
    
    def invert2(self, node):
        if node is None:
            return
        
        tmp = node.left
        node.left = node.right
        node.right = tmp
        
        self.invert2(node.left)
        self.invert2(node.right)