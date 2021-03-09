public class Main {
    public static void main (String[] args) {
        int[] nums = new int[]{6, 4, 7, 2, 9, 3, 1, 5, 8};
        printArray(nums);
        quickSort(nums,0,nums.length-1);
        printArray(nums);


    }

    public static void quickSort (int[] array,int start,int stop) {
        if (stop<start){
            return;
        }
        int i = start;
        int j = stop;
        int x = array[start];


        while (i<j) {
            while (i<j&&array[j]>x) {
               j--;
            }
            if (i<j){
                array[i] = array[j];
                i++;
            }
            while (i<j&&array[i]<x) {
               i++;
            }
            if (i<j){
                array[j] = array[i];
                j--;
            }

        }
        if (i==j) {
            array[i] = x;
        }

        printArray(array);
        quickSort(array,start,i-1);
        quickSort(array,i+1,stop);
    }

    public static void printArray (int[] array) {
        System.out.print("[");
        for (int i=0; i<array.length; i++){
            System.out.print(array[i]+",");
        }
        System.out.print("]\n");

    }
}
