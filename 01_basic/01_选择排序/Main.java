// 选择排序
public class Main {
    public static void main (String[] args) {
        int[] nums = new int[]{6, 4, 7, 2, 9, 3, 1, 5, 8};
        int len = nums.length;
        for (int i=0;i<len-1;i++){
            int minPos = i;

            for (int j=i+1;j<len ;j++ ) {
                if (nums[j]<nums[minPos]){
                    minPos = j;
                }

            }
            System.out.println("minPos:"+minPos);
            int tmp = nums[i];
            nums[i] =  nums[minPos];
            nums[minPos] = tmp;
        }

        for (int i=0; i<len; i++){
            System.out.print(nums[i]+",");
        }
    }
}
