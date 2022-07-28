"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def postorder(self, root: 'Node') -> List[int]:
        output = []
        if root is not None:
            self.helper(root, output)
        return output
    
    def helper(self, root, output):
        if root is None:
            return
        
        if root.children:
            for c in root.children:
                self.helper(c, output)
        
        output.append(root.val)
        