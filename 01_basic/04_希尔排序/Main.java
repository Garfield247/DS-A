public class Main {

    public static void main (String[] args) {
        int[] nums = new int[]{213, 42, 32, 432, 12, 543, 256, 65, 856, 74, 234, 532, 43, 92, 63, 433};
        shellSort(nums);
    }
    private static void shellSort(int[] array) {
        int len = array.length;
        for (int gap=len/2;gap>0;gap/=2){
            System.out.println(gap);
            for (int i=gap;i<len;i++){
                for (int j=i;j>=gap;j-=gap){
                    if (array[j]<array[j-gap]){
                        //swap(array[j],array[j-1]);
                        swap(array,gap,j);
                        printArray(array);
                    }
                }
            }
        }
        printArray(array);

    }

    //这里是这样写会是值传递,原array并不会发生改变
    //public static void swap(int a,int b) {
        //int tmp = a;
        //a = b;
        //b = tmp;
    //}

    // 这样写才是引用传递,能够正常对调array的两个元素
    public static void swap(int[]array,int gap,int index) {
        int tmp = array[index];
        array[index] = array[index-gap];
        array[index-gap] = tmp;
    }
    public static void printArray(int[] array) {
        System.out.print("[");
        for (int i = 0; i < array.length; i++) {
            System.out.print(array[i]+",");
        }
        System.out.print("]\n");

    }
}
