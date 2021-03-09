def main():
    nums = [213, 42, 32, 432, 12, 543, 256, 65, 856, 74, 234, 532, 43, 92, 63, 433]


    shellAort(nums)

def shellAort(array):
    gap = int(len(array)/2)
    while gap>0:
        for i in range(gap,len(array),gap):
            for j in range(i,gap-1,-gap):
                if array[j]<array[j-gap]:
                    array[j],array[j-1] = array[j-1],array[j]
        gap = int(gap/2)
    print(array)

if __name__ == "__main__":
    main()
