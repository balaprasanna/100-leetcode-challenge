# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        
        out= None
        while head != None:
            remaining = head.next
            head.next = out
            out = head
            head = remaining
            
            # 1-2-3-4-5
            # 2-3-4-5 (head)
            # 1-null (output)
            
        return out