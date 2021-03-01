def main():
    list1 = [6, 4, 7, 2, 9, 3, 1, 5, 8]
    print("list1:",list1)

    for i in range(len(list1)-1):
        # 找出当前最小值
        minPos = i
        for j in range(i+1,len(list1)):
            if list1[j] < list1[minPos]:
                minPos = j

        print(minPos)
        # 内容调换
        list1[i],list1[minPos] = list1[minPos],list1[i]

    print(list1)

if __name__ == "__main__":
    main()
