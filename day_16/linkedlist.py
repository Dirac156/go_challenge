class Node:
    def __init__(self,data):
        self.data = data
        self.next = None 
class Solution: 
    def display(self,head):
        current = head
        while current:
            print(current.data,end=' ')
            current = current.next

    def insert(self,head,data): 
        """ insert new value into the linked list """
        if head is None:
            head = Node(data)
        else:
            a = head
            while a.next:
                a = a.next
            a.next = Node(data)
        return head

        # print(data, end=" ")

mylist= Solution()
T=int(input())
head=None
for i in range(T):
    data=int(input())
    head=mylist.insert(head,data)    
mylist.display(head); 	  