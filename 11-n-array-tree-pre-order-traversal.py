"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        out = []
        if root is None:
            return out
        else:
            self.helper(root, out)
        return out
        
    def helper(self, root, out):
        if root is None:
            return
        
        if root.children is None:
            return
        
        out.append(root.val)
        for c in root.children:
            self.helper(c, out)
        
        