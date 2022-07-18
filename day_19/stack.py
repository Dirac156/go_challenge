import sys

class Solution:
    # Write your code here
    def __init__(self):
        self.stack = []
        self.queue = []
    
    def pushCharacter(self, char):
        """ Add element at the beginiing of the stack LIFO """
        self.stack.insert(0, char)
    
        
    def enqueueCharacter(self, char):
        """ Add element at the end of the stack FIFO """
        self.queue.append(char)
        
    def popCharacter(self):
        """ Remove first item inside of the stack LIFO """
        item = self.stack[0]
        del self.stack[0]
        return item
    
    def dequeueCharacter(self):
        """ Remove first item inside of the queue FIFO """
        item = self.queue[0]
        del self.queue[0]
        return item


# read the string s
s=input()
#Create the Solution class object
obj=Solution()   

l=len(s)
# push/enqueue all the characters of string s to stack
for i in range(l):
    obj.pushCharacter(s[i])
    obj.enqueueCharacter(s[i])
    
isPalindrome=True
'''
pop the top character from stack
dequeue the first character from queue
compare both the characters
''' 
for i in range(l // 2):
    if obj.popCharacter()!=obj.dequeueCharacter():
        isPalindrome=False
        break
#finally print whether string s is palindrome or not.
if isPalindrome:
    print("The word, "+s+", is a palindrome.")
else:
    print("The word, "+s+", is not a palindrome.")    