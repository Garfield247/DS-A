def main():
    nums = [6, 4, 7, 2, 9, 3, 1, 5, 8]
    print(nums)
    quickSort(nums,0,len(nums)-1)
    print(nums)

def quickSort(array, start, stop):
    if stop < start:
        return

    i = start
    j = stop
    x = array[start]

    while i < j:
        while i<j and array[j]>x:
            j-=1
        if i<j:
            array[i] = array[j]
            i+=1

        while i<j and array[i]<x:
            i+=1
        if i<j:
            array[j] = array[i]
            j-=1
    if i == j:
        array[i] = x

    print(array)
    quickSort(array,start,i-1)
    quickSort(array,i+1,stop)


if __name__ == "__main__":
    main()
