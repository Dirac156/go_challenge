class Difference:
    def __init__(self, a):
        self.__elements = a
        self.maximumDifference = 0

    # Add your code here
    def computeDifference(self):
        """ find maximum absolute value in an array """
        arr = self.__elements
        # sort the array
        arr.sort()
        # find the first  - last element 
        self.maximumDifference = abs(arr[0] - arr[len(arr) - 1]) 
        

# End of Difference class

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)