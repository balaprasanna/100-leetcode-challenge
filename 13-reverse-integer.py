class Solution:
    def reverseInt(self, x: int) -> bool:
        y = x
        rev_y = 0
        
        while y > 0:
            rem = y % 10
            y = int(y / 10)
            rev_y = (rev_y*10) + rem
        
        return rev_y
        
    
    