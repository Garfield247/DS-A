
def merge(array,start,mid,stop):

    temp = [0] * (stop - start + 1)
    print(temp)
    i = start
    j = mid
    k = 0

    while i < mid and j <= stop:
        if array[i] <= array[j]:
            temp[k] = array[i]
            k += 1
            i += 1
        else:
            temp[k] = array[j]
            k += 1
            j += 1

    while i < mid:
        temp[k] = array[i]
        k += 1
        i += 1

    while j <= stop:
        temp[k] = array[j]
        k += 1
        j += 1

    print(temp)
    for x in range(len(temp)):
        array[x+start] = temp[x]




def mysort(array, start, stop):
    if start == stop:
        return

    mid = start + int((stop - start) / 2)

    mysort(array, start, mid)
    mysort(array, mid+1, stop)

    print(array, start, mid+1, stop)
    merge(array, start, mid+1, stop)


def main():
    # nums = [2, 4, 6, 7, 1, 3, 5, 8, 9]
    nums = [6, 4, 7, 2, 9, 3, 1, 5, 8]
    mysort(nums, 0, len(nums) - 1)
    print(nums)

if __name__ == "__main__":
    main()
