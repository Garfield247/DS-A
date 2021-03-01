public class Main {
    public static void main (String[] args) {
        int[] nums = new int[]{6, 4, 7, 2, 9, 3, 1, 5, 8};
        int len = nums.length;


        for (int i =0;i<len-1;i++){
            for ( int j=1;j<len-1-i;j++ ){
                if ( nums[j]>nums[j+1] ){
                    int tmp = nums[j];
                    nums[j] = nums[j+1];
                    nums[j+1] = tmp;
                }
            }
        }

        for (int i = 1; i<len; i++){
            System.out.print(nums[i]+",");
        }
    }

}
