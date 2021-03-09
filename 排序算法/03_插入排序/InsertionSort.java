public class InsertionSort {

    public static void main (String[] args) {
        int[] nums = new int[]{6, 4, 7, 2, 9, 3, 1, 5, 8};
        int len = nums.length;
        printArray(nums);

        for (int i=1;i<len;i++){
            for (int j=i;j>=1;j--){
                if (nums[j]<nums[j-1]){
                    moveLeft(nums,j);
                }
            }
            printArray(nums);
        }
    }

    public static void printArray( int[] arr) {

        System.out.print("[");
        for (int i=0; i<arr.length ;i++){
            System.out.print(arr[i]+",");
        }
        System.out.println("]");


    }

    private static void moveLeft(int[] arr, int index) {

        int tmp = arr[index];
        arr[index]=arr[index-1];
        arr[index-1] = tmp;
    }

}
