def main():
    nums = [6, 4, 7, 2, 9, 3, 1, 5, 8]
    length = len(nums)
    print(nums)

    for i in range(1,length):
        j = i
        for j in range(i,0,-1):
            if nums[j]<nums[j-1]:
                nums[j],nums[j-1] = nums[j-1],nums[j]
        print(nums)
    print(nums)



if __name__ == "__main__":
    main()
