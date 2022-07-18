class Node:
    def __init__(self,data):
        self.data = data
        self.next = None 
class Solution: 
    def insert(self,head,data):
            p = Node(data)           
            if head==None:
                head=p
            elif head.next==None:
                head.next=p
            else:
                start=head
                while(start.next!=None):
                    start=start.next
                start.next=p
            return head  
    def display(self,head):
        current = head
        while current:
            print(current.data,end=' ')
            current = current.next

    def removeNode(self, node, PreviousNode, NextNode):
        """ remove current node """
        PreviousNode.next = NextNode
    
    def removeDuplicates(self,head):
        #Write your code here
        node = head
        
        while node:
            innerNode = node.next
            previousNode = node
            while innerNode:
                if node.data ==innerNode.data:
                    self.removeNode(innerNode, previousNode, innerNode.next)
                innerNode = innerNode.next
            node = node.next
            
        return head

mylist= Solution()
T=int(input())
head=None
for i in range(T):
    data=int(input())
    head=mylist.insert(head,data)    
head=mylist.removeDuplicates(head)
mylist.display(head); 