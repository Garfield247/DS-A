public class MergeSort {

    public static void main (String[] args) {
        int[] nums = new int[]{2,4,6,7,1,3,5,8,9};
        sort(nums,0,nums.length-1);
        printArray(nums);

    }

    public static void merge (int[] array,int leftPoint,int midPoint,int rightBound) {
        int mid = midPoint-1;
        int len =rightBound-leftPoint+1;
        int[] temp = new int[len];
        int i = leftPoint;
        int j = midPoint;
        int k =0;

        while (i<=mid && j<=rightBound) {
            if (array[i]<=array[j]){
                System.out.printf("i|array[%d]:%d\n",i,array[i]);
                temp[k++] = array[i++];
            }else{
                System.out.printf("j|array[%d]:%d\n",j,array[j]);
                temp[k++] = array[j++];
            }
        }
        while (i<=mid){
            temp[k++] = array[i++];
        }
        while (j<=rightBound){
            temp[k++] = array[j++];
        }

        for (int x=0;x<len;x++){
            array[x+leftPoint] = temp[x];
        }

    }
    public static void sort(int[] array,int leftPoint,int rightBound) {
        if (leftPoint == rightBound){
            return;
        }
        int midPoint = leftPoint + (rightBound-leftPoint)/2;
        sort(array,leftPoint,midPoint);
        sort(array,midPoint+1,rightBound);

        printArray(array);
        System.out.printf("leftPoint:%d | midPoint:%d | rightBound:%d | \n",leftPoint,midPoint+1,rightBound);
        sliceArray(array,leftPoint,rightBound);
        merge(array,leftPoint,midPoint+1,rightBound);

    }

    public static void printArray (int[] array) {
        System.out.print("[");
        for (int index = 0; index < array.length;index++){
            System.out.print(array[index]+",");
        }
        System.out.print("]\n");
    }
    public static void sliceArray (int[] array ,int start ,int stop) {
        System.out.print("[");
        for (int index = start; index < stop;index++){
            System.out.print(array[index]+",");
        }
        System.out.print("]\n");
    }
}
