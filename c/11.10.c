#include <stdbool.h>
void wiggleSort(int *nums, int numsSize) {
  if (numsSize <= 1) {
    return;
  }
  bool flag = true;
  if (nums[0] < nums[1]) {
    flag = true;
  } else {
    flag = false;
  }
  for (int i = 2; i < numsSize; i++) {
    if ((flag && nums[i - 1] < nums[i]) || (!flag && nums[i - 1] > nums[i])) {
      int tem = nums[i];
      nums[i] = nums[i - 1];
      nums[i - 1] = tem;
    }
    flag = !flag;
  }
}