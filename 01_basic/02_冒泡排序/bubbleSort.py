def main():
    nums = [6, 4, 7, 2, 9, 3, 1, 5, 8]

    for i in range(len(nums)-1):
        for j in range(len(nums)-1-i):
            if nums[j] > nums[j+1]:
                nums[j],nums[j+1] = nums[j+1],nums[j]
        print(nums)

    print(nums)


if __name__ == "__main__":
    main()
