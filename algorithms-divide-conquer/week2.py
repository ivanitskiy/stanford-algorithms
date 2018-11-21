"""
            [3, 2, 4, 9, 10, 1]
             /               \
        [3, 2, 4]        [9, 10, 1]
       /      \            /    \
    [3]      [2, 4]     [9]   [10, 1]

then merge two arrays (a=[3], b=[2,4]) -> [2,3,4])
result = []
i,j=0
if a[i] > b[j] then result.append(b[j]) j++
if b[j] > a[i] then result.append(a[i]) i++
"""


def merge(a,b):
    la = len(a)
    lb = len(b)
    i = 0
    j = 0
#    inversions = 0
    retval = []
    print a,b
    for k in range(la+lb):
        if (i < la) and (j < lb):
            if a[i] < b[j]:
                retval.append(a[i])
                i += 1
            elif b[j] < a[i]:
                retval.append(b[j])
                print "inversions", a[i],b[j]
#                inversions +=1
                j += 1
        elif (i < la):
            retval.append(a[i])
            i += 1
        elif (j < lb):
            retval.append(b[j])
            j += 1
    return retval

def mergeSort(array):
    l = len(array)
    if l > 1:
        a1 = mergeSort(array[:l/2])
        a2 = mergeSort(array[l/2:])
        return merge(a1, a2)
    return array
# print mergeSort([9,3,1])
print mergeSort([3,2,4,9,10,1, 5,6,7,8])
# merge([2,3],[4,9])
